package dagql

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/dagger/dagql/idproto"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
)

type Resolver interface {
	Resolve(context.Context, *ast.FieldDefinition, map[string]Literal) (any, error)
}

type FieldSpec struct {
	// Name is the name of the field.
	// Name string

	// Args is the list of arguments that the field accepts.
	Args []ArgSpec

	// Type is the type of the field's result.
	Type *ast.Type

	// Meta indicates that the field has no impact on the field's result.
	Meta bool

	// Pure indicates that the field is a pure function of its arguments, and
	// thus can be cached indefinitely.
	Pure bool
}

type ArgSpec struct {
	// Name is the name of the argument.
	Name string
	// Type is the type of the argument.
	Type *ast.Type
	// Default is the default value of the argument.
	Default *Literal
}

type Literal struct {
	*idproto.Literal
}

func (lit Literal) ToAST() *ast.Value {
	switch x := lit.Value.(type) {
	case *idproto.Literal_Int:
		return &ast.Value{
			Kind: ast.IntValue,
			Raw:  fmt.Sprintf("%d", lit.GetInt()),
		}
	case *idproto.Literal_String_:
		return &ast.Value{
			Kind: ast.IntValue,
			Raw:  fmt.Sprintf("%d", lit.GetInt()),
		}
	default:
		panic(fmt.Errorf("cannot convert %T to *ast.Value", x))
	}
}

type Selector struct {
	Field string
	Args  map[string]Literal
}

// Per the GraphQL spec, a Node always has an ID.
type Node interface {
	ID() *idproto.ID

	Typed
	Resolver
}

type TypeResolver interface {
	isType()
}

// var IDResolver = ScalarResolver[*idproto.ID]{
// 	ToResponse: func(value *idproto.ID) (any, error) {
// 		proto, err := proto.Marshal(value)
// 		if err != nil {
// 			return nil, err
// 		}
// 		compressed := new(bytes.Buffer)
// 		lw := lz4.NewWriter(compressed)
// 		if _, err := lw.Write(proto); err != nil {
// 			return nil, err
// 		}
// 		if err := lw.Close(); err != nil {
// 			return nil, err
// 		}
// 		log.Printf("compressed: %q", compressed.String())
// 		uncompressed := new(bytes.Buffer)
// 		lz4.NewReader(bytes.NewBuffer(compressed.Bytes())).WriteTo(uncompressed)
// 		log.Printf("uncompressed: %q", uncompressed.String())
// 		return base64.URLEncoding.EncodeToString(compressed.Bytes()), nil
// 	},
// 	FromQuery: func(lit ast.Value) (*idproto.ID, error) {
// 		switch x := lit.(type) {
// 		case *ast.StringValue:
// 			bytes, err := base64.URLEncoding.DecodeString(x.Value)
// 			if err != nil {
// 				return nil, err
// 			}
// 			var idproto idproto.ID
// 			if err := proto.Unmarshal(bytes, &idproto); err != nil {
// 				return nil, err
// 			}
// 			return &idproto, nil
// 		default:
// 			return nil, fmt.Errorf("cannot convert %T to *idproto.ID", x)
// 		}
// 	},
// }

type Nullable interface {
	NullableValue() any
}

type Optional[T any] struct {
	Value T
	Valid bool
}

func (n Optional[T]) NullableValue() any {
	return n.Value
}

func Func[T any, A any, R Typed](fn func(ctx context.Context, self T, args A) (R, error)) Field[T] {
	var argSpecs []ArgSpec
	var zeroArgs A
	argsType := reflect.TypeOf(zeroArgs)
	if argsType != nil {
		if argsType.Kind() != reflect.Struct {
			panic(fmt.Sprintf("args must be a struct, got %T", zeroArgs))
		}

		for i := 0; i < argsType.NumField(); i++ {
			field := argsType.Field(i)
			argName := field.Tag.Get("arg")
			if argName == "" {
				argName = strcase.ToLowerCamel(field.Name)
			}

			var argDefault *Literal
			defaultJSON := []byte(field.Tag.Get("default")) // TODO: would make more sense to GraphQL-Unmarshal this
			if len(defaultJSON) > 0 {
				dec := json.NewDecoder(bytes.NewReader(defaultJSON))
				dec.UseNumber()

				var defaultAny any
				if err := dec.Decode(&defaultAny); err != nil {
					panic(err)
				}

				argDefault = &Literal{idproto.LiteralValue(defaultAny)}
			}

			argType, err := TypeOf(reflect.New(field.Type).Interface())
			if err != nil {
				panic(err)
			}

			argSpecs = append(argSpecs, ArgSpec{
				Name:    argName,
				Type:    argType,
				Default: argDefault,
			})
		}
	}

	var zeroRet R
	retType, err := TypeOf(zeroRet)
	if err != nil {
		panic(err)
	}

	return Field[T]{
		Spec: FieldSpec{
			Args: argSpecs,
			Type: retType,
		},
		Func: func(ctx context.Context, self T, argVals map[string]Literal) (any, error) {
			var args A

			argsVal := reflect.ValueOf(&args)

			for i, arg := range argSpecs {
				argVal, ok := argVals[arg.Name]
				if !ok {
					if arg.Default != nil {
						argVal = *arg.Default
					} else {
						return nil, fmt.Errorf("missing required argument: %q", arg.Name)
					}
				}
				field := argsType.Field(i)
				arg := reflect.New(field.Type).Interface()

				if um, ok := arg.(Unmarshaler); ok {
					if err := um.UnmarshalLiteral(argVal.Literal); err != nil {
						return nil, err
					}
				} else {
					return nil, fmt.Errorf("cannot unmarshal %T", arg)
				}

				argsVal.Elem().Field(i).Set(reflect.ValueOf(arg).Elem())
			}

			return fn(ctx, self, args)
		},
	}
}

type Marshaler interface {
	MarshalLiteral() (*idproto.Literal, error)
}

type Unmarshaler interface {
	UnmarshalLiteral(*idproto.Literal) error
}

type ScalarResolver[T Marshaler] struct{}

func (ScalarResolver[T]) isType() {}

// Class creates Nodes (a.k.a. "objects").
//
// (The metaphor is a bit of a stretch, but it's accurate enough and odd enough
// to distignuish itself.)
// type Class interface {
// 	Instantiate(*idproto.ID, any) (Node, error)
// 	Call(context.Context, Node, string, map[string]Literal) (any, error)
// }

type Class[T Typed] struct {
	Fields ObjectFields[T]
}

func (r Class[T]) isType() {}

type ClassType interface {
	Instantiate(*idproto.ID, any) (Node, error)
}

var _ ClassType = Class[Typed]{}

func (cls Class[T]) Instantiate(id *idproto.ID, val any) (Node, error) {
	return ObjectNode[T]{
		Constructor: id,
		Self:        val.(T), // TODO error
		Class:       cls,
	}, nil
}

func (cls Class[T]) Call(ctx context.Context, node ObjectNode[T], fieldName string, args map[string]Literal) (any, error) {
	field, ok := cls.Fields[fieldName]
	if !ok {
		return nil, fmt.Errorf("no such field: %q", fieldName)
	}
	if field.NodeFunc != nil {
		return field.NodeFunc(ctx, node, args)
	}
	return field.Func(ctx, node.Self, args)
}

type ObjectNode[T Typed] struct {
	Constructor *idproto.ID
	Self        T
	Class       Class[T]
}

var _ Node = ObjectNode[Typed]{}

func (o ObjectNode[T]) TypeName() string {
	return o.Self.TypeName()
}

type ObjectFields[T Typed] map[string]Field[T]

type Field[T any] struct {
	Spec     FieldSpec
	Func     func(ctx context.Context, self T, args map[string]Literal) (any, error)
	NodeFunc func(ctx context.Context, self Node, args map[string]Literal) (any, error)
}

var _ Node = ObjectNode[Typed]{}

func (r ObjectNode[T]) ID() *idproto.ID {
	return r.Constructor
}

var _ Resolver = ObjectNode[Typed]{}

func (r ObjectNode[T]) Resolve(ctx context.Context, field *ast.FieldDefinition, givenArgs map[string]Literal) (any, error) {
	args := make(map[string]Literal, len(field.Arguments))
	for _, arg := range field.Arguments {
		val, ok := givenArgs[arg.Name]
		if ok {
			args[arg.Name] = val
		} else {
			if arg.DefaultValue != nil {
				val, err := arg.DefaultValue.Value(nil)
				if err != nil {
					return nil, err
				}
				args[arg.Name] = Literal{idproto.LiteralValue(val)}
			} else if arg.Type.NonNull {
				return nil, fmt.Errorf("missing required argument: %q", arg.Name)
			}
		}
	}
	return r.Class.Call(ctx, r, field.Name, args)
}

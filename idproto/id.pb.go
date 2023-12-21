// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: id.proto

package idproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ID represents a GraphQL value of a certain type, constructed by evaluating
// its contained pipeline. In other words, it represents a
// constructor-addressed value, which may be an object, an array, or a scalar
// value.
//
// It may be binary=>base64-encoded to be used as a GraphQL ID value for
// objects. Alternatively it may be stored in a database and referred to via an
// RFC-6920 ni://sha-256;... URI.
type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent ID, if any.
	Parent *ID `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The GraphQL type of the value.
	Type *Type `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// GraphQL field name.
	Field string `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
	// GraphQL field arguments, always in alphabetical order.
	Args []*Argument `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	// If true, this Selector is not reproducible.
	//
	// TODO: do we need to refer to session/client IDs or anything here? Or is
	// that all internal? Forcing function is whether this is used as an
	// in-memory query cache key. But the query cache might be made per-session
	// or even per-client instead anyway! What buys us the most?
	Tainted bool `protobuf:"varint,5,opt,name=tainted,proto3" json:"tainted,omitempty"`
	// If true, this Selector may be omitted from the pipeline without changing
	// the ultimate result.
	//
	// This is used to prevent meta-queries like 'pipeline' and 'withFocus' from
	// busting cache keys when desired.
	//
	// It is worth noting that we don't store meta information at this level and
	// continue to force metadata to be set via GraphQL queries. It makes IDs
	// always easy to evaluate.
	Meta bool `protobuf:"varint,6,opt,name=meta,proto3" json:"meta,omitempty"`
	// If the field returns a list, this is the index of the element to select.
	// Note that this defaults to zero, as IDs always refer to
	//
	// Here we're teetering dangerously close to full blown attribute path
	// selection, but we're intentionally limiting ourselves instead to cover
	// only the common case of returning a list of objects. The only case not
	// handled is a nested list. Don't do that; have a type instead.
	Nth int64 `protobuf:"varint,7,opt,name=nth,proto3" json:"nth,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_id_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_id_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_id_proto_rawDescGZIP(), []int{0}
}

func (x *ID) GetParent() *ID {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *ID) GetType() *Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *ID) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *ID) GetArgs() []*Argument {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *ID) GetTainted() bool {
	if x != nil {
		return x.Tainted
	}
	return false
}

func (x *ID) GetMeta() bool {
	if x != nil {
		return x.Meta
	}
	return false
}

func (x *ID) GetNth() int64 {
	if x != nil {
		return x.Nth
	}
	return 0
}

// A named value passed to a GraphQL field or contained in an input object.
type Argument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value *Literal `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Argument) Reset() {
	*x = Argument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_id_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Argument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Argument) ProtoMessage() {}

func (x *Argument) ProtoReflect() protoreflect.Message {
	mi := &file_id_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Argument.ProtoReflect.Descriptor instead.
func (*Argument) Descriptor() ([]byte, []int) {
	return file_id_proto_rawDescGZIP(), []int{1}
}

func (x *Argument) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Argument) GetValue() *Literal {
	if x != nil {
		return x.Value
	}
	return nil
}

// A value passed to an argument or contained in a list.
type Literal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Literal_Id
	//	*Literal_Null
	//	*Literal_Bool
	//	*Literal_Enum
	//	*Literal_Int
	//	*Literal_Float
	//	*Literal_String_
	//	*Literal_List
	//	*Literal_Object
	Value isLiteral_Value `protobuf_oneof:"value"`
}

func (x *Literal) Reset() {
	*x = Literal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_id_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Literal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Literal) ProtoMessage() {}

func (x *Literal) ProtoReflect() protoreflect.Message {
	mi := &file_id_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Literal.ProtoReflect.Descriptor instead.
func (*Literal) Descriptor() ([]byte, []int) {
	return file_id_proto_rawDescGZIP(), []int{2}
}

func (m *Literal) GetValue() isLiteral_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Literal) GetId() *ID {
	if x, ok := x.GetValue().(*Literal_Id); ok {
		return x.Id
	}
	return nil
}

func (x *Literal) GetNull() bool {
	if x, ok := x.GetValue().(*Literal_Null); ok {
		return x.Null
	}
	return false
}

func (x *Literal) GetBool() bool {
	if x, ok := x.GetValue().(*Literal_Bool); ok {
		return x.Bool
	}
	return false
}

func (x *Literal) GetEnum() string {
	if x, ok := x.GetValue().(*Literal_Enum); ok {
		return x.Enum
	}
	return ""
}

func (x *Literal) GetInt() int64 {
	if x, ok := x.GetValue().(*Literal_Int); ok {
		return x.Int
	}
	return 0
}

func (x *Literal) GetFloat() float64 {
	if x, ok := x.GetValue().(*Literal_Float); ok {
		return x.Float
	}
	return 0
}

func (x *Literal) GetString_() string {
	if x, ok := x.GetValue().(*Literal_String_); ok {
		return x.String_
	}
	return ""
}

func (x *Literal) GetList() *List {
	if x, ok := x.GetValue().(*Literal_List); ok {
		return x.List
	}
	return nil
}

func (x *Literal) GetObject() *Object {
	if x, ok := x.GetValue().(*Literal_Object); ok {
		return x.Object
	}
	return nil
}

type isLiteral_Value interface {
	isLiteral_Value()
}

type Literal_Id struct {
	Id *ID `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type Literal_Null struct {
	Null bool `protobuf:"varint,2,opt,name=null,proto3,oneof"`
}

type Literal_Bool struct {
	Bool bool `protobuf:"varint,3,opt,name=bool,proto3,oneof"`
}

type Literal_Enum struct {
	Enum string `protobuf:"bytes,4,opt,name=enum,proto3,oneof"`
}

type Literal_Int struct {
	Int int64 `protobuf:"varint,5,opt,name=int,proto3,oneof"`
}

type Literal_Float struct {
	Float float64 `protobuf:"fixed64,6,opt,name=float,proto3,oneof"`
}

type Literal_String_ struct {
	String_ string `protobuf:"bytes,7,opt,name=string,proto3,oneof"`
}

type Literal_List struct {
	List *List `protobuf:"bytes,8,opt,name=list,proto3,oneof"`
}

type Literal_Object struct {
	Object *Object `protobuf:"bytes,9,opt,name=object,proto3,oneof"`
}

func (*Literal_Id) isLiteral_Value() {}

func (*Literal_Null) isLiteral_Value() {}

func (*Literal_Bool) isLiteral_Value() {}

func (*Literal_Enum) isLiteral_Value() {}

func (*Literal_Int) isLiteral_Value() {}

func (*Literal_Float) isLiteral_Value() {}

func (*Literal_String_) isLiteral_Value() {}

func (*Literal_List) isLiteral_Value() {}

func (*Literal_Object) isLiteral_Value() {}

// A list of values.
type List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Literal `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *List) Reset() {
	*x = List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_id_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*List) ProtoMessage() {}

func (x *List) ProtoReflect() protoreflect.Message {
	mi := &file_id_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use List.ProtoReflect.Descriptor instead.
func (*List) Descriptor() ([]byte, []int) {
	return file_id_proto_rawDescGZIP(), []int{3}
}

func (x *List) GetValues() []*Literal {
	if x != nil {
		return x.Values
	}
	return nil
}

// A series of named values.
type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Argument `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_id_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_id_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_id_proto_rawDescGZIP(), []int{4}
}

func (x *Object) GetValues() []*Argument {
	if x != nil {
		return x.Values
	}
	return nil
}

// A GraphQL type.
type Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NamedType string `protobuf:"bytes,1,opt,name=namedType,proto3" json:"namedType,omitempty"`
	Elem      *Type  `protobuf:"bytes,2,opt,name=elem,proto3" json:"elem,omitempty"`
	NonNull   bool   `protobuf:"varint,3,opt,name=nonNull,proto3" json:"nonNull,omitempty"`
}

func (x *Type) Reset() {
	*x = Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_id_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	mi := &file_id_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Type.ProtoReflect.Descriptor instead.
func (*Type) Descriptor() ([]byte, []int) {
	return file_id_proto_rawDescGZIP(), []int{5}
}

func (x *Type) GetNamedType() string {
	if x != nil {
		return x.NamedType
	}
	return ""
}

func (x *Type) GetElem() *Type {
	if x != nil {
		return x.Elem
	}
	return nil
}

func (x *Type) GetNonNull() bool {
	if x != nil {
		return x.NonNull
	}
	return false
}

var File_id_proto protoreflect.FileDescriptor

var file_id_proto_rawDesc = []byte{
	0x0a, 0x08, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x61, 0x67, 0x67,
	0x65, 0x72, 0x22, 0xc6, 0x01, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x64, 0x61, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x49, 0x44, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x72, 0x67,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x61, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74, 0x61,
	0x69, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x74, 0x68,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x74, 0x68, 0x22, 0x45, 0x0a, 0x08, 0x41,
	0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x86, 0x02, 0x0a, 0x07, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x1c,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x64, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x49, 0x44, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04,
	0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x75,
	0x6c, 0x6c, 0x12, 0x14, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x12,
	0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x03, 0x69,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x00, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2f, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x74,
	0x65, 0x72, 0x61, 0x6c, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x32, 0x0a, 0x06,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x22, 0x60, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x65, 0x6c, 0x65, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x65, 0x6c, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x6e, 0x4e,
	0x75, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6e, 0x6f, 0x6e, 0x4e, 0x75,
	0x6c, 0x6c, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_id_proto_rawDescOnce sync.Once
	file_id_proto_rawDescData = file_id_proto_rawDesc
)

func file_id_proto_rawDescGZIP() []byte {
	file_id_proto_rawDescOnce.Do(func() {
		file_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_id_proto_rawDescData)
	})
	return file_id_proto_rawDescData
}

var file_id_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_id_proto_goTypes = []interface{}{
	(*ID)(nil),       // 0: dagger.ID
	(*Argument)(nil), // 1: dagger.Argument
	(*Literal)(nil),  // 2: dagger.Literal
	(*List)(nil),     // 3: dagger.List
	(*Object)(nil),   // 4: dagger.Object
	(*Type)(nil),     // 5: dagger.Type
}
var file_id_proto_depIdxs = []int32{
	0,  // 0: dagger.ID.parent:type_name -> dagger.ID
	5,  // 1: dagger.ID.type:type_name -> dagger.Type
	1,  // 2: dagger.ID.args:type_name -> dagger.Argument
	2,  // 3: dagger.Argument.value:type_name -> dagger.Literal
	0,  // 4: dagger.Literal.id:type_name -> dagger.ID
	3,  // 5: dagger.Literal.list:type_name -> dagger.List
	4,  // 6: dagger.Literal.object:type_name -> dagger.Object
	2,  // 7: dagger.List.values:type_name -> dagger.Literal
	1,  // 8: dagger.Object.values:type_name -> dagger.Argument
	5,  // 9: dagger.Type.elem:type_name -> dagger.Type
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_id_proto_init() }
func file_id_proto_init() {
	if File_id_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_id_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_id_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Argument); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_id_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Literal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_id_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*List); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_id_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_id_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Type); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_id_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Literal_Id)(nil),
		(*Literal_Null)(nil),
		(*Literal_Bool)(nil),
		(*Literal_Enum)(nil),
		(*Literal_Int)(nil),
		(*Literal_Float)(nil),
		(*Literal_String_)(nil),
		(*Literal_List)(nil),
		(*Literal_Object)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_id_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_id_proto_goTypes,
		DependencyIndexes: file_id_proto_depIdxs,
		MessageInfos:      file_id_proto_msgTypes,
	}.Build()
	File_id_proto = out.File
	file_id_proto_rawDesc = nil
	file_id_proto_goTypes = nil
	file_id_proto_depIdxs = nil
}

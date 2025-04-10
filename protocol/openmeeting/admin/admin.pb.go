// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: admin.proto

package admin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserLoginReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Account       string                 `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	Password      string                 `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserLoginReq) Reset() {
	*x = UserLoginReq{}
	mi := &file_admin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginReq) ProtoMessage() {}

func (x *UserLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginReq.ProtoReflect.Descriptor instead.
func (*UserLoginReq) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{0}
}

func (x *UserLoginReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLoginResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	UserID        string                 `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserLoginResp) Reset() {
	*x = UserLoginResp{}
	mi := &file_admin_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserLoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginResp) ProtoMessage() {}

func (x *UserLoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginResp.ProtoReflect.Descriptor instead.
func (*UserLoginResp) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{1}
}

func (x *UserLoginResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserLoginResp) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type UserRegisterReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Nickname      string                 `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Account       string                 `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRegisterReq) Reset() {
	*x = UserRegisterReq{}
	mi := &file_admin_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterReq) ProtoMessage() {}

func (x *UserRegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterReq.ProtoReflect.Descriptor instead.
func (*UserRegisterReq) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{2}
}

func (x *UserRegisterReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserRegisterReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserRegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserRegisterResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRegisterResp) Reset() {
	*x = UserRegisterResp{}
	mi := &file_admin_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterResp) ProtoMessage() {}

func (x *UserRegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterResp.ProtoReflect.Descriptor instead.
func (*UserRegisterResp) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{3}
}

var File_admin_proto protoreflect.FileDescriptor

const file_admin_proto_rawDesc = "" +
	"\n" +
	"\vadmin.proto\x12\x11openmeeting.admin\"D\n" +
	"\fuserLoginReq\x12\x18\n" +
	"\aaccount\x18\x04 \x01(\tR\aaccount\x12\x1a\n" +
	"\bpassword\x18\x05 \x01(\tR\bpassword\"=\n" +
	"\ruserLoginResp\x12\x14\n" +
	"\x05token\x18\x02 \x01(\tR\x05token\x12\x16\n" +
	"\x06userID\x18\x03 \x01(\tR\x06userID\"c\n" +
	"\x0fuserRegisterReq\x12\x1a\n" +
	"\bnickname\x18\x01 \x01(\tR\bnickname\x12\x18\n" +
	"\aaccount\x18\x02 \x01(\tR\aaccount\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\"\x12\n" +
	"\x10userRegisterRespB+Z)github.com/eos/protocol/openmeeting/adminb\x06proto3"

var (
	file_admin_proto_rawDescOnce sync.Once
	file_admin_proto_rawDescData []byte
)

func file_admin_proto_rawDescGZIP() []byte {
	file_admin_proto_rawDescOnce.Do(func() {
		file_admin_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_admin_proto_rawDesc), len(file_admin_proto_rawDesc)))
	})
	return file_admin_proto_rawDescData
}

var file_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_admin_proto_goTypes = []any{
	(*UserLoginReq)(nil),     // 0: openmeeting.admin.userLoginReq
	(*UserLoginResp)(nil),    // 1: openmeeting.admin.userLoginResp
	(*UserRegisterReq)(nil),  // 2: openmeeting.admin.userRegisterReq
	(*UserRegisterResp)(nil), // 3: openmeeting.admin.userRegisterResp
}
var file_admin_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_proto_init() }
func file_admin_proto_init() {
	if File_admin_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_admin_proto_rawDesc), len(file_admin_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_proto_goTypes,
		DependencyIndexes: file_admin_proto_depIdxs,
		MessageInfos:      file_admin_proto_msgTypes,
	}.Build()
	File_admin_proto = out.File
	file_admin_proto_goTypes = nil
	file_admin_proto_depIdxs = nil
}

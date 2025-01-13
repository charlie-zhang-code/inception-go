// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.19.4
// source: proto/rpc/member.proto

package pb

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

type CodeMessageResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int64                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CodeMessageResult) Reset() {
	*x = CodeMessageResult{}
	mi := &file_proto_rpc_member_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CodeMessageResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeMessageResult) ProtoMessage() {}

func (x *CodeMessageResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_member_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeMessageResult.ProtoReflect.Descriptor instead.
func (*CodeMessageResult) Descriptor() ([]byte, []int) {
	return file_proto_rpc_member_proto_rawDescGZIP(), []int{0}
}

func (x *CodeMessageResult) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CodeMessageResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MemberData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                             //用户唯一标识符，主键
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`                  //用户登录时使用的用户名
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`                  //用户账户的密码，应存储加密后的值
	Nickname      string                 `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`                  //用户在系统中显示的名字，可选
	Avatar        string                 `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`                      //用户头像的URL或路径
	Quote         string                 `protobuf:"bytes,6,opt,name=quote,proto3" json:"quote,omitempty"`                        //用户设置的个人座右铭或签名
	Birthday      string                 `protobuf:"bytes,7,opt,name=birthday,proto3" json:"birthday,omitempty"`                  //用户的出生日期，用于个性化服务或统计
	Gender        int64                  `protobuf:"varint,8,opt,name=gender,proto3" json:"gender,omitempty"`                     //用户的性别，0表示保密，1表示男性，2表示女性
	Email         string                 `protobuf:"bytes,9,opt,name=email,proto3" json:"email,omitempty"`                        //用户的电子邮件地址，可用于找回密码或接收通知
	Telephone     string                 `protobuf:"bytes,10,opt,name=telephone,proto3" json:"telephone,omitempty"`               //用户的联系电话，可用于身份验证或联系用户
	Status        int64                  `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`                    //状态，0表示禁用，1表示正常启用
	Deleted       int64                  `protobuf:"varint,12,opt,name=deleted,proto3" json:"deleted,omitempty"`                  //逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复
	Remark        string                 `protobuf:"bytes,13,opt,name=remark,proto3" json:"remark,omitempty"`                     //对记录的备注信息，如特殊说明等
	CreateAt      string                 `protobuf:"bytes,14,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"` //记录创建的时间戳
	CreateBy      string                 `protobuf:"bytes,15,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"` //创建该记录的用户标识符
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MemberData) Reset() {
	*x = MemberData{}
	mi := &file_proto_rpc_member_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemberData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberData) ProtoMessage() {}

func (x *MemberData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_member_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberData.ProtoReflect.Descriptor instead.
func (*MemberData) Descriptor() ([]byte, []int) {
	return file_proto_rpc_member_proto_rawDescGZIP(), []int{1}
}

func (x *MemberData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MemberData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *MemberData) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *MemberData) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *MemberData) GetQuote() string {
	if x != nil {
		return x.Quote
	}
	return ""
}

func (x *MemberData) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *MemberData) GetGender() int64 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *MemberData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *MemberData) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *MemberData) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MemberData) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

func (x *MemberData) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *MemberData) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

func (x *MemberData) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

// 删除用户
type DeleteMemberIds struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ids           []int64                `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteMemberIds) Reset() {
	*x = DeleteMemberIds{}
	mi := &file_proto_rpc_member_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMemberIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemberIds) ProtoMessage() {}

func (x *DeleteMemberIds) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_member_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemberIds.ProtoReflect.Descriptor instead.
func (*DeleteMemberIds) Descriptor() ([]byte, []int) {
	return file_proto_rpc_member_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteMemberIds) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

// 查询用户详情（通过id）
type QueryMemberDetailById struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryMemberDetailById) Reset() {
	*x = QueryMemberDetailById{}
	mi := &file_proto_rpc_member_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryMemberDetailById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMemberDetailById) ProtoMessage() {}

func (x *QueryMemberDetailById) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_member_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMemberDetailById.ProtoReflect.Descriptor instead.
func (*QueryMemberDetailById) Descriptor() ([]byte, []int) {
	return file_proto_rpc_member_proto_rawDescGZIP(), []int{3}
}

func (x *QueryMemberDetailById) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 查询用户详情（通过用户名）
type QueryMemberDetailByUsername struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryMemberDetailByUsername) Reset() {
	*x = QueryMemberDetailByUsername{}
	mi := &file_proto_rpc_member_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryMemberDetailByUsername) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMemberDetailByUsername) ProtoMessage() {}

func (x *QueryMemberDetailByUsername) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_member_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMemberDetailByUsername.ProtoReflect.Descriptor instead.
func (*QueryMemberDetailByUsername) Descriptor() ([]byte, []int) {
	return file_proto_rpc_member_proto_rawDescGZIP(), []int{4}
}

func (x *QueryMemberDetailByUsername) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// 分页查询用户列表
type QueryPageMemberList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int64                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`      //第几页
	Size          int64                  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`      //每页的数量
	Keyword       string                 `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"` //关键字
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryPageMemberList) Reset() {
	*x = QueryPageMemberList{}
	mi := &file_proto_rpc_member_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryPageMemberList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPageMemberList) ProtoMessage() {}

func (x *QueryPageMemberList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_member_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPageMemberList.ProtoReflect.Descriptor instead.
func (*QueryPageMemberList) Descriptor() ([]byte, []int) {
	return file_proto_rpc_member_proto_rawDescGZIP(), []int{5}
}

func (x *QueryPageMemberList) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *QueryPageMemberList) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *QueryPageMemberList) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

type QueryPageMemberListResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int64                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"` //总数
	Size          int64                  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`   //每页的数量
	List          []*MemberData          `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`    //列表
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryPageMemberListResult) Reset() {
	*x = QueryPageMemberListResult{}
	mi := &file_proto_rpc_member_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryPageMemberListResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPageMemberListResult) ProtoMessage() {}

func (x *QueryPageMemberListResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_member_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPageMemberListResult.ProtoReflect.Descriptor instead.
func (*QueryPageMemberListResult) Descriptor() ([]byte, []int) {
	return file_proto_rpc_member_proto_rawDescGZIP(), []int{6}
}

func (x *QueryPageMemberListResult) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *QueryPageMemberListResult) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *QueryPageMemberListResult) GetList() []*MemberData {
	if x != nil {
		return x.List
	}
	return nil
}

var File_proto_rpc_member_proto protoreflect.FileDescriptor

var file_proto_rpc_member_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x41, 0x0a, 0x11, 0x43, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x8a, 0x03, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79,
	0x22, 0x23, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x49, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39,
	0x0a, 0x1b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x57, 0x0a, 0x13, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x6d, 0x0a, 0x19, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x32, 0xd2, 0x03, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x44, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x2e,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x73, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x2e, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00,
	0x12, 0x56, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x2e,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x1a, 0x12, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rpc_member_proto_rawDescOnce sync.Once
	file_proto_rpc_member_proto_rawDescData = file_proto_rpc_member_proto_rawDesc
)

func file_proto_rpc_member_proto_rawDescGZIP() []byte {
	file_proto_rpc_member_proto_rawDescOnce.Do(func() {
		file_proto_rpc_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rpc_member_proto_rawDescData)
	})
	return file_proto_rpc_member_proto_rawDescData
}

var file_proto_rpc_member_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_rpc_member_proto_goTypes = []any{
	(*CodeMessageResult)(nil),           // 0: member.CodeMessageResult
	(*MemberData)(nil),                  // 1: member.MemberData
	(*DeleteMemberIds)(nil),             // 2: member.DeleteMemberIds
	(*QueryMemberDetailById)(nil),       // 3: member.QueryMemberDetailById
	(*QueryMemberDetailByUsername)(nil), // 4: member.QueryMemberDetailByUsername
	(*QueryPageMemberList)(nil),         // 5: member.QueryPageMemberList
	(*QueryPageMemberListResult)(nil),   // 6: member.QueryPageMemberListResult
}
var file_proto_rpc_member_proto_depIdxs = []int32{
	1, // 0: member.QueryPageMemberListResult.list:type_name -> member.MemberData
	1, // 1: member.Member.CreateMember:input_type -> member.MemberData
	2, // 2: member.Member.DeleteMember:input_type -> member.DeleteMemberIds
	1, // 3: member.Member.UpdateMember:input_type -> member.MemberData
	3, // 4: member.Member.GetMemberDetailById:input_type -> member.QueryMemberDetailById
	4, // 5: member.Member.GetMemberDetailByUsername:input_type -> member.QueryMemberDetailByUsername
	5, // 6: member.Member.GetKeywordPageMemberList:input_type -> member.QueryPageMemberList
	0, // 7: member.Member.CreateMember:output_type -> member.CodeMessageResult
	0, // 8: member.Member.DeleteMember:output_type -> member.CodeMessageResult
	0, // 9: member.Member.UpdateMember:output_type -> member.CodeMessageResult
	1, // 10: member.Member.GetMemberDetailById:output_type -> member.MemberData
	1, // 11: member.Member.GetMemberDetailByUsername:output_type -> member.MemberData
	6, // 12: member.Member.GetKeywordPageMemberList:output_type -> member.QueryPageMemberListResult
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_rpc_member_proto_init() }
func file_proto_rpc_member_proto_init() {
	if File_proto_rpc_member_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_rpc_member_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rpc_member_proto_goTypes,
		DependencyIndexes: file_proto_rpc_member_proto_depIdxs,
		MessageInfos:      file_proto_rpc_member_proto_msgTypes,
	}.Build()
	File_proto_rpc_member_proto = out.File
	file_proto_rpc_member_proto_rawDesc = nil
	file_proto_rpc_member_proto_goTypes = nil
	file_proto_rpc_member_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: api/test/v1/test.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Sex int32

const (
	//未知
	Sex_UN Sex = 0
	//男
	Sex_MAN Sex = 1
	//女
	Sex_WOMAN Sex = 2
)

// Enum value maps for Sex.
var (
	Sex_name = map[int32]string{
		0: "UN",
		1: "MAN",
		2: "WOMAN",
	}
	Sex_value = map[string]int32{
		"UN":    0,
		"MAN":   1,
		"WOMAN": 2,
	}
)

func (x Sex) Enum() *Sex {
	p := new(Sex)
	*p = x
	return p
}

func (x Sex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sex) Descriptor() protoreflect.EnumDescriptor {
	return file_api_test_v1_test_proto_enumTypes[0].Descriptor()
}

func (Sex) Type() protoreflect.EnumType {
	return &file_api_test_v1_test_proto_enumTypes[0]
}

func (x Sex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sex.Descriptor instead.
func (Sex) EnumDescriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{0}
}

//创建请求
type CreateTestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Test `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateTestReply) Reset() {
	*x = CreateTestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestReply) ProtoMessage() {}

func (x *CreateTestReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestReply.ProtoReflect.Descriptor instead.
func (*CreateTestReply) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTestReply) GetData() *Test {
	if x != nil {
		return x.Data
	}
	return nil
}

//更新请求
type UpdateTestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Test `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateTestReply) Reset() {
	*x = UpdateTestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTestReply) ProtoMessage() {}

func (x *UpdateTestReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTestReply.ProtoReflect.Descriptor instead.
func (*UpdateTestReply) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateTestReply) GetData() *Test {
	if x != nil {
		return x.Data
	}
	return nil
}

//批量删除请求
type DeleteTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteTestRequest) Reset() {
	*x = DeleteTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTestRequest) ProtoMessage() {}

func (x *DeleteTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTestRequest.ProtoReflect.Descriptor instead.
func (*DeleteTestRequest) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteTestRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

//删除结果
type DeleteTestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTestReply) Reset() {
	*x = DeleteTestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTestReply) ProtoMessage() {}

func (x *DeleteTestReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTestReply.ProtoReflect.Descriptor instead.
func (*DeleteTestReply) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{3}
}

//查询单个数据
type GetTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTestRequest) Reset() {
	*x = GetTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestRequest) ProtoMessage() {}

func (x *GetTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestRequest.ProtoReflect.Descriptor instead.
func (*GetTestRequest) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{4}
}

func (x *GetTestRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

//列表查询条件
type ListTestOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTestOption) Reset() {
	*x = ListTestOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTestOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTestOption) ProtoMessage() {}

func (x *ListTestOption) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTestOption.ProtoReflect.Descriptor instead.
func (*ListTestOption) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{5}
}

//列表查询返回
type ListTestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Test `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	//数据总计
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	//页码
	Page uint32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	//分页大小
	PageSize uint32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListTestReply) Reset() {
	*x = ListTestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTestReply) ProtoMessage() {}

func (x *ListTestReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTestReply.ProtoReflect.Descriptor instead.
func (*ListTestReply) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{6}
}

func (x *ListTestReply) GetList() []*Test {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ListTestReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListTestReply) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListTestReply) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// Test结构体(请在此处定义数据结构)
type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 简介
	Desc string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	// 性别
	Sex Sex `protobuf:"varint,4,opt,name=sex,proto3,enum=api.test.v1.Sex" json:"sex,omitempty"`
	// 测试字段
	Ui uint32 `protobuf:"varint,5,opt,name=ui,proto3" json:"ui,omitempty"`
	// 测试字段
	I3 int32 `protobuf:"varint,6,opt,name=i3,proto3" json:"i3,omitempty"`
	// 测试字段
	Si3 int32 `protobuf:"zigzag32,7,opt,name=si3,proto3" json:"si3,omitempty"`
	// 测试字段
	Sf int32 `protobuf:"fixed32,8,opt,name=sf,proto3" json:"sf,omitempty"`
	// 测试字段
	F float32 `protobuf:"fixed32,9,opt,name=f,proto3" json:"f,omitempty"`
	// 测试时间字段
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,22,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// 测试时间字段
	CreateTime *timestamp.Timestamp `protobuf:"bytes,23,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_v1_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_v1_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_api_test_v1_test_proto_rawDescGZIP(), []int{7}
}

func (x *Test) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Test) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Test) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Test) GetSex() Sex {
	if x != nil {
		return x.Sex
	}
	return Sex_UN
}

func (x *Test) GetUi() uint32 {
	if x != nil {
		return x.Ui
	}
	return 0
}

func (x *Test) GetI3() int32 {
	if x != nil {
		return x.I3
	}
	return 0
}

func (x *Test) GetSi3() int32 {
	if x != nil {
		return x.Si3
	}
	return 0
}

func (x *Test) GetSf() int32 {
	if x != nil {
		return x.Sf
	}
	return 0
}

func (x *Test) GetF() float32 {
	if x != nil {
		return x.F
	}
	return 0
}

func (x *Test) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Test) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

var File_api_test_v1_test_proto protoreflect.FileDescriptor

var file_api_test_v1_test_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x38,
	0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x25, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22,
	0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7d, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xac, 0x02, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x22, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x78, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x75, 0x69,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x75, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x33,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x33, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69,
	0x33, 0x18, 0x07, 0x20, 0x01, 0x28, 0x11, 0x52, 0x03, 0x73, 0x69, 0x33, 0x12, 0x0e, 0x0a, 0x02,
	0x73, 0x66, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x02, 0x73, 0x66, 0x12, 0x0c, 0x0a, 0x01,
	0x66, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x66, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x2a, 0x21, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x12, 0x06, 0x0a, 0x02, 0x55,
	0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x57, 0x4f, 0x4d, 0x41, 0x4e, 0x10, 0x02, 0x32, 0xd7, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x39, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x08,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x42, 0x55, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x65, 0x61, 0x6c, 0x6f, 0x74, 0x7a, 0x2f, 0x77, 0x68, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_test_v1_test_proto_rawDescOnce sync.Once
	file_api_test_v1_test_proto_rawDescData = file_api_test_v1_test_proto_rawDesc
)

func file_api_test_v1_test_proto_rawDescGZIP() []byte {
	file_api_test_v1_test_proto_rawDescOnce.Do(func() {
		file_api_test_v1_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_test_v1_test_proto_rawDescData)
	})
	return file_api_test_v1_test_proto_rawDescData
}

var file_api_test_v1_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_test_v1_test_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_test_v1_test_proto_goTypes = []interface{}{
	(Sex)(0),                    // 0: api.test.v1.Sex
	(*CreateTestReply)(nil),     // 1: api.test.v1.CreateTestReply
	(*UpdateTestReply)(nil),     // 2: api.test.v1.UpdateTestReply
	(*DeleteTestRequest)(nil),   // 3: api.test.v1.DeleteTestRequest
	(*DeleteTestReply)(nil),     // 4: api.test.v1.DeleteTestReply
	(*GetTestRequest)(nil),      // 5: api.test.v1.GetTestRequest
	(*ListTestOption)(nil),      // 6: api.test.v1.ListTestOption
	(*ListTestReply)(nil),       // 7: api.test.v1.ListTestReply
	(*Test)(nil),                // 8: api.test.v1.Test
	(*timestamp.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_api_test_v1_test_proto_depIdxs = []int32{
	8,  // 0: api.test.v1.CreateTestReply.data:type_name -> api.test.v1.Test
	8,  // 1: api.test.v1.UpdateTestReply.data:type_name -> api.test.v1.Test
	8,  // 2: api.test.v1.ListTestReply.list:type_name -> api.test.v1.Test
	0,  // 3: api.test.v1.Test.sex:type_name -> api.test.v1.Sex
	9,  // 4: api.test.v1.Test.update_time:type_name -> google.protobuf.Timestamp
	9,  // 5: api.test.v1.Test.create_time:type_name -> google.protobuf.Timestamp
	8,  // 6: api.test.v1.TestService.CreateTest:input_type -> api.test.v1.Test
	8,  // 7: api.test.v1.TestService.UpdateTest:input_type -> api.test.v1.Test
	3,  // 8: api.test.v1.TestService.DeleteTest:input_type -> api.test.v1.DeleteTestRequest
	5,  // 9: api.test.v1.TestService.GetTest:input_type -> api.test.v1.GetTestRequest
	6,  // 10: api.test.v1.TestService.ListTest:input_type -> api.test.v1.ListTestOption
	1,  // 11: api.test.v1.TestService.CreateTest:output_type -> api.test.v1.CreateTestReply
	2,  // 12: api.test.v1.TestService.UpdateTest:output_type -> api.test.v1.UpdateTestReply
	4,  // 13: api.test.v1.TestService.DeleteTest:output_type -> api.test.v1.DeleteTestReply
	8,  // 14: api.test.v1.TestService.GetTest:output_type -> api.test.v1.Test
	7,  // 15: api.test.v1.TestService.ListTest:output_type -> api.test.v1.ListTestReply
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_test_v1_test_proto_init() }
func file_api_test_v1_test_proto_init() {
	if File_api_test_v1_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_test_v1_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTestReply); i {
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
		file_api_test_v1_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTestReply); i {
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
		file_api_test_v1_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTestRequest); i {
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
		file_api_test_v1_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTestReply); i {
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
		file_api_test_v1_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestRequest); i {
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
		file_api_test_v1_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTestOption); i {
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
		file_api_test_v1_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTestReply); i {
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
		file_api_test_v1_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_test_v1_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_test_v1_test_proto_goTypes,
		DependencyIndexes: file_api_test_v1_test_proto_depIdxs,
		EnumInfos:         file_api_test_v1_test_proto_enumTypes,
		MessageInfos:      file_api_test_v1_test_proto_msgTypes,
	}.Build()
	File_api_test_v1_test_proto = out.File
	file_api_test_v1_test_proto_rawDesc = nil
	file_api_test_v1_test_proto_goTypes = nil
	file_api_test_v1_test_proto_depIdxs = nil
}

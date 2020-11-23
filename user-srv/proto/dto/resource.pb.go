// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: course/user-srv/proto/dto/resource.proto

package dto

import (
	proto "github.com/golang/protobuf/proto"
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

type ResourceDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	// id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	//
	// 名称|菜单或按钮
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	//
	// 页面|路由
	Page string `protobuf:"bytes,3,opt,name=page,proto3" json:"page,omitempty"`
	//
	// 请求|接口
	Request string `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
	//
	// 父id
	Parent   string         `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent,omitempty"`
	Children []*ResourceDto `protobuf:"bytes,6,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *ResourceDto) Reset() {
	*x = ResourceDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_user_srv_proto_dto_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceDto) ProtoMessage() {}

func (x *ResourceDto) ProtoReflect() protoreflect.Message {
	mi := &file_course_user_srv_proto_dto_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceDto.ProtoReflect.Descriptor instead.
func (*ResourceDto) Descriptor() ([]byte, []int) {
	return file_course_user_srv_proto_dto_resource_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceDto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResourceDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceDto) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *ResourceDto) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *ResourceDto) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ResourceDto) GetChildren() []*ResourceDto {
	if x != nil {
		return x.Children
	}
	return nil
}

type ResourceDtoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows []*ResourceDto `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *ResourceDtoList) Reset() {
	*x = ResourceDtoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_user_srv_proto_dto_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceDtoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceDtoList) ProtoMessage() {}

func (x *ResourceDtoList) ProtoReflect() protoreflect.Message {
	mi := &file_course_user_srv_proto_dto_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceDtoList.ProtoReflect.Descriptor instead.
func (*ResourceDtoList) Descriptor() ([]byte, []int) {
	return file_course_user_srv_proto_dto_resource_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceDtoList) GetRows() []*ResourceDto {
	if x != nil {
		return x.Rows
	}
	return nil
}

type RoleDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Desc        string   `protobuf:"bytes,3,opt,name=Desc,proto3" json:"Desc,omitempty"`
	ResourceIds []string `protobuf:"bytes,4,rep,name=resourceIds,proto3" json:"resourceIds,omitempty"`
	UserIds     []string `protobuf:"bytes,5,rep,name=userIds,proto3" json:"userIds,omitempty"`
}

func (x *RoleDto) Reset() {
	*x = RoleDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_user_srv_proto_dto_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleDto) ProtoMessage() {}

func (x *RoleDto) ProtoReflect() protoreflect.Message {
	mi := &file_course_user_srv_proto_dto_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleDto.ProtoReflect.Descriptor instead.
func (*RoleDto) Descriptor() ([]byte, []int) {
	return file_course_user_srv_proto_dto_resource_proto_rawDescGZIP(), []int{2}
}

func (x *RoleDto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleDto) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *RoleDto) GetResourceIds() []string {
	if x != nil {
		return x.ResourceIds
	}
	return nil
}

func (x *RoleDto) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type RoleUserDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	RoleId string `protobuf:"bytes,3,opt,name=roleId,proto3" json:"roleId,omitempty"`
}

func (x *RoleUserDto) Reset() {
	*x = RoleUserDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_user_srv_proto_dto_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleUserDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleUserDto) ProtoMessage() {}

func (x *RoleUserDto) ProtoReflect() protoreflect.Message {
	mi := &file_course_user_srv_proto_dto_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleUserDto.ProtoReflect.Descriptor instead.
func (*RoleUserDto) Descriptor() ([]byte, []int) {
	return file_course_user_srv_proto_dto_resource_proto_rawDescGZIP(), []int{3}
}

func (x *RoleUserDto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleUserDto) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RoleUserDto) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

var File_course_user_srv_proto_dto_resource_proto protoreflect.FileDescriptor

var file_course_user_srv_proto_dto_resource_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x72,
	0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0xa6, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x74, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x74, 0x6f, 0x52,
	0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x38, 0x0a, 0x0f, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04,
	0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x74, 0x6f, 0x52, 0x04, 0x72,
	0x6f, 0x77, 0x73, 0x22, 0x7d, 0x0a, 0x07, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x44, 0x65, 0x73, 0x63, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x22, 0x4d, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x74,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c,
	0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49,
	0x64, 0x42, 0x1b, 0x5a, 0x19, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2d, 0x73, 0x72, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_course_user_srv_proto_dto_resource_proto_rawDescOnce sync.Once
	file_course_user_srv_proto_dto_resource_proto_rawDescData = file_course_user_srv_proto_dto_resource_proto_rawDesc
)

func file_course_user_srv_proto_dto_resource_proto_rawDescGZIP() []byte {
	file_course_user_srv_proto_dto_resource_proto_rawDescOnce.Do(func() {
		file_course_user_srv_proto_dto_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_course_user_srv_proto_dto_resource_proto_rawDescData)
	})
	return file_course_user_srv_proto_dto_resource_proto_rawDescData
}

var file_course_user_srv_proto_dto_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_course_user_srv_proto_dto_resource_proto_goTypes = []interface{}{
	(*ResourceDto)(nil),     // 0: user.ResourceDto
	(*ResourceDtoList)(nil), // 1: user.ResourceDtoList
	(*RoleDto)(nil),         // 2: user.RoleDto
	(*RoleUserDto)(nil),     // 3: user.RoleUserDto
}
var file_course_user_srv_proto_dto_resource_proto_depIdxs = []int32{
	0, // 0: user.ResourceDto.children:type_name -> user.ResourceDto
	0, // 1: user.ResourceDtoList.rows:type_name -> user.ResourceDto
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_course_user_srv_proto_dto_resource_proto_init() }
func file_course_user_srv_proto_dto_resource_proto_init() {
	if File_course_user_srv_proto_dto_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_course_user_srv_proto_dto_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceDto); i {
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
		file_course_user_srv_proto_dto_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceDtoList); i {
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
		file_course_user_srv_proto_dto_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleDto); i {
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
		file_course_user_srv_proto_dto_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleUserDto); i {
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
			RawDescriptor: file_course_user_srv_proto_dto_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_course_user_srv_proto_dto_resource_proto_goTypes,
		DependencyIndexes: file_course_user_srv_proto_dto_resource_proto_depIdxs,
		MessageInfos:      file_course_user_srv_proto_dto_resource_proto_msgTypes,
	}.Build()
	File_course_user_srv_proto_dto_resource_proto = out.File
	file_course_user_srv_proto_dto_resource_proto_rawDesc = nil
	file_course_user_srv_proto_dto_resource_proto_goTypes = nil
	file_course_user_srv_proto_dto_resource_proto_depIdxs = nil
}

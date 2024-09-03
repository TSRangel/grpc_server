// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: application/proto/game.proto

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

type CreateGameInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Platform string  `protobuf:"bytes,2,opt,name=Platform,proto3" json:"Platform,omitempty"`
	Price    float64 `protobuf:"fixed64,3,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *CreateGameInput) Reset() {
	*x = CreateGameInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_proto_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGameInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameInput) ProtoMessage() {}

func (x *CreateGameInput) ProtoReflect() protoreflect.Message {
	mi := &file_application_proto_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameInput.ProtoReflect.Descriptor instead.
func (*CreateGameInput) Descriptor() ([]byte, []int) {
	return file_application_proto_game_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGameInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGameInput) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *CreateGameInput) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type BaseGameOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Platform string  `protobuf:"bytes,3,opt,name=Platform,proto3" json:"Platform,omitempty"`
	Price    float64 `protobuf:"fixed64,4,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *BaseGameOutput) Reset() {
	*x = BaseGameOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_proto_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseGameOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseGameOutput) ProtoMessage() {}

func (x *BaseGameOutput) ProtoReflect() protoreflect.Message {
	mi := &file_application_proto_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseGameOutput.ProtoReflect.Descriptor instead.
func (*BaseGameOutput) Descriptor() ([]byte, []int) {
	return file_application_proto_game_proto_rawDescGZIP(), []int{1}
}

func (x *BaseGameOutput) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *BaseGameOutput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BaseGameOutput) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *BaseGameOutput) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type BlankInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BlankInput) Reset() {
	*x = BlankInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_proto_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlankInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlankInput) ProtoMessage() {}

func (x *BlankInput) ProtoReflect() protoreflect.Message {
	mi := &file_application_proto_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlankInput.ProtoReflect.Descriptor instead.
func (*BlankInput) Descriptor() ([]byte, []int) {
	return file_application_proto_game_proto_rawDescGZIP(), []int{2}
}

type ListAllGamesOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Games []*BaseGameOutput `protobuf:"bytes,1,rep,name=Games,proto3" json:"Games,omitempty"`
}

func (x *ListAllGamesOutput) Reset() {
	*x = ListAllGamesOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_proto_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllGamesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllGamesOutput) ProtoMessage() {}

func (x *ListAllGamesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_application_proto_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllGamesOutput.ProtoReflect.Descriptor instead.
func (*ListAllGamesOutput) Descriptor() ([]byte, []int) {
	return file_application_proto_game_proto_rawDescGZIP(), []int{3}
}

func (x *ListAllGamesOutput) GetGames() []*BaseGameOutput {
	if x != nil {
		return x.Games
	}
	return nil
}

type GetGameInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *GetGameInput) Reset() {
	*x = GetGameInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_proto_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameInput) ProtoMessage() {}

func (x *GetGameInput) ProtoReflect() protoreflect.Message {
	mi := &file_application_proto_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameInput.ProtoReflect.Descriptor instead.
func (*GetGameInput) Descriptor() ([]byte, []int) {
	return file_application_proto_game_proto_rawDescGZIP(), []int{4}
}

func (x *GetGameInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_application_proto_game_proto protoreflect.FileDescriptor

var file_application_proto_game_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x57, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x66, 0x0a, 0x0e, 0x42,
	0x61, 0x73, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x22, 0x3e, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x61, 0x6d, 0x65,
	0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x47, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x47, 0x61, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x05, 0x47, 0x61, 0x6d, 0x65,
	0x73, 0x22, 0x22, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xc8, 0x02, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x61, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x43,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61,
	0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x6c, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x00, 0x28, 0x01, 0x12, 0x4e, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c,
	0x47, 0x61, 0x6d, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00,
	0x42, 0x10, 0x5a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_application_proto_game_proto_rawDescOnce sync.Once
	file_application_proto_game_proto_rawDescData = file_application_proto_game_proto_rawDesc
)

func file_application_proto_game_proto_rawDescGZIP() []byte {
	file_application_proto_game_proto_rawDescOnce.Do(func() {
		file_application_proto_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_application_proto_game_proto_rawDescData)
	})
	return file_application_proto_game_proto_rawDescData
}

var file_application_proto_game_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_application_proto_game_proto_goTypes = []any{
	(*CreateGameInput)(nil),    // 0: pb.CreateGameInput
	(*BaseGameOutput)(nil),     // 1: pb.BaseGameOutput
	(*BlankInput)(nil),         // 2: pb.BlankInput
	(*ListAllGamesOutput)(nil), // 3: pb.ListAllGamesOutput
	(*GetGameInput)(nil),       // 4: pb.GetGameInput
}
var file_application_proto_game_proto_depIdxs = []int32{
	1, // 0: pb.ListAllGamesOutput.Games:type_name -> pb.BaseGameOutput
	0, // 1: pb.GameService.CreateGame:input_type -> pb.CreateGameInput
	0, // 2: pb.GameService.CreateGameStream:input_type -> pb.CreateGameInput
	0, // 3: pb.GameService.CreateGameStreamBidirectional:input_type -> pb.CreateGameInput
	2, // 4: pb.GameService.ListAllGames:input_type -> pb.BlankInput
	4, // 5: pb.GameService.GetGame:input_type -> pb.GetGameInput
	1, // 6: pb.GameService.CreateGame:output_type -> pb.BaseGameOutput
	3, // 7: pb.GameService.CreateGameStream:output_type -> pb.ListAllGamesOutput
	1, // 8: pb.GameService.CreateGameStreamBidirectional:output_type -> pb.BaseGameOutput
	3, // 9: pb.GameService.ListAllGames:output_type -> pb.ListAllGamesOutput
	1, // 10: pb.GameService.GetGame:output_type -> pb.BaseGameOutput
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_application_proto_game_proto_init() }
func file_application_proto_game_proto_init() {
	if File_application_proto_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_application_proto_game_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateGameInput); i {
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
		file_application_proto_game_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BaseGameOutput); i {
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
		file_application_proto_game_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BlankInput); i {
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
		file_application_proto_game_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListAllGamesOutput); i {
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
		file_application_proto_game_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetGameInput); i {
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
			RawDescriptor: file_application_proto_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_application_proto_game_proto_goTypes,
		DependencyIndexes: file_application_proto_game_proto_depIdxs,
		MessageInfos:      file_application_proto_game_proto_msgTypes,
	}.Build()
	File_application_proto_game_proto = out.File
	file_application_proto_game_proto_rawDesc = nil
	file_application_proto_game_proto_goTypes = nil
	file_application_proto_game_proto_depIdxs = nil
}

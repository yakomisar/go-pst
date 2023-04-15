// go-pst is a library for reading Personal Storage Table (.pst) files (written in Go/Golang).
//
// Copyright 2023 Marten Mooij
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate msgp -tests=false

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: extracted_entity.proto

package properties

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

type ExtractedEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains an XML document with a single AddressSet element.
	ExtractedAddresses *string `protobuf:"bytes,1,opt,name=extracted_addresses,json=extractedAddresses,proto3,oneof" json:"extracted_addresses,omitempty"`
	// Contains an XML document with a single ContactSet element.
	ExtractedContacts *string `protobuf:"bytes,2,opt,name=extracted_contacts,json=extractedContacts,proto3,oneof" json:"extracted_contacts,omitempty"`
	// Contains an XML document with a single EmailSet element.
	ExtractedEmails *string `protobuf:"bytes,3,opt,name=extracted_emails,json=extractedEmails,proto3,oneof" json:"extracted_emails,omitempty"`
	// Contains an XML document with a single MeetingSet element.
	ExtractedMeetings *string `protobuf:"bytes,4,opt,name=extracted_meetings,json=extractedMeetings,proto3,oneof" json:"extracted_meetings,omitempty"`
	// Contains an XML document with a single PhoneSet element.
	ExtractedPhones *string `protobuf:"bytes,5,opt,name=extracted_phones,json=extractedPhones,proto3,oneof" json:"extracted_phones,omitempty"`
	// Contains an XML document with a single TaskSet element.
	ExtractedTasks *string `protobuf:"bytes,6,opt,name=extracted_tasks,json=extractedTasks,proto3,oneof" json:"extracted_tasks,omitempty"`
	// Contains an XML document with a single UrlSet element.
	ExtractedUrls *string `protobuf:"bytes,7,opt,name=extracted_urls,json=extractedUrls,proto3,oneof" json:"extracted_urls,omitempty"`
}

func (x *ExtractedEntity) Reset() {
	*x = ExtractedEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extracted_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractedEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractedEntity) ProtoMessage() {}

func (x *ExtractedEntity) ProtoReflect() protoreflect.Message {
	mi := &file_extracted_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractedEntity.ProtoReflect.Descriptor instead.
func (*ExtractedEntity) Descriptor() ([]byte, []int) {
	return file_extracted_entity_proto_rawDescGZIP(), []int{0}
}

func (x *ExtractedEntity) GetExtractedAddresses() string {
	if x != nil && x.ExtractedAddresses != nil {
		return *x.ExtractedAddresses
	}
	return ""
}

func (x *ExtractedEntity) GetExtractedContacts() string {
	if x != nil && x.ExtractedContacts != nil {
		return *x.ExtractedContacts
	}
	return ""
}

func (x *ExtractedEntity) GetExtractedEmails() string {
	if x != nil && x.ExtractedEmails != nil {
		return *x.ExtractedEmails
	}
	return ""
}

func (x *ExtractedEntity) GetExtractedMeetings() string {
	if x != nil && x.ExtractedMeetings != nil {
		return *x.ExtractedMeetings
	}
	return ""
}

func (x *ExtractedEntity) GetExtractedPhones() string {
	if x != nil && x.ExtractedPhones != nil {
		return *x.ExtractedPhones
	}
	return ""
}

func (x *ExtractedEntity) GetExtractedTasks() string {
	if x != nil && x.ExtractedTasks != nil {
		return *x.ExtractedTasks
	}
	return ""
}

func (x *ExtractedEntity) GetExtractedUrls() string {
	if x != nil && x.ExtractedUrls != nil {
		return *x.ExtractedUrls
	}
	return ""
}

var File_extracted_entity_proto protoreflect.FileDescriptor

var file_extracted_entity_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x04, 0x0a, 0x0f, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x13,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x12, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x32, 0x0a, 0x12, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x11, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x64, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x73, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x12, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x11, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x4d,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x06, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x55, 0x72, 0x6c,
	0x73, 0x88, 0x01, 0x01, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x42, 0x15, 0x0a, 0x13,
	0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x73, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42,
	0x13, 0x0a, 0x11, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x73, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x64, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6f, 0x69, 0x6a, 0x74,
	0x65, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x73, 0x74, 0x3b, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_extracted_entity_proto_rawDescOnce sync.Once
	file_extracted_entity_proto_rawDescData = file_extracted_entity_proto_rawDesc
)

func file_extracted_entity_proto_rawDescGZIP() []byte {
	file_extracted_entity_proto_rawDescOnce.Do(func() {
		file_extracted_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_extracted_entity_proto_rawDescData)
	})
	return file_extracted_entity_proto_rawDescData
}

var file_extracted_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_extracted_entity_proto_goTypes = []interface{}{
	(*ExtractedEntity)(nil), // 0: ExtractedEntity
}
var file_extracted_entity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_extracted_entity_proto_init() }
func file_extracted_entity_proto_init() {
	if File_extracted_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_extracted_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractedEntity); i {
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
	file_extracted_entity_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_extracted_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_extracted_entity_proto_goTypes,
		DependencyIndexes: file_extracted_entity_proto_depIdxs,
		MessageInfos:      file_extracted_entity_proto_msgTypes,
	}.Build()
	File_extracted_entity_proto = out.File
	file_extracted_entity_proto_rawDesc = nil
	file_extracted_entity_proto_goTypes = nil
	file_extracted_entity_proto_depIdxs = nil
}

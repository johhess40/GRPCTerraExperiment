// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: greet.proto

package proto

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

type GreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName  string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Occupation string `protobuf:"bytes,3,opt,name=occupation,proto3" json:"occupation,omitempty"`
	Location   string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *GreetRequest) Reset() {
	*x = GreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequest) ProtoMessage() {}

func (x *GreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequest.ProtoReflect.Descriptor instead.
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return file_greet_proto_rawDescGZIP(), []int{0}
}

func (x *GreetRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GreetRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GreetRequest) GetOccupation() string {
	if x != nil {
		return x.Occupation
	}
	return ""
}

func (x *GreetRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type GreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultFirstName  string `protobuf:"bytes,1,opt,name=result_first_name,json=resultFirstName,proto3" json:"result_first_name,omitempty"`
	ResultLastName   string `protobuf:"bytes,2,opt,name=result_last_name,json=resultLastName,proto3" json:"result_last_name,omitempty"`
	ResultOccupation string `protobuf:"bytes,3,opt,name=result_occupation,json=resultOccupation,proto3" json:"result_occupation,omitempty"`
	ResultLocation   string `protobuf:"bytes,4,opt,name=result_location,json=resultLocation,proto3" json:"result_location,omitempty"`
}

func (x *GreetResponse) Reset() {
	*x = GreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponse) ProtoMessage() {}

func (x *GreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponse.ProtoReflect.Descriptor instead.
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return file_greet_proto_rawDescGZIP(), []int{1}
}

func (x *GreetResponse) GetResultFirstName() string {
	if x != nil {
		return x.ResultFirstName
	}
	return ""
}

func (x *GreetResponse) GetResultLastName() string {
	if x != nil {
		return x.ResultLastName
	}
	return ""
}

func (x *GreetResponse) GetResultOccupation() string {
	if x != nil {
		return x.ResultOccupation
	}
	return ""
}

func (x *GreetResponse) GetResultLocation() string {
	if x != nil {
		return x.ResultLocation
	}
	return ""
}

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum uint32 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_greet_proto_rawDescGZIP(), []int{2}
}

func (x *SumRequest) GetSum() uint32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type TerraRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider     string            `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	ResourceName string            `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	NumberNeeded string            `protobuf:"bytes,3,opt,name=number_needed,json=numberNeeded,proto3" json:"number_needed,omitempty"`
	Prefix       string            `protobuf:"bytes,4,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Tags         map[string]string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TerraRequest) Reset() {
	*x = TerraRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerraRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerraRequest) ProtoMessage() {}

func (x *TerraRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerraRequest.ProtoReflect.Descriptor instead.
func (*TerraRequest) Descriptor() ([]byte, []int) {
	return file_greet_proto_rawDescGZIP(), []int{3}
}

func (x *TerraRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *TerraRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *TerraRequest) GetNumberNeeded() string {
	if x != nil {
		return x.NumberNeeded
	}
	return ""
}

func (x *TerraRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *TerraRequest) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type TerraResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider            string            `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	ResourceName        string            `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	NumberToBeRequested string            `protobuf:"bytes,3,opt,name=number_to_be_requested,json=numberToBeRequested,proto3" json:"number_to_be_requested,omitempty"`
	PlannedNamespace    string            `protobuf:"bytes,4,opt,name=planned_namespace,json=plannedNamespace,proto3" json:"planned_namespace,omitempty"`
	TagsToBeApplied     map[string]string `protobuf:"bytes,5,rep,name=tags_to_be_applied,json=tagsToBeApplied,proto3" json:"tags_to_be_applied,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TerraResponse) Reset() {
	*x = TerraResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerraResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerraResponse) ProtoMessage() {}

func (x *TerraResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerraResponse.ProtoReflect.Descriptor instead.
func (*TerraResponse) Descriptor() ([]byte, []int) {
	return file_greet_proto_rawDescGZIP(), []int{4}
}

func (x *TerraResponse) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *TerraResponse) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *TerraResponse) GetNumberToBeRequested() string {
	if x != nil {
		return x.NumberToBeRequested
	}
	return ""
}

func (x *TerraResponse) GetPlannedNamespace() string {
	if x != nil {
		return x.PlannedNamespace
	}
	return ""
}

func (x *TerraResponse) GetTagsToBeApplied() map[string]string {
	if x != nil {
		return x.TagsToBeApplied
	}
	return nil
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result uint32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_greet_proto_rawDescGZIP(), []int{5}
}

func (x *SumResponse) GetResult() uint32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_greet_proto protoreflect.FileDescriptor

var file_greet_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x22, 0x86, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbb, 0x01,
	0x0a, 0x0d, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f,
	0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1e, 0x0a, 0x0a, 0x53,
	0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x22, 0xf8, 0x01, 0x0a, 0x0c,
	0x54, 0x65, 0x72, 0x72, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4e, 0x65, 0x65, 0x64,
	0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x61,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x37, 0x0a,
	0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xce, 0x02, 0x0a, 0x0d, 0x54, 0x65, 0x72, 0x72, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x62, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x54, 0x6f, 0x42, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x2b,
	0x0a, 0x11, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x6c, 0x61, 0x6e, 0x6e,
	0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x12, 0x74,
	0x61, 0x67, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x62, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e,
	0x54, 0x65, 0x72, 0x72, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x61,
	0x67, 0x73, 0x54, 0x6f, 0x42, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0f, 0x74, 0x61, 0x67, 0x73, 0x54, 0x6f, 0x42, 0x65, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x64, 0x1a, 0x42, 0x0a, 0x14, 0x54, 0x61, 0x67, 0x73, 0x54, 0x6f, 0x42, 0x65, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x42,
	0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32,
	0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x3a, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2c, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e,
	0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x46,
	0x0a, 0x10, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x54, 0x65, 0x72, 0x72, 0x61, 0x12, 0x13, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x68, 0x65, 0x73, 0x73, 0x34, 0x30, 0x2f, 0x41, 0x6c,
	0x70, 0x68, 0x61, 0x67, 0x52, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greet_proto_rawDescOnce sync.Once
	file_greet_proto_rawDescData = file_greet_proto_rawDesc
)

func file_greet_proto_rawDescGZIP() []byte {
	file_greet_proto_rawDescOnce.Do(func() {
		file_greet_proto_rawDescData = protoimpl.X.CompressGZIP(file_greet_proto_rawDescData)
	})
	return file_greet_proto_rawDescData
}

var file_greet_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_greet_proto_goTypes = []interface{}{
	(*GreetRequest)(nil),  // 0: greet.GreetRequest
	(*GreetResponse)(nil), // 1: greet.GreetResponse
	(*SumRequest)(nil),    // 2: greet.SumRequest
	(*TerraRequest)(nil),  // 3: greet.TerraRequest
	(*TerraResponse)(nil), // 4: greet.TerraResponse
	(*SumResponse)(nil),   // 5: greet.SumResponse
	nil,                   // 6: greet.TerraRequest.TagsEntry
	nil,                   // 7: greet.TerraResponse.TagsToBeAppliedEntry
}
var file_greet_proto_depIdxs = []int32{
	6, // 0: greet.TerraRequest.tags:type_name -> greet.TerraRequest.TagsEntry
	7, // 1: greet.TerraResponse.tags_to_be_applied:type_name -> greet.TerraResponse.TagsToBeAppliedEntry
	0, // 2: greet.GreetService.Greet:input_type -> greet.GreetRequest
	2, // 3: greet.SumService.Sum:input_type -> greet.SumRequest
	3, // 4: greet.TerraformService.Terra:input_type -> greet.TerraRequest
	1, // 5: greet.GreetService.Greet:output_type -> greet.GreetResponse
	5, // 6: greet.SumService.Sum:output_type -> greet.SumResponse
	4, // 7: greet.TerraformService.Terra:output_type -> greet.TerraResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_greet_proto_init() }
func file_greet_proto_init() {
	if File_greet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRequest); i {
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
		file_greet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetResponse); i {
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
		file_greet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_greet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerraRequest); i {
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
		file_greet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerraResponse); i {
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
		file_greet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
			RawDescriptor: file_greet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_greet_proto_goTypes,
		DependencyIndexes: file_greet_proto_depIdxs,
		MessageInfos:      file_greet_proto_msgTypes,
	}.Build()
	File_greet_proto = out.File
	file_greet_proto_rawDesc = nil
	file_greet_proto_goTypes = nil
	file_greet_proto_depIdxs = nil
}
// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/artifacts_store.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Artifacts store access specification
type ArtifactsStore struct {
	// UUID
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Optional name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Optional description
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Optional a readme text describing this entity
	Readme string `protobuf:"bytes,4,opt,name=readme,proto3" json:"readme,omitempty"`
	// Optional Tags of this entity
	Tags []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	// Optional time when the entityt was created
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Optional last time the entity was updated
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Optional a flag to freeze an access
	Frozen bool `protobuf:"varint,8,opt,name=frozen,proto3" json:"frozen,omitempty"`
	// Optional a flag to disable an access
	Disabled bool `protobuf:"varint,9,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Optional if the entity has been deleted
	Deleted bool `protobuf:"varint,10,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// Optional the k8s secret to use
	K8SSecret string `protobuf:"bytes,11,opt,name=k8s_secret,json=k8sSecret,proto3" json:"k8s_secret,omitempty"`
	// Optional type of the store
	Type string `protobuf:"bytes,12,opt,name=type,proto3" json:"type,omitempty"`
	// Optional mounth path
	MountPath string `protobuf:"bytes,13,opt,name=mount_path,json=mountPath,proto3" json:"mount_path,omitempty"`
	// Optional host path
	HostPath string `protobuf:"bytes,14,opt,name=host_path,json=hostPath,proto3" json:"host_path,omitempty"`
	// Optional volume claim
	VolumeClaim string `protobuf:"bytes,15,opt,name=volume_claim,json=volumeClaim,proto3" json:"volume_claim,omitempty"`
	// Optional bucket
	Bucket string `protobuf:"bytes,16,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// Optional flag to set this store to read only mode
	ReadOnly             bool     `protobuf:"varint,17,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactsStore) Reset()         { *m = ArtifactsStore{} }
func (m *ArtifactsStore) String() string { return proto.CompactTextString(m) }
func (*ArtifactsStore) ProtoMessage()    {}
func (*ArtifactsStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_53cb0be1952660cd, []int{0}
}

func (m *ArtifactsStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactsStore.Unmarshal(m, b)
}
func (m *ArtifactsStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactsStore.Marshal(b, m, deterministic)
}
func (m *ArtifactsStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactsStore.Merge(m, src)
}
func (m *ArtifactsStore) XXX_Size() int {
	return xxx_messageInfo_ArtifactsStore.Size(m)
}
func (m *ArtifactsStore) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactsStore.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactsStore proto.InternalMessageInfo

func (m *ArtifactsStore) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ArtifactsStore) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArtifactsStore) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ArtifactsStore) GetReadme() string {
	if m != nil {
		return m.Readme
	}
	return ""
}

func (m *ArtifactsStore) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ArtifactsStore) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ArtifactsStore) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *ArtifactsStore) GetFrozen() bool {
	if m != nil {
		return m.Frozen
	}
	return false
}

func (m *ArtifactsStore) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *ArtifactsStore) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *ArtifactsStore) GetK8SSecret() string {
	if m != nil {
		return m.K8SSecret
	}
	return ""
}

func (m *ArtifactsStore) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ArtifactsStore) GetMountPath() string {
	if m != nil {
		return m.MountPath
	}
	return ""
}

func (m *ArtifactsStore) GetHostPath() string {
	if m != nil {
		return m.HostPath
	}
	return ""
}

func (m *ArtifactsStore) GetVolumeClaim() string {
	if m != nil {
		return m.VolumeClaim
	}
	return ""
}

func (m *ArtifactsStore) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *ArtifactsStore) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

// Request data to create/update artifacts store
type ArtifactsStoreBodyRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Artifact store body
	ArtifactStore        *ArtifactsStore `protobuf:"bytes,2,opt,name=artifact_store,json=artifactStore,proto3" json:"artifact_store,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ArtifactsStoreBodyRequest) Reset()         { *m = ArtifactsStoreBodyRequest{} }
func (m *ArtifactsStoreBodyRequest) String() string { return proto.CompactTextString(m) }
func (*ArtifactsStoreBodyRequest) ProtoMessage()    {}
func (*ArtifactsStoreBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_53cb0be1952660cd, []int{1}
}

func (m *ArtifactsStoreBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactsStoreBodyRequest.Unmarshal(m, b)
}
func (m *ArtifactsStoreBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactsStoreBodyRequest.Marshal(b, m, deterministic)
}
func (m *ArtifactsStoreBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactsStoreBodyRequest.Merge(m, src)
}
func (m *ArtifactsStoreBodyRequest) XXX_Size() int {
	return xxx_messageInfo_ArtifactsStoreBodyRequest.Size(m)
}
func (m *ArtifactsStoreBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactsStoreBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactsStoreBodyRequest proto.InternalMessageInfo

func (m *ArtifactsStoreBodyRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ArtifactsStoreBodyRequest) GetArtifactStore() *ArtifactsStore {
	if m != nil {
		return m.ArtifactStore
	}
	return nil
}

// Contains list artifacts stores
type ListArtifactsStoresResponse struct {
	// Count of the entities
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// List of all entities
	Results []*ArtifactsStore `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	// Previous page
	Previous string `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
	// Next page
	Next                 string   `protobuf:"bytes,4,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListArtifactsStoresResponse) Reset()         { *m = ListArtifactsStoresResponse{} }
func (m *ListArtifactsStoresResponse) String() string { return proto.CompactTextString(m) }
func (*ListArtifactsStoresResponse) ProtoMessage()    {}
func (*ListArtifactsStoresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_53cb0be1952660cd, []int{2}
}

func (m *ListArtifactsStoresResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListArtifactsStoresResponse.Unmarshal(m, b)
}
func (m *ListArtifactsStoresResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListArtifactsStoresResponse.Marshal(b, m, deterministic)
}
func (m *ListArtifactsStoresResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListArtifactsStoresResponse.Merge(m, src)
}
func (m *ListArtifactsStoresResponse) XXX_Size() int {
	return xxx_messageInfo_ListArtifactsStoresResponse.Size(m)
}
func (m *ListArtifactsStoresResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListArtifactsStoresResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListArtifactsStoresResponse proto.InternalMessageInfo

func (m *ListArtifactsStoresResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListArtifactsStoresResponse) GetResults() []*ArtifactsStore {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ListArtifactsStoresResponse) GetPrevious() string {
	if m != nil {
		return m.Previous
	}
	return ""
}

func (m *ListArtifactsStoresResponse) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

func init() {
	proto.RegisterType((*ArtifactsStore)(nil), "v1.ArtifactsStore")
	proto.RegisterType((*ArtifactsStoreBodyRequest)(nil), "v1.ArtifactsStoreBodyRequest")
	proto.RegisterType((*ListArtifactsStoresResponse)(nil), "v1.ListArtifactsStoresResponse")
}

func init() { proto.RegisterFile("v1/artifacts_store.proto", fileDescriptor_53cb0be1952660cd) }

var fileDescriptor_53cb0be1952660cd = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0xbf, 0x5b, 0x77, 0xb7, 0x80, 0x85, 0x90, 0xe9, 0x0a, 0x11, 0x7a, 0xea, 0x01, 0xa5,
	0xea, 0x72, 0xd9, 0x3d, 0x16, 0xae, 0x48, 0xa0, 0x2c, 0xf7, 0xc8, 0x8d, 0xa7, 0x6d, 0xd4, 0x24,
	0x0e, 0xf6, 0xb8, 0x50, 0xfe, 0xc5, 0xfe, 0x63, 0x34, 0x76, 0xb2, 0xa2, 0x07, 0xb4, 0xb7, 0x79,
	0x6f, 0xe6, 0x8d, 0x67, 0x3c, 0x8f, 0x89, 0xd3, 0x7a, 0x25, 0x0d, 0xe6, 0x3b, 0x99, 0xa1, 0x4d,
	0x2d, 0x6a, 0x03, 0x71, 0x6d, 0x34, 0x6a, 0xde, 0x3d, 0xad, 0xe7, 0xef, 0xf7, 0x5a, 0xef, 0x0b,
	0x58, 0x79, 0x66, 0xeb, 0x76, 0x2b, 0xcc, 0x4b, 0xb0, 0x28, 0xcb, 0x3a, 0x14, 0x2d, 0x1e, 0xfb,
	0x6c, 0xb6, 0x69, 0xe5, 0x0f, 0xa4, 0xe6, 0x9c, 0xf5, 0x9d, 0xcb, 0x95, 0xe8, 0x44, 0x9d, 0xe5,
	0x24, 0xf1, 0x31, 0x71, 0x95, 0x2c, 0x41, 0x74, 0x03, 0x47, 0x31, 0x8f, 0xd8, 0x54, 0x81, 0xcd,
	0x4c, 0x5e, 0x63, 0xae, 0x2b, 0xd1, 0xf3, 0xa9, 0x7f, 0x29, 0xfe, 0x86, 0x0d, 0x0d, 0x48, 0x55,
	0x82, 0xe8, 0xfb, 0x64, 0x83, 0xa8, 0x1b, 0xca, 0xbd, 0x15, 0x83, 0xa8, 0x47, 0xdd, 0x28, 0xe6,
	0xf7, 0x8c, 0x65, 0x06, 0x24, 0x82, 0x4a, 0x25, 0x8a, 0x61, 0xd4, 0x59, 0x4e, 0x6f, 0xe7, 0x71,
	0x18, 0x3f, 0x6e, 0xc7, 0x8f, 0x7f, 0xb4, 0xe3, 0x27, 0x93, 0xa6, 0x7a, 0x83, 0x24, 0x75, 0xb5,
	0x6a, 0xa5, 0xa3, 0xe7, 0xa5, 0x4d, 0xf5, 0x06, 0x69, 0xc2, 0x9d, 0xd1, 0x7f, 0xa0, 0x12, 0xe3,
	0xa8, 0xb3, 0x1c, 0x27, 0x0d, 0xe2, 0x73, 0x36, 0x56, 0xb9, 0x95, 0xdb, 0x02, 0x94, 0x98, 0xf8,
	0xcc, 0x13, 0xe6, 0x82, 0x8d, 0x14, 0x14, 0x80, 0xa0, 0x04, 0xf3, 0xa9, 0x16, 0xf2, 0x77, 0x8c,
	0x1d, 0xef, 0x6c, 0x6a, 0x21, 0x33, 0x80, 0x62, 0xea, 0x77, 0x9e, 0x1c, 0xef, 0xec, 0x83, 0x27,
	0xfc, 0xda, 0xe7, 0x1a, 0xc4, 0x55, 0xf8, 0x44, 0x8a, 0x49, 0x52, 0x6a, 0x57, 0x61, 0x5a, 0x4b,
	0x3c, 0x88, 0xeb, 0x20, 0xf1, 0xcc, 0x77, 0x89, 0x07, 0x7e, 0xc3, 0x26, 0x07, 0x6d, 0x9b, 0xec,
	0xcc, 0x67, 0xc7, 0x44, 0xf8, 0xe4, 0x07, 0x76, 0x75, 0xd2, 0x85, 0x2b, 0x21, 0xcd, 0x0a, 0x99,
	0x97, 0xe2, 0x45, 0xb8, 0x40, 0xe0, 0xbe, 0x10, 0x45, 0xfb, 0x6d, 0x5d, 0x76, 0x04, 0x14, 0x2f,
	0xc3, 0x05, 0x02, 0xa2, 0xbe, 0x74, 0x8b, 0x54, 0x57, 0xc5, 0x59, 0xbc, 0x0a, 0x0b, 0x12, 0xf1,
	0xad, 0x2a, 0xce, 0x8b, 0x82, 0xbd, 0xbd, 0xb4, 0xc4, 0x67, 0xad, 0xce, 0x09, 0xfc, 0x74, 0x60,
	0x91, 0xbf, 0x66, 0x03, 0xfd, 0xab, 0x02, 0xd3, 0xd8, 0x23, 0x00, 0x7e, 0xcf, 0x66, 0xad, 0x09,
	0x83, 0x07, 0xbd, 0x53, 0xa6, 0xb7, 0x3c, 0x3e, 0xad, 0xe3, 0xcb, 0x66, 0xc9, 0x75, 0x5b, 0xe9,
	0xe1, 0xe2, 0xb1, 0xc3, 0x6e, 0xbe, 0xe6, 0x16, 0x2f, 0xab, 0x6c, 0x02, 0xb6, 0xd6, 0x95, 0x05,
	0x7a, 0x30, 0xa3, 0xff, 0xf0, 0x0f, 0x0e, 0x92, 0x00, 0xf8, 0x47, 0x36, 0x32, 0x60, 0x5d, 0x81,
	0x56, 0x74, 0xa3, 0xde, 0x7f, 0x5e, 0x6a, 0x4b, 0xe8, 0x9c, 0xb5, 0x81, 0x53, 0xae, 0x9d, 0x6d,
	0x7c, 0xfa, 0x84, 0xbd, 0xb5, 0xe1, 0x37, 0x36, 0x16, 0xf5, 0xf1, 0x76, 0xe8, 0x5d, 0xf3, 0xe9,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0x66, 0x6a, 0x76, 0x5d, 0x03, 0x00, 0x00,
}

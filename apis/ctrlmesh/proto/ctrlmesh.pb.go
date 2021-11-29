// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: apis/ctrlmesh/proto/ctrlmesh.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ProxyStatusV1 struct {
	SelfInfo             *SelfInfo    `protobuf:"bytes,1,opt,name=selfInfo,proto3" json:"selfInfo,omitempty"`
	CurrentSpec          *ProxySpecV1 `protobuf:"bytes,2,opt,name=currentSpec,proto3" json:"currentSpec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProxyStatusV1) Reset()         { *m = ProxyStatusV1{} }
func (m *ProxyStatusV1) String() string { return proto.CompactTextString(m) }
func (*ProxyStatusV1) ProtoMessage()    {}
func (*ProxyStatusV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_015120ea8d6d1ea6, []int{0}
}
func (m *ProxyStatusV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyStatusV1.Unmarshal(m, b)
}
func (m *ProxyStatusV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyStatusV1.Marshal(b, m, deterministic)
}
func (m *ProxyStatusV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyStatusV1.Merge(m, src)
}
func (m *ProxyStatusV1) XXX_Size() int {
	return xxx_messageInfo_ProxyStatusV1.Size(m)
}
func (m *ProxyStatusV1) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyStatusV1.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyStatusV1 proto.InternalMessageInfo

func (m *ProxyStatusV1) GetSelfInfo() *SelfInfo {
	if m != nil {
		return m.SelfInfo
	}
	return nil
}

func (m *ProxyStatusV1) GetCurrentSpec() *ProxySpecV1 {
	if m != nil {
		return m.CurrentSpec
	}
	return nil
}

type SelfInfo struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelfInfo) Reset()         { *m = SelfInfo{} }
func (m *SelfInfo) String() string { return proto.CompactTextString(m) }
func (*SelfInfo) ProtoMessage()    {}
func (*SelfInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_015120ea8d6d1ea6, []int{1}
}
func (m *SelfInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelfInfo.Unmarshal(m, b)
}
func (m *SelfInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelfInfo.Marshal(b, m, deterministic)
}
func (m *SelfInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelfInfo.Merge(m, src)
}
func (m *SelfInfo) XXX_Size() int {
	return xxx_messageInfo_SelfInfo.Size(m)
}
func (m *SelfInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SelfInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SelfInfo proto.InternalMessageInfo

func (m *SelfInfo) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SelfInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ProxySpecV1 struct {
	Meta                 *SpecMetaV1           `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Route                *RouteV1              `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	Endpoints            []*EndpointV1         `protobuf:"bytes,3,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	ControlInstruction   *ControlInstructionV1 `protobuf:"bytes,4,opt,name=controlInstruction,proto3" json:"controlInstruction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ProxySpecV1) Reset()         { *m = ProxySpecV1{} }
func (m *ProxySpecV1) String() string { return proto.CompactTextString(m) }
func (*ProxySpecV1) ProtoMessage()    {}
func (*ProxySpecV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_015120ea8d6d1ea6, []int{2}
}
func (m *ProxySpecV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxySpecV1.Unmarshal(m, b)
}
func (m *ProxySpecV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxySpecV1.Marshal(b, m, deterministic)
}
func (m *ProxySpecV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxySpecV1.Merge(m, src)
}
func (m *ProxySpecV1) XXX_Size() int {
	return xxx_messageInfo_ProxySpecV1.Size(m)
}
func (m *ProxySpecV1) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxySpecV1.DiscardUnknown(m)
}

var xxx_messageInfo_ProxySpecV1 proto.InternalMessageInfo

func (m *ProxySpecV1) GetMeta() *SpecMetaV1 {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ProxySpecV1) GetRoute() *RouteV1 {
	if m != nil {
		return m.Route
	}
	return nil
}

func (m *ProxySpecV1) GetEndpoints() []*EndpointV1 {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ProxySpecV1) GetControlInstruction() *ControlInstructionV1 {
	if m != nil {
		return m.ControlInstruction
	}
	return nil
}

type SpecMetaV1 struct {
	VAppName             string   `protobuf:"bytes,1,opt,name=vAppName,proto3" json:"vAppName,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpecMetaV1) Reset()         { *m = SpecMetaV1{} }
func (m *SpecMetaV1) String() string { return proto.CompactTextString(m) }
func (*SpecMetaV1) ProtoMessage()    {}
func (*SpecMetaV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_015120ea8d6d1ea6, []int{3}
}
func (m *SpecMetaV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecMetaV1.Unmarshal(m, b)
}
func (m *SpecMetaV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecMetaV1.Marshal(b, m, deterministic)
}
func (m *SpecMetaV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecMetaV1.Merge(m, src)
}
func (m *SpecMetaV1) XXX_Size() int {
	return xxx_messageInfo_SpecMetaV1.Size(m)
}
func (m *SpecMetaV1) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecMetaV1.DiscardUnknown(m)
}

var xxx_messageInfo_SpecMetaV1 proto.InternalMessageInfo

func (m *SpecMetaV1) GetVAppName() string {
	if m != nil {
		return m.VAppName
	}
	return ""
}

func (m *SpecMetaV1) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type RouteV1 struct {
	Subset                string                `protobuf:"bytes,1,opt,name=subset,proto3" json:"subset,omitempty"`
	GlobalLimits          []*MatchLimitRuleV1   `protobuf:"bytes,2,rep,name=globalLimits,proto3" json:"globalLimits,omitempty"`
	SubsetLimits          []*SubsetLimitV1      `protobuf:"bytes,3,rep,name=subsetLimits,proto3" json:"subsetLimits,omitempty"`
	SubsetPublicResources []*APIGroupResourceV1 `protobuf:"bytes,4,rep,name=subsetPublicResources,proto3" json:"subsetPublicResources,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}              `json:"-"`
	XXX_unrecognized      []byte                `json:"-"`
	XXX_sizecache         int32                 `json:"-"`
}

func (m *RouteV1) Reset()         { *m = RouteV1{} }
func (m *RouteV1) String() string { return proto.CompactTextString(m) }
func (*RouteV1) ProtoMessage()    {}
func (*RouteV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_015120ea8d6d1ea6, []int{4}
}
func (m *RouteV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteV1.Unmarshal(m, b)
}
func (m *RouteV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteV1.Marshal(b, m, deterministic)
}
func (m *RouteV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteV1.Merge(m, src)
}
func (m *RouteV1) XXX_Size() int {
	return xxx_messageInfo_RouteV1.Size(m)
}
func (m *RouteV1) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteV1.DiscardUnknown(m)
}

var xxx_messageInfo_RouteV1 proto.InternalMessageInfo

func (m *RouteV1) GetSubset() string {
	if m != nil {
		return m.Subset
	}
	return ""
}

func (m *RouteV1) GetGlobalLimits() []*MatchLimitRuleV1 {
	if m != nil {
		return m.GlobalLimits
	}
	return nil
}

func (m *RouteV1) GetSubsetLimits() []*SubsetLimitV1 {
	if m != nil {
		return m.SubsetLimits
	}
	return nil
}

func (m *RouteV1) GetSubsetPublicResources() []*APIGroupResourceV1 {
	if m != nil {
		return m.SubsetPublicResources
	}
	return nil
}

type SubsetLimitV1 struct {
	Subset               string              `protobuf:"bytes,1,opt,name=subset,proto3" json:"subset,omitempty"`
	Limits               []*MatchLimitRuleV1 `protobuf:"bytes,2,rep,name=limits,proto3" json:"limits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SubsetLimitV1) Reset()         { *m = SubsetLimitV1{} }
func (m *SubsetLimitV1) String() string { return proto.CompactTextString(m) }
func (*SubsetLimitV1) ProtoMessage()    {}
func (*SubsetLimitV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_015120ea8d6d1ea6, []int{5}
}
func (m *SubsetLimitV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubsetLimitV1.Unmarshal(m, b)
}
func (m *SubsetLimitV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubsetLimitV1.Marshal(b, m, deterministic)
}
func (m *SubsetLimitV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubsetLimitV1.Merge(m, src)
}
func (m *SubsetLimitV1) XXX_Size() int {
	return xxx_messageInfo_SubsetLimitV1.Size(m)
}
func (m *SubsetLimitV1) XXX_DiscardUnknown() {
	xxx_messageInfo_SubsetLimitV1.DiscardUnknown(m)
}

var xxx_messageInfo_SubsetLimitV1 proto.InternalMessageInfo

func (m *SubsetLimitV1) GetSubset() string {
	if m != nil {
		return m.Subset
	}
	return ""
}

func (m *SubsetLimitV1) GetLimits() []*MatchLimitRuleV1 {
	if m != nil {
		return m.Limits
	}
	return nil
}

type MatchLimitRuleV1 struct {
	Namespaces []string `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	// TODO
	//string objectSelector = 2;
	Resources            []*APIGroupResourceV1 `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MatchLimitRuleV1) Reset()         { *m = MatchLimitRuleV1{} }
func (m *MatchLimitRuleV1) String() string { return proto.CompactTextString(m) }
func (*MatchLimitRuleV1) ProtoMessage()    {}
func (*MatchLimitRuleV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_015120ea8d6d1ea6, []int{6}
}
func (m *MatchLimitRuleV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchLimitRuleV1.Unmarshal(m, b)
}
func (m *MatchLimitRuleV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchLimitRuleV1.Marshal(b, m, deterministic)
}
func (m *MatchLimitRuleV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchLimitRuleV1.Merge(m, src)
}
func (m *MatchLimitRuleV1) XXX_Size() int {
	return xxx_messageInfo_MatchLimitRuleV1.Size(m)
}
func (m *MatchLimitRuleV1) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchLimitRuleV1.DiscardUnknown(m)
}

var xxx_messageInfo_MatchLimitRuleV1 proto.InternalMessageInfo

func (m *MatchLimitRuleV1) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *MatchLimitRuleV1) GetResources() []*APIGroupResourceV1 {
	if m != nil {
		return m.Resources
	}
	return nil
}

type APIGroupResourceV1 struct {
	ApiGroups            []string `protobuf:"bytes,1,rep,name=apiGroups,proto3" json:"apiGroups,omitempty"`
	Resources            []string `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIGroupResourceV1) Reset()         { *m = APIGroupResourceV1{} }
func (m *APIGroupResourceV1) String() string { return proto.CompactTextString(m) }
func (*APIGroupResourceV1) ProtoMessage()    {}
func (*APIGroupResourceV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_015120ea8d6d1ea6, []int{7}
}
func (m *APIGroupResourceV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIGroupResourceV1.Unmarshal(m, b)
}
func (m *APIGroupResourceV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIGroupResourceV1.Marshal(b, m, deterministic)
}
func (m *APIGroupResourceV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIGroupResourceV1.Merge(m, src)
}
func (m *APIGroupResourceV1) XXX_Size() int {
	return xxx_messageInfo_APIGroupResourceV1.Size(m)
}
func (m *APIGroupResourceV1) XXX_DiscardUnknown() {
	xxx_messageInfo_APIGroupResourceV1.DiscardUnknown(m)
}

var xxx_messageInfo_APIGroupResourceV1 proto.InternalMessageInfo

func (m *APIGroupResourceV1) GetApiGroups() []string {
	if m != nil {
		return m.ApiGroups
	}
	return nil
}

func (m *APIGroupResourceV1) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

type EndpointV1 struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Subset               string   `protobuf:"bytes,2,opt,name=subset,proto3" json:"subset,omitempty"`
	Ip                   string   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndpointV1) Reset()         { *m = EndpointV1{} }
func (m *EndpointV1) String() string { return proto.CompactTextString(m) }
func (*EndpointV1) ProtoMessage()    {}
func (*EndpointV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_015120ea8d6d1ea6, []int{8}
}
func (m *EndpointV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointV1.Unmarshal(m, b)
}
func (m *EndpointV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointV1.Marshal(b, m, deterministic)
}
func (m *EndpointV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointV1.Merge(m, src)
}
func (m *EndpointV1) XXX_Size() int {
	return xxx_messageInfo_EndpointV1.Size(m)
}
func (m *EndpointV1) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointV1.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointV1 proto.InternalMessageInfo

func (m *EndpointV1) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EndpointV1) GetSubset() string {
	if m != nil {
		return m.Subset
	}
	return ""
}

func (m *EndpointV1) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type ControlInstructionV1 struct {
	BlockLeaderElection  bool     `protobuf:"varint,1,opt,name=blockLeaderElection,proto3" json:"blockLeaderElection,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ControlInstructionV1) Reset()         { *m = ControlInstructionV1{} }
func (m *ControlInstructionV1) String() string { return proto.CompactTextString(m) }
func (*ControlInstructionV1) ProtoMessage()    {}
func (*ControlInstructionV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_015120ea8d6d1ea6, []int{9}
}
func (m *ControlInstructionV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlInstructionV1.Unmarshal(m, b)
}
func (m *ControlInstructionV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlInstructionV1.Marshal(b, m, deterministic)
}
func (m *ControlInstructionV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlInstructionV1.Merge(m, src)
}
func (m *ControlInstructionV1) XXX_Size() int {
	return xxx_messageInfo_ControlInstructionV1.Size(m)
}
func (m *ControlInstructionV1) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlInstructionV1.DiscardUnknown(m)
}

var xxx_messageInfo_ControlInstructionV1 proto.InternalMessageInfo

func (m *ControlInstructionV1) GetBlockLeaderElection() bool {
	if m != nil {
		return m.BlockLeaderElection
	}
	return false
}

func init() {
	proto.RegisterType((*ProxyStatusV1)(nil), "proto.ProxyStatusV1")
	proto.RegisterType((*SelfInfo)(nil), "proto.SelfInfo")
	proto.RegisterType((*ProxySpecV1)(nil), "proto.ProxySpecV1")
	proto.RegisterType((*SpecMetaV1)(nil), "proto.SpecMetaV1")
	proto.RegisterType((*RouteV1)(nil), "proto.RouteV1")
	proto.RegisterType((*SubsetLimitV1)(nil), "proto.SubsetLimitV1")
	proto.RegisterType((*MatchLimitRuleV1)(nil), "proto.MatchLimitRuleV1")
	proto.RegisterType((*APIGroupResourceV1)(nil), "proto.APIGroupResourceV1")
	proto.RegisterType((*EndpointV1)(nil), "proto.EndpointV1")
	proto.RegisterType((*ControlInstructionV1)(nil), "proto.ControlInstructionV1")
}

func init() {
	proto.RegisterFile("apis/ctrlmesh/proto/ctrlmesh.proto", fileDescriptor_015120ea8d6d1ea6)
}

var fileDescriptor_015120ea8d6d1ea6 = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0xec, 0xa4, 0x6d, 0x3c, 0xfd, 0xf9, 0x60, 0x28, 0x60, 0x0a, 0x42, 0x91, 0x05, 0x52,
	0x24, 0xa4, 0xa6, 0x2e, 0x48, 0x20, 0x91, 0x9b, 0x82, 0x2a, 0x1a, 0xd1, 0x94, 0x68, 0x23, 0x59,
	0x88, 0x3b, 0x67, 0xbb, 0x6d, 0xac, 0x3a, 0xde, 0xd5, 0xee, 0x1a, 0xc1, 0x9b, 0xf2, 0x0a, 0xbc,
	0x05, 0xf2, 0x7a, 0xfd, 0x13, 0x1a, 0x04, 0x57, 0xf6, 0xcc, 0x39, 0x73, 0x66, 0xe6, 0x68, 0x77,
	0x21, 0x88, 0x45, 0xa2, 0x86, 0x54, 0xcb, 0x74, 0xc9, 0xd4, 0x62, 0x28, 0x24, 0xd7, 0xbc, 0x0e,
	0x0f, 0x4d, 0x88, 0x1b, 0xe6, 0x13, 0x48, 0xd8, 0x9d, 0x4a, 0xfe, 0xed, 0xfb, 0x4c, 0xc7, 0x3a,
	0x57, 0x51, 0x88, 0x2f, 0xa0, 0xa7, 0x58, 0x7a, 0x35, 0xce, 0xae, 0xb8, 0xef, 0xf4, 0x9d, 0xc1,
	0xf6, 0xf1, 0xff, 0x65, 0xc5, 0xe1, 0xcc, 0xa6, 0x49, 0x4d, 0xc0, 0x57, 0xb0, 0x4d, 0x73, 0x29,
	0x59, 0xa6, 0x67, 0x82, 0x51, 0xdf, 0x35, 0x7c, 0xb4, 0xfc, 0x52, 0x57, 0x30, 0x1a, 0x85, 0xa4,
	0x4d, 0x0b, 0x46, 0xd0, 0xab, 0xb4, 0xf0, 0x09, 0x78, 0x59, 0xbc, 0x64, 0x4a, 0xc4, 0x94, 0x99,
	0x7e, 0x1e, 0x69, 0x12, 0x88, 0xd0, 0x2d, 0x02, 0x23, 0xec, 0x11, 0xf3, 0x1f, 0xfc, 0x70, 0x60,
	0xbb, 0x25, 0x8d, 0xcf, 0xa1, 0xbb, 0x64, 0x3a, 0xb6, 0xc3, 0xde, 0xad, 0x86, 0x15, 0x8c, 0x4e,
	0x98, 0x8e, 0xa3, 0x90, 0x18, 0x18, 0x9f, 0xc1, 0x86, 0xe4, 0xb9, 0x66, 0x76, 0xc8, 0x3d, 0xcb,
	0x23, 0x45, 0x2e, 0x0a, 0x49, 0x09, 0xe2, 0x10, 0x3c, 0x96, 0x5d, 0x0a, 0x9e, 0x64, 0x5a, 0xf9,
	0x9d, 0x7e, 0xa7, 0xa5, 0x78, 0x6a, 0xf3, 0x51, 0x48, 0x1a, 0x0e, 0x7e, 0x04, 0xa4, 0x3c, 0xd3,
	0x92, 0xa7, 0xe3, 0x4c, 0x69, 0x99, 0x53, 0x9d, 0xf0, 0xcc, 0xef, 0x9a, 0x1e, 0x8f, 0x6d, 0xe5,
	0xfb, 0x5b, 0x84, 0x28, 0x24, 0x6b, 0xca, 0x82, 0x11, 0x40, 0x33, 0x37, 0x1e, 0x40, 0xef, 0xeb,
	0x89, 0x10, 0x17, 0x85, 0x01, 0xa5, 0x33, 0x75, 0x5c, 0x18, 0xb3, 0x88, 0xd5, 0xa2, 0x32, 0xa6,
	0xf8, 0x0f, 0x7e, 0x3a, 0xb0, 0x65, 0xd7, 0xc1, 0x07, 0xb0, 0xa9, 0xf2, 0xb9, 0x62, 0xda, 0x56,
	0xda, 0x08, 0xdf, 0xc2, 0xce, 0x75, 0xca, 0xe7, 0x71, 0x7a, 0x9e, 0x2c, 0x13, 0xad, 0x7c, 0xd7,
	0xac, 0xf8, 0xd0, 0x0e, 0x3a, 0x89, 0x35, 0x5d, 0x18, 0x84, 0xe4, 0x69, 0xe1, 0xca, 0x0a, 0x19,
	0xdf, 0xc0, 0x4e, 0x29, 0x63, 0x8b, 0x4b, 0x7f, 0xf6, 0x2b, 0xc7, 0x1b, 0xa8, 0xa8, 0x6c, 0x33,
	0xf1, 0x13, 0xdc, 0x2f, 0xe3, 0x69, 0x3e, 0x4f, 0x13, 0x4a, 0x98, 0xe2, 0xb9, 0xa4, 0x4c, 0xf9,
	0x5d, 0x23, 0xf1, 0xc8, 0x4a, 0x9c, 0x4c, 0xc7, 0x1f, 0x24, 0xcf, 0x45, 0x85, 0x47, 0x21, 0x59,
	0x5f, 0x17, 0x7c, 0x86, 0xdd, 0x95, 0x7e, 0x7f, 0x5c, 0x78, 0x08, 0x9b, 0xe9, 0x3f, 0xad, 0x6a,
	0x69, 0xc1, 0x0d, 0xdc, 0xf9, 0x1d, 0xc3, 0xa7, 0x00, 0xf5, 0x99, 0x54, 0xbe, 0xd3, 0xef, 0x0c,
	0x3c, 0xd2, 0xca, 0xe0, 0x6b, 0xf0, 0x64, 0xbd, 0x52, 0xe7, 0x6f, 0x2b, 0x35, 0xdc, 0x60, 0x0a,
	0x78, 0x9b, 0x50, 0xdc, 0x89, 0x58, 0x24, 0x26, 0x5b, 0x75, 0x6b, 0x12, 0x05, 0xda, 0x34, 0x73,
	0x4b, 0xb4, 0x51, 0x3c, 0x03, 0x68, 0x0e, 0x6a, 0x7d, 0x7f, 0x9c, 0xe6, 0xfe, 0xb4, 0x9c, 0x72,
	0x57, 0x9c, 0xda, 0x03, 0x37, 0x11, 0x7e, 0xc7, 0xe4, 0xdc, 0x44, 0x04, 0x67, 0xb0, 0xbf, 0xee,
	0xe0, 0xe2, 0x11, 0xdc, 0x9b, 0xa7, 0x9c, 0xde, 0x9c, 0xb3, 0xf8, 0x92, 0xc9, 0xd3, 0x94, 0x95,
	0x47, 0xbe, 0x68, 0xd1, 0x23, 0xeb, 0xa0, 0xe3, 0x0b, 0xd8, 0xb3, 0x4a, 0x29, 0x93, 0x13, 0xa6,
	0x16, 0x38, 0x02, 0x20, 0xec, 0x3a, 0x51, 0x9a, 0xc9, 0x28, 0xc4, 0xfd, 0x95, 0x07, 0xc3, 0x3e,
	0x44, 0x07, 0x6b, 0x9e, 0x91, 0xe0, 0xbf, 0x81, 0x73, 0xe4, 0xbc, 0xdb, 0xfa, 0x52, 0x3e, 0x5e,
	0xf3, 0x4d, 0xf3, 0x79, 0xf9, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x7d, 0x0b, 0x34, 0xf0, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ControllerMeshClient is the client API for ControllerMesh service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControllerMeshClient interface {
	RegisterV1(ctx context.Context, opts ...grpc.CallOption) (ControllerMesh_RegisterV1Client, error)
}

type controllerMeshClient struct {
	cc *grpc.ClientConn
}

func NewControllerMeshClient(cc *grpc.ClientConn) ControllerMeshClient {
	return &controllerMeshClient{cc}
}

func (c *controllerMeshClient) RegisterV1(ctx context.Context, opts ...grpc.CallOption) (ControllerMesh_RegisterV1Client, error) {
	stream, err := c.cc.NewStream(ctx, &_ControllerMesh_serviceDesc.Streams[0], "/proto.ControllerMesh/RegisterV1", opts...)
	if err != nil {
		return nil, err
	}
	x := &controllerMeshRegisterV1Client{stream}
	return x, nil
}

type ControllerMesh_RegisterV1Client interface {
	Send(*ProxyStatusV1) error
	Recv() (*ProxySpecV1, error)
	grpc.ClientStream
}

type controllerMeshRegisterV1Client struct {
	grpc.ClientStream
}

func (x *controllerMeshRegisterV1Client) Send(m *ProxyStatusV1) error {
	return x.ClientStream.SendMsg(m)
}

func (x *controllerMeshRegisterV1Client) Recv() (*ProxySpecV1, error) {
	m := new(ProxySpecV1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ControllerMeshServer is the server API for ControllerMesh service.
type ControllerMeshServer interface {
	RegisterV1(ControllerMesh_RegisterV1Server) error
}

// UnimplementedControllerMeshServer can be embedded to have forward compatible implementations.
type UnimplementedControllerMeshServer struct {
}

func (*UnimplementedControllerMeshServer) RegisterV1(srv ControllerMesh_RegisterV1Server) error {
	return status.Errorf(codes.Unimplemented, "method RegisterV1 not implemented")
}

func RegisterControllerMeshServer(s *grpc.Server, srv ControllerMeshServer) {
	s.RegisterService(&_ControllerMesh_serviceDesc, srv)
}

func _ControllerMesh_RegisterV1_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControllerMeshServer).RegisterV1(&controllerMeshRegisterV1Server{stream})
}

type ControllerMesh_RegisterV1Server interface {
	Send(*ProxySpecV1) error
	Recv() (*ProxyStatusV1, error)
	grpc.ServerStream
}

type controllerMeshRegisterV1Server struct {
	grpc.ServerStream
}

func (x *controllerMeshRegisterV1Server) Send(m *ProxySpecV1) error {
	return x.ServerStream.SendMsg(m)
}

func (x *controllerMeshRegisterV1Server) Recv() (*ProxyStatusV1, error) {
	m := new(ProxyStatusV1)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ControllerMesh_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ControllerMesh",
	HandlerType: (*ControllerMeshServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterV1",
			Handler:       _ControllerMesh_RegisterV1_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "apis/ctrlmesh/proto/ctrlmesh.proto",
}

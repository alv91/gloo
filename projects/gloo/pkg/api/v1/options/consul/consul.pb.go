// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/consul/consul.proto

package consul

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	options "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// Upstream Spec for Consul Upstreams
// consul Upstreams represent a set of one or more addressable pods for a consul Service
// the Gloo consul Upstream maps to a single service port. Because consul Services support multiple ports,
// Gloo requires that a different upstream be created for each port
// consul Upstreams are typically generated automatically by Gloo from the consul API
type UpstreamSpec struct {
	// The name of the Consul Service
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Deprecated: This field was renamed to subset_tags. If subset_tags is used, this field is
	// ignored. Otherwise, the behavior is the same as subset_tags field below.
	ServiceTags []string `protobuf:"bytes,2,rep,name=service_tags,json=serviceTags,proto3" json:"service_tags,omitempty"`
	// Gloo will segment instances based off of these tags. This allows you to set routes that route
	// to a subset of the instances of the service
	SubsetTags []string `protobuf:"bytes,6,rep,name=subset_tags,json=subsetTags,proto3" json:"subset_tags,omitempty"`
	// The list of service tags Gloo should search for on a service instance
	// before deciding whether or not to include the instance as part of this
	// upstream. Empty list means that all service instances with the same service name will be
	// included. When not empty, only service instances that match all of the tags (subset match) will be selected
	// for this upstream.
	InstanceTags []string `protobuf:"bytes,7,rep,name=instance_tags,json=instanceTags,proto3" json:"instance_tags,omitempty"`
	// The opposite of instanceTags, this is a list of service tags that gloo should ensure are not
	// in a service instance before including it in an upstream.
	InstanceBlacklistTags []string `protobuf:"bytes,8,rep,name=instance_blacklist_tags,json=instanceBlacklistTags,proto3" json:"instance_blacklist_tags,omitempty"`
	// An optional Service Spec describing the service listening at this address
	ServiceSpec *options.ServiceSpec `protobuf:"bytes,3,opt,name=service_spec,json=serviceSpec,proto3" json:"service_spec,omitempty"`
	// Is this consul service connect enabled.
	ConnectEnabled bool `protobuf:"varint,4,opt,name=connect_enabled,json=connectEnabled,proto3" json:"connect_enabled,omitempty"`
	// The data centers in which the service instance represented by this upstream is registered.
	DataCenters          []string `protobuf:"bytes,5,rep,name=data_centers,json=dataCenters,proto3" json:"data_centers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5077911f8bc0ad, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (m *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(m, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

func (m *UpstreamSpec) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *UpstreamSpec) GetServiceTags() []string {
	if m != nil {
		return m.ServiceTags
	}
	return nil
}

func (m *UpstreamSpec) GetSubsetTags() []string {
	if m != nil {
		return m.SubsetTags
	}
	return nil
}

func (m *UpstreamSpec) GetInstanceTags() []string {
	if m != nil {
		return m.InstanceTags
	}
	return nil
}

func (m *UpstreamSpec) GetInstanceBlacklistTags() []string {
	if m != nil {
		return m.InstanceBlacklistTags
	}
	return nil
}

func (m *UpstreamSpec) GetServiceSpec() *options.ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

func (m *UpstreamSpec) GetConnectEnabled() bool {
	if m != nil {
		return m.ConnectEnabled
	}
	return false
}

func (m *UpstreamSpec) GetDataCenters() []string {
	if m != nil {
		return m.DataCenters
	}
	return nil
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "consul.options.gloo.solo.io.UpstreamSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/consul/consul.proto", fileDescriptor_3c5077911f8bc0ad)
}

var fileDescriptor_3c5077911f8bc0ad = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x4e, 0xeb, 0x30,
	0x14, 0x87, 0x95, 0xb6, 0xb7, 0xb7, 0xd7, 0xed, 0x05, 0x29, 0x02, 0x11, 0x15, 0x09, 0x52, 0x18,
	0xe8, 0x42, 0x22, 0xfe, 0x88, 0x15, 0xa9, 0x80, 0x04, 0x0b, 0x43, 0x0b, 0x0b, 0x4b, 0xe5, 0xb8,
	0x47, 0xc1, 0x34, 0xf1, 0xb1, 0x62, 0xb7, 0xea, 0x23, 0xf1, 0x08, 0x3c, 0x0f, 0xe2, 0x15, 0xd8,
	0x51, 0x6c, 0xa7, 0x74, 0xa8, 0x04, 0x53, 0xe2, 0xcf, 0xdf, 0x39, 0xc9, 0xf9, 0xd9, 0xe4, 0x36,
	0xe5, 0xfa, 0x79, 0x96, 0x44, 0x0c, 0xf3, 0x58, 0x61, 0x86, 0xc7, 0x1c, 0xe3, 0x34, 0x43, 0x8c,
	0x65, 0x81, 0x2f, 0xc0, 0xb4, 0xb2, 0x2b, 0x2a, 0x79, 0x3c, 0x3f, 0x89, 0x51, 0x6a, 0x8e, 0x42,
	0xc5, 0x0c, 0x85, 0x9a, 0x65, 0xee, 0x11, 0xc9, 0x02, 0x35, 0xfa, 0xbb, 0x6e, 0xe5, 0x9c, 0xa8,
	0xac, 0x8b, 0xca, 0x96, 0x11, 0xc7, 0xee, 0x56, 0x8a, 0x29, 0x1a, 0x2f, 0x2e, 0xdf, 0x6c, 0x49,
	0xd7, 0x87, 0x85, 0xb6, 0x10, 0x16, 0xda, 0xb1, 0xf3, 0x9f, 0xbf, 0xae, 0xa0, 0x98, 0x73, 0x06,
	0x63, 0x25, 0x81, 0xd9, 0xaa, 0x83, 0x8f, 0x1a, 0xe9, 0x3c, 0x4a, 0xa5, 0x0b, 0xa0, 0xf9, 0x48,
	0x02, 0xf3, 0x7b, 0xa4, 0x53, 0x69, 0x82, 0xe6, 0x10, 0x78, 0xa1, 0xd7, 0xff, 0x37, 0x6c, 0x3b,
	0x76, 0x4f, 0x73, 0x58, 0x55, 0x34, 0x4d, 0x55, 0x50, 0x0b, 0xeb, 0x2b, 0xca, 0x03, 0x4d, 0x95,
	0xbf, 0x4f, 0xda, 0x6a, 0x96, 0x28, 0xd0, 0xd6, 0x68, 0x1a, 0x83, 0x58, 0x64, 0x84, 0x43, 0xf2,
	0x9f, 0x0b, 0xa5, 0xa9, 0xa8, 0x9a, 0xfc, 0x35, 0x4a, 0xa7, 0x82, 0x46, 0xba, 0x20, 0x3b, 0x4b,
	0x29, 0xc9, 0x28, 0x9b, 0x66, 0x5c, 0xb9, 0x8e, 0x2d, 0xa3, 0x6f, 0x57, 0xdb, 0x83, 0x6a, 0xd7,
	0xd4, 0x5d, 0x7f, 0xff, 0x60, 0x39, 0x6a, 0x50, 0x0f, 0xbd, 0x7e, 0xfb, 0xb4, 0xb7, 0x36, 0xe1,
	0x68, 0x64, 0xcd, 0x72, 0xf8, 0xe5, 0x0c, 0x26, 0x89, 0x23, 0xb2, 0xc9, 0x50, 0x08, 0x60, 0x7a,
	0x0c, 0x82, 0x26, 0x19, 0x4c, 0x82, 0x46, 0xe8, 0xf5, 0x5b, 0xc3, 0x0d, 0x87, 0x6f, 0x2c, 0x2d,
	0xf3, 0x98, 0x50, 0x4d, 0xc7, 0x0c, 0x84, 0x86, 0x42, 0x05, 0x7f, 0x6c, 0x1e, 0x25, 0xbb, 0xb2,
	0x68, 0x70, 0xf7, 0xf6, 0xd9, 0xf0, 0x5e, 0xdf, 0xf7, 0xbc, 0xa7, 0xcb, 0xdf, 0xdd, 0x1b, 0x39,
	0x4d, 0xd7, 0xdf, 0x9d, 0xa4, 0x69, 0x0e, 0xee, 0xec, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x8c,
	0xa8, 0x40, 0x81, 0x02, 0x00, 0x00,
}

func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ServiceName != that1.ServiceName {
		return false
	}
	if len(this.ServiceTags) != len(that1.ServiceTags) {
		return false
	}
	for i := range this.ServiceTags {
		if this.ServiceTags[i] != that1.ServiceTags[i] {
			return false
		}
	}
	if len(this.SubsetTags) != len(that1.SubsetTags) {
		return false
	}
	for i := range this.SubsetTags {
		if this.SubsetTags[i] != that1.SubsetTags[i] {
			return false
		}
	}
	if len(this.InstanceTags) != len(that1.InstanceTags) {
		return false
	}
	for i := range this.InstanceTags {
		if this.InstanceTags[i] != that1.InstanceTags[i] {
			return false
		}
	}
	if len(this.InstanceBlacklistTags) != len(that1.InstanceBlacklistTags) {
		return false
	}
	for i := range this.InstanceBlacklistTags {
		if this.InstanceBlacklistTags[i] != that1.InstanceBlacklistTags[i] {
			return false
		}
	}
	if !this.ServiceSpec.Equal(that1.ServiceSpec) {
		return false
	}
	if this.ConnectEnabled != that1.ConnectEnabled {
		return false
	}
	if len(this.DataCenters) != len(that1.DataCenters) {
		return false
	}
	for i := range this.DataCenters {
		if this.DataCenters[i] != that1.DataCenters[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

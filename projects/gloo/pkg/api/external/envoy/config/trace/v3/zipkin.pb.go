// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/trace/v3/zipkin.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/annotations"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/udpa/annotations"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

// Available Zipkin collector endpoint versions.
type ZipkinConfig_CollectorEndpointVersion int32

const (
	// Zipkin API v1, JSON over HTTP.
	// [#comment: The default implementation of Zipkin client before this field is added was only v1
	// and the way user configure this was by not explicitly specifying the version. Consequently,
	// before this is added, the corresponding Zipkin collector expected to receive v1 payload.
	// Hence the motivation of adding HTTP_JSON_V1 as the default is to avoid a breaking change when
	// user upgrading Envoy with this change. Furthermore, we also immediately deprecate this field,
	// since in Zipkin realm this v1 version is considered to be not preferable anymore.]
	ZipkinConfig_DEPRECATED_AND_UNAVAILABLE_DO_NOT_USE ZipkinConfig_CollectorEndpointVersion = 0 // Deprecated: Do not use.
	// Zipkin API v2, JSON over HTTP.
	ZipkinConfig_HTTP_JSON ZipkinConfig_CollectorEndpointVersion = 1
	// Zipkin API v2, protobuf over HTTP.
	ZipkinConfig_HTTP_PROTO ZipkinConfig_CollectorEndpointVersion = 2
)

var ZipkinConfig_CollectorEndpointVersion_name = map[int32]string{
	0: "DEPRECATED_AND_UNAVAILABLE_DO_NOT_USE",
	1: "HTTP_JSON",
	2: "HTTP_PROTO",
}

var ZipkinConfig_CollectorEndpointVersion_value = map[string]int32{
	"DEPRECATED_AND_UNAVAILABLE_DO_NOT_USE": 0,
	"HTTP_JSON":                             1,
	"HTTP_PROTO":                            2,
}

func (x ZipkinConfig_CollectorEndpointVersion) String() string {
	return proto.EnumName(ZipkinConfig_CollectorEndpointVersion_name, int32(x))
}

func (ZipkinConfig_CollectorEndpointVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dba010877df2c16a, []int{0, 0}
}

// Configuration for the Zipkin tracer.
// [#extension: envoy.tracers.zipkin]
// [#next-free-field: 6]
type ZipkinConfig struct {
	// The upstream that hosts the Zipkin collectors. Note that the
	// Zipkin cluster must be defined in the :ref:`Bootstrap static cluster
	// resources <envoy_api_field_config.bootstrap.v3.Bootstrap.StaticResources.clusters>`.
	CollectorUpstreamRef *core.ResourceRef `protobuf:"bytes,1,opt,name=collector_upstream_ref,json=collectorUpstreamRef,proto3" json:"collector_upstream_ref,omitempty"`
	// The API endpoint of the Zipkin service where the spans will be sent. When
	// using a standard Zipkin installation, the API endpoint is typically
	// /api/v1/spans, which is the default value.
	CollectorEndpoint string `protobuf:"bytes,2,opt,name=collector_endpoint,json=collectorEndpoint,proto3" json:"collector_endpoint,omitempty"`
	// Determines whether a 128bit trace id will be used when creating a new
	// trace instance. The default value is false, which will result in a 64 bit trace id being used.
	TraceId_128Bit bool `protobuf:"varint,3,opt,name=trace_id_128bit,json=traceId128bit,proto3" json:"trace_id_128bit,omitempty"`
	// Determines whether client and server spans will share the same span context.
	// The default value is true.
	SharedSpanContext *types.BoolValue `protobuf:"bytes,4,opt,name=shared_span_context,json=sharedSpanContext,proto3" json:"shared_span_context,omitempty"`
	// Determines the selected collector endpoint version. By default, the ``HTTP_JSON_V1`` will be
	// used.
	CollectorEndpointVersion ZipkinConfig_CollectorEndpointVersion `protobuf:"varint,5,opt,name=collector_endpoint_version,json=collectorEndpointVersion,proto3,enum=envoy.config.trace.v3.ZipkinConfig_CollectorEndpointVersion" json:"collector_endpoint_version,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                              `json:"-"`
	XXX_unrecognized         []byte                                `json:"-"`
	XXX_sizecache            int32                                 `json:"-"`
}

func (m *ZipkinConfig) Reset()         { *m = ZipkinConfig{} }
func (m *ZipkinConfig) String() string { return proto.CompactTextString(m) }
func (*ZipkinConfig) ProtoMessage()    {}
func (*ZipkinConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_dba010877df2c16a, []int{0}
}
func (m *ZipkinConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZipkinConfig.Unmarshal(m, b)
}
func (m *ZipkinConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZipkinConfig.Marshal(b, m, deterministic)
}
func (m *ZipkinConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZipkinConfig.Merge(m, src)
}
func (m *ZipkinConfig) XXX_Size() int {
	return xxx_messageInfo_ZipkinConfig.Size(m)
}
func (m *ZipkinConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ZipkinConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ZipkinConfig proto.InternalMessageInfo

func (m *ZipkinConfig) GetCollectorUpstreamRef() *core.ResourceRef {
	if m != nil {
		return m.CollectorUpstreamRef
	}
	return nil
}

func (m *ZipkinConfig) GetCollectorEndpoint() string {
	if m != nil {
		return m.CollectorEndpoint
	}
	return ""
}

func (m *ZipkinConfig) GetTraceId_128Bit() bool {
	if m != nil {
		return m.TraceId_128Bit
	}
	return false
}

func (m *ZipkinConfig) GetSharedSpanContext() *types.BoolValue {
	if m != nil {
		return m.SharedSpanContext
	}
	return nil
}

func (m *ZipkinConfig) GetCollectorEndpointVersion() ZipkinConfig_CollectorEndpointVersion {
	if m != nil {
		return m.CollectorEndpointVersion
	}
	return ZipkinConfig_DEPRECATED_AND_UNAVAILABLE_DO_NOT_USE
}

func init() {
	proto.RegisterEnum("envoy.config.trace.v3.ZipkinConfig_CollectorEndpointVersion", ZipkinConfig_CollectorEndpointVersion_name, ZipkinConfig_CollectorEndpointVersion_value)
	proto.RegisterType((*ZipkinConfig)(nil), "envoy.config.trace.v3.ZipkinConfig")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/trace/v3/zipkin.proto", fileDescriptor_dba010877df2c16a)
}

var fileDescriptor_dba010877df2c16a = []byte{
	// 642 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x4f, 0x13, 0x4d,
	0x18, 0xfe, 0xb6, 0x1f, 0xf0, 0xc1, 0xf0, 0x81, 0x65, 0x45, 0x2d, 0x4d, 0x6c, 0x2a, 0xf8, 0xa3,
	0x1c, 0x9c, 0x09, 0xad, 0x31, 0xc6, 0x78, 0x69, 0x69, 0x13, 0x21, 0xa4, 0x6d, 0x96, 0x96, 0x03,
	0x97, 0xcd, 0x74, 0x3b, 0x5d, 0x46, 0x96, 0x79, 0x27, 0xb3, 0xb3, 0x2b, 0x70, 0xf2, 0x68, 0xe2,
	0xc1, 0xab, 0xf1, 0xe8, 0x89, 0x3f, 0xc0, 0x93, 0x89, 0x47, 0x13, 0xaf, 0xfe, 0x0b, 0xfe, 0x09,
	0x5e, 0x4c, 0x3c, 0x18, 0xb3, 0xb3, 0x0b, 0x8a, 0x40, 0xc2, 0x6d, 0x66, 0x9e, 0xe7, 0x79, 0x7f,
	0x3c, 0xef, 0xbc, 0xa8, 0xef, 0x73, 0xbd, 0x13, 0x0d, 0xb0, 0x07, 0x7b, 0x24, 0x84, 0x00, 0xee,
	0x73, 0x20, 0x7e, 0x00, 0x40, 0xa4, 0x82, 0x67, 0xcc, 0xd3, 0x61, 0x7a, 0xa3, 0x92, 0x13, 0xb6,
	0xaf, 0x99, 0x12, 0x34, 0x20, 0x4c, 0xc4, 0x70, 0x40, 0x3c, 0x10, 0x23, 0xee, 0x13, 0xad, 0xa8,
	0xc7, 0x48, 0x5c, 0x23, 0x87, 0x5c, 0xee, 0x72, 0x81, 0xa5, 0x02, 0x0d, 0xf6, 0x35, 0xc3, 0xc1,
	0x29, 0x07, 0x1b, 0x0e, 0x8e, 0x6b, 0xc5, 0x92, 0x0f, 0xe0, 0x07, 0x8c, 0x18, 0xd2, 0x20, 0x1a,
	0x91, 0xe7, 0x8a, 0x4a, 0xc9, 0x54, 0x98, 0xca, 0x8a, 0x4b, 0x69, 0x68, 0x2a, 0x04, 0x68, 0xaa,
	0x39, 0x88, 0x90, 0x0c, 0x99, 0x54, 0xcc, 0x33, 0x97, 0x8c, 0x54, 0x8a, 0x86, 0x92, 0x9e, 0xe2,
	0xec, 0x71, 0x5f, 0x51, 0xcd, 0x32, 0xfc, 0xe6, 0x19, 0x3c, 0xd4, 0x54, 0x47, 0xc7, 0x39, 0x6e,
	0x9d, 0x81, 0x63, 0xa6, 0x42, 0x0e, 0x82, 0x0b, 0x3f, 0xa3, 0xdc, 0x88, 0x69, 0xc0, 0x87, 0x54,
	0x33, 0x72, 0x7c, 0xc8, 0x80, 0x05, 0x63, 0xd1, 0x2e, 0xd7, 0xc6, 0x90, 0x78, 0x85, 0x28, 0x36,
	0xca, 0xa0, 0x79, 0x1f, 0x7c, 0x30, 0x47, 0x92, 0x9c, 0xd2, 0xd7, 0xc5, 0x77, 0x63, 0xe8, 0xff,
	0x6d, 0x63, 0xcc, 0xaa, 0xb1, 0xc2, 0xee, 0xa0, 0xeb, 0x1e, 0x04, 0x01, 0xf3, 0x34, 0x28, 0x37,
	0x92, 0xa1, 0x56, 0x8c, 0xee, 0xb9, 0x8a, 0x8d, 0x0a, 0x56, 0xd9, 0xaa, 0x4c, 0x57, 0x17, 0xb0,
	0x07, 0x8a, 0xe1, 0x24, 0x0f, 0xe6, 0x80, 0x1d, 0x16, 0x42, 0xa4, 0x3c, 0xe6, 0xb0, 0x91, 0x33,
	0x7f, 0x22, 0xec, 0x67, 0x3a, 0x87, 0x8d, 0xec, 0x87, 0xc8, 0xfe, 0x1d, 0x90, 0x89, 0xa1, 0x04,
	0x2e, 0x74, 0x21, 0x57, 0xb6, 0x2a, 0x53, 0x8d, 0xff, 0x7e, 0x34, 0xc6, 0x54, 0x2e, 0x6f, 0x39,
	0x73, 0x27, 0x94, 0x56, 0xc6, 0xb0, 0xef, 0xa2, 0x2b, 0x66, 0x2c, 0x2e, 0x1f, 0xba, 0x2b, 0xd5,
	0x47, 0x03, 0xae, 0x0b, 0xff, 0x96, 0xad, 0xca, 0xa4, 0x33, 0x63, 0x9e, 0xd7, 0x86, 0xe9, 0xa3,
	0xbd, 0x8e, 0xae, 0x86, 0x3b, 0x54, 0xb1, 0xa1, 0x1b, 0x4a, 0x2a, 0x5c, 0x0f, 0x84, 0x66, 0xfb,
	0xba, 0x30, 0x66, 0xaa, 0x2d, 0xe2, 0x74, 0xa0, 0xf8, 0x78, 0xa0, 0xb8, 0x01, 0x10, 0x6c, 0xd1,
	0x20, 0x62, 0xce, 0x5c, 0x2a, 0xdb, 0x94, 0x34, 0xe9, 0x3d, 0x11, 0xd9, 0x87, 0xa8, 0x78, 0xb6,
	0x56, 0x37, 0xb3, 0xbf, 0x30, 0x5e, 0xb6, 0x2a, 0xb3, 0xd5, 0x27, 0xf8, 0xdc, 0xaf, 0x83, 0xff,
	0x74, 0x11, 0xaf, 0xfe, 0xdd, 0xce, 0x56, 0x1a, 0xc3, 0x29, 0x78, 0x17, 0x20, 0x8b, 0x02, 0x15,
	0x2e, 0x52, 0xd9, 0x35, 0x74, 0xa7, 0xd9, 0xea, 0x3a, 0xad, 0xd5, 0x7a, 0xaf, 0xd5, 0x74, 0xeb,
	0xed, 0xa6, 0xdb, 0x6f, 0xd7, 0xb7, 0xea, 0x6b, 0x1b, 0xf5, 0xc6, 0x46, 0xcb, 0x6d, 0x76, 0xdc,
	0x76, 0xa7, 0xe7, 0xf6, 0x37, 0x5b, 0xf9, 0x7f, 0x8a, 0x93, 0x47, 0xdf, 0xdf, 0xbf, 0xca, 0x59,
	0x93, 0x96, 0x3d, 0x83, 0xa6, 0x9e, 0xf6, 0x7a, 0x5d, 0x77, 0x7d, 0xb3, 0xd3, 0xce, 0x5b, 0xf6,
	0x2c, 0x42, 0xe6, 0xda, 0x75, 0x3a, 0xbd, 0x4e, 0x3e, 0xf7, 0x78, 0xf9, 0xed, 0xa7, 0x97, 0xa5,
	0xdb, 0x68, 0xf1, 0xbc, 0x6e, 0xaa, 0xa7, 0xba, 0x69, 0x7c, 0xb4, 0x8e, 0xbe, 0x96, 0xac, 0x6f,
	0x6f, 0x7e, 0xbe, 0x1e, 0x5f, 0xb6, 0xef, 0xa5, 0x82, 0x64, 0xd9, 0x44, 0x52, 0x5c, 0x98, 0x8a,
	0x54, 0x88, 0xb3, 0x05, 0x8b, 0x1f, 0xd0, 0x40, 0xee, 0xd0, 0x0f, 0x2f, 0x3e, 0x7f, 0x99, 0xc8,
	0xe5, 0x73, 0x68, 0x89, 0x43, 0x6a, 0x99, 0x54, 0xb0, 0x7f, 0x70, 0xbe, 0x7b, 0x8d, 0xe9, 0x34,
	0x61, 0x37, 0x19, 0x53, 0xd7, 0xda, 0xde, 0xb8, 0xdc, 0xd6, 0xcb, 0x5d, 0xff, 0x12, 0x9b, 0x3f,
	0x98, 0x30, 0xd3, 0xaf, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x22, 0x6a, 0xc9, 0x96, 0x4c, 0x04,
	0x00, 0x00,
}

func (this *ZipkinConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ZipkinConfig)
	if !ok {
		that2, ok := that.(ZipkinConfig)
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
	if !this.CollectorUpstreamRef.Equal(that1.CollectorUpstreamRef) {
		return false
	}
	if this.CollectorEndpoint != that1.CollectorEndpoint {
		return false
	}
	if this.TraceId_128Bit != that1.TraceId_128Bit {
		return false
	}
	if !this.SharedSpanContext.Equal(that1.SharedSpanContext) {
		return false
	}
	if this.CollectorEndpointVersion != that1.CollectorEndpointVersion {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

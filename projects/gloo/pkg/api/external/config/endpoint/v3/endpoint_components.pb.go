// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/endpoint/v3/endpoint_components.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/config/core/v3"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// Upstream host identifier.
type Endpoint struct {
	// The upstream host address.
	//
	// .. attention::
	//
	//   The form of host address depends on the given cluster type. For STATIC or EDS,
	//   it is expected to be a direct IP address (or something resolvable by the
	//   specified :ref:`resolver <envoy_api_field_config.core.v3.SocketAddress.resolver_name>`
	//   in the Address). For LOGICAL or STRICT DNS, it is expected to be hostname,
	//   and will be resolved via DNS.
	Address *v3.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// The optional health check configuration is used as configuration for the
	// health checker to contact the health checked host.
	//
	// .. attention::
	//
	//   This takes into effect only for upstream clusters with
	//   :ref:`active health checking <arch_overview_health_checking>` enabled.
	HealthCheckConfig *Endpoint_HealthCheckConfig `protobuf:"bytes,2,opt,name=health_check_config,json=healthCheckConfig,proto3" json:"health_check_config,omitempty"`
	// The hostname associated with this endpoint. This hostname is not used for routing or address
	// resolution. If provided, it will be associated with the endpoint, and can be used for features
	// that require a hostname, like
	// :ref:`auto_host_rewrite <envoy_api_field_config.route.v3.RouteAction.auto_host_rewrite>`.
	Hostname             string   `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_446c756397f8034c, []int{0}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetAddress() *v3.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetHealthCheckConfig() *Endpoint_HealthCheckConfig {
	if m != nil {
		return m.HealthCheckConfig
	}
	return nil
}

func (m *Endpoint) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

// The optional health check configuration.
type Endpoint_HealthCheckConfig struct {
	// Optional alternative health check port value.
	//
	// By default the health check address port of an upstream host is the same
	// as the host's serving address port. This provides an alternative health
	// check port. Setting this with a non-zero value allows an upstream host
	// to have different health check address port.
	PortValue uint32 `protobuf:"varint,1,opt,name=port_value,json=portValue,proto3" json:"port_value,omitempty"`
	// By default, the host header for L7 health checks is controlled by cluster level configuration
	// (see: :ref:`host <envoy_api_field_config.core.v3.HealthCheck.HttpHealthCheck.host>` and
	// :ref:`authority <envoy_api_field_config.core.v3.HealthCheck.GrpcHealthCheck.authority>`). Setting this
	// to a non-empty value allows overriding the cluster level configuration for a specific
	// endpoint.
	Hostname             string   `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint_HealthCheckConfig) Reset()         { *m = Endpoint_HealthCheckConfig{} }
func (m *Endpoint_HealthCheckConfig) String() string { return proto.CompactTextString(m) }
func (*Endpoint_HealthCheckConfig) ProtoMessage()    {}
func (*Endpoint_HealthCheckConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_446c756397f8034c, []int{0, 0}
}
func (m *Endpoint_HealthCheckConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Unmarshal(m, b)
}
func (m *Endpoint_HealthCheckConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Marshal(b, m, deterministic)
}
func (m *Endpoint_HealthCheckConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint_HealthCheckConfig.Merge(m, src)
}
func (m *Endpoint_HealthCheckConfig) XXX_Size() int {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Size(m)
}
func (m *Endpoint_HealthCheckConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint_HealthCheckConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint_HealthCheckConfig proto.InternalMessageInfo

func (m *Endpoint_HealthCheckConfig) GetPortValue() uint32 {
	if m != nil {
		return m.PortValue
	}
	return 0
}

func (m *Endpoint_HealthCheckConfig) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

// An Endpoint that Envoy can route traffic to.
// [#next-free-field: 6]
type LbEndpoint struct {
	Endpoint *Endpoint `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// The optional load balancing weight of the upstream host; at least 1.
	// Envoy uses the load balancing weight in some of the built in load
	// balancers. The load balancing weight for an endpoint is divided by the sum
	// of the weights of all endpoints in the endpoint's locality to produce a
	// percentage of traffic for the endpoint. This percentage is then further
	// weighted by the endpoint's locality's load balancing weight from
	// LocalityLbEndpoints. If unspecified, each host is presumed to have equal
	// weight in a locality.
	LoadBalancingWeight  *types.UInt32Value `protobuf:"bytes,4,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LbEndpoint) Reset()         { *m = LbEndpoint{} }
func (m *LbEndpoint) String() string { return proto.CompactTextString(m) }
func (*LbEndpoint) ProtoMessage()    {}
func (*LbEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_446c756397f8034c, []int{1}
}
func (m *LbEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LbEndpoint.Unmarshal(m, b)
}
func (m *LbEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LbEndpoint.Marshal(b, m, deterministic)
}
func (m *LbEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LbEndpoint.Merge(m, src)
}
func (m *LbEndpoint) XXX_Size() int {
	return xxx_messageInfo_LbEndpoint.Size(m)
}
func (m *LbEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_LbEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_LbEndpoint proto.InternalMessageInfo

func (m *LbEndpoint) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *LbEndpoint) GetLoadBalancingWeight() *types.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

// A group of endpoints belonging to a Locality.
// One can have multiple LocalityLbEndpoints for a locality, but this is
// generally only done if the different groups need to have different load
// balancing weights or different priorities.
type LocalityLbEndpoints struct {
	// Identifies location of where the upstream hosts run.
	Locality *Locality `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	// The group of endpoints belonging to the locality specified.
	LbEndpoints []*LbEndpoint `protobuf:"bytes,2,rep,name=lb_endpoints,json=lbEndpoints,proto3" json:"lb_endpoints,omitempty"`
	// Optional: Per priority/region/zone/sub_zone weight; at least 1. The load
	// balancing weight for a locality is divided by the sum of the weights of all
	// localities  at the same priority level to produce the effective percentage
	// of traffic for the locality.
	//
	// Locality weights are only considered when :ref:`locality weighted load
	// balancing <arch_overview_load_balancing_locality_weighted_lb>` is
	// configured. These weights are ignored otherwise. If no weights are
	// specified when locality weighted load balancing is enabled, the locality is
	// assigned no load.
	LoadBalancingWeight *types.UInt32Value `protobuf:"bytes,3,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	// Optional: the priority for this LocalityLbEndpoints. If unspecified this will
	// default to the highest priority (0).
	//
	// Under usual circumstances, Envoy will only select endpoints for the highest
	// priority (0). In the event all endpoints for a particular priority are
	// unavailable/unhealthy, Envoy will fail over to selecting endpoints for the
	// next highest priority group.
	//
	// Priorities should range from 0 (highest) to N (lowest) without skipping.
	Priority uint32 `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
	// Optional: Per locality proximity value which indicates how close this
	// locality is from the source locality. This value only provides ordering
	// information (lower the value, closer it is to the source locality).
	// This will be consumed by load balancing schemes that need proximity order
	// to determine where to route the requests.
	// [#not-implemented-hide:]
	Proximity            *types.UInt32Value `protobuf:"bytes,6,opt,name=proximity,proto3" json:"proximity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LocalityLbEndpoints) Reset()         { *m = LocalityLbEndpoints{} }
func (m *LocalityLbEndpoints) String() string { return proto.CompactTextString(m) }
func (*LocalityLbEndpoints) ProtoMessage()    {}
func (*LocalityLbEndpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_446c756397f8034c, []int{2}
}
func (m *LocalityLbEndpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalityLbEndpoints.Unmarshal(m, b)
}
func (m *LocalityLbEndpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalityLbEndpoints.Marshal(b, m, deterministic)
}
func (m *LocalityLbEndpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalityLbEndpoints.Merge(m, src)
}
func (m *LocalityLbEndpoints) XXX_Size() int {
	return xxx_messageInfo_LocalityLbEndpoints.Size(m)
}
func (m *LocalityLbEndpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalityLbEndpoints.DiscardUnknown(m)
}

var xxx_messageInfo_LocalityLbEndpoints proto.InternalMessageInfo

func (m *LocalityLbEndpoints) GetLocality() *Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLbEndpoints() []*LbEndpoint {
	if m != nil {
		return m.LbEndpoints
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLoadBalancingWeight() *types.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

func (m *LocalityLbEndpoints) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *LocalityLbEndpoints) GetProximity() *types.UInt32Value {
	if m != nil {
		return m.Proximity
	}
	return nil
}

// Identifies location of where either Envoy runs or where upstream hosts run.
type Locality struct {
	// Region this :ref:`zone <envoy_api_field_config.core.v3.Locality.zone>` belongs to.
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// Defines the local service zone where Envoy is running. Though optional, it
	// should be set if discovery service routing is used and the discovery
	// service exposes :ref:`zone data <envoy_api_field_config.endpoint.v3.LocalityLbEndpoints.locality>`,
	// either in this message or via :option:`--service-zone`. The meaning of zone
	// is context dependent, e.g. `Availability Zone (AZ)
	// <https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html>`_
	// on AWS, `Zone <https://cloud.google.com/compute/docs/regions-zones/>`_ on
	// GCP, etc.
	Zone string `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	// When used for locality of upstream hosts, this field further splits zone
	// into smaller chunks of sub-zones so they can be load balanced
	// independently.
	SubZone              string   `protobuf:"bytes,3,opt,name=sub_zone,json=subZone,proto3" json:"sub_zone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Locality) Reset()         { *m = Locality{} }
func (m *Locality) String() string { return proto.CompactTextString(m) }
func (*Locality) ProtoMessage()    {}
func (*Locality) Descriptor() ([]byte, []int) {
	return fileDescriptor_446c756397f8034c, []int{3}
}
func (m *Locality) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Locality.Unmarshal(m, b)
}
func (m *Locality) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Locality.Marshal(b, m, deterministic)
}
func (m *Locality) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Locality.Merge(m, src)
}
func (m *Locality) XXX_Size() int {
	return xxx_messageInfo_Locality.Size(m)
}
func (m *Locality) XXX_DiscardUnknown() {
	xxx_messageInfo_Locality.DiscardUnknown(m)
}

var xxx_messageInfo_Locality proto.InternalMessageInfo

func (m *Locality) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Locality) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *Locality) GetSubZone() string {
	if m != nil {
		return m.SubZone
	}
	return ""
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "envoy.config.endpoint.v3.Endpoint")
	proto.RegisterType((*Endpoint_HealthCheckConfig)(nil), "envoy.config.endpoint.v3.Endpoint.HealthCheckConfig")
	proto.RegisterType((*LbEndpoint)(nil), "envoy.config.endpoint.v3.LbEndpoint")
	proto.RegisterType((*LocalityLbEndpoints)(nil), "envoy.config.endpoint.v3.LocalityLbEndpoints")
	proto.RegisterType((*Locality)(nil), "envoy.config.endpoint.v3.Locality")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/endpoint/v3/endpoint_components.proto", fileDescriptor_446c756397f8034c)
}

var fileDescriptor_446c756397f8034c = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0x56, 0xda, 0xfe, 0x5d, 0xeb, 0xfe, 0x5c, 0xcc, 0x05, 0x14, 0x2a, 0x36, 0xaa, 0x88, 0x8b,
	0xde, 0xe0, 0x48, 0x2d, 0x12, 0x12, 0x17, 0x48, 0x6c, 0x42, 0x30, 0x69, 0x42, 0x10, 0x04, 0x48,
	0xbb, 0x89, 0x9c, 0xd4, 0x4b, 0xc2, 0x5c, 0x1f, 0xcb, 0x71, 0xbb, 0x8e, 0xf7, 0xe0, 0x96, 0x6b,
	0x1e, 0x81, 0x47, 0xe0, 0x39, 0x78, 0x07, 0xee, 0x91, 0x9d, 0x38, 0xa3, 0x8c, 0xb1, 0x89, 0x3b,
	0x9f, 0x93, 0xef, 0xfb, 0xfc, 0x7d, 0xe7, 0x44, 0x46, 0x2c, 0x2b, 0x74, 0xbe, 0x4c, 0x48, 0x0a,
	0x8b, 0xb0, 0x04, 0x0e, 0x0f, 0x0a, 0x08, 0x33, 0x0e, 0x10, 0x4a, 0x05, 0x1f, 0x58, 0xaa, 0xcb,
	0xaa, 0xa2, 0xb2, 0x08, 0xd9, 0x5a, 0x33, 0x25, 0x28, 0x0f, 0x99, 0x58, 0xc1, 0x59, 0x98, 0x82,
	0x38, 0x2e, 0xb2, 0x90, 0x89, 0xb9, 0x84, 0x42, 0xe8, 0x70, 0x35, 0x6b, 0xce, 0x71, 0x0a, 0x0b,
	0x09, 0x82, 0x09, 0x5d, 0x12, 0xa9, 0x40, 0x03, 0xf6, 0x2d, 0x87, 0x54, 0x1c, 0xe2, 0x70, 0x64,
	0x35, 0x1b, 0xdd, 0xcc, 0x20, 0x03, 0x0b, 0x0a, 0xcd, 0xa9, 0xc2, 0x8f, 0x30, 0x5b, 0xeb, 0xaa,
	0xc9, 0xd6, 0xba, 0xee, 0xed, 0x66, 0x00, 0x19, 0x67, 0xa1, 0xad, 0x92, 0xe5, 0x71, 0x78, 0xaa,
	0xa8, 0x94, 0x4c, 0xd5, 0x77, 0x8c, 0x82, 0x0d, 0x5f, 0x29, 0x28, 0x66, 0x3c, 0xd1, 0xf9, 0x5c,
	0xb1, 0xd2, 0x61, 0xee, 0xfd, 0x11, 0x93, 0xd0, 0x92, 0x55, 0x80, 0xe0, 0x53, 0x0b, 0xf5, 0x9e,
	0xd5, 0xf6, 0xf0, 0x23, 0xb4, 0x55, 0xd3, 0x7d, 0x6f, 0xec, 0x4d, 0x06, 0xd3, 0x1d, 0xb2, 0x91,
	0xc3, 0xf0, 0xc9, 0x6a, 0x46, 0x9e, 0x56, 0xa0, 0xc8, 0xa1, 0xf1, 0x1c, 0x0d, 0x73, 0x46, 0xb9,
	0xce, 0xe3, 0x34, 0x67, 0xe9, 0x49, 0x5c, 0xe1, 0xfd, 0x96, 0x15, 0x79, 0x48, 0x2e, 0x1b, 0x06,
	0x71, 0x37, 0x93, 0x17, 0x96, 0xbd, 0x6f, 0xc8, 0xfb, 0x16, 0x16, 0x6d, 0xe7, 0xbf, 0xb7, 0xf0,
	0x08, 0xf5, 0x72, 0x28, 0xb5, 0xa0, 0x0b, 0xe6, 0xb7, 0xc7, 0xde, 0xa4, 0x1f, 0x35, 0xf5, 0xe8,
	0x25, 0xda, 0xbe, 0xa0, 0x81, 0x77, 0x10, 0x92, 0xa0, 0x74, 0xbc, 0xa2, 0x7c, 0xc9, 0x6c, 0xa4,
	0x1b, 0x51, 0xdf, 0x74, 0xde, 0x99, 0xc6, 0x86, 0x5e, 0x6b, 0x53, 0x2f, 0xf8, 0xec, 0x21, 0x74,
	0x98, 0x34, 0x93, 0x79, 0x82, 0x7a, 0xce, 0x77, 0x3d, 0x9a, 0xe0, 0xea, 0x54, 0x51, 0xc3, 0xc1,
	0xaf, 0xd0, 0x2d, 0x0e, 0x74, 0x1e, 0x27, 0x94, 0x53, 0x91, 0x16, 0x22, 0x8b, 0x4f, 0x59, 0x91,
	0xe5, 0xda, 0xef, 0x58, 0xb1, 0xbb, 0xa4, 0xda, 0x35, 0x71, 0xbb, 0x26, 0x6f, 0x0f, 0x84, 0x9e,
	0x4d, 0xad, 0xcf, 0x68, 0x68, 0xa8, 0x7b, 0x8e, 0xf9, 0xde, 0x12, 0x83, 0x6f, 0x2d, 0x34, 0x3c,
	0x84, 0x94, 0xf2, 0x42, 0x9f, 0x9d, 0x1b, 0x2d, 0x8d, 0x53, 0x5e, 0xb7, 0xaf, 0x76, 0xea, 0x04,
	0xa2, 0x86, 0x83, 0x9f, 0xa3, 0xff, 0x79, 0x12, 0x3b, 0x50, 0xe9, 0xb7, 0xc6, 0xed, 0xc9, 0x60,
	0x7a, 0xff, 0x2f, 0x1a, 0xcd, 0xe5, 0xd1, 0x80, 0xff, 0x62, 0xe4, 0xd2, 0xc8, 0xed, 0x7f, 0x8c,
	0x6c, 0xf6, 0x25, 0x55, 0x01, 0xca, 0x44, 0xfb, 0xcf, 0x2e, 0xb3, 0xa9, 0xf1, 0x63, 0xd4, 0x97,
	0x0a, 0xd6, 0xc5, 0xc2, 0x7c, 0xec, 0x5e, 0xe3, 0x86, 0x73, 0x78, 0xf0, 0x1a, 0xf5, 0xdc, 0x20,
	0xf0, 0x6d, 0xd4, 0x55, 0x2c, 0x2b, 0x40, 0xd8, 0xe1, 0xf5, 0xa3, 0xba, 0xc2, 0x18, 0x75, 0x3e,
	0x82, 0x70, 0xff, 0x89, 0x3d, 0xe3, 0x3b, 0xa8, 0x57, 0x2e, 0x93, 0xd8, 0xf6, 0xab, 0xff, 0x71,
	0xab, 0x5c, 0x26, 0x47, 0x20, 0xd8, 0xde, 0x9b, 0xaf, 0x3f, 0x3a, 0xde, 0x97, 0xef, 0xbb, 0xde,
	0xd1, 0xc1, 0xf5, 0x1e, 0x1c, 0x79, 0x92, 0x6d, 0x3e, 0x3a, 0x17, 0x9f, 0x9b, 0xa4, 0x6b, 0x83,
	0xcc, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xa0, 0x0c, 0x21, 0xc4, 0x04, 0x00, 0x00,
}

func (this *Endpoint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Endpoint)
	if !ok {
		that2, ok := that.(Endpoint)
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
	if !this.Address.Equal(that1.Address) {
		return false
	}
	if !this.HealthCheckConfig.Equal(that1.HealthCheckConfig) {
		return false
	}
	if this.Hostname != that1.Hostname {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Endpoint_HealthCheckConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Endpoint_HealthCheckConfig)
	if !ok {
		that2, ok := that.(Endpoint_HealthCheckConfig)
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
	if this.PortValue != that1.PortValue {
		return false
	}
	if this.Hostname != that1.Hostname {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LbEndpoint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LbEndpoint)
	if !ok {
		that2, ok := that.(LbEndpoint)
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
	if !this.Endpoint.Equal(that1.Endpoint) {
		return false
	}
	if !this.LoadBalancingWeight.Equal(that1.LoadBalancingWeight) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LocalityLbEndpoints) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LocalityLbEndpoints)
	if !ok {
		that2, ok := that.(LocalityLbEndpoints)
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
	if !this.Locality.Equal(that1.Locality) {
		return false
	}
	if len(this.LbEndpoints) != len(that1.LbEndpoints) {
		return false
	}
	for i := range this.LbEndpoints {
		if !this.LbEndpoints[i].Equal(that1.LbEndpoints[i]) {
			return false
		}
	}
	if !this.LoadBalancingWeight.Equal(that1.LoadBalancingWeight) {
		return false
	}
	if this.Priority != that1.Priority {
		return false
	}
	if !this.Proximity.Equal(that1.Proximity) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Locality) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Locality)
	if !ok {
		that2, ok := that.(Locality)
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
	if this.Region != that1.Region {
		return false
	}
	if this.Zone != that1.Zone {
		return false
	}
	if this.SubZone != that1.SubZone {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

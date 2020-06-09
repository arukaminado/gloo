// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/failover/failover.proto

package failover

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
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

type Failover struct {
	// Types that are valid to be assigned to Method:
	//	*Failover_Simple
	Method               isFailover_Method `protobuf_oneof:"method"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Failover) Reset()         { *m = Failover{} }
func (m *Failover) String() string { return proto.CompactTextString(m) }
func (*Failover) ProtoMessage()    {}
func (*Failover) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8087bcbe8320ae9, []int{0}
}
func (m *Failover) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Failover.Unmarshal(m, b)
}
func (m *Failover) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Failover.Marshal(b, m, deterministic)
}
func (m *Failover) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Failover.Merge(m, src)
}
func (m *Failover) XXX_Size() int {
	return xxx_messageInfo_Failover.Size(m)
}
func (m *Failover) XXX_DiscardUnknown() {
	xxx_messageInfo_Failover.DiscardUnknown(m)
}

var xxx_messageInfo_Failover proto.InternalMessageInfo

type isFailover_Method interface {
	isFailover_Method()
	Equal(interface{}) bool
}

type Failover_Simple struct {
	Simple *Simple `protobuf:"bytes,1,opt,name=simple,proto3,oneof" json:"simple,omitempty"`
}

func (*Failover_Simple) isFailover_Method() {}

func (m *Failover) GetMethod() isFailover_Method {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *Failover) GetSimple() *Simple {
	if x, ok := m.GetMethod().(*Failover_Simple); ok {
		return x.Simple
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Failover) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Failover_Simple)(nil),
	}
}

type Simple struct {
	PrioritizedLocalities []*Simple_PrioritizedLocality `protobuf:"bytes,1,rep,name=prioritized_localities,json=prioritizedLocalities,proto3" json:"prioritized_localities,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                      `json:"-"`
	XXX_unrecognized      []byte                        `json:"-"`
	XXX_sizecache         int32                         `json:"-"`
}

func (m *Simple) Reset()         { *m = Simple{} }
func (m *Simple) String() string { return proto.CompactTextString(m) }
func (*Simple) ProtoMessage()    {}
func (*Simple) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8087bcbe8320ae9, []int{1}
}
func (m *Simple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Simple.Unmarshal(m, b)
}
func (m *Simple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Simple.Marshal(b, m, deterministic)
}
func (m *Simple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Simple.Merge(m, src)
}
func (m *Simple) XXX_Size() int {
	return xxx_messageInfo_Simple.Size(m)
}
func (m *Simple) XXX_DiscardUnknown() {
	xxx_messageInfo_Simple.DiscardUnknown(m)
}

var xxx_messageInfo_Simple proto.InternalMessageInfo

func (m *Simple) GetPrioritizedLocalities() []*Simple_PrioritizedLocality {
	if m != nil {
		return m.PrioritizedLocalities
	}
	return nil
}

type Simple_PrioritizedLocality struct {
	Endpoints            []*LocalityLbEndpoints `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Simple_PrioritizedLocality) Reset()         { *m = Simple_PrioritizedLocality{} }
func (m *Simple_PrioritizedLocality) String() string { return proto.CompactTextString(m) }
func (*Simple_PrioritizedLocality) ProtoMessage()    {}
func (*Simple_PrioritizedLocality) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8087bcbe8320ae9, []int{1, 0}
}
func (m *Simple_PrioritizedLocality) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Simple_PrioritizedLocality.Unmarshal(m, b)
}
func (m *Simple_PrioritizedLocality) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Simple_PrioritizedLocality.Marshal(b, m, deterministic)
}
func (m *Simple_PrioritizedLocality) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Simple_PrioritizedLocality.Merge(m, src)
}
func (m *Simple_PrioritizedLocality) XXX_Size() int {
	return xxx_messageInfo_Simple_PrioritizedLocality.Size(m)
}
func (m *Simple_PrioritizedLocality) XXX_DiscardUnknown() {
	xxx_messageInfo_Simple_PrioritizedLocality.DiscardUnknown(m)
}

var xxx_messageInfo_Simple_PrioritizedLocality proto.InternalMessageInfo

func (m *Simple_PrioritizedLocality) GetEndpoints() []*LocalityLbEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

// Represents a single instance of an upstream
type Address struct {
	// Address (hostname or IP)
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	// Port the instance is listening on
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8087bcbe8320ae9, []int{2}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Address) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

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
	Address *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
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
	Hostname             string                `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	UpstreamSslConfig    *v1.UpstreamSslConfig `protobuf:"bytes,4,opt,name=upstream_ssl_config,json=upstreamSslConfig,proto3" json:"upstream_ssl_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8087bcbe8320ae9, []int{3}
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

func (m *Endpoint) GetAddress() *Address {
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

func (m *Endpoint) GetUpstreamSslConfig() *v1.UpstreamSslConfig {
	if m != nil {
		return m.UpstreamSslConfig
	}
	return nil
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
	return fileDescriptor_d8087bcbe8320ae9, []int{3, 0}
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
	return fileDescriptor_d8087bcbe8320ae9, []int{4}
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
	Locality *v3.Locality `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
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
	LoadBalancingWeight  *types.UInt32Value `protobuf:"bytes,3,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LocalityLbEndpoints) Reset()         { *m = LocalityLbEndpoints{} }
func (m *LocalityLbEndpoints) String() string { return proto.CompactTextString(m) }
func (*LocalityLbEndpoints) ProtoMessage()    {}
func (*LocalityLbEndpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8087bcbe8320ae9, []int{5}
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

func (m *LocalityLbEndpoints) GetLocality() *v3.Locality {
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

func init() {
	proto.RegisterType((*Failover)(nil), "failover.options.gloo.solo.io.Failover")
	proto.RegisterType((*Simple)(nil), "failover.options.gloo.solo.io.Simple")
	proto.RegisterType((*Simple_PrioritizedLocality)(nil), "failover.options.gloo.solo.io.Simple.PrioritizedLocality")
	proto.RegisterType((*Address)(nil), "failover.options.gloo.solo.io.Address")
	proto.RegisterType((*Endpoint)(nil), "failover.options.gloo.solo.io.Endpoint")
	proto.RegisterType((*Endpoint_HealthCheckConfig)(nil), "failover.options.gloo.solo.io.Endpoint.HealthCheckConfig")
	proto.RegisterType((*LbEndpoint)(nil), "failover.options.gloo.solo.io.LbEndpoint")
	proto.RegisterType((*LocalityLbEndpoints)(nil), "failover.options.gloo.solo.io.LocalityLbEndpoints")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/failover/failover.proto", fileDescriptor_d8087bcbe8320ae9)
}

var fileDescriptor_d8087bcbe8320ae9 = []byte{
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6e, 0x13, 0x3d,
	0x14, 0xfe, 0x27, 0xad, 0xd2, 0xd4, 0xfd, 0xbb, 0xa8, 0xf3, 0xf7, 0x57, 0x14, 0xd1, 0x8b, 0x82,
	0x80, 0xb2, 0xc0, 0x56, 0x93, 0x15, 0x6c, 0x80, 0x56, 0xa0, 0x22, 0x45, 0x50, 0x4d, 0xd5, 0x56,
	0x62, 0x33, 0xf2, 0xcc, 0xb8, 0x33, 0xa6, 0xce, 0x1c, 0xcb, 0x76, 0xd2, 0x96, 0xd7, 0x61, 0xc3,
	0x23, 0xf0, 0x3c, 0xf0, 0x04, 0x48, 0xb0, 0x47, 0xe3, 0xb9, 0x94, 0xde, 0x48, 0xc4, 0xce, 0x3e,
	0x73, 0xbe, 0xcb, 0x39, 0x3e, 0x67, 0xd0, 0x71, 0x22, 0x6c, 0x3a, 0x0e, 0x49, 0x04, 0x23, 0x6a,
	0x40, 0xc2, 0x13, 0x01, 0x34, 0x91, 0x00, 0x54, 0x69, 0xf8, 0xc0, 0x23, 0x6b, 0x8a, 0x1b, 0x53,
	0x82, 0x4e, 0xb6, 0x29, 0xcf, 0x2c, 0xd7, 0x4a, 0x0b, 0xc3, 0x29, 0x28, 0x2b, 0x20, 0x33, 0xf4,
	0x84, 0x09, 0x09, 0x13, 0xae, 0xeb, 0x03, 0x51, 0x1a, 0x2c, 0xe0, 0xb5, 0xfa, 0x5e, 0x66, 0x92,
	0x9c, 0x88, 0xe4, 0x1a, 0x44, 0x40, 0xf7, 0xbf, 0x04, 0x12, 0x70, 0x99, 0x34, 0x3f, 0x15, 0xa0,
	0x2e, 0xe6, 0xe7, 0xb6, 0x08, 0xf2, 0x73, 0x5b, 0xc6, 0xd6, 0x13, 0x80, 0x44, 0x72, 0xea, 0x6e,
	0xe1, 0xf8, 0x84, 0x9e, 0x69, 0xa6, 0x14, 0xd7, 0xa6, 0xfc, 0x7e, 0xff, 0x6e, 0xbb, 0xc6, 0xc8,
	0x32, 0x69, 0x83, 0x67, 0x13, 0xb8, 0xa0, 0x11, 0x64, 0x27, 0x22, 0xa1, 0x11, 0x68, 0x4e, 0x27,
	0x03, 0x1a, 0x32, 0xc3, 0x8b, 0x84, 0xde, 0x21, 0x6a, 0xbd, 0x2e, 0x0d, 0xe3, 0xe7, 0xa8, 0x69,
	0xc4, 0x48, 0x49, 0xde, 0xf1, 0x36, 0xbd, 0xad, 0xa5, 0xfe, 0x03, 0xf2, 0xc7, 0x5a, 0xc8, 0x81,
	0x4b, 0xde, 0xfb, 0xc7, 0x2f, 0x61, 0x3b, 0x2d, 0xd4, 0x1c, 0x71, 0x9b, 0x42, 0xdc, 0xfb, 0xe6,
	0xa1, 0x66, 0xf1, 0x19, 0x2b, 0xf4, 0xbf, 0xd2, 0x02, 0xb4, 0xb0, 0xe2, 0x23, 0x8f, 0x03, 0x09,
	0x11, 0x93, 0xc2, 0x0a, 0x6e, 0x3a, 0xde, 0xe6, 0xdc, 0xd6, 0x52, 0xff, 0xe9, 0x4c, 0x2a, 0x64,
	0xff, 0x92, 0x63, 0x58, 0x50, 0x5c, 0xf8, 0xab, 0xea, 0x46, 0x50, 0x70, 0xd3, 0x4d, 0x50, 0xfb,
	0x96, 0x6c, 0xbc, 0x8f, 0x16, 0x79, 0x16, 0x2b, 0x10, 0x99, 0x35, 0x9d, 0x86, 0xd3, 0xee, 0x4f,
	0xd1, 0xae, 0xb0, 0xc3, 0xf0, 0x55, 0x85, 0xf4, 0x2f, 0x49, 0x7a, 0xdb, 0x68, 0xe1, 0x65, 0x1c,
	0x6b, 0x6e, 0x0c, 0xc6, 0x68, 0x9e, 0xc5, 0xb1, 0x76, 0x9d, 0x5b, 0xf4, 0xdd, 0x39, 0x8f, 0x29,
	0xd0, 0xb6, 0xd3, 0xd8, 0xf4, 0xb6, 0x96, 0x7d, 0x77, 0xee, 0xfd, 0x68, 0xa0, 0x56, 0xc5, 0x85,
	0x5f, 0xa0, 0x05, 0x56, 0xe0, 0xcb, 0x8e, 0x3f, 0x9c, 0xe2, 0xa7, 0x54, 0xf3, 0x2b, 0x18, 0x16,
	0xa8, 0x9d, 0x72, 0x26, 0x6d, 0x1a, 0x44, 0x29, 0x8f, 0x4e, 0x83, 0xe2, 0xa1, 0x9d, 0xe2, 0xf4,
	0xce, 0x56, 0x3e, 0xc8, 0x9e, 0xa3, 0xd8, 0xcd, 0x19, 0x76, 0x1d, 0x81, 0xbf, 0x92, 0x5e, 0x0f,
	0xe1, 0x2e, 0x6a, 0xa5, 0x60, 0x6c, 0xc6, 0x46, 0xbc, 0x33, 0xe7, 0xaa, 0xac, 0xef, 0xf8, 0x1d,
	0x6a, 0x8f, 0x95, 0xb1, 0x9a, 0xb3, 0x51, 0x60, 0x8c, 0xac, 0x6c, 0xcc, 0x3b, 0x1b, 0x1b, 0x57,
	0x55, 0x0f, 0xcb, 0xc4, 0x03, 0x23, 0x2b, 0xb1, 0xf1, 0xf5, 0x50, 0xf7, 0x2d, 0x5a, 0xb9, 0x61,
	0x0a, 0xaf, 0x21, 0x94, 0xf7, 0x30, 0x98, 0x30, 0x39, 0x2e, 0x66, 0x74, 0xd9, 0x5f, 0xcc, 0x23,
	0x47, 0x79, 0xe0, 0x8a, 0xc1, 0xc6, 0x55, 0x83, 0xbd, 0x4f, 0x1e, 0x42, 0x97, 0x8f, 0x88, 0x77,
	0x51, 0xab, 0x7a, 0xc5, 0xb2, 0xf3, 0x8f, 0x66, 0xec, 0x95, 0x5f, 0x03, 0xf1, 0x3e, 0x5a, 0x95,
	0xc0, 0xe2, 0x20, 0x64, 0x92, 0x65, 0x91, 0xc8, 0x92, 0xe0, 0x8c, 0x8b, 0x24, 0xb5, 0x65, 0xd9,
	0xf7, 0x48, 0xb1, 0xc0, 0xa4, 0x5a, 0x60, 0x72, 0xf8, 0x26, 0xb3, 0x83, 0xbe, 0x33, 0xeb, 0xb7,
	0x73, 0xe8, 0x4e, 0x85, 0x3c, 0x76, 0xc0, 0xde, 0x77, 0x0f, 0xb5, 0x6f, 0x19, 0x39, 0xfc, 0x0c,
	0xb5, 0xca, 0xb5, 0xb9, 0x28, 0xed, 0xae, 0x13, 0xb7, 0xd8, 0xa4, 0x68, 0x34, 0xc9, 0x17, 0x9b,
	0x4c, 0x06, 0xf5, 0xbc, 0xfa, 0x75, 0x3e, 0x1e, 0xa2, 0x7f, 0x65, 0x18, 0x5c, 0x1f, 0xfc, 0xc7,
	0xd3, 0x06, 0xbf, 0x56, 0xf7, 0x97, 0xe4, 0x6f, 0x4e, 0xee, 0xac, 0x79, 0xee, 0x2f, 0x6b, 0xde,
	0x39, 0xfa, 0xf2, 0x73, 0xde, 0xfb, 0xfc, 0x75, 0xdd, 0x7b, 0x3f, 0x9c, 0xed, 0x97, 0xac, 0x4e,
	0x93, 0x19, 0x7e, 0xcb, 0x61, 0xd3, 0x59, 0x18, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x61, 0x9e,
	0xb2, 0xe3, 0xe9, 0x05, 0x00, 0x00,
}

func (this *Failover) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Failover)
	if !ok {
		that2, ok := that.(Failover)
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
	if that1.Method == nil {
		if this.Method != nil {
			return false
		}
	} else if this.Method == nil {
		return false
	} else if !this.Method.Equal(that1.Method) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Failover_Simple) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Failover_Simple)
	if !ok {
		that2, ok := that.(Failover_Simple)
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
	if !this.Simple.Equal(that1.Simple) {
		return false
	}
	return true
}
func (this *Simple) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Simple)
	if !ok {
		that2, ok := that.(Simple)
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
	if len(this.PrioritizedLocalities) != len(that1.PrioritizedLocalities) {
		return false
	}
	for i := range this.PrioritizedLocalities {
		if !this.PrioritizedLocalities[i].Equal(that1.PrioritizedLocalities[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Simple_PrioritizedLocality) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Simple_PrioritizedLocality)
	if !ok {
		that2, ok := that.(Simple_PrioritizedLocality)
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
	if len(this.Endpoints) != len(that1.Endpoints) {
		return false
	}
	for i := range this.Endpoints {
		if !this.Endpoints[i].Equal(that1.Endpoints[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Address) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Address)
	if !ok {
		that2, ok := that.(Address)
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
	if this.Addr != that1.Addr {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
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
	if !this.UpstreamSslConfig.Equal(that1.UpstreamSslConfig) {
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

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
	core "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core"
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

//
//Failover configuration for an upstream.
//
//Failover allows for optional fallback endpoints in the case that the primary set of endpoints is deemed
//unhealthy. As failover requires knowledge of the health of each set of endpoints, active or passive
//health checks must be configured on an upstream using failover in order for it to work properly.
type Failover struct {
	// The method by which to configure failover for the upstream
	//
	// Types that are valid to be assigned to Method:
	//	*Failover_Explicit
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

type Failover_Explicit struct {
	Explicit *ExplicitFailover `protobuf:"bytes,1,opt,name=explicit,proto3,oneof" json:"explicit,omitempty"`
}

func (*Failover_Explicit) isFailover_Method() {}

func (m *Failover) GetMethod() isFailover_Method {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *Failover) GetExplicit() *ExplicitFailover {
	if x, ok := m.GetMethod().(*Failover_Explicit); ok {
		return x.Explicit
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Failover) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Failover_Explicit)(nil),
	}
}

//
//ExplicitFailover closely resembles the Envoy config which this is translated to, with one notable exception.
//The priorities are not defined on the `LocalityLbEndpoints` but rather inferred from the list of
//`PrioritizedLocality`. More information on envoy prioritization can be found
//[here](https://www.envoyproxy.io/docs/envoy/v1.14.1/intro/arch_overview/upstream/load_balancing/priority#arch-overview-load-balancing-priority-levels).
//In practice this means that the priority of a given set of `LocalityLbEndpoints` is determined by it's index in
//the list, first being `0` through `n-1`.
//
type ExplicitFailover struct {
	// Identifies location of where the parent upstream hosts run.
	Locality *v3.Locality `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	// PrioritizedLocality is an implicitly prioritized list of lists of `LocalityLbEndpoints`. The priority of each
	// list of `LocalityLbEndpoints` is determined by it's index in the list.
	PrioritizedLocalities []*ExplicitFailover_PrioritizedLocality `protobuf:"bytes,2,rep,name=prioritized_localities,json=prioritizedLocalities,proto3" json:"prioritized_localities,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                                `json:"-"`
	XXX_unrecognized      []byte                                  `json:"-"`
	XXX_sizecache         int32                                   `json:"-"`
}

func (m *ExplicitFailover) Reset()         { *m = ExplicitFailover{} }
func (m *ExplicitFailover) String() string { return proto.CompactTextString(m) }
func (*ExplicitFailover) ProtoMessage()    {}
func (*ExplicitFailover) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8087bcbe8320ae9, []int{1}
}
func (m *ExplicitFailover) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExplicitFailover.Unmarshal(m, b)
}
func (m *ExplicitFailover) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExplicitFailover.Marshal(b, m, deterministic)
}
func (m *ExplicitFailover) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExplicitFailover.Merge(m, src)
}
func (m *ExplicitFailover) XXX_Size() int {
	return xxx_messageInfo_ExplicitFailover.Size(m)
}
func (m *ExplicitFailover) XXX_DiscardUnknown() {
	xxx_messageInfo_ExplicitFailover.DiscardUnknown(m)
}

var xxx_messageInfo_ExplicitFailover proto.InternalMessageInfo

func (m *ExplicitFailover) GetLocality() *v3.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *ExplicitFailover) GetPrioritizedLocalities() []*ExplicitFailover_PrioritizedLocality {
	if m != nil {
		return m.PrioritizedLocalities
	}
	return nil
}

type ExplicitFailover_PrioritizedLocality struct {
	Endpoints            []*LocalityLbEndpoints `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ExplicitFailover_PrioritizedLocality) Reset()         { *m = ExplicitFailover_PrioritizedLocality{} }
func (m *ExplicitFailover_PrioritizedLocality) String() string { return proto.CompactTextString(m) }
func (*ExplicitFailover_PrioritizedLocality) ProtoMessage()    {}
func (*ExplicitFailover_PrioritizedLocality) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8087bcbe8320ae9, []int{1, 0}
}
func (m *ExplicitFailover_PrioritizedLocality) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExplicitFailover_PrioritizedLocality.Unmarshal(m, b)
}
func (m *ExplicitFailover_PrioritizedLocality) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExplicitFailover_PrioritizedLocality.Marshal(b, m, deterministic)
}
func (m *ExplicitFailover_PrioritizedLocality) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExplicitFailover_PrioritizedLocality.Merge(m, src)
}
func (m *ExplicitFailover_PrioritizedLocality) XXX_Size() int {
	return xxx_messageInfo_ExplicitFailover_PrioritizedLocality.Size(m)
}
func (m *ExplicitFailover_PrioritizedLocality) XXX_DiscardUnknown() {
	xxx_messageInfo_ExplicitFailover_PrioritizedLocality.DiscardUnknown(m)
}

var xxx_messageInfo_ExplicitFailover_PrioritizedLocality proto.InternalMessageInfo

func (m *ExplicitFailover_PrioritizedLocality) GetEndpoints() []*LocalityLbEndpoints {
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
	Hostname             string                  `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	UpstreamSslConfig    *core.UpstreamSslConfig `protobuf:"bytes,4,opt,name=upstream_ssl_config,json=upstreamSslConfig,proto3" json:"upstream_ssl_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
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

func (m *Endpoint) GetUpstreamSslConfig() *core.UpstreamSslConfig {
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
	proto.RegisterType((*ExplicitFailover)(nil), "failover.options.gloo.solo.io.ExplicitFailover")
	proto.RegisterType((*ExplicitFailover_PrioritizedLocality)(nil), "failover.options.gloo.solo.io.ExplicitFailover.PrioritizedLocality")
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
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x69, 0x37, 0x6d, 0x99, 0xcb, 0x24, 0xe6, 0x32, 0x54, 0x55, 0x6c, 0x4c, 0x15, 0x82,
	0x72, 0xc0, 0xd6, 0xda, 0x13, 0x9c, 0x60, 0xd5, 0xd0, 0x90, 0x0a, 0x9a, 0x02, 0xdb, 0x24, 0x2e,
	0x91, 0x93, 0x78, 0x89, 0x99, 0x1b, 0x5b, 0xb6, 0xdb, 0xfd, 0xf8, 0x77, 0xb8, 0x70, 0xe7, 0xc2,
	0xdf, 0xc3, 0x7f, 0xc0, 0x81, 0x03, 0x37, 0x14, 0xc7, 0xc9, 0x58, 0xb7, 0xd1, 0x01, 0x37, 0xfb,
	0xe5, 0xfb, 0xbe, 0xef, 0xf3, 0x9e, 0x1d, 0x83, 0x83, 0x84, 0x99, 0x74, 0x1c, 0xa2, 0x48, 0x8c,
	0xb0, 0x16, 0x5c, 0x3c, 0x65, 0x02, 0x27, 0x5c, 0x08, 0x2c, 0x95, 0xf8, 0x48, 0x23, 0xa3, 0x8b,
	0x1d, 0x91, 0x0c, 0x4f, 0x36, 0x31, 0xcd, 0x0c, 0x55, 0x52, 0x31, 0x4d, 0xb1, 0x90, 0x86, 0x89,
	0x4c, 0xe3, 0x43, 0xc2, 0xb8, 0x98, 0x50, 0x55, 0x2d, 0x90, 0x54, 0xc2, 0x08, 0xb8, 0x56, 0xed,
	0x9d, 0x12, 0xe5, 0x46, 0x28, 0xaf, 0x81, 0x98, 0x68, 0xdf, 0x4d, 0x44, 0x22, 0xac, 0x12, 0xe7,
	0xab, 0x22, 0xa9, 0x0d, 0xe9, 0x89, 0x29, 0x82, 0xf4, 0xc4, 0xb8, 0xd8, 0x7a, 0x22, 0x44, 0xc2,
	0x29, 0xb6, 0xbb, 0x70, 0x7c, 0x88, 0x8f, 0x15, 0x91, 0x92, 0x2a, 0xed, 0xbe, 0x77, 0xaf, 0xc7,
	0x8d, 0x84, 0xa2, 0x58, 0x6b, 0xee, 0x94, 0x0f, 0x68, 0x36, 0x11, 0xa7, 0x38, 0x12, 0xd9, 0x21,
	0x4b, 0x8a, 0x8f, 0x93, 0x3e, 0x0e, 0x89, 0xa6, 0x85, 0xa0, 0x13, 0x01, 0xef, 0x95, 0xa3, 0x86,
	0x6f, 0x80, 0x47, 0x4f, 0x24, 0x67, 0x11, 0x33, 0xad, 0xda, 0x46, 0xad, 0xdb, 0xe8, 0x61, 0xf4,
	0xc7, 0x96, 0xd0, 0xb6, 0x93, 0x97, 0x16, 0x3b, 0xb7, 0xfc, 0xca, 0x62, 0xcb, 0x03, 0x0b, 0x23,
	0x6a, 0x52, 0x11, 0x77, 0xbe, 0xd4, 0xc1, 0x9d, 0x69, 0x29, 0x7c, 0x0e, 0x3c, 0x2e, 0x22, 0xc2,
	0x99, 0x39, 0x75, 0xd5, 0xd6, 0x91, 0xa5, 0x45, 0x05, 0x2d, 0xca, 0x69, 0xd1, 0xa4, 0x8f, 0x86,
	0x4e, 0xe5, 0x57, 0x7a, 0x78, 0x06, 0xee, 0x49, 0xc5, 0x84, 0x62, 0x86, 0x9d, 0xd1, 0x38, 0x70,
	0x71, 0x46, 0x75, 0xab, 0xbe, 0x31, 0xd7, 0x6d, 0xf4, 0x06, 0x7f, 0xc9, 0x8d, 0x76, 0xcf, 0xdd,
	0xaa, 0x72, 0xab, 0xf2, 0x52, 0x90, 0x51, 0xdd, 0x4e, 0x40, 0xf3, 0x0a, 0x35, 0xdc, 0x05, 0x4b,
	0x34, 0x8b, 0xa5, 0x60, 0x99, 0x29, 0x29, 0x7a, 0x33, 0x28, 0xca, 0xdc, 0x61, 0xb8, 0x5d, 0x66,
	0xfa, 0xe7, 0x26, 0x9d, 0x4d, 0xb0, 0xf8, 0x32, 0x8e, 0x15, 0xd5, 0x1a, 0x42, 0x30, 0x4f, 0xe2,
	0x58, 0xd9, 0x39, 0x2d, 0xf9, 0x76, 0x9d, 0xc7, 0xa4, 0x50, 0xa6, 0x55, 0xdf, 0xa8, 0x75, 0x97,
	0x7d, 0xbb, 0xee, 0xfc, 0xac, 0x03, 0xaf, 0xf4, 0x82, 0x2f, 0xc0, 0x22, 0x29, 0xf2, 0xdd, 0x7c,
	0x1f, 0xcd, 0xe0, 0x71, 0xd5, 0xfc, 0x32, 0x0d, 0x32, 0xd0, 0x4c, 0x29, 0xe1, 0x26, 0x0d, 0xa2,
	0x94, 0x46, 0x47, 0x41, 0x71, 0x30, 0xb6, 0x62, 0xa3, 0xf7, 0x6c, 0xd6, 0x8c, 0x1d, 0x07, 0xda,
	0xb1, 0x16, 0x83, 0xdc, 0x61, 0x60, 0x0d, 0xfc, 0x95, 0x74, 0x3a, 0x04, 0xdb, 0xc0, 0x4b, 0x85,
	0x36, 0x19, 0x19, 0xd1, 0xd6, 0x9c, 0xed, 0xb2, 0xda, 0xc3, 0xf7, 0xa0, 0x39, 0x96, 0xda, 0x28,
	0x4a, 0x46, 0x81, 0xd6, 0xbc, 0xc4, 0x98, 0xb7, 0x18, 0x0f, 0x8b, 0x7b, 0x72, 0xa1, 0xf4, 0x9e,
	0x53, 0xbf, 0xd3, 0xbc, 0xac, 0x38, 0x9e, 0x0e, 0xb5, 0xdf, 0x82, 0x95, 0x4b, 0x64, 0x70, 0x0d,
	0x80, 0x7c, 0x90, 0xc1, 0x84, 0xf0, 0x31, 0xb5, 0x63, 0x5b, 0xf6, 0x97, 0xf2, 0xc8, 0x7e, 0x1e,
	0xb8, 0x40, 0x59, 0xbf, 0x48, 0xd9, 0xf9, 0x54, 0x03, 0xe0, 0xfc, 0x24, 0xe1, 0x00, 0x78, 0xe5,
	0x51, 0xba, 0xf1, 0x3f, 0xbe, 0xe1, 0xc0, 0xfc, 0x2a, 0x11, 0xee, 0x82, 0x55, 0x2e, 0x48, 0x1c,
	0x84, 0x84, 0x93, 0x2c, 0x62, 0x59, 0x12, 0x1c, 0x53, 0x96, 0xa4, 0xc6, 0xf5, 0x7e, 0x1f, 0x15,
	0x0f, 0x05, 0x2a, 0x1f, 0x0a, 0xb4, 0xf7, 0x3a, 0x33, 0xfd, 0x9e, 0x85, 0xf5, 0x9b, 0x79, 0xea,
	0x56, 0x99, 0x79, 0x60, 0x13, 0x3b, 0xdf, 0x6b, 0xa0, 0x79, 0xc5, 0xbd, 0xfb, 0xaf, 0xbf, 0x71,
	0x08, 0x6e, 0xf3, 0x30, 0x98, 0xbe, 0xfd, 0x4f, 0x66, 0xdd, 0xfe, 0xaa, 0xba, 0xdf, 0xe0, 0xbf,
	0x91, 0x5c, 0xdb, 0xf3, 0xdc, 0x3f, 0xf6, 0xbc, 0xb5, 0xff, 0xf5, 0xc7, 0x7c, 0xed, 0xf3, 0xb7,
	0xf5, 0xda, 0x87, 0xe1, 0xcd, 0x9e, 0x7e, 0x79, 0x94, 0xdc, 0xe0, 0xf9, 0x0f, 0x17, 0x2c, 0x42,
	0xff, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x89, 0x9e, 0xa7, 0x51, 0x06, 0x00, 0x00,
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
func (this *Failover_Explicit) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Failover_Explicit)
	if !ok {
		that2, ok := that.(Failover_Explicit)
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
	if !this.Explicit.Equal(that1.Explicit) {
		return false
	}
	return true
}
func (this *ExplicitFailover) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ExplicitFailover)
	if !ok {
		that2, ok := that.(ExplicitFailover)
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
func (this *ExplicitFailover_PrioritizedLocality) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ExplicitFailover_PrioritizedLocality)
	if !ok {
		that2, ok := that.(ExplicitFailover_PrioritizedLocality)
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

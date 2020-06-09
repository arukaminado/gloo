// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/rest/rest.proto

package rest

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	transformation "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation"
	transformation1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation"
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

type ServiceSpec struct {
	Transformations      map[string]*transformation.TransformationTemplate `protobuf:"bytes,1,rep,name=transformations,proto3" json:"transformations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SwaggerInfo          *ServiceSpec_SwaggerInfo                          `protobuf:"bytes,2,opt,name=swagger_info,json=swaggerInfo,proto3" json:"swagger_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                          `json:"-"`
	XXX_unrecognized     []byte                                            `json:"-"`
	XXX_sizecache        int32                                             `json:"-"`
}

func (m *ServiceSpec) Reset()         { *m = ServiceSpec{} }
func (m *ServiceSpec) String() string { return proto.CompactTextString(m) }
func (*ServiceSpec) ProtoMessage()    {}
func (*ServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_1abca825b1c7f5e5, []int{0}
}
func (m *ServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSpec.Unmarshal(m, b)
}
func (m *ServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSpec.Marshal(b, m, deterministic)
}
func (m *ServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSpec.Merge(m, src)
}
func (m *ServiceSpec) XXX_Size() int {
	return xxx_messageInfo_ServiceSpec.Size(m)
}
func (m *ServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSpec proto.InternalMessageInfo

func (m *ServiceSpec) GetTransformations() map[string]*transformation.TransformationTemplate {
	if m != nil {
		return m.Transformations
	}
	return nil
}

func (m *ServiceSpec) GetSwaggerInfo() *ServiceSpec_SwaggerInfo {
	if m != nil {
		return m.SwaggerInfo
	}
	return nil
}

type ServiceSpec_SwaggerInfo struct {
	// Types that are valid to be assigned to SwaggerSpec:
	//	*ServiceSpec_SwaggerInfo_Url
	//	*ServiceSpec_SwaggerInfo_Inline
	SwaggerSpec          isServiceSpec_SwaggerInfo_SwaggerSpec `protobuf_oneof:"swagger_spec"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ServiceSpec_SwaggerInfo) Reset()         { *m = ServiceSpec_SwaggerInfo{} }
func (m *ServiceSpec_SwaggerInfo) String() string { return proto.CompactTextString(m) }
func (*ServiceSpec_SwaggerInfo) ProtoMessage()    {}
func (*ServiceSpec_SwaggerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1abca825b1c7f5e5, []int{0, 1}
}
func (m *ServiceSpec_SwaggerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSpec_SwaggerInfo.Unmarshal(m, b)
}
func (m *ServiceSpec_SwaggerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSpec_SwaggerInfo.Marshal(b, m, deterministic)
}
func (m *ServiceSpec_SwaggerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSpec_SwaggerInfo.Merge(m, src)
}
func (m *ServiceSpec_SwaggerInfo) XXX_Size() int {
	return xxx_messageInfo_ServiceSpec_SwaggerInfo.Size(m)
}
func (m *ServiceSpec_SwaggerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSpec_SwaggerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSpec_SwaggerInfo proto.InternalMessageInfo

type isServiceSpec_SwaggerInfo_SwaggerSpec interface {
	isServiceSpec_SwaggerInfo_SwaggerSpec()
	Equal(interface{}) bool
}

type ServiceSpec_SwaggerInfo_Url struct {
	Url string `protobuf:"bytes,1,opt,name=url,proto3,oneof" json:"url,omitempty"`
}
type ServiceSpec_SwaggerInfo_Inline struct {
	Inline string `protobuf:"bytes,2,opt,name=inline,proto3,oneof" json:"inline,omitempty"`
}

func (*ServiceSpec_SwaggerInfo_Url) isServiceSpec_SwaggerInfo_SwaggerSpec()    {}
func (*ServiceSpec_SwaggerInfo_Inline) isServiceSpec_SwaggerInfo_SwaggerSpec() {}

func (m *ServiceSpec_SwaggerInfo) GetSwaggerSpec() isServiceSpec_SwaggerInfo_SwaggerSpec {
	if m != nil {
		return m.SwaggerSpec
	}
	return nil
}

func (m *ServiceSpec_SwaggerInfo) GetUrl() string {
	if x, ok := m.GetSwaggerSpec().(*ServiceSpec_SwaggerInfo_Url); ok {
		return x.Url
	}
	return ""
}

func (m *ServiceSpec_SwaggerInfo) GetInline() string {
	if x, ok := m.GetSwaggerSpec().(*ServiceSpec_SwaggerInfo_Inline); ok {
		return x.Inline
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ServiceSpec_SwaggerInfo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ServiceSpec_SwaggerInfo_Url)(nil),
		(*ServiceSpec_SwaggerInfo_Inline)(nil),
	}
}

// This is only for upstream with REST service spec
type DestinationSpec struct {
	FunctionName           string                                 `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	Parameters             *transformation1.Parameters            `protobuf:"bytes,2,opt,name=parameters,proto3" json:"parameters,omitempty"`
	ResponseTransformation *transformation.TransformationTemplate `protobuf:"bytes,3,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                               `json:"-"`
	XXX_unrecognized       []byte                                 `json:"-"`
	XXX_sizecache          int32                                  `json:"-"`
}

func (m *DestinationSpec) Reset()         { *m = DestinationSpec{} }
func (m *DestinationSpec) String() string { return proto.CompactTextString(m) }
func (*DestinationSpec) ProtoMessage()    {}
func (*DestinationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_1abca825b1c7f5e5, []int{1}
}
func (m *DestinationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestinationSpec.Unmarshal(m, b)
}
func (m *DestinationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestinationSpec.Marshal(b, m, deterministic)
}
func (m *DestinationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestinationSpec.Merge(m, src)
}
func (m *DestinationSpec) XXX_Size() int {
	return xxx_messageInfo_DestinationSpec.Size(m)
}
func (m *DestinationSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DestinationSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DestinationSpec proto.InternalMessageInfo

func (m *DestinationSpec) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *DestinationSpec) GetParameters() *transformation1.Parameters {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *DestinationSpec) GetResponseTransformation() *transformation.TransformationTemplate {
	if m != nil {
		return m.ResponseTransformation
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceSpec)(nil), "rest.options.gloo.solo.io.ServiceSpec")
	proto.RegisterMapType((map[string]*transformation.TransformationTemplate)(nil), "rest.options.gloo.solo.io.ServiceSpec.TransformationsEntry")
	proto.RegisterType((*ServiceSpec_SwaggerInfo)(nil), "rest.options.gloo.solo.io.ServiceSpec.SwaggerInfo")
	proto.RegisterType((*DestinationSpec)(nil), "rest.options.gloo.solo.io.DestinationSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/rest/rest.proto", fileDescriptor_1abca825b1c7f5e5)
}

var fileDescriptor_1abca825b1c7f5e5 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0x26, 0x0d, 0x54, 0xc2, 0x29, 0x14, 0x59, 0x15, 0x84, 0x1c, 0xd0, 0xaa, 0x5c, 0xf6, 0x82,
	0x0d, 0xcb, 0x05, 0x81, 0x38, 0x50, 0x95, 0x3f, 0x21, 0x01, 0xca, 0x2e, 0x17, 0x2e, 0x2b, 0x37,
	0x9a, 0xa4, 0xa6, 0x89, 0xc7, 0xb2, 0xbd, 0x61, 0xf7, 0x25, 0x78, 0x0e, 0x1e, 0x81, 0xe7, 0xe1,
	0x11, 0x90, 0xb8, 0xa3, 0x38, 0x29, 0xbb, 0x89, 0xb6, 0xa2, 0xe2, 0x12, 0xcd, 0x37, 0xce, 0xf7,
	0x7d, 0x9e, 0x19, 0x0f, 0x39, 0x2e, 0xa4, 0x3b, 0x5d, 0x9c, 0xb0, 0x0c, 0x2b, 0x6e, 0xb1, 0xc4,
	0x07, 0x12, 0x79, 0x51, 0x22, 0x72, 0x6d, 0xf0, 0x0b, 0x64, 0xce, 0xb6, 0x48, 0x68, 0xc9, 0xeb,
	0x47, 0x1c, 0xb5, 0x93, 0xa8, 0x2c, 0x37, 0x60, 0x9d, 0xff, 0x30, 0x6d, 0xd0, 0x21, 0xbd, 0xeb,
	0xe3, 0xee, 0x94, 0x35, 0x0c, 0xd6, 0x88, 0x31, 0x89, 0xc9, 0x41, 0x81, 0x05, 0xfa, 0xbf, 0x78,
	0x13, 0xb5, 0x84, 0x84, 0xc2, 0xd2, 0xb5, 0x49, 0x58, 0x76, 0x22, 0xc9, 0xec, 0x02, 0x5f, 0x58,
	0x3a, 0x30, 0x4a, 0x94, 0x1c, 0x54, 0x8d, 0x2b, 0x0f, 0x95, 0xf5, 0xd7, 0x70, 0x46, 0x28, 0x9b,
	0xa3, 0xa9, 0x44, 0x63, 0x3c, 0x80, 0x9d, 0xea, 0x8b, 0x7f, 0x57, 0x33, 0x90, 0xd1, 0xc2, 0x88,
	0x0a, 0x1c, 0x18, 0xdb, 0x4a, 0x1c, 0x7e, 0x0b, 0x49, 0x34, 0x05, 0x53, 0xcb, 0x0c, 0xa6, 0x1a,
	0x32, 0x0a, 0x64, 0xbf, 0x4f, 0xb1, 0x71, 0x30, 0x0a, 0xc7, 0xd1, 0xe4, 0x19, 0xbb, 0xb0, 0x0f,
	0x6c, 0x43, 0x80, 0xcd, 0xfa, 0xec, 0x97, 0xca, 0x99, 0x55, 0x3a, 0xd4, 0xa4, 0x9f, 0xc8, 0x9e,
	0xfd, 0x2a, 0x8a, 0x02, 0xcc, 0x5c, 0xaa, 0x1c, 0xe3, 0x9d, 0x51, 0x30, 0x8e, 0x26, 0x93, 0x4b,
	0x7a, 0x4c, 0x5b, 0xea, 0x5b, 0x95, 0x63, 0x1a, 0xd9, 0x35, 0x48, 0x1c, 0x39, 0xd8, 0xe6, 0x4f,
	0x6f, 0x91, 0xf0, 0x0c, 0x56, 0x71, 0x30, 0x0a, 0xc6, 0xd7, 0xd3, 0x26, 0xa4, 0xaf, 0xc8, 0xb5,
	0x5a, 0x94, 0x0b, 0xe8, 0x9c, 0x1f, 0x32, 0x3f, 0x00, 0x26, 0xb4, 0x64, 0xf5, 0x84, 0xe5, 0xb2,
	0x74, 0x60, 0xd8, 0xa9, 0x73, 0x7a, 0x50, 0xd0, 0x0c, 0x2a, 0x5d, 0x0a, 0x07, 0x69, 0x4b, 0x7f,
	0xba, 0xf3, 0x24, 0x48, 0xde, 0x91, 0x68, 0xe3, 0x46, 0x94, 0x92, 0x70, 0x61, 0xca, 0xd6, 0xec,
	0xcd, 0x95, 0xb4, 0x01, 0x34, 0x26, 0xbb, 0x52, 0x95, 0x52, 0xb5, 0x7e, 0x4d, 0xba, 0xc3, 0x47,
	0x37, 0xd7, 0x9d, 0xb0, 0x1a, 0xb2, 0xc3, 0x5f, 0x01, 0xd9, 0x3f, 0x06, 0xeb, 0xa4, 0xf2, 0x7e,
	0x7e, 0x28, 0xf7, 0xc9, 0x8d, 0x7c, 0xa1, 0xb2, 0x06, 0xcf, 0x95, 0xa8, 0xa0, 0x2b, 0x64, 0xef,
	0x3c, 0xf9, 0x5e, 0x54, 0x40, 0x3f, 0x10, 0xb2, 0x9e, 0x6e, 0x57, 0x16, 0x67, 0x83, 0x77, 0xb3,
	0xb5, 0xb5, 0x1f, 0xff, 0xd2, 0xd2, 0x0d, 0x09, 0x2a, 0xc9, 0x1d, 0x03, 0x56, 0xa3, 0xb2, 0x30,
	0xef, 0xcb, 0xc4, 0xe1, 0x7f, 0x36, 0xed, 0xf6, 0xb9, 0x60, 0xff, 0xfc, 0xe8, 0xf5, 0x8f, 0xdf,
	0x57, 0x83, 0xef, 0x3f, 0xef, 0x05, 0x9f, 0x9f, 0x5f, 0x6e, 0x67, 0xf5, 0x59, 0xb1, 0x6d, 0x6f,
	0x4f, 0x76, 0xfd, 0xab, 0x7e, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xc1, 0xb6, 0xf9, 0xfb,
	0x03, 0x00, 0x00,
}

func (this *ServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec)
	if !ok {
		that2, ok := that.(ServiceSpec)
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
	if len(this.Transformations) != len(that1.Transformations) {
		return false
	}
	for i := range this.Transformations {
		if !this.Transformations[i].Equal(that1.Transformations[i]) {
			return false
		}
	}
	if !this.SwaggerInfo.Equal(that1.SwaggerInfo) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ServiceSpec_SwaggerInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_SwaggerInfo)
	if !ok {
		that2, ok := that.(ServiceSpec_SwaggerInfo)
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
	if that1.SwaggerSpec == nil {
		if this.SwaggerSpec != nil {
			return false
		}
	} else if this.SwaggerSpec == nil {
		return false
	} else if !this.SwaggerSpec.Equal(that1.SwaggerSpec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ServiceSpec_SwaggerInfo_Url) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_SwaggerInfo_Url)
	if !ok {
		that2, ok := that.(ServiceSpec_SwaggerInfo_Url)
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
	if this.Url != that1.Url {
		return false
	}
	return true
}
func (this *ServiceSpec_SwaggerInfo_Inline) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_SwaggerInfo_Inline)
	if !ok {
		that2, ok := that.(ServiceSpec_SwaggerInfo_Inline)
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
	if this.Inline != that1.Inline {
		return false
	}
	return true
}
func (this *DestinationSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec)
	if !ok {
		that2, ok := that.(DestinationSpec)
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
	if this.FunctionName != that1.FunctionName {
		return false
	}
	if !this.Parameters.Equal(that1.Parameters) {
		return false
	}
	if !this.ResponseTransformation.Equal(that1.ResponseTransformation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

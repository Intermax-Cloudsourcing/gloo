// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/aws/aws.proto

package aws

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

type DestinationSpec_InvocationStyle int32

const (
	DestinationSpec_SYNC  DestinationSpec_InvocationStyle = 0
	DestinationSpec_ASYNC DestinationSpec_InvocationStyle = 1
)

var DestinationSpec_InvocationStyle_name = map[int32]string{
	0: "SYNC",
	1: "ASYNC",
}

var DestinationSpec_InvocationStyle_value = map[string]int32{
	"SYNC":  0,
	"ASYNC": 1,
}

func (x DestinationSpec_InvocationStyle) String() string {
	return proto.EnumName(DestinationSpec_InvocationStyle_name, int32(x))
}

func (DestinationSpec_InvocationStyle) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7e8a9525eed72921, []int{2, 0}
}

// Upstream Spec for AWS Lambda Upstreams
// AWS Upstreams represent a collection of Lambda Functions for a particular AWS Account (IAM Role or User account)
// in a particular region
type UpstreamSpec struct {
	// The AWS Region where the desired Lambda Functions exist
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// A [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an AWS Secret
	// AWS Secrets can be created with `glooctl secret create aws ...`
	// If the secret is created manually, it must conform to the following structure:
	//  ```
	//  access_key: <aws access key>
	//  secret_key: <aws secret key>
	//  session_token: <(optional) aws session token>
	//  ```
	SecretRef *core.ResourceRef `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
	// The list of Lambda Functions contained within this region.
	// This list will be automatically populated by Gloo if discovery is enabled for AWS Lambda Functions
	LambdaFunctions      []*LambdaFunctionSpec `protobuf:"bytes,3,rep,name=lambda_functions,json=lambdaFunctions,proto3" json:"lambda_functions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8a9525eed72921, []int{0}
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

func (m *UpstreamSpec) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *UpstreamSpec) GetSecretRef() *core.ResourceRef {
	if m != nil {
		return m.SecretRef
	}
	return nil
}

func (m *UpstreamSpec) GetLambdaFunctions() []*LambdaFunctionSpec {
	if m != nil {
		return m.LambdaFunctions
	}
	return nil
}

// Each Lambda Function Spec contains data necessary for Gloo to invoke Lambda functions:
// - name of the function
// - qualifier for the function
type LambdaFunctionSpec struct {
	// the logical name gloo should associate with this function. if left empty, it will default to
	// lambda_function_name+qualifier
	LogicalName string `protobuf:"bytes,1,opt,name=logical_name,json=logicalName,proto3" json:"logical_name,omitempty"`
	// The Name of the Lambda Function as it appears in the AWS Lambda Portal
	LambdaFunctionName string `protobuf:"bytes,2,opt,name=lambda_function_name,json=lambdaFunctionName,proto3" json:"lambda_function_name,omitempty"`
	// The Qualifier for the Lambda Function. Qualifiers act as a kind of version
	// for Lambda Functions. See https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html for more info.
	Qualifier            string   `protobuf:"bytes,3,opt,name=qualifier,proto3" json:"qualifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LambdaFunctionSpec) Reset()         { *m = LambdaFunctionSpec{} }
func (m *LambdaFunctionSpec) String() string { return proto.CompactTextString(m) }
func (*LambdaFunctionSpec) ProtoMessage()    {}
func (*LambdaFunctionSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8a9525eed72921, []int{1}
}
func (m *LambdaFunctionSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LambdaFunctionSpec.Unmarshal(m, b)
}
func (m *LambdaFunctionSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LambdaFunctionSpec.Marshal(b, m, deterministic)
}
func (m *LambdaFunctionSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LambdaFunctionSpec.Merge(m, src)
}
func (m *LambdaFunctionSpec) XXX_Size() int {
	return xxx_messageInfo_LambdaFunctionSpec.Size(m)
}
func (m *LambdaFunctionSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_LambdaFunctionSpec.DiscardUnknown(m)
}

var xxx_messageInfo_LambdaFunctionSpec proto.InternalMessageInfo

func (m *LambdaFunctionSpec) GetLogicalName() string {
	if m != nil {
		return m.LogicalName
	}
	return ""
}

func (m *LambdaFunctionSpec) GetLambdaFunctionName() string {
	if m != nil {
		return m.LambdaFunctionName
	}
	return ""
}

func (m *LambdaFunctionSpec) GetQualifier() string {
	if m != nil {
		return m.Qualifier
	}
	return ""
}

// Each Lambda Function Spec contains data necessary for Gloo to invoke Lambda functions
type DestinationSpec struct {
	// The Logical Name of the LambdaFunctionSpec to be invoked.
	LogicalName string `protobuf:"bytes,1,opt,name=logical_name,json=logicalName,proto3" json:"logical_name,omitempty"`
	// Can be either Sync or Async.
	InvocationStyle DestinationSpec_InvocationStyle `protobuf:"varint,2,opt,name=invocation_style,json=invocationStyle,proto3,enum=aws.options.gloo.solo.io.DestinationSpec_InvocationStyle" json:"invocation_style,omitempty"`
	// de-jsonify response bodies returned from aws lambda
	ResponseTransformation bool     `protobuf:"varint,5,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *DestinationSpec) Reset()         { *m = DestinationSpec{} }
func (m *DestinationSpec) String() string { return proto.CompactTextString(m) }
func (*DestinationSpec) ProtoMessage()    {}
func (*DestinationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8a9525eed72921, []int{2}
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

func (m *DestinationSpec) GetLogicalName() string {
	if m != nil {
		return m.LogicalName
	}
	return ""
}

func (m *DestinationSpec) GetInvocationStyle() DestinationSpec_InvocationStyle {
	if m != nil {
		return m.InvocationStyle
	}
	return DestinationSpec_SYNC
}

func (m *DestinationSpec) GetResponseTransformation() bool {
	if m != nil {
		return m.ResponseTransformation
	}
	return false
}

func init() {
	proto.RegisterEnum("aws.options.gloo.solo.io.DestinationSpec_InvocationStyle", DestinationSpec_InvocationStyle_name, DestinationSpec_InvocationStyle_value)
	proto.RegisterType((*UpstreamSpec)(nil), "aws.options.gloo.solo.io.UpstreamSpec")
	proto.RegisterType((*LambdaFunctionSpec)(nil), "aws.options.gloo.solo.io.LambdaFunctionSpec")
	proto.RegisterType((*DestinationSpec)(nil), "aws.options.gloo.solo.io.DestinationSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/aws/aws.proto", fileDescriptor_7e8a9525eed72921)
}

var fileDescriptor_7e8a9525eed72921 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x65, 0x9b, 0xb6, 0x6a, 0x26, 0x15, 0x89, 0x56, 0x55, 0x71, 0x2b, 0x84, 0x42, 0x0e, 0x28,
	0x07, 0x58, 0x43, 0x38, 0x00, 0x12, 0x17, 0x0a, 0xaa, 0x84, 0x84, 0x7a, 0x70, 0x40, 0x08, 0x2e,
	0xd6, 0x66, 0x3b, 0x36, 0x4b, 0x6d, 0xcf, 0xb2, 0xbb, 0x29, 0xe5, 0x0b, 0xf8, 0x15, 0x3e, 0x01,
	0x7e, 0x87, 0x7f, 0xe0, 0xc4, 0x05, 0x79, 0xed, 0x80, 0x12, 0x88, 0x44, 0x0f, 0x96, 0x66, 0xdf,
	0x7b, 0x33, 0xef, 0x8d, 0x35, 0x70, 0x94, 0x6b, 0xff, 0x6e, 0x3e, 0x13, 0x8a, 0xca, 0xd8, 0x51,
	0x41, 0x77, 0x34, 0xc5, 0x79, 0x41, 0x14, 0x1b, 0x4b, 0xef, 0x51, 0x79, 0xd7, 0xbc, 0xa4, 0xd1,
	0xf1, 0xf9, 0xbd, 0x98, 0x8c, 0xd7, 0x54, 0xb9, 0x58, 0x7e, 0x0c, 0x9f, 0x30, 0x96, 0x3c, 0xf1,
	0xa8, 0x2e, 0x5b, 0x4a, 0xd4, 0x72, 0x51, 0x4f, 0x12, 0x9a, 0x0e, 0xf7, 0x72, 0xca, 0x29, 0x88,
	0xe2, 0xba, 0x6a, 0xf4, 0x87, 0x1c, 0x2f, 0x7c, 0x03, 0xe2, 0x85, 0x6f, 0xb1, 0x83, 0x60, 0x7e,
	0xa6, 0xfd, 0xc2, 0xca, 0x62, 0xd6, 0x50, 0xa3, 0x6f, 0x0c, 0x76, 0x5f, 0x19, 0xe7, 0x2d, 0xca,
	0x72, 0x6a, 0x50, 0xf1, 0x7d, 0xd8, 0xb6, 0x98, 0x6b, 0xaa, 0x22, 0x36, 0x64, 0xe3, 0x6e, 0xd2,
	0xbe, 0xf8, 0x43, 0x00, 0x87, 0xca, 0xa2, 0x4f, 0x2d, 0x66, 0xd1, 0xc6, 0x90, 0x8d, 0x7b, 0x93,
	0x03, 0xa1, 0xc8, 0xe2, 0x22, 0x90, 0x48, 0xd0, 0xd1, 0xdc, 0x2a, 0x4c, 0x30, 0x4b, 0xba, 0x8d,
	0x38, 0xc1, 0x8c, 0xbf, 0x86, 0x41, 0x21, 0xcb, 0xd9, 0xa9, 0x4c, 0xb3, 0x79, 0xa5, 0xc2, 0x22,
	0x51, 0x67, 0xd8, 0x19, 0xf7, 0x26, 0xb7, 0xc5, 0xba, 0xe5, 0xc4, 0x8b, 0xd0, 0x71, 0xdc, 0x36,
	0xd4, 0xc9, 0x92, 0x7e, 0xb1, 0x84, 0xb9, 0xd1, 0x67, 0x06, 0xfc, 0x6f, 0x1d, 0xbf, 0x09, 0xbb,
	0x05, 0xe5, 0x5a, 0xc9, 0x22, 0xad, 0x64, 0x89, 0xed, 0x1e, 0xbd, 0x16, 0x3b, 0x91, 0x25, 0xf2,
	0xbb, 0xb0, 0xb7, 0x12, 0xa9, 0x91, 0x6e, 0x04, 0x29, 0x5f, 0x36, 0x0a, 0x1d, 0xd7, 0xa1, 0xfb,
	0x61, 0x2e, 0x0b, 0x9d, 0x69, 0xb4, 0x51, 0x27, 0xc8, 0xfe, 0x00, 0xa3, 0x9f, 0x0c, 0xfa, 0xcf,
	0xd0, 0x79, 0x5d, 0xc9, 0xcb, 0xc4, 0x38, 0x85, 0x81, 0xae, 0xce, 0x49, 0x85, 0xa6, 0xd4, 0xf9,
	0x4f, 0x45, 0x13, 0xe1, 0xea, 0xe4, 0xd1, 0xfa, 0x3f, 0xb3, 0xe2, 0x23, 0x9e, 0xff, 0x9e, 0x30,
	0xad, 0x07, 0x24, 0x7d, 0xbd, 0x0c, 0xf0, 0x07, 0x70, 0xcd, 0xa2, 0x33, 0x54, 0x39, 0x4c, 0xbd,
	0x95, 0x95, 0xcb, 0xc8, 0x96, 0x81, 0x8f, 0xb6, 0x86, 0x6c, 0xbc, 0x93, 0xec, 0x2f, 0xe8, 0x97,
	0x4b, 0xec, 0xe8, 0x16, 0xf4, 0x57, 0x86, 0xf3, 0x1d, 0xd8, 0x9c, 0xbe, 0x39, 0x79, 0x3a, 0xb8,
	0xc2, 0xbb, 0xb0, 0xf5, 0x24, 0x94, 0xec, 0xe8, 0xf8, 0xeb, 0x8f, 0x4d, 0xf6, 0xe5, 0xfb, 0x0d,
	0xf6, 0xf6, 0xf1, 0xff, 0x1d, 0xbc, 0x39, 0xcb, 0xff, 0x71, 0xf4, 0xb3, 0xed, 0x70, 0x92, 0xf7,
	0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x38, 0x97, 0xbd, 0x37, 0x03, 0x00, 0x00,
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
	if this.Region != that1.Region {
		return false
	}
	if !this.SecretRef.Equal(that1.SecretRef) {
		return false
	}
	if len(this.LambdaFunctions) != len(that1.LambdaFunctions) {
		return false
	}
	for i := range this.LambdaFunctions {
		if !this.LambdaFunctions[i].Equal(that1.LambdaFunctions[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LambdaFunctionSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LambdaFunctionSpec)
	if !ok {
		that2, ok := that.(LambdaFunctionSpec)
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
	if this.LogicalName != that1.LogicalName {
		return false
	}
	if this.LambdaFunctionName != that1.LambdaFunctionName {
		return false
	}
	if this.Qualifier != that1.Qualifier {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
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
	if this.LogicalName != that1.LogicalName {
		return false
	}
	if this.InvocationStyle != that1.InvocationStyle {
		return false
	}
	if this.ResponseTransformation != that1.ResponseTransformation {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

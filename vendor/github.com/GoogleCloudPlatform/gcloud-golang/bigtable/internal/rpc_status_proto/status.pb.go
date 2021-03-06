// Code generated by protoc-gen-go.
// source: github.com/GoogleCloudPlatform/gcloud-golang/bigtable/internal/rpc_status_proto/status.proto
// DO NOT EDIT!

/*
Package google_rpc is a generated protocol buffer package.

It is generated from these files:
	github.com/GoogleCloudPlatform/gcloud-golang/bigtable/internal/rpc_status_proto/status.proto

It has these top-level messages:
	Status
*/
package google_rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// The `Status` type defines a logical error model that is suitable for different
// programming environments, including REST APIs and RPC APIs. It is used by
// [gRPC](https://github.com/grpc). The error model is designed to be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error message,
// and error details. The error code should be an enum value of
// [google.rpc.Code][google.rpc.Code], but it may accept additional error codes if needed.  The
// error message should be a developer-facing English message that helps
// developers *understand* and *resolve* the error. If a localized user-facing
// error message is needed, put the localized message in the error details or
// localize it in the client. The optional error details may contain arbitrary
// information about the error. There is a predefined set of error detail types
// in the package `google.rpc` which can be used for common error conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error model, but it
// is not necessarily the actual wire format. When the `Status` message is
// exposed in different client libraries and different wire protocols, it can be
// mapped differently. For example, it will likely be mapped to some exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety of
// environments, either with or without APIs, to provide a
// consistent developer experience across different environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the client,
//     it may embed the `Status` in the normal response to indicate the partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step may
//     have a `Status` message for error reporting purpose.
//
// - Batch operations. If a client uses batch request and batch response, the
//     `Status` message should be used directly inside batch response, one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous operation
//     results in its response, the status of those operations should be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message `Status` could
//     be used directly after any stripping needed for security/privacy reasons.
type Status struct {
	// The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code].
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	// A developer-facing error message, which should be in English. Any
	// user-facing error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// A list of messages that carry the error details.  There will be a
	// common set of message types for APIs to use.
	Details []*google_protobuf.Any `protobuf:"bytes,3,rep,name=details" json:"details,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Status) GetDetails() []*google_protobuf.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "google.rpc.Status")
}

var fileDescriptor0 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x8e, 0xb1, 0x6e, 0x85, 0x20,
	0x18, 0x85, 0x43, 0x6d, 0x35, 0xc5, 0xa4, 0x03, 0xe9, 0x40, 0xbb, 0xd4, 0x74, 0x72, 0x82, 0xa4,
	0x7d, 0x82, 0xba, 0x74, 0x35, 0xf6, 0x01, 0x0c, 0x20, 0x12, 0x13, 0xe4, 0x37, 0x80, 0x83, 0x6f,
	0xdf, 0x5c, 0xd0, 0xdc, 0xed, 0x9c, 0xf0, 0x1d, 0xbe, 0x1f, 0xff, 0x1a, 0x00, 0x63, 0x35, 0x33,
	0x60, 0x85, 0x33, 0x0c, 0xbc, 0xe1, 0xca, 0xc2, 0x3e, 0x71, 0xb9, 0x98, 0x28, 0xa4, 0xd5, 0x7c,
	0x71, 0x51, 0x7b, 0x27, 0x2c, 0xf7, 0x9b, 0x1a, 0x43, 0x14, 0x71, 0x0f, 0xe3, 0xe6, 0x21, 0x02,
	0xcf, 0x85, 0xa5, 0x42, 0xf0, 0xf9, 0x91, 0xdf, 0xd4, 0xfb, 0x5b, 0xce, 0x3c, 0xbd, 0xc8, 0x7d,
	0xe6, 0xc2, 0x1d, 0x19, 0xfb, 0x9c, 0x71, 0xf9, 0x97, 0x66, 0x84, 0xe0, 0x47, 0x05, 0x93, 0xa6,
	0xa8, 0x41, 0xed, 0xd3, 0x90, 0x32, 0xa1, 0xb8, 0x5a, 0x75, 0x08, 0xc2, 0x68, 0xfa, 0xd0, 0xa0,
	0xf6, 0x79, 0xb8, 0x2a, 0x61, 0xb8, 0x9a, 0x74, 0x14, 0x8b, 0x0d, 0xb4, 0x68, 0x8a, 0xb6, 0xfe,
	0x7a, 0x65, 0xa7, 0xf0, 0x92, 0xb0, 0x1f, 0x77, 0x0c, 0x17, 0xd4, 0x7d, 0xe0, 0x17, 0x05, 0x2b,
	0xbb, 0x1f, 0xd5, 0xd5, 0xd9, 0xdb, 0xdf, 0xf0, 0x1e, 0xc9, 0x32, 0xed, 0xbe, 0xff, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x77, 0xd3, 0x68, 0xaf, 0x01, 0x01, 0x00, 0x00,
}

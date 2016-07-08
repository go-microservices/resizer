// Code generated by protoc-gen-go.
// source: google.golang.org/cloud/datastore/internal/type_proto/latlng.proto
// DO NOT EDIT!

/*
Package google_type is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/cloud/datastore/internal/type_proto/latlng.proto

It has these top-level messages:
	LatLng
*/
package google_type

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An object representing a latitude/longitude pair. This is expressed as a pair
// of doubles representing degrees latitude and degrees longitude. Unless
// specified otherwise, this must conform to the
// <a href="http://www.unoosa.org/pdf/icg/2012/template/WGS_84.pdf">WGS84
// standard</a>. Values must be within normalized ranges.
//
// Example of normalization code in Python:
//
//     def NormalizeLongitude(longitude):
//       """Wraps decimal degrees longitude to [-180.0, 180.0]."""
//       q, r = divmod(longitude, 360.0)
//       if r > 180.0 or (r == 180.0 and q <= -1.0):
//         return r - 360.0
//       return r
//
//     def NormalizeLatLng(latitude, longitude):
//       """Wraps decimal degrees latitude and longitude to
//       [-180.0, 180.0] and [-90.0, 90.0], respectively."""
//       r = latitude % 360.0
//       if r <= 90.0:
//         return r, NormalizeLongitude(longitude)
//       elif r >= 270.0:
//         return r - 360, NormalizeLongitude(longitude)
//       else:
//         return 180 - r, NormalizeLongitude(longitude + 180.0)
//
//     assert 180.0 == NormalizeLongitude(180.0)
//     assert -180.0 == NormalizeLongitude(-180.0)
//     assert -179.0 == NormalizeLongitude(181.0)
//     assert (0.0, 0.0) == NormalizeLatLng(360.0, 0.0)
//     assert (0.0, 0.0) == NormalizeLatLng(-360.0, 0.0)
//     assert (85.0, 180.0) == NormalizeLatLng(95.0, 0.0)
//     assert (-85.0, -170.0) == NormalizeLatLng(-95.0, 10.0)
//     assert (90.0, 10.0) == NormalizeLatLng(90.0, 10.0)
//     assert (-90.0, -10.0) == NormalizeLatLng(-90.0, -10.0)
//     assert (0.0, -170.0) == NormalizeLatLng(-180.0, 10.0)
//     assert (0.0, -170.0) == NormalizeLatLng(180.0, 10.0)
//     assert (-90.0, 10.0) == NormalizeLatLng(270.0, 10.0)
//     assert (90.0, 10.0) == NormalizeLatLng(-270.0, 10.0)
type LatLng struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *LatLng) Reset()                    { *m = LatLng{} }
func (m *LatLng) String() string            { return proto.CompactTextString(m) }
func (*LatLng) ProtoMessage()               {}
func (*LatLng) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*LatLng)(nil), "google.type.LatLng")
}

func init() {
	proto.RegisterFile("google.golang.org/cloud/datastore/internal/type_proto/latlng.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x72, 0x4a, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4b, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0xcb, 0x2f, 0x4a, 0xd7, 0x4f, 0xce,
	0xc9, 0x2f, 0x4d, 0xd1, 0x4f, 0x49, 0x2c, 0x49, 0x2c, 0x2e, 0xc9, 0x2f, 0x4a, 0xd5, 0xcf, 0xcc,
	0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x8d, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x49, 0x2c, 0xc9, 0x01, 0x2a, 0x07, 0x73, 0x84, 0xb8, 0xa1, 0x66, 0x80, 0xe4,
	0x95, 0x9c, 0xb8, 0xd8, 0x7c, 0x12, 0x4b, 0x7c, 0xf2, 0xd2, 0x85, 0xa4, 0xb8, 0x38, 0x80, 0xca,
	0x32, 0x4b, 0x4a, 0x53, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x18, 0x83, 0xe0, 0x7c, 0x21, 0x19,
	0x2e, 0xce, 0x9c, 0xfc, 0xbc, 0x74, 0x88, 0x24, 0x13, 0x58, 0x12, 0x21, 0xe0, 0xc4, 0xb6, 0x88,
	0x89, 0xd9, 0x3d, 0x24, 0x20, 0x89, 0x0d, 0x6c, 0xbe, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc0,
	0x8d, 0x0e, 0xf3, 0xa5, 0x00, 0x00, 0x00,
}

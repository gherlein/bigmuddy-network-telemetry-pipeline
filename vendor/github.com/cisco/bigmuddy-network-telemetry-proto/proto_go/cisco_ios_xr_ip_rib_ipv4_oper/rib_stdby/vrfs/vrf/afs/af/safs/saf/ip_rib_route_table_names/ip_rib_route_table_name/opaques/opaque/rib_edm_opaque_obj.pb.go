// Code generated by protoc-gen-go.
// source: rib_edm_opaque_obj.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_opaques_opaque is a generated protocol buffer package.

It is generated from these files:
	rib_edm_opaque_obj.proto

It has these top-level messages:
	RibEdmOpaqueObj_KEYS
	RibEdmOpaqueObj
*/
package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_opaques_opaque

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

// Informaton of an opaque data
type RibEdmOpaqueObj_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
	OpaqueClientid uint32 `protobuf:"varint,5,opt,name=opaque_clientid,json=opaqueClientid" json:"opaque_clientid,omitempty"`
	OpaqueProtoid  uint32 `protobuf:"varint,6,opt,name=opaque_protoid,json=opaqueProtoid" json:"opaque_protoid,omitempty"`
	OpaqueKeyType  uint32 `protobuf:"varint,7,opt,name=opaque_key_type,json=opaqueKeyType" json:"opaque_key_type,omitempty"`
	OpaqueKeySize  uint32 `protobuf:"varint,8,opt,name=opaque_key_size,json=opaqueKeySize" json:"opaque_key_size,omitempty"`
	OpaqueDataSize uint32 `protobuf:"varint,9,opt,name=opaque_data_size,json=opaqueDataSize" json:"opaque_data_size,omitempty"`
	OpaqueString   string `protobuf:"bytes,10,opt,name=opaque_string,json=opaqueString" json:"opaque_string,omitempty"`
}

func (m *RibEdmOpaqueObj_KEYS) Reset()                    { *m = RibEdmOpaqueObj_KEYS{} }
func (m *RibEdmOpaqueObj_KEYS) String() string            { return proto.CompactTextString(m) }
func (*RibEdmOpaqueObj_KEYS) ProtoMessage()               {}
func (*RibEdmOpaqueObj_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RibEdmOpaqueObj_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *RibEdmOpaqueObj_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *RibEdmOpaqueObj_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *RibEdmOpaqueObj_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *RibEdmOpaqueObj_KEYS) GetOpaqueClientid() uint32 {
	if m != nil {
		return m.OpaqueClientid
	}
	return 0
}

func (m *RibEdmOpaqueObj_KEYS) GetOpaqueProtoid() uint32 {
	if m != nil {
		return m.OpaqueProtoid
	}
	return 0
}

func (m *RibEdmOpaqueObj_KEYS) GetOpaqueKeyType() uint32 {
	if m != nil {
		return m.OpaqueKeyType
	}
	return 0
}

func (m *RibEdmOpaqueObj_KEYS) GetOpaqueKeySize() uint32 {
	if m != nil {
		return m.OpaqueKeySize
	}
	return 0
}

func (m *RibEdmOpaqueObj_KEYS) GetOpaqueDataSize() uint32 {
	if m != nil {
		return m.OpaqueDataSize
	}
	return 0
}

func (m *RibEdmOpaqueObj_KEYS) GetOpaqueString() string {
	if m != nil {
		return m.OpaqueString
	}
	return ""
}

type RibEdmOpaqueObj struct {
	Key  []byte `protobuf:"bytes,50,opt,name=key,proto3" json:"key,omitempty"`
	Data []byte `protobuf:"bytes,51,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *RibEdmOpaqueObj) Reset()                    { *m = RibEdmOpaqueObj{} }
func (m *RibEdmOpaqueObj) String() string            { return proto.CompactTextString(m) }
func (*RibEdmOpaqueObj) ProtoMessage()               {}
func (*RibEdmOpaqueObj) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RibEdmOpaqueObj) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RibEdmOpaqueObj) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*RibEdmOpaqueObj_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.opaques.opaque.rib_edm_opaque_obj_KEYS")
	proto.RegisterType((*RibEdmOpaqueObj)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.opaques.opaque.rib_edm_opaque_obj")
}

func init() { proto.RegisterFile("rib_edm_opaque_obj.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4f, 0xc2, 0x30,
	0x18, 0xc6, 0x33, 0x41, 0x3e, 0x1a, 0x40, 0xd2, 0x0b, 0xe5, 0x46, 0x30, 0xea, 0x4e, 0x3d, 0x88,
	0x27, 0xaf, 0xea, 0x89, 0xc4, 0x18, 0xe0, 0xe2, 0xa9, 0xe9, 0xd8, 0x3b, 0x53, 0xf9, 0xe8, 0x6c,
	0xcb, 0xe2, 0x38, 0xe9, 0x7f, 0x6e, 0xf6, 0xb6, 0x1a, 0x02, 0xf1, 0xb2, 0x76, 0xbf, 0xe7, 0xf7,
	0x6e, 0x4f, 0x37, 0xc2, 0x8c, 0x4a, 0x04, 0xa4, 0x1b, 0xa1, 0x73, 0xf9, 0xb1, 0x03, 0xa1, 0x93,
	0x77, 0x9e, 0x1b, 0xed, 0x34, 0xfd, 0x8a, 0x96, 0xca, 0x2e, 0xb5, 0x50, 0xda, 0x8a, 0x4f, 0x23,
	0x54, 0x2e, 0x2a, 0x55, 0xe5, 0xc5, 0x9d, 0xd0, 0x39, 0x18, 0x5e, 0xdd, 0x59, 0x97, 0x26, 0x25,
	0x2f, 0x4c, 0x66, 0xab, 0x0b, 0x97, 0x99, 0xe5, 0x32, 0xe3, 0xb6, 0x5a, 0xad, 0xcc, 0x78, 0x98,
	0x31, 0x7a, 0xe7, 0x40, 0x38, 0x99, 0xac, 0x41, 0x6c, 0xe5, 0x06, 0xec, 0x7f, 0x01, 0xf7, 0x25,
	0x6c, 0x58, 0xc7, 0xdf, 0x35, 0x32, 0x38, 0xed, 0x27, 0xa6, 0x4f, 0xaf, 0x73, 0x3a, 0x24, 0xad,
	0xc2, 0x64, 0x38, 0xc6, 0xa2, 0x51, 0x14, 0xb7, 0x67, 0xcd, 0xc2, 0x64, 0xcf, 0x72, 0x03, 0x74,
	0x40, 0x9a, 0x32, 0x24, 0x67, 0x98, 0x34, 0xa4, 0x0f, 0x86, 0xa4, 0x65, 0x7f, 0x93, 0x9a, 0x9f,
	0xb1, 0x21, 0x8a, 0x49, 0xff, 0xb8, 0x0d, 0xab, 0xa3, 0xd2, 0x43, 0xbe, 0xa8, 0x30, 0x9a, 0x37,
	0xe4, 0x22, 0x74, 0x59, 0xae, 0x15, 0x6c, 0x9d, 0x4a, 0xd9, 0xf9, 0x28, 0x8a, 0xbb, 0xb3, 0x9e,
	0xc7, 0x0f, 0x81, 0xd2, 0x2b, 0x12, 0x88, 0xc0, 0x0f, 0xaa, 0x52, 0xd6, 0x40, 0xaf, 0xeb, 0xe9,
	0x8b, 0x87, 0xf4, 0xfa, 0xef, 0x79, 0x2b, 0x28, 0x85, 0x2b, 0x73, 0x60, 0xcd, 0x43, 0x6f, 0x0a,
	0xe5, 0xa2, 0xcc, 0xe1, 0xc8, 0xb3, 0x6a, 0x0f, 0xac, 0x75, 0xe4, 0xcd, 0xd5, 0x1e, 0x4f, 0x12,
	0xbc, 0x54, 0x3a, 0xe9, 0xc5, 0xf6, 0x61, 0xc1, 0x47, 0xe9, 0x24, 0x9a, 0x97, 0x24, 0x8c, 0x0a,
	0xeb, 0x8c, 0xda, 0xbe, 0x31, 0x82, 0x07, 0xee, 0x78, 0x38, 0x47, 0x36, 0xbe, 0x27, 0xf4, 0xf4,
	0x17, 0xd0, 0x3e, 0xa9, 0xad, 0xa0, 0x64, 0xb7, 0xa3, 0x28, 0xee, 0xcc, 0xaa, 0x2d, 0xa5, 0xa4,
	0x5e, 0xbd, 0x8f, 0x4d, 0x10, 0xe1, 0x3e, 0x69, 0xe0, 0xc1, 0x27, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x3c, 0xea, 0x66, 0x03, 0x65, 0x02, 0x00, 0x00,
}

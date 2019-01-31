// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helper/storagepacker/types.proto

package storagepacker

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Item represents a entry that gets inserted into the storage packer
type Item struct {
	// ID is the UUID to identify the item
	ID string `sentinel:"" protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// message is the contents of the item
	Message              *any.Any `sentinel:"" protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0e98c66c4f51b7f, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Item) GetMessage() *any.Any {
	if m != nil {
		return m.Message
	}
	return nil
}

// Bucket is a construct to hold multiple items within itself. This
// abstraction contains multiple buckets of the same kind within itself and
// shares amont them the items that get inserted. When the bucket as a whole
// gets too big to hold more items, the contained buckets gets pushed out only
// to become independent buckets. Hence, this can grow infinitely in terms of
// storage space for items that get inserted.
type Bucket struct {
	// Key is the storage path where the bucket gets stored
	Key string `sentinel:"" protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Items holds the items contained within this bucket. Used by v1.
	Items []*Item `sentinel:"" protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	// ItemMap stores a mapping of item ID to message. Used by v2.
	ItemMap map[string]*any.Any `sentinel:"" protobuf:"bytes,3,rep,name=item_map,json=itemMap,proto3" json:"item_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Buckets are the buckets contained within this bucket
	Buckets map[string]*Bucket `sentinel:"" protobuf:"bytes,4,rep,name=buckets,proto3" json:"buckets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Sharded indicates if the contained buckets are pushed out or not
	Sharded bool `sentinel:"" protobuf:"varint,5,opt,name=sharded,proto3" json:"sharded,omitempty"`
	// Size of this bucket in number of bytes
	Size                 uint32   `sentinel:"" protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0e98c66c4f51b7f, []int{1}
}

func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (m *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(m, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Bucket) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Bucket) GetItemMap() map[string]*any.Any {
	if m != nil {
		return m.ItemMap
	}
	return nil
}

func (m *Bucket) GetBuckets() map[string]*Bucket {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *Bucket) GetSharded() bool {
	if m != nil {
		return m.Sharded
	}
	return false
}

func (m *Bucket) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*Item)(nil), "storagepacker.Item")
	proto.RegisterType((*Bucket)(nil), "storagepacker.Bucket")
	proto.RegisterMapType((map[string]*Bucket)(nil), "storagepacker.Bucket.BucketsEntry")
	proto.RegisterMapType((map[string]*any.Any)(nil), "storagepacker.Bucket.ItemMapEntry")
}

func init() { proto.RegisterFile("helper/storagepacker/types.proto", fileDescriptor_c0e98c66c4f51b7f) }

var fileDescriptor_c0e98c66c4f51b7f = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4b, 0xeb, 0x40,
	0x10, 0xc6, 0x49, 0xd2, 0x36, 0x7d, 0xd3, 0xf6, 0xf1, 0xd8, 0xa7, 0xb0, 0xf6, 0x14, 0x7a, 0x8a,
	0x0a, 0xbb, 0x58, 0x2f, 0x22, 0x7a, 0xb0, 0xa0, 0xe0, 0x41, 0xd0, 0x1c, 0xbd, 0xc8, 0x26, 0x19,
	0x93, 0xd0, 0xa4, 0x09, 0xbb, 0x9b, 0x42, 0xfc, 0x3f, 0xfc, 0x7f, 0xa5, 0xd9, 0x16, 0x5a, 0x09,
	0x3d, 0xed, 0x0c, 0xf3, 0x7d, 0xbf, 0xfd, 0x66, 0x59, 0xf0, 0x52, 0xcc, 0x2b, 0x94, 0x5c, 0xe9,
	0x52, 0x8a, 0x04, 0x2b, 0x11, 0x2d, 0x51, 0x72, 0xdd, 0x54, 0xa8, 0x58, 0x25, 0x4b, 0x5d, 0x92,
	0xc9, 0xc1, 0x68, 0x7a, 0x96, 0x94, 0x65, 0x92, 0x23, 0x6f, 0x87, 0x61, 0xfd, 0xc9, 0xc5, 0xaa,
	0x31, 0xca, 0xd9, 0x13, 0xf4, 0x9e, 0x35, 0x16, 0xe4, 0x2f, 0xd8, 0x59, 0x4c, 0x2d, 0xcf, 0xf2,
	0xff, 0x04, 0x76, 0x16, 0x13, 0x06, 0x6e, 0x81, 0x4a, 0x89, 0x04, 0xa9, 0xed, 0x59, 0xfe, 0x68,
	0x7e, 0xc2, 0x0c, 0x84, 0xed, 0x20, 0xec, 0x61, 0xd5, 0x04, 0x3b, 0xd1, 0xec, 0xdb, 0x81, 0xc1,
	0xa2, 0x8e, 0x96, 0xa8, 0xc9, 0x3f, 0x70, 0x96, 0xd8, 0x6c, 0x59, 0x9b, 0x92, 0x9c, 0x43, 0x3f,
	0xd3, 0x58, 0x28, 0x6a, 0x7b, 0x8e, 0x3f, 0x9a, 0xff, 0x67, 0x07, 0xf1, 0xd8, 0x26, 0x40, 0x60,
	0x14, 0xe4, 0x1e, 0x86, 0x9b, 0xe2, 0xa3, 0x10, 0x15, 0x75, 0x5a, 0xf5, 0xec, 0x97, 0xda, 0xdc,
	0xd2, 0x9a, 0x5e, 0x44, 0xf5, 0xb8, 0xd2, 0xb2, 0x09, 0xdc, 0xcc, 0x74, 0xe4, 0x0e, 0xdc, 0xb0,
	0x9d, 0x2b, 0xda, 0x3b, 0xe6, 0x36, 0x87, 0xda, 0xba, 0xb7, 0x16, 0x42, 0xc1, 0x55, 0xa9, 0x90,
	0x31, 0xc6, 0xb4, 0xef, 0x59, 0xfe, 0x30, 0xd8, 0xb5, 0x84, 0x40, 0x4f, 0x65, 0x5f, 0x48, 0x07,
	0x9e, 0xe5, 0x4f, 0x82, 0xb6, 0x9e, 0xbe, 0xc2, 0x78, 0x3f, 0x44, 0xc7, 0xde, 0x17, 0xd0, 0x5f,
	0x8b, 0xbc, 0x3e, 0xfe, 0x84, 0x46, 0x72, 0x6b, 0xdf, 0x58, 0xd3, 0x37, 0x18, 0xef, 0x07, 0xeb,
	0x20, 0x5e, 0x1e, 0x12, 0x4f, 0x3b, 0xb7, 0xdb, 0x43, 0x2e, 0xae, 0xde, 0x79, 0x92, 0xe9, 0xb4,
	0x0e, 0x59, 0x54, 0x16, 0x3c, 0x15, 0x2a, 0xcd, 0xa2, 0x52, 0x56, 0x7c, 0x2d, 0xea, 0x5c, 0xf3,
	0xae, 0x8f, 0x14, 0x0e, 0xda, 0x78, 0xd7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x29, 0x1c,
	0x91, 0x67, 0x02, 0x00, 0x00,
}

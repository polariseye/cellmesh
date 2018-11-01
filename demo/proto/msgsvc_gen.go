// Generated by github.com/davyxu/protoplus
// DO NOT EDIT!
package proto

import (
	"github.com/davyxu/protoplus/proto"
	"unsafe"
)

var (
	_ *proto.Buffer
	_ unsafe.Pointer
)

type ClientID struct {
	ID    int64  // 客户端在网关上的SessionID
	SvcID string // 客户端在哪个网关
}

func (self *ClientID) String() string { return proto.CompactTextString(self) }

func (self *ClientID) Size() (ret int) {

	ret += proto.SizeInt64(0, self.ID)

	ret += proto.SizeString(1, self.SvcID)

	return
}

func (self *ClientID) Marshal(buffer *proto.Buffer) error {

	proto.MarshalInt64(buffer, 0, self.ID)

	proto.MarshalString(buffer, 1, self.SvcID)

	return nil
}

func (self *ClientID) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalInt64(buffer, wt, &self.ID)
	case 1:
		return proto.UnmarshalString(buffer, wt, &self.SvcID)

	}

	return proto.ErrUnknownField
}

type BindBackendACK struct {
	ID int64
}

func (self *BindBackendACK) String() string { return proto.CompactTextString(self) }

func (self *BindBackendACK) Size() (ret int) {

	ret += proto.SizeInt64(0, self.ID)

	return
}

func (self *BindBackendACK) Marshal(buffer *proto.Buffer) error {

	proto.MarshalInt64(buffer, 0, self.ID)

	return nil
}

func (self *BindBackendACK) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalInt64(buffer, wt, &self.ID)

	}

	return proto.ErrUnknownField
}

type CloseClientACK struct {
	ID  []int64
	All bool
}

func (self *CloseClientACK) String() string { return proto.CompactTextString(self) }

func (self *CloseClientACK) Size() (ret int) {

	ret += proto.SizeInt64Slice(0, self.ID)

	ret += proto.SizeBool(1, self.All)

	return
}

func (self *CloseClientACK) Marshal(buffer *proto.Buffer) error {

	proto.MarshalInt64Slice(buffer, 0, self.ID)

	proto.MarshalBool(buffer, 1, self.All)

	return nil
}

func (self *CloseClientACK) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalInt64Slice(buffer, wt, &self.ID)
	case 1:
		return proto.UnmarshalBool(buffer, wt, &self.All)

	}

	return proto.ErrUnknownField
}

type ClientClosedACK struct {
	ID ClientID
}

func (self *ClientClosedACK) String() string { return proto.CompactTextString(self) }

func (self *ClientClosedACK) Size() (ret int) {

	ret += proto.SizeStruct(0, &self.ID)

	return
}

func (self *ClientClosedACK) Marshal(buffer *proto.Buffer) error {

	proto.MarshalStruct(buffer, 0, &self.ID)

	return nil
}

func (self *ClientClosedACK) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalStruct(buffer, wt, &self.ID)

	}

	return proto.ErrUnknownField
}

type SubscribeChannelREQ struct {
	Channel string
}

func (self *SubscribeChannelREQ) String() string { return proto.CompactTextString(self) }

func (self *SubscribeChannelREQ) Size() (ret int) {

	ret += proto.SizeString(0, self.Channel)

	return
}

func (self *SubscribeChannelREQ) Marshal(buffer *proto.Buffer) error {

	proto.MarshalString(buffer, 0, self.Channel)

	return nil
}

func (self *SubscribeChannelREQ) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalString(buffer, wt, &self.Channel)

	}

	return proto.ErrUnknownField
}

type SubscribeChannelACK struct {
	Channel string
}

func (self *SubscribeChannelACK) String() string { return proto.CompactTextString(self) }

func (self *SubscribeChannelACK) Size() (ret int) {

	ret += proto.SizeString(0, self.Channel)

	return
}

func (self *SubscribeChannelACK) Marshal(buffer *proto.Buffer) error {

	proto.MarshalString(buffer, 0, self.Channel)

	return nil
}

func (self *SubscribeChannelACK) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalString(buffer, wt, &self.Channel)

	}

	return proto.ErrUnknownField
}

type SvcStatusACK struct {
	SvcID     string
	UserCount int32
}

func (self *SvcStatusACK) String() string { return proto.CompactTextString(self) }

func (self *SvcStatusACK) Size() (ret int) {

	ret += proto.SizeString(0, self.SvcID)

	ret += proto.SizeInt32(1, self.UserCount)

	return
}

func (self *SvcStatusACK) Marshal(buffer *proto.Buffer) error {

	proto.MarshalString(buffer, 0, self.SvcID)

	proto.MarshalInt32(buffer, 1, self.UserCount)

	return nil
}

func (self *SvcStatusACK) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalString(buffer, wt, &self.SvcID)
	case 1:
		return proto.UnmarshalInt32(buffer, wt, &self.UserCount)

	}

	return proto.ErrUnknownField
}

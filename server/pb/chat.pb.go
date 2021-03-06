// Code generated by protoc-gen-go.
// source: chat.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	chat.proto
	echo.proto
	item.proto
	role.proto
	task.proto

It has these top-level messages:
	MQChat
*/
package pb

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MQChat struct {
	ChatType         *uint32 `protobuf:"varint,1,req,name=chatType" json:"chatType,omitempty"`
	TargetId         *uint32 `protobuf:"varint,2,req,name=targetId" json:"targetId,omitempty"`
	Content          *string `protobuf:"bytes,3,req,name=content" json:"content,omitempty"`
	From             *uint32 `protobuf:"varint,4,opt,name=from" json:"from,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MQChat) Reset()         { *m = MQChat{} }
func (m *MQChat) String() string { return proto.CompactTextString(m) }
func (*MQChat) ProtoMessage()    {}

func (m *MQChat) GetChatType() uint32 {
	if m != nil && m.ChatType != nil {
		return *m.ChatType
	}
	return 0
}

func (m *MQChat) GetTargetId() uint32 {
	if m != nil && m.TargetId != nil {
		return *m.TargetId
	}
	return 0
}

func (m *MQChat) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func (m *MQChat) GetFrom() uint32 {
	if m != nil && m.From != nil {
		return *m.From
	}
	return 0
}

func init() {
}

// Code generated by protoc-gen-go.
// source: task.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Task struct {
	Id               *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Progress         *uint32 `protobuf:"varint,2,req,name=progress" json:"progress,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}

func (m *Task) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Task) GetProgress() uint32 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

func init() {
}

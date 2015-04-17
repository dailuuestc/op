// Code generated by protoc-gen-go.
// source: role.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MRRoleBasic struct {
	Basic            *RoleBasic `protobuf:"bytes,1,req,name=basic" json:"basic,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *MRRoleBasic) Reset()         { *m = MRRoleBasic{} }
func (m *MRRoleBasic) String() string { return proto.CompactTextString(m) }
func (*MRRoleBasic) ProtoMessage()    {}

func (m *MRRoleBasic) GetBasic() *RoleBasic {
	if m != nil {
		return m.Basic
	}
	return nil
}

type MRTasklist struct {
	Tasklist         []*Task `protobuf:"bytes,2,rep,name=tasklist" json:"tasklist,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MRTasklist) Reset()         { *m = MRTasklist{} }
func (m *MRTasklist) String() string { return proto.CompactTextString(m) }
func (*MRTasklist) ProtoMessage()    {}

func (m *MRTasklist) GetTasklist() []*Task {
	if m != nil {
		return m.Tasklist
	}
	return nil
}

type MRItemlist struct {
	Itemlist         []*Item `protobuf:"bytes,3,rep,name=itemlist" json:"itemlist,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MRItemlist) Reset()         { *m = MRItemlist{} }
func (m *MRItemlist) String() string { return proto.CompactTextString(m) }
func (*MRItemlist) ProtoMessage()    {}

func (m *MRItemlist) GetItemlist() []*Item {
	if m != nil {
		return m.Itemlist
	}
	return nil
}

func init() {
}
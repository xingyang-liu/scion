// Code generated by capnpc-go. DO NOT EDIT.

package proto

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type CtrlExtnDataList struct{ capnp.Struct }

// CtrlExtnDataList_TypeID is the unique identifier for the type CtrlExtnDataList.
const CtrlExtnDataList_TypeID = 0xcd601dbe8602c166

func NewCtrlExtnDataList(s *capnp.Segment) (CtrlExtnDataList, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return CtrlExtnDataList{st}, err
}

func NewRootCtrlExtnDataList(s *capnp.Segment) (CtrlExtnDataList, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return CtrlExtnDataList{st}, err
}

func ReadRootCtrlExtnDataList(msg *capnp.Message) (CtrlExtnDataList, error) {
	root, err := msg.RootPtr()
	return CtrlExtnDataList{root.Struct()}, err
}

func (s CtrlExtnDataList) String() string {
	str, _ := text.Marshal(0xcd601dbe8602c166, s.Struct)
	return str
}

func (s CtrlExtnDataList) Items() (CtrlExtnData_List, error) {
	p, err := s.Struct.Ptr(0)
	return CtrlExtnData_List{List: p.List()}, err
}

func (s CtrlExtnDataList) HasItems() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CtrlExtnDataList) SetItems(v CtrlExtnData_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewItems sets the items field to a newly
// allocated CtrlExtnData_List, preferring placement in s's segment.
func (s CtrlExtnDataList) NewItems(n int32) (CtrlExtnData_List, error) {
	l, err := NewCtrlExtnData_List(s.Struct.Segment(), n)
	if err != nil {
		return CtrlExtnData_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// CtrlExtnDataList_List is a list of CtrlExtnDataList.
type CtrlExtnDataList_List struct{ capnp.List }

// NewCtrlExtnDataList creates a new list of CtrlExtnDataList.
func NewCtrlExtnDataList_List(s *capnp.Segment, sz int32) (CtrlExtnDataList_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return CtrlExtnDataList_List{l}, err
}

func (s CtrlExtnDataList_List) At(i int) CtrlExtnDataList { return CtrlExtnDataList{s.List.Struct(i)} }

func (s CtrlExtnDataList_List) Set(i int, v CtrlExtnDataList) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s CtrlExtnDataList_List) String() string {
	str, _ := text.MarshalList(0xcd601dbe8602c166, s.List)
	return str
}

// CtrlExtnDataList_Promise is a wrapper for a CtrlExtnDataList promised by a client call.
type CtrlExtnDataList_Promise struct{ *capnp.Pipeline }

func (p CtrlExtnDataList_Promise) Struct() (CtrlExtnDataList, error) {
	s, err := p.Pipeline.Struct()
	return CtrlExtnDataList{s}, err
}

type CtrlExtnData struct{ capnp.Struct }

// CtrlExtnData_TypeID is the unique identifier for the type CtrlExtnData.
const CtrlExtnData_TypeID = 0xbd3bde9856b0a03f

func NewCtrlExtnData(s *capnp.Segment) (CtrlExtnData, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return CtrlExtnData{st}, err
}

func NewRootCtrlExtnData(s *capnp.Segment) (CtrlExtnData, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return CtrlExtnData{st}, err
}

func ReadRootCtrlExtnData(msg *capnp.Message) (CtrlExtnData, error) {
	root, err := msg.RootPtr()
	return CtrlExtnData{root.Struct()}, err
}

func (s CtrlExtnData) String() string {
	str, _ := text.Marshal(0xbd3bde9856b0a03f, s.Struct)
	return str
}

func (s CtrlExtnData) Type() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s CtrlExtnData) HasType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CtrlExtnData) SetType(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s CtrlExtnData) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s CtrlExtnData) HasData() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s CtrlExtnData) SetData(v []byte) error {
	return s.Struct.SetData(1, v)
}

// CtrlExtnData_List is a list of CtrlExtnData.
type CtrlExtnData_List struct{ capnp.List }

// NewCtrlExtnData creates a new list of CtrlExtnData.
func NewCtrlExtnData_List(s *capnp.Segment, sz int32) (CtrlExtnData_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return CtrlExtnData_List{l}, err
}

func (s CtrlExtnData_List) At(i int) CtrlExtnData { return CtrlExtnData{s.List.Struct(i)} }

func (s CtrlExtnData_List) Set(i int, v CtrlExtnData) error { return s.List.SetStruct(i, v.Struct) }

func (s CtrlExtnData_List) String() string {
	str, _ := text.MarshalList(0xbd3bde9856b0a03f, s.List)
	return str
}

// CtrlExtnData_Promise is a wrapper for a CtrlExtnData promised by a client call.
type CtrlExtnData_Promise struct{ *capnp.Pipeline }

func (p CtrlExtnData_Promise) Struct() (CtrlExtnData, error) {
	s, err := p.Pipeline.Struct()
	return CtrlExtnData{s}, err
}

const schema_fa3f5dec2b78a085 = "x\xda2`gu`2d\xd5Wc`\x08\xde\xc3" +
	"\xc8\xca\xf6\xdf~\xc1\x86\xb0\x19\xf7\xac\xf72\x08\x0a0" +
	"\xfeo]P\xa1\xfd&\xd6\xfe\x17\x03+\x13;\x03\x83" +
	"\xf0_\xa1W\xc2\x9c\xc2 \x16\xabp9\x03\xe3\xff\xb4" +
	"\x83Lm\xfbd\x13\xce\xa2\xa9edg`0N\x14" +
	"\xe6b\x14.\x04+\xce\x15\xb6g\xf8\x0f\x85\x1b\xfe'" +
	"\x97\x14\xe5\xc4\xa7V\x940\xe5\xe9%'\x16\xe4\x15X" +
	"9\x97\x14\xe5\xb8V\x94\xe4\xf1\xbb$\x96$\x0602" +
	"\x06r0\xb300\xb0020\x08jj10\x04" +
	"\xaa03\x06\x1a01\x0a22\x8a0\x82\x04uA" +
	"\x82\x1a\xcc\x8c\x81&L\x8c\xfc%\x95\x05\xa9\x8c\xbc\x0c" +
	"L\x8c\xbc\x0c\x8c\xfc)\x89%\x890\x0e\xdc\x1eft" +
	"{@\xd6\xf8d\x16\x9700\x80,c\x81[\xc6k" +
	"\xc4\xc0\x10\xc8\xc1\xcc\x18\xa8\xc2\xc4(\x9fY\x92\x9a[" +
	"\xcc\xc8\xc7\xc0\x18\xc0\xcc\xc8(\x80\x08\x15\x06F\x90 " +
	" \x00\x00\xff\xff\xa8WK\xcc"

func init() {
	schemas.Register(schema_fa3f5dec2b78a085,
		0xbd3bde9856b0a03f,
		0xcd601dbe8602c166)
}
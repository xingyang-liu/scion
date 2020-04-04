package seg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/scionproto/scion/go/proto"
	"io"
)

var _ proto.Cerealizable = (*Lppair)(nil)
var _ proto.Cerealizable = (*Lpcluster)(nil)
var _ proto.Cerealizable = (*Lnpcluster)(nil)
var _ proto.Cerealizable = (*Latencyinfo)(nil)
var _ proto.Cerealizable = (*LinkmetricsExtn)(nil)

type Lppair struct {
	Interface  uint16
	Interdelay uint16
}

func (lp *Lppair) ProtoId() proto.ProtoIdType {
	return proto.Lppair_TypeID
}

func (lp *Lppair) String() string {
	if lp == nil {
		return "nil"
	}
	return fmt.Sprintf("(Interface %v: Delay %v)", lp.Interface, lp.Interdelay)
}

type Lpcluster struct {
	ClusterDelay    uint16
	LInterfacePairs []*Lppair
}

func (lpc *Lpcluster) ProtoId() proto.ProtoIdType {
	return proto.Lpcluster_TypeID
}

func (lpc *Lpcluster) String() string {
	if lpc == nil {
		return "nil"
	}
	pairs := ""
	for _, pair := range lpc.LInterfacePairs {
		pairs += fmt.Sprintf("%v,", pair)
	}
	return fmt.Sprintf("ClusterDelay %v: [%v]", lpc.ClusterDelay, pairs)
}

type Lnpcluster struct {
	ClusterDelay uint16
	Interfaces   []uint16
}

func (lnpc *Lnpcluster) ProtoId() proto.ProtoIdType {
	return proto.Lnpcluster_TypeID
}

func (lnpc *Lnpcluster) String() string {
	if lnpc == nil {
		return "nil"
	}
	interfaces := ""
	for _, intf := range lnpc.Interfaces {
		interfaces += fmt.Sprintf("%v,", intf)
	}
	return fmt.Sprintf("ClusterDelay %v: [%v]", lnpc.ClusterDelay, interfaces)
}

type Latencyinfo struct {
	LnpClusters    []Lnpcluster
	LpClusters     []Lpcluster
	EgressLatency  uint16
	InToOutLatency uint16
}

func (li *Latencyinfo) ProtoId() proto.ProtoIdType {
	return proto.Latencyinfo_TypeID
}

func (li *Latencyinfo) String() string {
	if li == nil {
		return "nil"
	}
	lpclusters := "\n  "
	for _, lpc := range li.LpClusters {
		lpclusters += fmt.Sprintf("  %v\n  ", lpc)
	}
	lnpclusters := "    "
	for _, lnpc := range li.LnpClusters {
		lnpclusters += fmt.Sprintf("%v\n    ", lnpc)
	}
	return fmt.Sprintf("\n  EgressLatency: %v\n  InToOutLatency: %v\n  LpClusters: [%v]\n  LnpClusters: [%v]\n",
		li.EgressLatency, li.InToOutLatency, lpclusters, lnpclusters)
}

type LinkmetricsExtn struct {
	Set   bool
	LInfo *Latencyinfo
}

func (lm *LinkmetricsExtn) ProtoId() proto.ProtoIdType {
	return proto.LinkmetricsExtn_TypeID
}

func (lm *LinkmetricsExtn) String() string {
	if lm == nil {
		return "nil"
	}
	return fmt.Sprintf("Set: %v\nLatencyInfo: %v", lm.Set, lm.LInfo)
}

func (lm *LinkmetricsExtn) WriteTo(w io.Writer) (int64, error) {
	var total int64
	var buf bytes.Buffer
	err := proto.SerializeTo(lm, &buf)
	if err != nil {
		return 0, err
	}

	blen := make([]byte, 4)
	binary.LittleEndian.PutUint32(blen, uint32(buf.Len()))
	n, err := w.Write(blen)
	if err != nil {
		return 0, err
	}
	total += int64(n)

	n, err = w.Write(buf.Bytes())
	if err != nil {
		return total, err
	}
	total += int64(n)
	return total, nil
}

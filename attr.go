package main

const (
	attrTypeOrigin          = 1
	attrTypeAsPath          = 2
	attrTypeNextHop         = 3
	attrTypeMED             = 4
	attrTypeLocalPref       = 5
	attrTypeAtomicAggregate = 6
	attrTypeAggregator      = 7
)

const (
	attrFlagOptional    = 1 << 7
	attrFlagTransitive  = 1 << 6
	attrFlagPartial     = 1 << 5
	attrFlagExtendedLen = 1 << 4
)

//----------------------------------------------------------------------------------------------------------------------
// Attr
//----------------------------------------------------------------------------------------------------------------------

// Attr is the interface for all attributes
type Attr interface {
	Encode(buf *MsgBuffer)
}

//----------------------------------------------------------------------------------------------------------------------
// OriginAttr
//----------------------------------------------------------------------------------------------------------------------

// OriginAttr contains a decoded origin attribute
type OriginAttr uint8

// Origin attribute values
const (
	OriginAttrIgp        = 0
	OriginAttrEgp        = 1
	OriginAttrIncomplete = 2
)

// Encode encodes an origin attribute
func (a *OriginAttr) Encode(buf *MsgBuffer) {
	buf.AppendByte(attrFlagTransitive)
	buf.AppendByte(attrTypeOrigin)
	buf.AppendByte(uint8(*a))
}

//----------------------------------------------------------------------------------------------------------------------
// AsPathAttr
//----------------------------------------------------------------------------------------------------------------------

// ASPathAttr contains a decoded as-path
type ASPathAttr struct {
	ASPath ASPath // TODO: Maybe not a good idea: same name for var and type
}

// Encode encodes an as-path attribute
func (p *ASPathAttr) Encode(buf *MsgBuffer) {
	p.ASPath.Encode(buf)
}

//----------------------------------------------------------------------------------------------------------------------
// NextHopAttr
//----------------------------------------------------------------------------------------------------------------------

// NextHopAttr contains a decoded next-hop attribute
type NextHopAttr IPv4Address

// Encode encodes a next-hop attribute
func (a *NextHopAttr) Encode(buf *MsgBuffer) {
	buf.AppendByte(attrFlagTransitive)
	buf.AppendByte(attrTypeNextHop)
	a.Encode(buf)
}

//----------------------------------------------------------------------------------------------------------------------
// MEDAttr
//----------------------------------------------------------------------------------------------------------------------

// MEDAttr contains a decoded multi-exit discriminator attribute
type MEDAttr uint32

// Encode encodes a MED attribute
func (a *MEDAttr) Encode(buf *MsgBuffer) {
	buf.AppendByte(attrFlagTransitive)
	buf.AppendByte(attrTypeMED)
	buf.AppendDoubleWord(uint32(*a))
}

//----------------------------------------------------------------------------------------------------------------------
// LocalPrefAttr
//----------------------------------------------------------------------------------------------------------------------

// LocalPrefAttr contains a decoded local preference attribute
type LocalPrefAttr uint32

// Encode encodes a local preference attribute
func (a *LocalPrefAttr) Encode(buf *MsgBuffer) {
	buf.AppendByte(attrFlagTransitive)
	buf.AppendByte(attrTypeMED)
	buf.AppendDoubleWord(uint32(*a))
}

//----------------------------------------------------------------------------------------------------------------------
// AtomicAggregateAttr
//----------------------------------------------------------------------------------------------------------------------

// AtomicAggregateAttr contains a decoded atomic aggregate attribute
type AtomicAggregateAttr struct{}

// Encode encodes an atomic aggregate attribute
func (a *AtomicAggregateAttr) Encode(buf *MsgBuffer) {
	buf.AppendByte(attrFlagTransitive)
	buf.AppendByte(attrTypeAtomicAggregate)
}

//----------------------------------------------------------------------------------------------------------------------
// AggregatorAttr
//----------------------------------------------------------------------------------------------------------------------

// AggregatorAttr contains a decoded aggregator attribute
type AggregatorAttr IPv4Address

// Encode encodes an aggregator attribute
func (a *AggregatorAttr) Encode(buf *MsgBuffer) {
	buf.AppendByte(attrFlagTransitive)
	buf.AppendByte(attrTypeAggregator)
	a.Encode(buf)
}

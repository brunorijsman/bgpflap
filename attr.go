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
func (attr *OriginAttr) Encode(buf *MsgBuffer) {
	buf.AppendByte(attrFlagTransitive)
	buf.AppendByte(attrTypeOrigin)
	buf.AppendByte(1) // Attribute data length
	buf.AppendByte(uint8(*attr))
}

//----------------------------------------------------------------------------------------------------------------------
// AsPathAttr
//----------------------------------------------------------------------------------------------------------------------

// ASPathAttr contains a decoded as-path
type ASPathAttr struct {
	ASPath
}

// Encode encodes an as-path attribute
func (attr *ASPathAttr) Encode(buf *MsgBuffer) {
	len := attr.ASPath.EncodedLen()
	flags := attrFlagTransitive
	if len > 255 {
		flags |= attrFlagExtendedLen
	}
	buf.AppendByte(uint8(flags))
	buf.AppendByte(attrTypeAsPath)
	if len > 255 {
		// TODO: Stil need to check for overflow - doesn't fit in 16 bits
		buf.AppendWord(uint16(len))
	} else {
		buf.AppendByte(uint8(len))
	}
	attr.ASPath.Encode(buf)
}

//----------------------------------------------------------------------------------------------------------------------
// NextHopAttr
//----------------------------------------------------------------------------------------------------------------------

// NextHopAttr contains a decoded next-hop attribute
type NextHopAttr struct {
	IPv4Address
}

// Encode encodes a next-hop attribute
func (attr *NextHopAttr) Encode(buf *MsgBuffer) {
	buf.AppendByte(attrFlagTransitive)
	buf.AppendByte(attrTypeNextHop)
	buf.AppendByte(4) // Attribute data length
	attr.IPv4Address.Encode(buf)
}

//----------------------------------------------------------------------------------------------------------------------
// MEDAttr
//----------------------------------------------------------------------------------------------------------------------

// MEDAttr contains a decoded multi-exit discriminator attribute
type MEDAttr uint32

// Encode encodes a MED attribute
func (attr *MEDAttr) Encode(buf *MsgBuffer) {
	buf.AppendByte(attrFlagOptional)
	buf.AppendByte(attrTypeMED)
	buf.AppendByte(4) // Attribute data length
	buf.AppendDoubleWord(uint32(*attr))
}

//----------------------------------------------------------------------------------------------------------------------
// LocalPrefAttr
//----------------------------------------------------------------------------------------------------------------------

// LocalPrefAttr contains a decoded local preference attribute
type LocalPrefAttr uint32

// Encode encodes a local preference attribute
func (attr *LocalPrefAttr) Encode(buf *MsgBuffer) {
	buf.AppendByte(attrFlagTransitive)
	buf.AppendByte(attrTypeLocalPref)
	buf.AppendByte(4) // Attribute data length
	buf.AppendDoubleWord(uint32(*attr))
}

//----------------------------------------------------------------------------------------------------------------------
// AtomicAggregateAttr
//----------------------------------------------------------------------------------------------------------------------

// AtomicAggregateAttr contains a decoded atomic aggregate attribute
type AtomicAggregateAttr struct{}

// Encode encodes an atomic aggregate attribute
func (attr *AtomicAggregateAttr) Encode(buf *MsgBuffer) {
	buf.AppendByte(attrFlagTransitive)
	buf.AppendByte(attrTypeAtomicAggregate)
	buf.AppendByte(0) // Attribute data length
}

//----------------------------------------------------------------------------------------------------------------------
// AggregatorAttr
//----------------------------------------------------------------------------------------------------------------------

// AggregatorAttr contains a decoded aggregator attribute
type AggregatorAttr struct {
	IPv4Address
}

// Encode encodes an aggregator attribute
func (attr *AggregatorAttr) Encode(buf *MsgBuffer) {
	buf.AppendByte(attrFlagTransitive)
	buf.AppendByte(attrTypeAggregator)
	buf.AppendByte(4) // Attribute data length
	attr.IPv4Address.Encode(buf)
}

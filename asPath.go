package main

//----------------------------------------------------------------------------------------------------------------------
// ASNumber
//----------------------------------------------------------------------------------------------------------------------

// ASNumber represents an automous system number
type ASNumber uint16

// Encode encodes an autonomous system number
func (n *ASNumber) Encode(buf *MsgBuffer) {
	buf.AppendWord(uint16(*n))
}

//----------------------------------------------------------------------------------------------------------------------
// ASPathSegment
//----------------------------------------------------------------------------------------------------------------------

// AS path segment types
const (
	ASPathSegmentTypeSet      = 1
	ASPathSegmentTypeSequence = 2
)

// ASPathSegment contains an autonomous system path segment
type ASPathSegment struct {
	Type      uint8
	ASNumbers []ASNumber
}

// Encode encodes an autonomous system path segment
func (seg *ASPathSegment) Encode(buf *MsgBuffer) {
	buf.AppendByte(seg.Type)
	// TODO: Check for overflow - too many AS numbers in segment to fit in octet?
	buf.AppendByte(uint8(len(seg.ASNumbers)))
	for _, n := range seg.ASNumbers {
		buf.AppendWord(uint16(n))
	}
}

// EncodedLen returns what the length would be if the autonmous system path segment were to be encoded
func (seg *ASPathSegment) EncodedLen() int {
	return 2 + 2*len(seg.ASNumbers)
}

//----------------------------------------------------------------------------------------------------------------------
// ASPath
//----------------------------------------------------------------------------------------------------------------------

// ASPath contains an autonomous system path
type ASPath struct {
	Segments []ASPathSegment
}

// Encode encodes an autonomous system path
func (p *ASPath) Encode(buf *MsgBuffer) {
	for _, s := range p.Segments {
		s.Encode(buf)
	}
}

// EncodedLen returns what the length would be if the autonmous system path were to be encoded
func (p *ASPath) EncodedLen() int {
	l := 0
	for _, s := range p.Segments {
		l += s.EncodedLen()
	}
	return l
}

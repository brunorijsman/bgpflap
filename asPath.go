package main

//------------------------------------------------------------------------------
// ASNumber
//------------------------------------------------------------------------------

// ASNumber represents an automous system number
type ASNumber uint16

// Encode encodes an autonomous system number
func (n *ASNumber) Encode(buf *MsgBuffer) {
	buf.AppendWord(uint16(*n))
}

//------------------------------------------------------------------------------
// ASPathSegment
//------------------------------------------------------------------------------

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
func (s *ASPathSegment) Encode(buf *MsgBuffer) {
	buf.AppendByte(s.Type)
	// TODO: Check for overflow - too many AS numbers in segment to fit in octet?
	buf.AppendByte(uint8(len(s.ASNumbers)))
	for _, n := range s.ASNumbers {
		buf.AppendWord(uint16(n))
	}
}

//------------------------------------------------------------------------------
// ASPath
//------------------------------------------------------------------------------

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

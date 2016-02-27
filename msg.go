package main

const (
	bgpMessageTypeOpen         = 1
	bgpMessageTypeUpdate       = 2
	bgpMessageTypeNotification = 3
	bgpMessageTypeKeepAlive    = 4
)

//----------------------------------------------------------------------------------------------------------------------
// OpenMsg
//----------------------------------------------------------------------------------------------------------------------

// OpenMsg is a decoded open message
type OpenMsg struct {
	AsNumber   uint16
	HoldTime   uint16
	Identifier uint32
}

// Encode encodes an open message
func (m *OpenMsg) Encode() *MsgBuffer {
	buf := NewMsgBuffer()
	buf.AppendMarker()
	lenPos := buf.SkipWord()
	buf.AppendByte(bgpMessageTypeOpen)
	buf.AppendByte(bgpVersion)
	buf.AppendWord(m.AsNumber)
	buf.AppendWord(m.HoldTime)
	buf.AppendDoubleWord(m.Identifier)
	buf.AppendByte(0) // No optional paramaters
	buf.InsertWord(uint16(buf.Len()), lenPos)
	return buf
}

//----------------------------------------------------------------------------------------------------------------------
// OriginAttr
//----------------------------------------------------------------------------------------------------------------------

// PathAttr contains the decoded path attributes
type PathAttr struct {
}

// Encode encodes the path attributes
func (pa *PathAttr) Encode() *MsgBuffer {
	buf := NewMsgBuffer()
	buf.AppendMarker()
	lenPos := buf.SkipWord()
	buf.AppendByte(bgpMessageTypeUpdate)
	// @@@
	buf.InsertWord(uint16(buf.Len()), lenPos)
	return buf
}

//----------------------------------------------------------------------------------------------------------------------
// UpdateMsg
//----------------------------------------------------------------------------------------------------------------------

// UpdateMsg is a decoded update message
type UpdateMsg struct {
	Withdraws      []NLRI
	Attributes     []Attr
	Advertisements []NLRI
}

// Encode encodes an update message
func (msg *UpdateMsg) Encode() *MsgBuffer {

	// Encode header
	buf := NewMsgBuffer()
	buf.AppendMarker()
	totalLenPos := buf.SkipWord()
	buf.AppendByte(bgpMessageTypeUpdate)

	// Encode withdraws
	withdrawsLenPos := buf.SkipWord()
	withdrawsPos := buf.Pos()
	for _, nlri := range msg.Withdraws {
		nlri.Encode(buf)
	}
	withdrawsLen := buf.Pos() - withdrawsPos
	buf.InsertWord(uint16(withdrawsLen), withdrawsLenPos)

	// Encode attributes
	attributesLenPos := buf.SkipWord()
	attributesPos := buf.Pos()
	for _, attr := range msg.Attributes {
		attr.Encode(buf)
	}
	attributesLen := buf.Pos() - attributesPos
	buf.InsertWord(uint16(attributesLen), attributesLenPos)

	// Encode advertisements
	for _, nlri := range msg.Advertisements {
		nlri.Encode(buf)
	}

	// Fixup total length
	buf.InsertWord(uint16(buf.Len()), totalLenPos)
	return buf
}

//----------------------------------------------------------------------------------------------------------------------
// KeepAliveMsg
//----------------------------------------------------------------------------------------------------------------------

// KeepAliveMsg is a decoded keep-alive message
type KeepAliveMsg struct{}

// Encode encodes a keep-alive message
func (m *KeepAliveMsg) Encode() *MsgBuffer {
	buf := NewMsgBuffer()
	buf.AppendMarker()
	buf.AppendWord(19) // Length of keep-alive message
	buf.AppendByte(bgpMessageTypeKeepAlive)
	return buf
}

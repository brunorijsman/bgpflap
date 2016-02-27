package main

import "bytes"

// A MsgBuffer contains an encoded BGP message
// TODO: Make the type private...
type MsgBuffer struct {
	buffer *bytes.Buffer
}

// NewMsgBuffer creates a new message buffer
func NewMsgBuffer() *MsgBuffer {
	msgBuffer := MsgBuffer{
		buffer: new(bytes.Buffer),
	}
	return &msgBuffer
}

// Clear makes the buffer empty
func (m *MsgBuffer) Clear() {
	m.buffer.Reset()
}

// AppendByte appends a byte
func (m *MsgBuffer) AppendByte(b byte) {
	m.buffer.WriteByte(b)
}

// AppendWord appends a word
func (m *MsgBuffer) AppendWord(w uint16) {
	m.buffer.WriteByte(byte(w >> 8))
	m.buffer.WriteByte(byte(w & 0xff))
}

// SkipWord skips a word and returns the position of the skipped word
func (m *MsgBuffer) SkipWord() int {
	pos := m.buffer.Len()
	m.buffer.WriteByte(0)
	m.buffer.WriteByte(0)
	return pos
}

// Pos returns the current position in the buffer
func (m *MsgBuffer) Pos() int {
	return m.buffer.Len()
}

// AppendDoubleWord appends a double word
func (m *MsgBuffer) AppendDoubleWord(dw uint32) {
	m.buffer.WriteByte(byte((dw >> 24) & 0xff))
	m.buffer.WriteByte(byte((dw >> 16) & 0xff))
	m.buffer.WriteByte(byte((dw >> 8) & 0xff))
	m.buffer.WriteByte(byte(dw & 0xff))
}

// InsertWord inserts a word at a given position
func (m *MsgBuffer) InsertWord(w uint16, pos int) {
	m.buffer.Bytes()[pos] = byte(w >> 8)
	m.buffer.Bytes()[pos+1] = byte(w & 0xff)
}

// AppendMarker appends a marker sequence (16 x 255)
func (m *MsgBuffer) AppendMarker() {
	m.buffer.Write([]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255})
}

// Len returns the length
func (m *MsgBuffer) Len() int {
	return m.buffer.Len()
}

// Bytes returns the bytes slice for the contents of the buffer
func (m *MsgBuffer) Bytes() []byte {
	return m.buffer.Bytes()
}

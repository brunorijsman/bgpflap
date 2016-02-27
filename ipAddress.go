package main

import "fmt"

//------------------------------------------------------------------------------
// IPv4Address
//------------------------------------------------------------------------------

// IPv4Address contains an IPv4 address
type IPv4Address struct {
	addr uint32
}

// FromString parses an IPv4 address from a dotted decimal string
func (addr *IPv4Address) FromString(s string) error {
	var a, b, c, d int
	_, err := fmt.Sscanf(s, "%d.%d.%d.%d", &a, &b, &c, &d)
	if err != nil {
		return err
	}
	if (a < 0) || (a > 255) ||
		(b < 0) || (b > 255) ||
		(c < 0) || (c > 255) ||
		(d < 0) || (d > 255) {
		return fmt.Errorf("IPv4Address.FromString: decimal out of range in %s", s)
	}
	return nil
}

// Encode encodes an IPv4 address
func (addr *IPv4Address) Encode(buf *MsgBuffer) {
	buf.AppendDoubleWord(addr.addr)
}

//------------------------------------------------------------------------------
// IPv4Prefix
//------------------------------------------------------------------------------

// IPv4Prefix contains an IPv4 prefix
type IPv4Prefix struct {
	len  uint8
	addr uint32
}

// Encode encodes an IPv4 prefix
func (p *IPv4Prefix) Encode(buf *MsgBuffer) {
	buf.AppendByte(p.len)
	if p.len > 0 {
		buf.AppendByte(byte((p.addr >> 24) & 0xff))
	}
	if p.len > 8 {
		buf.AppendByte(byte((p.addr >> 16) & 0xff))
	}
	if p.len > 16 {
		buf.AppendByte(byte((p.addr >> 8) & 0xff))
	}
	if p.len > 24 {
		buf.AppendByte(byte(p.addr & 0xff))
	}
}

package main

import "fmt"

//----------------------------------------------------------------------------------------------------------------------
// IPv4Address
//----------------------------------------------------------------------------------------------------------------------

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
	addr.addr = uint32((a << 24) | (b << 16) | (c << 8) | d)
	return nil
}

// ToString converts an IPv4 address to a dotted decimal string
func (addr *IPv4Address) ToString() string {
	a := int((addr.addr >> 24) & 0xff)
	b := int((addr.addr >> 16) & 0xff)
	c := int((addr.addr >> 8) & 0xff)
	d := int(addr.addr & 0xff)
	return fmt.Sprintf("%d.%d.%d.%d", a, b, c, d)
}

// Encode encodes an IPv4 address
func (addr *IPv4Address) Encode(buf *MsgBuffer) {
	buf.AppendDoubleWord(addr.addr)
}

//----------------------------------------------------------------------------------------------------------------------
// IPv4Prefix
//----------------------------------------------------------------------------------------------------------------------

// IPv4Prefix contains an IPv4 prefix
type IPv4Prefix struct {
	len  uint8
	addr uint32
}

// FromString parses an IPv4 prefix from a x.x.x.x/x string
func (prefix *IPv4Prefix) FromString(s string) error {
	var a, b, c, d, len int
	_, err := fmt.Sscanf(s, "%d.%d.%d.%d/%d", &a, &b, &c, &d, &len)
	if err != nil {
		return err
	}
	if (a < 0) || (a > 255) ||
		(b < 0) || (b > 255) ||
		(c < 0) || (c > 255) ||
		(d < 0) || (d > 255) {
		return fmt.Errorf("IPv4Prefix.FromString: decimal out of range in %s", s)
	}
	if (len < 0) || (len > 32) {
		return fmt.Errorf("IPv4Prefix.FromString: prefix length out of range in %s",
			s)
	}
	prefix.len = uint8(len)
	prefix.addr = uint32((a << 24) | (b << 16) | (c << 8) | d)
	return nil
}

// Encode encodes an IPv4 prefix
func (prefix *IPv4Prefix) Encode(buf *MsgBuffer) {
	buf.AppendByte(prefix.len)
	if prefix.len > 0 {
		buf.AppendByte(byte((prefix.addr >> 24) & 0xff))
	}
	if prefix.len > 8 {
		buf.AppendByte(byte((prefix.addr >> 16) & 0xff))
	}
	if prefix.len > 16 {
		buf.AppendByte(byte((prefix.addr >> 8) & 0xff))
	}
	if prefix.len > 24 {
		buf.AppendByte(byte(prefix.addr & 0xff))
	}
}

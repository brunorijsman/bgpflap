package main

import (
	"reflect"
	"testing"
)

//----------------------------------------------------------------------------------------------------------------------
// IPv4Address
//----------------------------------------------------------------------------------------------------------------------

func TestIPv4AddressFromString(t *testing.T) {
	var addr IPv4Address
	err := addr.FromString("1.2.3.4")
	if err != nil {
		t.Errorf("FromString(\"1.2.3.4\") returned %v, expected nil", err)
	}
	err = addr.FromString("0.0.0.0")
	if err != nil {
		t.Errorf("FromString(\"0.0.0.0\") returned %v, expected nil", err)
	}
	err = addr.FromString("255.255.255.255")
	if err != nil {
		t.Errorf("FromString(\"255.255.255.255\") returned %v, expected nil", err)
	}
	err = addr.FromString("nonsense")
	if err == nil {
		t.Errorf("FromString(\"nonsense\") returned %v, expected !nil", err)
	}
	err = addr.FromString("1.1.1")
	if err == nil {
		t.Errorf("FromString(\"1.1.1\") returned %v, expected !nil", err)
	}
	err = addr.FromString("1.1.blah.1")
	if err == nil {
		t.Errorf("FromString(\"1.1.blah.1\") returned %v, expected !nil", err)
	}
	err = addr.FromString("1.1.-1.1")
	if err == nil {
		t.Errorf("FromString(\"1.1.-1.1\") returned %v, expected !nil", err)
	}
	err = addr.FromString("1.1.256.1")
	if err == nil {
		t.Errorf("FromString(\"1.1.256.1\") returned %v, expected !nil", err)
	}
}

func TestIPv4AddressToString(t *testing.T) {
	var addr IPv4Address
	err := addr.FromString("1.2.3.4")
	if err != nil {
		t.Fatalf("FromString(\"1.2.3.4\") returned %v, expected nil", err)
	}
	str := addr.ToString()
	if str != "1.2.3.4" {
		t.Errorf("ToString returned %v, expected \"1.2.3.4\"", str)
	}
}

func TestIPv4AddressEncode(t *testing.T) {
	var addr IPv4Address
	err := addr.FromString("1.2.3.4")
	if err != nil {
		t.Fatalf("FromString(\"1.2.3.4\") returned %v, expected nil", err)
	}
	buf := NewMsgBuffer()
	addr.Encode(buf)
	if !reflect.DeepEqual(buf.Bytes(), []byte{1, 2, 3, 4}) {
		t.Errorf("Encode(\"1.2.3.4\") returned %v, expected [1 2 3 4]", buf.Bytes())
	}
}

//----------------------------------------------------------------------------------------------------------------------
// IPv4Prefix
//----------------------------------------------------------------------------------------------------------------------

func TestIPv4PrefixFromString(t *testing.T) {
	var prefix IPv4Prefix
	err := prefix.FromString("1.2.3.0/24")
	if err != nil {
		t.Errorf("FromString(\"1.2.3.0/24\") returned %v, expected nil", err)
	}
	err = prefix.FromString("0.0.0.0/0")
	if err != nil {
		t.Errorf("FromString(\"0.0.0.0/0\") returned %v, expected nil", err)
	}
	err = prefix.FromString("255.255.255.255/32")
	if err != nil {
		t.Errorf("FromString(\"255.255.255.255/32\") returned %v, expected nil", err)
	}
	err = prefix.FromString("nonsense")
	if err == nil {
		t.Errorf("FromString(\"nonsense\") returned %v, expected !nil", err)
	}
	err = prefix.FromString("1.1.blah.1/24")
	if err == nil {
		t.Errorf("FromString(\"1.1.blah.1/24\") returned %v, expected !nil", err)
	}
	err = prefix.FromString("1.1.blah.1")
	if err == nil {
		t.Errorf("FromString(\"1.1.blah.1/24\") returned %v, expected !nil", err)
	}
	err = prefix.FromString("1.1.-1.1/24")
	if err == nil {
		t.Errorf("FromString(\"1.1.-1.1/24\") returned %v, expected !nil", err)
	}
	err = prefix.FromString("1.1.256.1/24")
	if err == nil {
		t.Errorf("FromString(\"1.1.256.1/24\") returned %v, expected !nil", err)
	}
	err = prefix.FromString("1.1.1.1/blah")
	if err == nil {
		t.Errorf("FromString(\"1.1.1.1/blah\") returned %v, expected !nil", err)
	}
	err = prefix.FromString("1.1.1.1/-1")
	if err == nil {
		t.Errorf("FromString(\"1.1.1.1/-1\") returned %v, expected !nil", err)
	}
	err = prefix.FromString("1.1.1.1/33")
	if err == nil {
		t.Errorf("FromString(\"1.1.1.1/33\") returned %v, expected !nil", err)
	}
}

func TestIPv4PrefixToString(t *testing.T) {
	var prefix IPv4Prefix
	err := prefix.FromString("1.2.3.0/24")
	if err != nil {
		t.Fatalf("FromString(\"1.2.3.0/24\") returned %v, expected nil", err)
	}
	str := prefix.ToString()
	if str != "1.2.3.0/24" {
		t.Errorf("ToString returned %v, expected \"1.2.3.0/24\"", str)
	}
}

func TestIPv4PrefixEncode(t *testing.T) {
	var prefix IPv4Prefix
	err := prefix.FromString("0.0.0.0/0")
	if err != nil {
		t.Fatalf("FromString(\"0.0.0.0/0\") returned %v, expected nil", err)
	}
	buf := NewMsgBuffer()
	prefix.Encode(buf)
	if !reflect.DeepEqual(buf.Bytes(), []byte{0}) {
		t.Errorf("Encode(\"0.0.0.0/0\") returned %v, expected [0]", buf.Bytes())
	}
	err = prefix.FromString("130.0.0.0/7")
	if err != nil {
		t.Fatalf("FromString(\"130.0.0.0/7\") returned %v, expected nil", err)
	}
	buf.Clear()
	prefix.Encode(buf)
	if !reflect.DeepEqual(buf.Bytes(), []byte{7, 130}) {
		t.Errorf("Encode(\"130.0.0.0/7\") returned %v, expected [7 130]", buf.Bytes())
	}
	err = prefix.FromString("1.128.0.0/9")
	if err != nil {
		t.Fatalf("FromString(\"1.128.0.0/9\") returned %v, expected nil", err)
	}
	buf.Clear()
	prefix.Encode(buf)
	if !reflect.DeepEqual(buf.Bytes(), []byte{9, 1, 128}) {
		t.Errorf("Encode(\"1.128.0.0/9\") returned %v, expected [9 1 128]", buf.Bytes())
	}
	err = prefix.FromString("1.2.3.0/24")
	if err != nil {
		t.Fatalf("FromString(\"1.2.3.0/24\") returned %v, expected nil", err)
	}
	buf.Clear()
	prefix.Encode(buf)
	if !reflect.DeepEqual(buf.Bytes(), []byte{24, 1, 2, 3}) {
		t.Errorf("Encode(\"1.2.3.0/24\") returned %v, expected [24 1 2 3]", buf.Bytes())
	}
	err = prefix.FromString("255.255.255.255/32")
	if err != nil {
		t.Fatalf("FromString(\"255.255.255.255/32\") returned %v, expected nil", err)
	}
	buf.Clear()
	prefix.Encode(buf)
	if !reflect.DeepEqual(buf.Bytes(), []byte{32, 255, 255, 255, 255}) {
		t.Errorf("Encode(\"255.255.255.255/32\") returned %v, expected [32 255 255 255 255]", buf.Bytes())
	}
}

package main

import "testing"

//------------------------------------------------------------------------------
// IPv4Address
//------------------------------------------------------------------------------

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
		t.Errorf("FromString(\"1.1.blah.1\") returned %v, expected !nil", err)
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

func TestIPv4AddressEncode(t *testing.T) {
	var addr IPv4Address
	err := addr.FromString("1.2.3.4")
	if err != nil {
		t.Fatalf("FromString(\"1.2.3.4\") returned %v, expected nil", err)
	}
	var buf MsgBuffer
	addr.Encode(&buf)
	/* if !reflect.DeepEqual(buf.Bytes(), []byte{1, 2, 3, 4}) {
		t.Errorf("Encode(\"1.2.3.4\") returned %v, expected {1,2,3,4}", buf.Bytes())
	} */
}

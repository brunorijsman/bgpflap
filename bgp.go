package main

import (
	"fmt"
	"log"
	"net"
)

const (
	bgpVersion = 4
	maxMsgSize = 4096
)

func receive(conn net.Conn) {

	buf := make([]byte, 16)

	for {

		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		for _, b := range buf[:n] {
			fmt.Printf("%d ", b)
		}
		fmt.Printf("\n")

	}

}

func sendOpenMsg(conn net.Conn) {
	open := OpenMsg{AsNumber: 1234, HoldTime: 90, Identifier: 0x01010101}
	buf := open.Encode()
	conn.Write(buf.Bytes())
}

func sendKeepAliveMsg(conn net.Conn) {
	keepAlive := KeepAliveMsg{}
	buf := keepAlive.Encode()
	conn.Write(buf.Bytes())
}

func prefixStr(i int) string {
	a := i / (256 * 256)
	rest := i % (256 * 256)
	b := rest / 256
	c := rest % 256
	return fmt.Sprintf("%d.%d.%d.0/24", a, b, c)
}

func sendAdvertiseUpdateMsg(conn net.Conn, i int) {
	var prefix IPv4Prefix
	prefix.FromString(prefixStr(i))
	var origin OriginAttr
	origin = OriginAttrEgp
	var asPath ASPathAttr
	var nextHop NextHopAttr
	nextHop.FromString("192.168.56.1")
	var localPref LocalPrefAttr
	update := UpdateMsg{
		Withdraws:      []NLRI{},
		Attributes:     []Attr{&origin, &asPath, &nextHop, &localPref},
		Advertisements: []NLRI{&prefix},
	}
	buf := update.Encode()
	conn.Write(buf.Bytes())
}

func sendWithdrawUpdateMsg(conn net.Conn, i int) {
	var prefix IPv4Prefix
	prefix.FromString(prefixStr(i))
	update := UpdateMsg{
		Withdraws:      []NLRI{&prefix},
		Attributes:     []Attr{},
		Advertisements: []NLRI{},
	}
	buf := update.Encode()
	conn.Write(buf.Bytes())
}

func main() {

	conn, err := net.Dial("tcp", "192.168.56.2:179")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected")

	go receive(conn)

	sendOpenMsg(conn)

	sendKeepAliveMsg(conn)

	nrPrefixes := 500000

	for {
		for i := 0; i < nrPrefixes; i++ {
			fmt.Printf("[A:%d] ", i)
			sendAdvertiseUpdateMsg(conn, i)
		}
		for i := 0; i < nrPrefixes; i++ {
			fmt.Printf("[W:%d] ", i)
			sendWithdrawUpdateMsg(conn, i)
		}
	}

	/*
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Disconnected")
	*/
}

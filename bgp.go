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

func sendUpdateMsg(conn net.Conn) {
	update := UpdateMsg{
		Withdraws:      []NLRI{},
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

	for {
	}

	/*
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Disconnected")
	*/
}

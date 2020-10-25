package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"

	"github.com/trunghai95/f1-telemetry/model"
)

func main() {
	s, err := net.ResolveUDPAddr("udp4", ":20777")
	if err != nil {
		log.Fatal(err)
	}

	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		log.Fatal(err)
	}

	packet := model.PacketSessionData{
		Header: model.PacketHeader{
			PacketID: uint8(15),
		},
		TrackID: int8(5),
	}

	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, packet)
	n, err := c.Write(b.Bytes())
	log.Println(n, err)
}

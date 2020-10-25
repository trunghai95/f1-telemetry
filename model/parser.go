package model

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/pkg/errors"
	"github.com/trunghai95/f1-telemetry/util"
)

func getPacketStruct(packetID uint8) interface{} {
	switch packetID {
	case PacketIDMotion:
		return &PacketMotionData{}
	case PacketIDSession:
		return &PacketSessionData{}
	case PacketIDLapData:
		return &PacketLapData{}
	case PacketIDEvent:
		return nil
	case PacketIDParticipants:
		return &PacketParticipantsData{}
	case PacketIDCarSetups:
		return &PacketCarSetupData{}
	case PacketIDCarTelemetry:
		return &PacketCarTelemetryData{}
	case PacketIDCarStatus:
		return &PacketCarStatusData{}
	case PacketIDFinalClassification:
		return &PacketFinalClassificationData{}
	case PacketIDLobbyInfo:
		return &PacketLobbyInfoData{}
	}
	return nil
}

// ParsePacket parses the byte array into one of the packet structure
// The result depends on the packet ID in the packet header
func ParsePacket(b []byte) (err error) {
	// Decode the header first to get the packet ID
	header := PacketHeader{}
	bufHeader := bytes.NewBuffer(b[:binary.Size(header)])
	err = binary.Read(bufHeader, binary.LittleEndian, &header)
	if err != nil {
		return errors.Wrap(err, "read header err")
	}

	packet := getPacketStruct(header.PacketID)
	if packet == nil {
		return errors.New("invalid packet ID")
	}
	buf := bytes.NewBuffer(b)
	err = binary.Read(buf, binary.LittleEndian, packet)
	if err != nil {
		return errors.Wrap(err, "read packet err")
	}

	log.Println(util.JSONEncode(packet))
	return
}

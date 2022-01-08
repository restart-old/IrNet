package IrNet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"sync"
)

func init() {
	packets = make(map[Byte]Packet)
}

var packets map[Byte]Packet
var packetsMu sync.RWMutex

func readPacket(id byte, b *bytes.Buffer) (Packet, bool) {
	packetsMu.RLock()
	defer packetsMu.RUnlock()
	pk, ok := packets[Byte(id)]
	if !ok {
		return nil, ok
	}
	pk.Marshal(b)
	return pk, true
}

func RegisterPacket(packet Packet) error {
	packetsMu.RLock()
	if _, ok := packets[packet.ID()]; ok {
		packetsMu.RUnlock()
		return fmt.Errorf("packet with the id %s is already registered", strconv.Itoa(int(packet.ID())))
	}
	packetsMu.RUnlock()
	packetsMu.Lock()
	defer packetsMu.Unlock()
	packets[packet.ID()] = packet
	return nil
}

type Packet interface {
	ID() Byte
	Content() []Type
	Marshal(b *bytes.Buffer)
}

type Ping struct {
	Time Long
}

func (p *Ping) Marshal(b *bytes.Buffer) {
	binary.Read(b, binary.BigEndian, &p.Time)
}

func (*Ping) ID() Byte { return 0x01 }
func (ping *Ping) Content() []Type {
	return []Type{
		Long(ping.Time),
	}
}

type Pong struct {
	PingTime, PongTime Long
}

func (p *Pong) Marshal(b *bytes.Buffer) {
	binary.Read(b, binary.BigEndian, &p.PingTime)
	binary.Read(b, binary.BigEndian, &p.PongTime)
}

func (*Pong) ID() Byte { return 0x02 }
func (pong *Pong) Content() []Type {
	return []Type{
		Long(pong.PingTime),
		Long(pong.PongTime),
	}
}

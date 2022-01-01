package IrNet

import (
	"bytes"
	"net"
)

type Conn struct {
	conn net.Conn
}

func (c Conn) WritePacket(pk Packet) error {
	conn := c.conn
	content, id := pk.Content(), pk.ID()

	var buf bytes.Buffer
	if err := id.write(&buf); err != nil {
		return err
	}
	if err := DefaultMagic.write(&buf); err != nil {
		return err
	}
	for _, t := range content {
		if err := t.write(&buf); err != nil {
			return err
		}
	}
	_, err := conn.Write(buf.Bytes())
	return err
}
func (c Conn) ReadPacket() (Packet, error)

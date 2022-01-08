package IrNet

import (
	"bytes"
	"net"
)

type Conn struct {
	net.Conn
}

func (c Conn) WritePacket(pk Packet) (int, error) {
	content, id := pk.Content(), pk.ID()

	var buf bytes.Buffer
	if err := id.write(&buf); err != nil {
		return 0, err
	}
	for _, t := range content {
		if err := t.write(&buf); err != nil {
			return 0, err
		}
	}
	n, err := c.Write(buf.Bytes())
	return n, err
}
func (c Conn) ReadPacket() (Packet, error) {
	var b [1024]byte
	n, err := c.Read(b[:])
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b[:n][0:])
	pk, ok := readPacket(b[0], buf)
	if !ok {
		return nil, nil
	}
	return pk, nil
}

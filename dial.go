package IrNet

import "net"

func Dial(address string) (Conn, error) {
	var v Conn
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return v, err
	}
	v.Conn = conn
	return v, nil
}

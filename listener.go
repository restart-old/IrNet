package IrNet

import "net"

type Listener struct {
	listener net.Listener
}

func (l Listener) Accept() (Conn, error) {
	var v Conn
	listener := l.listener
	conn, err := listener.Accept()
	if err != nil {
		return v, err
	}
	v.conn = conn
	return v, nil
}
func (l Listener) Close() error   { return l.listener.Close() }
func (l Listener) Addr() net.Addr { return l.listener.Addr() }

func Listen(address string) (Listener, error) {
	listener, err := net.Listen("tcp", address)
	return Listener{listener: listener}, err
}

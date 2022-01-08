package IrNet

import "net"

type Listener struct {
	net.Listener
}

func (l Listener) Accept() (Conn, error) {
	var v Conn
	listener := l.Listener
	conn, err := listener.Accept()
	if err != nil {
		return v, err
	}
	v.Conn = conn
	return v, nil
}
func (l Listener) Close() error   { return l.Listener.Close() }
func (l Listener) Addr() net.Addr { return l.Listener.Addr() }

func Listen(address string) (Listener, error) {
	listener, err := net.Listen("tcp", address)
	return Listener{Listener: listener}, err
}

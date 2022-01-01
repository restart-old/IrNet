package IrNet

type Packet interface {
	ID() Byte
	Content() []Type
}

type Ping struct {
	time Long
}

func (Ping) ID() Byte { return 0x01 }
func (ping Ping) Content() []Type {
	return []Type{
		Long(ping.time),
	}
}

type Pong struct {
	PingTime, PongTime Long
}

func (Pong) ID() Byte { return 0x02 }
func (pong Pong) Content() []Type {
	return []Type{
		Long(pong.PingTime),
		Long(pong.PongTime),
	}
}

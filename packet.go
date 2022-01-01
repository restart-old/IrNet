package IrNet

type Packet interface {
	ID() Byte
	Content() []Type
}

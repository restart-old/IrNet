package IrNet

import (
	"encoding/binary"
	"errors"
	"io"
)

const (
	invalidTypeWrite = "could not write: "
)

// Data types

type Type interface {
	valid() bool
	write(w io.Writer) error
}

// 0 to 255
type Byte byte

func (Byte) valid() bool { return true }
func (b Byte) write(w io.Writer) error {
	if !b.valid() {
		return errors.New(invalidTypeWrite + "invalid Byte value")
	}
	return binary.Write(w, binary.BigEndian, b)
}

// 0 to 1
type Boolean bool

func (Boolean) valid() bool { return true }
func (b Boolean) write(w io.Writer) error {
	if !b.valid() {
		return errors.New(invalidTypeWrite + "invalid Boolean value")
	}
	return binary.Write(w, binary.BigEndian, b)
}

// 	-2^63 to 2^63-1
type Long int64

func (Long) valid() bool { return true }
func (l Long) write(w io.Writer) error {
	if !l.valid() {
		return errors.New(invalidTypeWrite + "invalid Long value")
	}
	return binary.Write(w, binary.BigEndian, l)
}

// 49724e657450726f746f636f6c4279526573746172744655
// 24 Bytes
type Magic []byte

func (m Magic) compare(magic Magic) bool {
	for b, v := range m {
		if magic[b] != v {
			return false
		}
	}
	return true
}

func (m Magic) valid() bool { return m.compare(DefaultMagic) }
func (m Magic) write(w io.Writer) error {
	if !m.valid() {
		return errors.New(invalidTypeWrite + "invalid Magic: does not match the default magic")
	}
	return binary.Write(w, binary.BigEndian, m)
}

type Short int16

func (Short) valid() bool { return true }
func (s Short) write(w io.Writer) error {
	if !s.valid() {
		return errors.New(invalidTypeWrite + "invalid Short value")
	}
	return binary.Write(w, binary.BigEndian, s)
}

type UnsignedShort uint16

func (UnsignedShort) valid() bool { return true }
func (u UnsignedShort) write(w io.Writer) error {
	if !u.valid() {
		return errors.New(invalidTypeWrite + "invalid UnsignedShort value")
	}
	return binary.Write(w, binary.BigEndian, u)
}

type String struct {
	size uint16
	str  string
}

func NewString(s string) String {
	return String{
		size: uint16(len([]byte(s))),
		str:  s,
	}
}

func (s String) valid() bool { return int(s.size) == len(s.str) }
func (s String) write(w io.Writer) error {
	if !s.valid() {
		return errors.New(invalidTypeWrite + "invalid String: the size given does not match the length of the string")
	}
	err := binary.Write(w, binary.BigEndian, s.size)
	if err != nil {
		return err
	}
	return binary.Write(w, binary.BigEndian, []byte(s.str))
}

package megabytes

import "github.com/pivotal-golang/bytefmt"

// Megabytes - Megabytes value.
type Megabytes struct {
	Value uint64
}

// Set - Parses value using
// https://godoc.org/github.com/pivotal-golang/bytefmt#ToMegabytes
// and sets as struct value.
func (mb *Megabytes) Set(value string) (err error) {
	mb.Value, err = bytefmt.ToMegabytes(value)
	return
}

// String - Returns string representation of underlying value.
func (mb *Megabytes) String() string {
	return bytefmt.ByteSize(mb.Value * bytefmt.MEGABYTE)
}

// Bytes - Returns size in bytes.
func (mb *Megabytes) Bytes() uint64 {
	return mb.Value * bytefmt.MEGABYTE
}

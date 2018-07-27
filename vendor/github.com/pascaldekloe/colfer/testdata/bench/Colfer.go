package bench

// This file was generated by colf(1); DO NOT EDIT

import (
	"fmt"
	"io"
	"math"
)

// Colfer configuration attributes
var (
	// ColferSizeMax is the upper limit for serial byte sizes.
	ColferSizeMax = 16 * 1024 * 1024
)

// ColferMax signals an upper limit breach.
type ColferMax string

// Error honors the error interface.
func (m ColferMax) Error() string { return string(m) }

// ColferError signals a data mismatch as as a byte index.
type ColferError int

// Error honors the error interface.
func (i ColferError) Error() string {
	return fmt.Sprintf("colfer: unknown header at byte %d", i)
}

// ColferTail signals data continuation as a byte index.
type ColferTail int

// Error honors the error interface.
func (i ColferTail) Error() string {
	return fmt.Sprintf("colfer: data continuation at byte %d", i)
}

type Colfer struct {
	Key	int64
	Host	string
	Port	int32
	Size	int64
	Hash	uint64
	Ratio	float64
	Route	bool
}

// MarshalTo encodes o as Colfer into buf and returns the number of bytes written.
// If the buffer is too small, MarshalTo will panic.
func (o *Colfer) MarshalTo(buf []byte) int {
	var i int

	if v := o.Key; v != 0 {
		x := uint64(v)
		if v >= 0 {
			buf[i] = 0
		} else {
			x = ^x + 1
			buf[i] = 0 | 0x80
		}
		i++
		for n := 0; n < 8 && x >= 0x80; n++ {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if l := len(o.Host); l != 0 {
		buf[i] = 1
		i++
		x := uint(l)
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		copy(buf[i:], o.Host)
		i += l
	}

	if v := o.Port; v != 0 {
		x := uint32(v)
		if v >= 0 {
			buf[i] = 2
		} else {
			x = ^x + 1
			buf[i] = 2 | 0x80
		}
		i++
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if v := o.Size; v != 0 {
		x := uint64(v)
		if v >= 0 {
			buf[i] = 3
		} else {
			x = ^x + 1
			buf[i] = 3 | 0x80
		}
		i++
		for n := 0; n < 8 && x >= 0x80; n++ {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if x := o.Hash; x >= 1<<49 {
		buf[i] = 4 | 0x80
		buf[i+1], buf[i+2], buf[i+3], buf[i+4] = byte(x>>56), byte(x>>48), byte(x>>40), byte(x>>32)
		buf[i+5], buf[i+6], buf[i+7], buf[i+8] = byte(x>>24), byte(x>>16), byte(x>>8), byte(x)
		i += 9
	} else if x != 0 {
		buf[i] = 4
		i++
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if v := o.Ratio; v != 0.0 {
		buf[i] = 5
		x := math.Float64bits(v)
		buf[i+1], buf[i+2], buf[i+3], buf[i+4] = byte(x>>56), byte(x>>48), byte(x>>40), byte(x>>32)
		buf[i+5], buf[i+6], buf[i+7], buf[i+8] = byte(x>>24), byte(x>>16), byte(x>>8), byte(x)
		i += 9
	}

	if o.Route {
		buf[i] = 6
		i++
	}

	buf[i] = 0x7f
	i++
	return i
}

// MarshalLen returns the Colfer serial byte size.
// The error return option is bench.ColferMax.
func (o *Colfer) MarshalLen() (int, error) {
	l := 1

	if v := o.Key; v != 0 {
		l += 2
		x := uint64(v)
		if v < 0 {
			x = ^x + 1
		}
		for n := 0; n < 8 && x >= 0x80; n++ {
			x >>= 7
			l++
		}
	}

	if x := len(o.Host); x != 0 {
		l += x
		for x >= 0x80 {
			x >>= 7
			l++
		}
		l += 2
	}

	if v := o.Port; v != 0 {
		l += 2
		x := uint32(v)
		if v < 0 {
			x = ^x + 1
		}
		for x >= 0x80 {
			x >>= 7
			l++
		}
	}

	if v := o.Size; v != 0 {
		l += 2
		x := uint64(v)
		if v < 0 {
			x = ^x + 1
		}
		for n := 0; n < 8 && x >= 0x80; n++ {
			x >>= 7
			l++
		}
	}

	if x := o.Hash; x >= 1<<49 {
		l += 9
	} else if x != 0 {
		l += 2
		for x >= 0x80 {
			x >>= 7
			l++
		}
	}

	if o.Ratio != 0.0 {
		l += 9
	}

	if o.Route {
		l++
	}

	if l > ColferSizeMax {
		return l, ColferMax(fmt.Sprintf("colfer: struct testdata/bench.Colfer exceeds %d bytes", ColferSizeMax))
	}
	return l, nil
}

// MarshalBinary encodes o as Colfer conform encoding.BinaryMarshaler.
// The error return option is bench.ColferMax.
func (o *Colfer) MarshalBinary() (data []byte, err error) {
	l, err := o.MarshalLen()
	if err != nil {
		return nil, err
	}
	data = make([]byte, l)
	o.MarshalTo(data)
	return data, nil
}

// Unmarshal decodes data as Colfer and returns the number of bytes read.
// The error return options are io.EOF, bench.ColferError and bench.ColferMax.
func (o *Colfer) Unmarshal(data []byte) (int, error) {
	if len(data) > ColferSizeMax {
		n, err := o.Unmarshal(data[:ColferSizeMax])
		if err == io.EOF {
			return 0, ColferMax(fmt.Sprintf("colfer: struct testdata/bench.Colfer exceeds %d bytes", ColferSizeMax))
		}
		return n, err
	}

	if len(data) == 0 {
		return 0, io.EOF
	}
	header := data[0]
	i := 1

	if header == 0 {
		var x uint64
		for shift := uint(0); ; shift += 7 {
			if i >= len(data) {
				return 0, io.EOF
			}
			b := data[i]
			i++
			if shift == 56 || b < 0x80 {
				x |= uint64(b) << shift
				break
			}
			x |= (uint64(b) & 0x7f) << shift
		}
		o.Key = int64(x)

		if i >= len(data) {
			return 0, io.EOF
		}
		header = data[i]
		i++
	} else if header == 0|0x80 {
		var x uint64
		for shift := uint(0); ; shift += 7 {
			if i >= len(data) {
				return 0, io.EOF
			}
			b := data[i]
			i++
			if shift == 56 || b < 0x80 {
				x |= uint64(b) << shift
				break
			}
			x |= (uint64(b) & 0x7f) << shift
		}
		o.Key = int64(^x + 1)

		if i >= len(data) {
			return 0, io.EOF
		}
		header = data[i]
		i++
	}

	if header == 1 {
		var x uint32
		for shift := uint(0); ; shift += 7 {
			if i >= len(data) {
				return 0, io.EOF
			}
			b := data[i]
			i++
			if b < 0x80 {
				x |= uint32(b) << shift
				break
			}
			x |= (uint32(b) & 0x7f) << shift
		}
		to := i + int(x)
		if to >= len(data) {
			return 0, io.EOF
		}
		o.Host = string(data[i:to])

		header = data[to]
		i = to + 1
	}

	if header == 2 {
		var x uint32
		for shift := uint(0); ; shift += 7 {
			if i >= len(data) {
				return 0, io.EOF
			}
			b := data[i]
			i++
			if b < 0x80 {
				x |= uint32(b) << shift
				break
			}
			x |= (uint32(b) & 0x7f) << shift
		}
		o.Port = int32(x)

		if i >= len(data) {
			return 0, io.EOF
		}
		header = data[i]
		i++
	} else if header == 2|0x80 {
		var x uint32
		for shift := uint(0); ; shift += 7 {
			if i >= len(data) {
				return 0, io.EOF
			}
			b := data[i]
			i++
			if b < 0x80 {
				x |= uint32(b) << shift
				break
			}
			x |= (uint32(b) & 0x7f) << shift
		}
		o.Port = int32(^x + 1)

		if i >= len(data) {
			return 0, io.EOF
		}
		header = data[i]
		i++
	}

	if header == 3 {
		var x uint64
		for shift := uint(0); ; shift += 7 {
			if i >= len(data) {
				return 0, io.EOF
			}
			b := data[i]
			i++
			if shift == 56 || b < 0x80 {
				x |= uint64(b) << shift
				break
			}
			x |= (uint64(b) & 0x7f) << shift
		}
		o.Size = int64(x)

		if i >= len(data) {
			return 0, io.EOF
		}
		header = data[i]
		i++
	} else if header == 3|0x80 {
		var x uint64
		for shift := uint(0); ; shift += 7 {
			if i >= len(data) {
				return 0, io.EOF
			}
			b := data[i]
			i++
			if shift == 56 || b < 0x80 {
				x |= uint64(b) << shift
				break
			}
			x |= (uint64(b) & 0x7f) << shift
		}
		o.Size = int64(^x + 1)

		if i >= len(data) {
			return 0, io.EOF
		}
		header = data[i]
		i++
	}

	if header == 4 {
		var x uint64
		for shift := uint(0); ; shift += 7 {
			if i >= len(data) {
				return 0, io.EOF
			}
			b := data[i]
			i++
			if shift == 56 || b < 0x80 {
				x |= uint64(b) << shift
				break
			}
			x |= (uint64(b) & 0x7f) << shift
		}
		o.Hash = x

		if i >= len(data) {
			return 0, io.EOF
		}
		header = data[i]
		i++
	} else if header == 4|0x80 {
		if i+8 >= len(data) {
			return 0, io.EOF
		}
		o.Hash = uint64(data[i])<<56 | uint64(data[i+1])<<48 | uint64(data[i+2])<<40 | uint64(data[i+3])<<32 | uint64(data[i+4])<<24 | uint64(data[i+5])<<16 | uint64(data[i+6])<<8 | uint64(data[i+7])
		header = data[i+8]
		i += 9
	}

	if header == 5 {
		if i+8 >= len(data) {
			return 0, io.EOF
		}
		x := uint64(data[i])<<56 | uint64(data[i+1])<<48 | uint64(data[i+2])<<40 | uint64(data[i+3])<<32
		x |= uint64(data[i+4])<<24 | uint64(data[i+5])<<16 | uint64(data[i+6])<<8 | uint64(data[i+7])
		o.Ratio = math.Float64frombits(x)

		header = data[i+8]
		i += 9
	}

	if header == 6 {
		o.Route = true

		if i >= len(data) {
			return 0, io.EOF
		}
		header = data[i]
		i++
	}

	if header != 0x7f {
		return 0, ColferError(i - 1)
	}
	return i, nil
}

// UnmarshalBinary decodes data as Colfer conform encoding.BinaryUnmarshaler.
// The error return options are io.EOF, bench.ColferError, bench.ColferTail and bench.ColferMax.
func (o *Colfer) UnmarshalBinary(data []byte) error {
	i, err := o.Unmarshal(data)
	if err != nil {
		return err
	}
	if i != len(data) {
		return ColferTail(i)
	}
	return nil
}
package mystenbcs

// https://github.com/fardream/go-bcs/blob/main/bcs/uleb128.go

import (
	"encoding/binary"
	"fmt"
	"io"
)

// MaxUleb128Length is the max possible number of bytes for an ULEB128 encoded integer.
// Go's widest integer is uint64, so the length is 10.
const MaxUleb128Length = 10

// ULEB128SupportedTypes is a contraint interface that limits the input to
// [ULEB128Encode] and [ULEB128Decode] to signed and unsigned integers.
type ULEB128SupportedTypes interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~int8 | ~int16 | ~int32 | ~int64 | ~int
}

// ULEB128Encode converts an integer into []byte (see [wikipedia] and [bcs])
//
// This reuses [binary.PutUvarint] in standard library.
//
// [wikipedia]: https://en.wikipedia.org/wiki/LEB128
// [bcs]: https://github.com/diem/bcs#uleb128-encoded-integers
func ULEB128Encode[T ULEB128SupportedTypes](input T) []byte {
	result := make([]byte, 10)
	i := binary.PutUvarint(result, uint64(input))
	return result[:i]
}

// ULEB128Decode decodes [io.Reader] into an integer, returns the resulted value, the number of byte read, and a possible error.
//
// [binary.ReadUvarint] is not used here because
//   - it doesn't support returning the number of bytes read.
//   - it accepts only [io.ByteReader], but the recommended way of creating one from [bufio.NewReader] will read more than 1 byte at the
//     to fill the buffer.
func ULEB128Decode[T ULEB128SupportedTypes](r io.Reader) (T, int, error) {
	buf := make([]byte, 1)
	var v, shift T
	var n int
	for n < 10 {
		i, err := r.Read(buf)
		if i == 0 {
			return 0, n, fmt.Errorf("zero read in. possible EOF")
		}
		if err != nil {
			return 0, n, err
		}
		n += i

		d := T(buf[0])
		ld := d & 127
		if (ld<<shift)>>shift != ld {
			return v, n, fmt.Errorf("overflow at index %d: %v", n-1, ld)
		}

		ld <<= shift
		v = ld + v
		if v < ld {
			return v, n, fmt.Errorf("overflow after adding index %d: %v %v", n-1, ld, v)
		}
		if d <= 127 {
			return v, n, nil
		}

		shift += 7
	}

	return 0, n, fmt.Errorf("failed to find most significant bytes after reading %d bytes", n)
}

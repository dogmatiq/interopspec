package uuidpb

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"unicode/utf8"
)

// New returns a new randonly generated UUID.
func New() *UUID {
	var data [16]byte
	_, err := rand.Read(data[:])
	if err != nil {
		panic(err)
	}

	data[6] = (data[6] & 0x0f) | 0x40 // Version 4
	data[8] = (data[8] & 0x3f) | 0x80 // Variant is 10

	return &UUID{
		Upper: binary.BigEndian.Uint64(data[:8]),
		Lower: binary.BigEndian.Uint64(data[8:]),
	}
}

// FromBytes returns a UUID from a byte slice.
func FromBytes(data []byte) *UUID {
	if len(data) != 16 {
		panic("invalid UUID byte slice length")
	}

	return &UUID{
		Upper: binary.BigEndian.Uint64(data[:8]),
		Lower: binary.BigEndian.Uint64(data[8:]),
	}
}

// ToBytes returns the UUID as a byte slice.
func (x *UUID) ToBytes() []byte {
	var data [16]byte
	binary.BigEndian.PutUint64(data[:8], x.Upper)
	binary.BigEndian.PutUint64(data[8:], x.Lower)
	return data[:]
}

// ToString returns the UUID as an RFC 4122 string.
func (x *UUID) ToString() string {
	var str [36]byte
	uuid := x.ToBytes()

	hex.Encode(str[:], uuid[:4])
	str[8] = '-'
	hex.Encode(str[9:], uuid[4:6])
	str[13] = '-'
	hex.Encode(str[14:], uuid[6:8])
	str[18] = '-'
	hex.Encode(str[19:], uuid[8:10])
	str[23] = '-'
	hex.Encode(str[24:], uuid[10:])

	return string(str[:])
}

// Format implements the fmt.Formatter interface, allowing UUIDs to be formatted
// with functions from the fmt package.
func (x *UUID) Format(f fmt.State, verb rune) {
	fmt.Fprintf(
		f,
		formatString(f, verb),
		x.ToString(),
	)
}

// FormatString returns a string representing the fully qualified formatting
// directive captured by the State, followed by the argument verb.
//
// It is a copy of the fmt.FormatString() function in Go v1.20, and may be
// removed once this Go v1.19 support is dropped.
func formatString(state fmt.State, verb rune) string {
	var tmp [16]byte // Use a local buffer.
	b := append(tmp[:0], '%')
	for _, c := range " +-#0" { // All known flags
		if state.Flag(int(c)) { // The argument is an int for historical reasons.
			b = append(b, byte(c))
		}
	}
	if w, ok := state.Width(); ok {
		b = strconv.AppendInt(b, int64(w), 10)
	}
	if p, ok := state.Precision(); ok {
		b = append(b, '.')
		b = strconv.AppendInt(b, int64(p), 10)
	}
	b = utf8.AppendRune(b, verb)
	return string(b)
}

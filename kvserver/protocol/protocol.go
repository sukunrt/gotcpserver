package protocol

import (
	"encoding/binary"
)

func EncodeStrings(s ...string) []byte {
	res := make([]byte, 0)
	for _, v := range s {
		res = append(res, EncodeString(v)...)
	}
	return res
}

func EncodeString(s string) []byte {
	n := uint32(len([]byte(s)))
	header := make([]byte, 4)
	binary.LittleEndian.PutUint32(header, n)
	b := make([]byte, 4+n)
	copy(b, header)
	copy(b[4:], []byte(s))
	return b
}

func DecodeStrings(b []byte) []string {
	i, n := 0, len(b)
	res := make([]string, 0)
	for i < n {
		xn := int(binary.LittleEndian.Uint32(b[i:]))
		println((xn))
		res = append(res, string(b[i+4:i+4+xn]))
		i = i + 4 + xn
	}
	return res
}

func DecodeString(b []byte) string {
	return string(b[4:])
}

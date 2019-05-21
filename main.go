package gocodecs

type Coder struct {
	Encode func(string) []byte
	Decode func([]byte) string
}

func (e *Coder) Init() {
	e.Encode = encodeUTF8
	e.Decode = decodeUTF8
}

func encodeUTF8(s string) []byte {
	var b []byte
	b = []byte(s)
	return b
}

func decodeUTF8(b []byte) string {
	var s string
	s = string(b[:])
	return s
}

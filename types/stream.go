package types

import (
	"bytes"
	"io"
)

type Stream struct {
	r      io.Reader
	w      bytes.Buffer
	Option *Option
}

func NewWriter() *Stream {
	return &Stream{
		w:      bytes.Buffer{},
		Option: &Option{},
	}
}

func NewReader(r io.Reader) *Stream {
	return &Stream{
		r:      r,
		Option: &Option{},
	}
}

func (s *Stream) Bytes() []byte {
	return s.w.Bytes()
}

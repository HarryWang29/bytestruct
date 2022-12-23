package types

import (
	"encoding/hex"
	"fmt"
	"github.com/HarryWang29/bytestruct/errors"
	"reflect"
	"unsafe"
)

type float32Type struct {
	structName string
	fieldName  string
}

func newFloat32Decoder(structName, fieldName string) *float32Type {
	return &float32Type{
		structName: structName,
		fieldName:  fieldName,
	}
}

func (d *float32Type) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
	return &errors.UnmarshalTypeError{
		Value:  fmt.Sprintf("float %s", hex.EncodeToString(buf)),
		Type:   reflect.TypeOf(int8(0)),
		Offset: offset,
	}
}

func (d *float32Type) DecodeStream(s *Stream, p unsafe.Pointer) error {
	u, err := ParseFloat32(s.Option, s.r)
	if err != nil {
		return fmt.Errorf("parse float32 error: %w", err)
	}
	*(*float32)(p) = u
	return nil
}

func (d *float32Type) EncodeStream(stream *Stream, pointer unsafe.Pointer) error {
	return PutFloat32(stream.Option, &stream.w, *(*float32)(pointer))
}

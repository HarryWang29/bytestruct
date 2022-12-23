package types

import (
	"encoding/hex"
	"fmt"
	"github.com/HarryWang29/bytestruct/errors"
	"reflect"
	"unsafe"
)

type float64Type struct {
	structName string
	fieldName  string
}

func newFloat64Decoder(structName, fieldName string) *float64Type {
	return &float64Type{
		structName: structName,
		fieldName:  fieldName,
	}
}

func (d *float64Type) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
	return &errors.UnmarshalTypeError{
		Value:  fmt.Sprintf("float %s", hex.EncodeToString(buf)),
		Type:   reflect.TypeOf(int8(0)),
		Offset: offset,
	}
}

func (d *float64Type) DecodeStream(s *Stream, p unsafe.Pointer) error {
	u, err := ParseFloat64(s.Option, s.r)
	if err != nil {
		return fmt.Errorf("parse float64 error: %w", err)
	}
	*(*float64)(p) = u
	return nil
}

func (d *float64Type) EncodeStream(stream *Stream, pointer unsafe.Pointer) error {
	return PutFloat64(stream.Option, &stream.w, *(*float64)(pointer))
}

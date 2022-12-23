package types

import (
	"encoding/hex"
	"fmt"
	"github.com/HarryWang29/bytestruct/errors"
	"reflect"
	"unsafe"
)

type uint8Type struct {
	structName string
	fieldName  string
}

func newUint8Decoder(structName, fieldName string) *uint8Type {
	return &uint8Type{
		structName: structName,
		fieldName:  fieldName,
	}
}

func (d *uint8Type) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
	return &errors.UnmarshalTypeError{
		Value:  fmt.Sprintf("uint8 %s", hex.EncodeToString(buf)),
		Type:   reflect.TypeOf(uint8(0)),
		Offset: offset,
	}
}

func (d *uint8Type) DecodeStream(s *Stream, p unsafe.Pointer) error {
	u8, err := ParseUint8(s.Option, s.r)
	if err != nil {
		return fmt.Errorf("parse uint8 error: %w", err)
	}
	*(*uint8)(p) = uint8(u8)
	return nil
}

func (d *uint8Type) EncodeStream(stream *Stream, pointer unsafe.Pointer) error {
	return PutUint8(stream.Option, &stream.w, *(*uint8)(pointer))
}

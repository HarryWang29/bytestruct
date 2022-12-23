package types

import (
	"encoding/hex"
	"fmt"
	"github.com/HarryWang29/bytestruct/errors"
	"reflect"
	"unsafe"
)

type uint64Type struct {
	structName string
	fieldName  string
}

func newUint64Decoder(structName, fieldName string) *uint64Type {
	return &uint64Type{
		structName: structName,
		fieldName:  fieldName,
	}
}

func (d *uint64Type) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
	return &errors.UnmarshalTypeError{
		Value:  fmt.Sprintf("uint64 %s", hex.EncodeToString(buf)),
		Type:   reflect.TypeOf(uint64(0)),
		Offset: offset,
	}
}

func (d *uint64Type) DecodeStream(s *Stream, p unsafe.Pointer) error {
	u, err := ParseUint64(s.Option, s.r)
	if err != nil {
		return fmt.Errorf("parse uint64 error: %w", err)
	}
	*(*uint64)(p) = u
	return nil
}

func (d *uint64Type) EncodeStream(stream *Stream, pointer unsafe.Pointer) error {
	return PutUint64(stream.Option, &stream.w, *(*uint64)(pointer))
}

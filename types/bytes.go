package types

import (
	"encoding/hex"
	"fmt"
	"github.com/HarryWang29/bytestruct/errors"
	"reflect"
	"unsafe"
)

type bytesType struct {
	structName string
	fieldName  string
}

func newBytesDecoder(structName, fieldName string) *bytesType {
	return &bytesType{
		structName: structName,
		fieldName:  fieldName,
	}
}

func (d *bytesType) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
	return &errors.UnmarshalTypeError{
		Value:  fmt.Sprintf("bytes %s", hex.EncodeToString(buf)),
		Type:   reflect.TypeOf(""),
		Offset: offset,
		Struct: d.structName,
		Field:  d.fieldName,
	}
}

func (d *bytesType) DecodeStream(s *Stream, p unsafe.Pointer) error {
	bytes, err := ParseBytes(s.Option, s.r)
	if err != nil {
		return fmt.Errorf("parse bytes error: %w", err)
	}
	*(*[]byte)(p) = bytes
	return nil
}

func (d *bytesType) EncodeStream(stream *Stream, pointer unsafe.Pointer) error {
	return PutBytes(stream.Option, &stream.w, *(*[]byte)(pointer))
}

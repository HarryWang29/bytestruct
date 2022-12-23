package types

import (
	"encoding/hex"
	"fmt"
	"github.com/HarryWang29/bytestruct/errors"
	"reflect"
	"unsafe"
)

type stringType struct {
	structName string
	fieldName  string
}

func newStringDecoder(structName, fieldName string) *stringType {
	return &stringType{
		structName: structName,
		fieldName:  fieldName,
	}
}

func (d *stringType) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
	return &errors.UnmarshalTypeError{
		Value:  fmt.Sprintf("string %s", hex.EncodeToString(buf)),
		Type:   reflect.TypeOf(""),
		Offset: offset,
		Struct: d.structName,
		Field:  d.fieldName,
	}
}

func (d *stringType) DecodeStream(s *Stream, p unsafe.Pointer) error {
	str, err := ParseString(s.Option, s.r)
	if err != nil {
		return fmt.Errorf("parse string error: %w", err)
	}
	**(**string)(unsafe.Pointer(&p)) = *(*string)(unsafe.Pointer(&str))
	return nil
}

func (d *stringType) EncodeStream(stream *Stream, pointer unsafe.Pointer) error {
	return PutString(stream.Option, &stream.w, *(*string)(pointer))
}

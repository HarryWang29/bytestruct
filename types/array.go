package types

import (
	"github.com/HarryWang29/bytestruct/errors"
	"github.com/HarryWang29/bytestruct/runtime"
	"reflect"
	"unsafe"
)

type arrayType struct {
	elemType     *runtime.Type
	valueDecoder Codec
	size         uintptr
	alen         int
	structName   string
	fieldName    string
	zeroValue    unsafe.Pointer
}

func newArrayDecoder(dec Codec, elemType *runtime.Type, alen int, structName, fieldName string) *arrayType {
	zeroValuePtr := unsafe_New(elemType)
	zeroValue := **(**unsafe.Pointer)(unsafe.Pointer(&zeroValuePtr))
	return &arrayType{
		valueDecoder: dec,
		elemType:     elemType,
		size:         elemType.Size(),
		alen:         alen,
		structName:   structName,
		fieldName:    fieldName,
		zeroValue:    zeroValue,
	}
}

func (d *arrayType) errNumber(offset int64) *errors.UnmarshalTypeError {
	return &errors.UnmarshalTypeError{
		Value:  "array",
		Type:   reflect.SliceOf(runtime.RType2Type(d.elemType)),
		Struct: d.structName,
		Field:  d.fieldName,
		Offset: offset,
	}
}

func (d *arrayType) DecodeStream(s *Stream, p unsafe.Pointer) error {
	for idx := 0; idx < d.alen; idx++ {
		if err := d.valueDecoder.DecodeStream(s, unsafe.Pointer(uintptr(p)+uintptr(idx)*d.size)); err != nil {
			return err
		}
	}
	return nil
}

func (d *arrayType) EncodeStream(s *Stream, p unsafe.Pointer) error {
	for idx := 0; idx < d.alen; idx++ {
		if err := d.valueDecoder.EncodeStream(s, unsafe.Pointer(uintptr(p)+uintptr(idx)*d.size)); err != nil {
			return err
		}
	}
	return nil
}

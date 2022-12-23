package types

import (
	"encoding/hex"
	"fmt"
	"github.com/HarryWang29/bytestruct/errors"
	"github.com/HarryWang29/bytestruct/runtime"
	"reflect"
	"unsafe"
)

type hiter struct {
	key         unsafe.Pointer
	elem        unsafe.Pointer
	t           unsafe.Pointer
	h           unsafe.Pointer
	buckets     unsafe.Pointer
	bptr        unsafe.Pointer
	overflow    *[]unsafe.Pointer
	oldoverflow *[]unsafe.Pointer
	startBucket uintptr
	offset      uint8
	wrapped     bool
	B           uint8
	i           uint8
	bucket      uintptr
	checkBucket uintptr
}

type mapType struct {
	mapType                 *runtime.Type
	keyType                 *runtime.Type
	valueType               *runtime.Type
	canUseAssignFaststrType bool
	keyDecoder              Codec
	valueDecoder            Codec
	structName              string
	fieldName               string
}

func newMapDecoder(mapTyp *runtime.Type, keyType *runtime.Type, keyDec Codec, valueType *runtime.Type, valueDec Codec, structName, fieldName string) *mapType {
	return &mapType{
		mapType:                 mapTyp,
		keyDecoder:              keyDec,
		keyType:                 keyType,
		canUseAssignFaststrType: canUseAssignFaststrType(keyType, valueType),
		valueType:               valueType,
		valueDecoder:            valueDec,
		structName:              structName,
		fieldName:               fieldName,
	}
}

const (
	mapMaxElemSize = 128
)

// See detail: https://github.com/goccy/go-json/pull/283
func canUseAssignFaststrType(key *runtime.Type, value *runtime.Type) bool {
	indirectElem := value.Size() > mapMaxElemSize
	if indirectElem {
		return false
	}
	return key.Kind() == reflect.String
}

//go:linkname makemap reflect.makemap
func makemap(*runtime.Type, int) unsafe.Pointer

//nolint:golint
//go:linkname mapassign_faststr runtime.mapassign_faststr
//go:noescape
func mapassign_faststr(t *runtime.Type, m unsafe.Pointer, s string) unsafe.Pointer

//go:linkname mapassign reflect.mapassign
//go:noescape
func mapassign(t *runtime.Type, m unsafe.Pointer, k, v unsafe.Pointer)

//go:linkname maplen reflect.maplen
//go:noescape
func maplen(m unsafe.Pointer) int

//go:linkname mapiterinit reflect.mapiterinit
//go:noescape
func mapiterinit(t *runtime.Type, m unsafe.Pointer, it *hiter)

//go:linkname mapiterkey reflect.mapiterkey
//go:noescape
func mapiterkey(it *hiter) (key unsafe.Pointer)

//go:linkname mapiterelem reflect.mapiterelem
//go:noescape
func mapiterelem(it *hiter) (elem unsafe.Pointer)

//go:linkname mapiternext reflect.mapiternext
//go:noescape
func mapiternext(it *hiter)

func (d *mapType) mapassign(t *runtime.Type, m, k, v unsafe.Pointer) {
	if d.canUseAssignFaststrType {
		mapV := mapassign_faststr(t, m, *(*string)(k))
		typedmemmove(d.valueType, mapV, v)
	} else {
		mapassign(t, m, k, v)
	}
}

func (d *mapType) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
	return &errors.UnmarshalTypeError{
		Value:  fmt.Sprintf("map %s", hex.EncodeToString(buf)),
		Type:   reflect.TypeOf(int32(0)),
		Offset: offset,
	}
}

func (d *mapType) DecodeStream(s *Stream, p unsafe.Pointer) error {
	l, err := ParseMapLen(s.Option, s.r)
	if err != nil {
		return fmt.Errorf("parse map len error: %w", err)
	}

	mapValue := *(*unsafe.Pointer)(p)
	if mapValue == nil {
		//todo 空值时应会创建一个新的map，但是无法正常赋值
		mapValue = makemap(d.mapType, 0)
	}
	for i := 0; i < int(l); i++ {
		k := unsafe_New(d.keyType)
		if err := d.keyDecoder.DecodeStream(s, k); err != nil {
			return err
		}

		v := unsafe_New(d.valueType)
		if err := d.valueDecoder.DecodeStream(s, v); err != nil {
			return err
		}

		d.mapassign(d.mapType, mapValue, k, v)
	}
	return nil
}

func (d *mapType) EncodeStream(s *Stream, p unsafe.Pointer) error {
	mapValue := *(*unsafe.Pointer)(p)
	if mapValue == nil {
		return PutUint8(s.Option, &s.w, 0)
	}
	mapLen := maplen(mapValue)
	if err := PutUint8(s.Option, &s.w, uint8(mapLen)); err != nil {
		return err
	}
	var it hiter
	mapiterinit(d.mapType, mapValue, &it)
	for i := 0; i < mapLen; i++ {
		k := mapiterkey(&it)
		if err := d.keyDecoder.EncodeStream(s, k); err != nil {
			return err
		}

		v := mapiterelem(&it)
		if err := d.valueDecoder.EncodeStream(s, v); err != nil {
			return err
		}
		mapiternext(&it)
	}
	return nil
}

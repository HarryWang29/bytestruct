package types

import (
	"github.com/HarryWang29/bytestruct/errors"
	"github.com/HarryWang29/bytestruct/runtime"
	"reflect"
	"sync"
	"unsafe"
)

var (
	sliceTyp = runtime.Type2RType(
		reflect.TypeOf((*sliceHeader)(nil)).Elem(),
	)
	nilSlice = unsafe.Pointer(&sliceHeader{})
)

type sliceType struct {
	elemType          *runtime.Type
	isElemPointerType bool
	valueDecoder      Codec
	size              uintptr
	arrayPool         sync.Pool
	structName        string
	fieldName         string
}

type sliceHeader struct {
	data unsafe.Pointer
	len  int
	cap  int
}

const (
	defaultSliceCapacity = 2
)

func newSliceDecoder(dec Codec, elemType *runtime.Type, size uintptr, structName, fieldName string) *sliceType {
	return &sliceType{
		valueDecoder:      dec,
		elemType:          elemType,
		isElemPointerType: elemType.Kind() == reflect.Ptr || elemType.Kind() == reflect.Map,
		size:              size,
		arrayPool: sync.Pool{
			New: func() interface{} {
				return &sliceHeader{
					data: newArray(elemType, defaultSliceCapacity),
					len:  0,
					cap:  defaultSliceCapacity,
				}
			},
		},
		structName: structName,
		fieldName:  fieldName,
	}
}

func (d *sliceType) newSlice(src *sliceHeader) *sliceHeader {
	slice := d.arrayPool.Get().(*sliceHeader)
	if src.len > 0 {
		// copy original elem
		if slice.cap < src.cap {
			data := newArray(d.elemType, src.cap)
			slice = &sliceHeader{data: data, len: src.len, cap: src.cap}
		} else {
			slice.len = src.len
		}
		copySlice(d.elemType, *slice, *src)
	} else {
		slice.len = 0
	}
	return slice
}

func (d *sliceType) releaseSlice(p *sliceHeader) {
	d.arrayPool.Put(p)
}

//go:linkname copySlice reflect.typedslicecopy
func copySlice(elemType *runtime.Type, dst, src sliceHeader) int

//go:linkname newArray reflect.unsafe_NewArray
func newArray(*runtime.Type, int) unsafe.Pointer

//go:linkname typedmemmove reflect.typedmemmove
func typedmemmove(t *runtime.Type, dst, src unsafe.Pointer)

func (d *sliceType) errNumber(offset int64) *errors.UnmarshalTypeError {
	return &errors.UnmarshalTypeError{
		Value:  "number",
		Type:   reflect.SliceOf(runtime.RType2Type(d.elemType)),
		Struct: d.structName,
		Field:  d.fieldName,
		Offset: offset,
	}
}

func (d *sliceType) DecodeStream(s *Stream, p unsafe.Pointer) error {
	u8, err := ParseSliceLen(s.Option, s.r)
	if err != nil {
		return err
	}
	slice := d.newSlice((*sliceHeader)(p))
	idx := 0
	srcLen := slice.len
	capacity := slice.cap
	data := slice.data
	for ; idx < int(u8); idx++ {
		if capacity <= idx {
			src := sliceHeader{data: data, len: idx, cap: capacity}
			capacity *= 2
			data = newArray(d.elemType, capacity)
			dst := sliceHeader{data: data, len: idx, cap: capacity}
			copySlice(d.elemType, dst, src)
		}
		ep := unsafe.Pointer(uintptr(data) + uintptr(idx)*d.size)
		if srcLen <= idx {
			if d.isElemPointerType {
				**(**unsafe.Pointer)(unsafe.Pointer(&ep)) = nil // initialize elem pointer
			} else {
				// assign new element to the slice
				typedmemmove(d.elemType, ep, unsafe_New(d.elemType))
			}
			if err := d.valueDecoder.DecodeStream(s, ep); err != nil {
				return err
			}
		}
	}

	slice.cap = capacity
	slice.len = idx
	slice.data = data
	dst := (*sliceHeader)(p)
	dst.len = idx
	if dst.len > dst.cap {
		dst.data = newArray(d.elemType, dst.len)
		dst.cap = dst.len
	}
	copySlice(d.elemType, *dst, *slice)
	d.releaseSlice(slice)

	return nil
}

func (d *sliceType) EncodeStream(s *Stream, p unsafe.Pointer) error {
	vLen := (*sliceHeader)(p).len
	if err := PutUint8(s.Option, &s.w, uint8(vLen)); err != nil {
		return err
	}
	for i := 0; i < vLen; i++ {
		ep := unsafe.Pointer(uintptr((*sliceHeader)(p).data) + uintptr(i)*d.size)
		if err := d.valueDecoder.EncodeStream(s, ep); err != nil {
			return err
		}
	}
	return nil
}

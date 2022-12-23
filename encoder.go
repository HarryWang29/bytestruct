package bytestruct

import (
	"github.com/HarryWang29/bytestruct/runtime"
	"github.com/HarryWang29/bytestruct/types"
	"unsafe"
)

func (r *ByteStruct) Marshal(v interface{}) ([]byte, error) {
	r.s = types.NewWriter()
	r.s.Option = r.option
	header := (*emptyInterface)(unsafe.Pointer(&v))
	typ := header.typ
	ptr := uintptr(header.ptr)
	typeptr := uintptr(unsafe.Pointer(typ))
	// noescape trick for header.typ ( reflect.*rtype )
	copiedType := *(**runtime.Type)(unsafe.Pointer(&typeptr))

	if err := validateType(copiedType, ptr); err != nil {
		return nil, err
	}

	dec, err := types.CompileToGetDecoder(typ)
	if err != nil {
		return nil, err
	}
	s := r.s
	if err := dec.EncodeStream(s, header.ptr); err != nil {
		return nil, err
	}
	return s.Bytes(), nil
}

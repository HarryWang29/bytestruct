package bytestruct

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/HarryWang29/bytestruct/errors"
	"github.com/HarryWang29/bytestruct/runtime"
	"github.com/HarryWang29/bytestruct/types"
	"io"
	"reflect"
	"unsafe"
)

func (r *ByteStruct) UnmarshalFromBytes(data []byte, v interface{}) error {
	return r.Unmarshal(bytes.NewReader(data), v)
}

func (r *ByteStruct) UnmarshalFromHexString(data string, v interface{}) error {
	buf, err := hex.DecodeString(data)
	if err != nil {
		return fmt.Errorf("decode hex string error: %w", err)
	}
	return r.UnmarshalFromBytes(buf, v)
}

func (r *ByteStruct) Unmarshal(reader io.ReadSeeker, v interface{}) error {
	r.s = types.NewReader(reader)
	r.s.Option = r.option
	header := (*emptyInterface)(unsafe.Pointer(&v))
	typ := header.typ
	ptr := uintptr(header.ptr)
	typeptr := uintptr(unsafe.Pointer(typ))
	// noescape trick for header.typ ( reflect.*rtype )
	copiedType := *(**runtime.Type)(unsafe.Pointer(&typeptr))

	if err := validateType(copiedType, ptr); err != nil {
		return err
	}

	dec, err := types.CompileToGetDecoder(typ)
	if err != nil {
		return err
	}
	s := r.s
	if err := dec.DecodeStream(s, header.ptr); err != nil {
		return err
	}
	return nil
}

func validateType(typ *runtime.Type, p uintptr) error {
	if typ == nil || typ.Kind() != reflect.Ptr || p == 0 {
		return &errors.UnmarshalTypeError{Type: runtime.RType2Type(typ)}
	}
	return nil
}

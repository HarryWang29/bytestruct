package bytestruct

import (
    "github.com/HarryWang29/bytestruct/runtime"
    "github.com/HarryWang29/bytestruct/types"
    "reflect"
    "unsafe"
)

func (r *ByteStruct) Unmarshal(v interface{}, optFuncs ...OptionFunc) error {
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
    if err := r.s.PrepareForDecode(); err != nil {
        return err
    }
    s := r.s
    for _, optFunc := range optFuncs {
        optFunc(s.Option)
    }
    if err := dec.DecodeStream(s, header.ptr); err != nil {
        return err
    }
    s.Reset()
    return nil
}

func validateType(typ *runtime.Type, p uintptr) error {
    if typ == nil || typ.Kind() != reflect.Ptr || p == 0 {
        return &InvalidUnmarshalError{Type: runtime.RType2Type(typ)}
    }
    return nil
}

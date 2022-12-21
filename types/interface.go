package types

import (
    "github.com/HarryWang29/bytestruct/errors"
    "github.com/HarryWang29/bytestruct/runtime"
    "reflect"
    "unsafe"
)

type interfaceType struct {
    typ        *runtime.Type
    structName string
    fieldName  string
}

func newEmptyInterfaceDecoder(structName, fieldName string) *interfaceType {
    ifaceDecoder := &interfaceType{
        typ:        emptyInterfaceType,
        structName: structName,
        fieldName:  fieldName,
    }
    return ifaceDecoder
}

func newInterfaceDecoder(typ *runtime.Type, structName, fieldName string) *interfaceType {
    return &interfaceType{
        typ:        typ,
        structName: structName,
        fieldName:  fieldName,
    }
}

var (
    emptyInterfaceType = runtime.Type2RType(reflect.TypeOf((*interface{})(nil)).Elem())
    EmptyInterfaceType = emptyInterfaceType
    interfaceMapType   = runtime.Type2RType(
        reflect.TypeOf((*map[interface{}]interface{})(nil)).Elem(),
    )
    stringTyp = runtime.Type2RType(
        reflect.TypeOf(""),
    )
)

func (d *interfaceType) DecodeStream(s *Stream, p unsafe.Pointer) error {
    runtimeInterfaceValue := *(*interface{})(unsafe.Pointer(&emptyInterface{
        typ: d.typ,
        ptr: p,
    }))
    rv := reflect.ValueOf(runtimeInterfaceValue)
    iface := rv.Interface()
    ifaceHeader := (*emptyInterface)(unsafe.Pointer(&iface))
    typ := ifaceHeader.typ
    //if ifaceHeader.ptr == nil || d.typ == typ || typ == nil {
    //	// concrete type is empty interface
    //	return d.DecodeStream(EmptyInterface(s, depth, p), depth)
    //}
    //if typ.Kind() == reflect.Ptr && typ.Elem() == d.typ || typ.Kind() != reflect.Ptr {
    //	return d.DecodeStream(EmptyInterface(s, depth, p), depth)
    //}
    decoder, err := CompileToGetDecoder(typ)
    if err != nil {
        return err
    }
    return decoder.DecodeStream(s, ifaceHeader.ptr)
}

func (d *interfaceType) errUnmarshalType(typ reflect.Type, offset int64) *errors.UnmarshalTypeError {
    return &errors.UnmarshalTypeError{
        Value:  typ.String(),
        Type:   typ,
        Offset: offset,
        Struct: d.structName,
        Field:  d.fieldName,
    }
}

func (d *interfaceType) EncodeStream(s *Stream, p unsafe.Pointer) error {
    runtimeInterfaceValue := *(*interface{})(unsafe.Pointer(&emptyInterface{
        typ: d.typ,
        ptr: p,
    }))
    rv := reflect.ValueOf(runtimeInterfaceValue)
    iface := rv.Interface()
    ifaceHeader := (*emptyInterface)(unsafe.Pointer(&iface))
    typ := ifaceHeader.typ
    //if ifaceHeader.ptr == nil || d.typ == typ || typ == nil {
    //	// concrete type is empty interface
    //	return d.DecodeStream(EmptyInterface(s, depth, p), depth)
    //}
    //if typ.Kind() == reflect.Ptr && typ.Elem() == d.typ || typ.Kind() != reflect.Ptr {
    //	return d.DecodeStream(EmptyInterface(s, depth, p), depth)
    //}
    decoder, err := CompileToGetDecoder(typ)
    if err != nil {
        return err
    }
    return decoder.EncodeStream(s, ifaceHeader.ptr)
}

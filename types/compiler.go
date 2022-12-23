package types

import (
	"fmt"
	"github.com/HarryWang29/bytestruct/errors"
	"github.com/HarryWang29/bytestruct/runtime"
	"reflect"
	"unicode"
	"unsafe"
)

var (
	typeAddr      *runtime.TypeAddr
	cachedDecoder []Codec
)

func init() {
	typeAddr = runtime.AnalyzeTypeAddr()
	if typeAddr == nil {
		typeAddr = &runtime.TypeAddr{}
	}
	cachedDecoder = make([]Codec, typeAddr.AddrRange>>typeAddr.AddrShift+1)
}

func compileHead(typ *runtime.Type, structTypeToDecoder map[uintptr]Codec) (Codec, error) {
	return compile(typ.Elem(), "", "", structTypeToDecoder)
}

func compile(typ *runtime.Type, structName, fieldName string, structTypeToDecoder map[uintptr]Codec) (Codec, error) {
	switch typ.Kind() {
	case reflect.Ptr:
		return compilePtr(typ, structName, fieldName, structTypeToDecoder)
	case reflect.Struct:
		return compileStruct(typ, structName, fieldName, structTypeToDecoder)
	case reflect.Slice:
		elem := typ.Elem()
		if elem.Kind() == reflect.Uint8 {
			return compileBytes(structName, fieldName)
		}
		return compileSlice(typ, structName, fieldName, structTypeToDecoder)
	case reflect.Array:
		return compileArray(typ, structName, fieldName, structTypeToDecoder)
	case reflect.Map:
		return compileMap(typ, structName, fieldName, structTypeToDecoder)
	case reflect.Interface:
		return compileInterface(typ, structName, fieldName, structTypeToDecoder)
	case reflect.Int8:
		return compileInt8(structName, fieldName)
	case reflect.Int16:
		return compileInt16(structName, fieldName)
	case reflect.Int32:
		return compileInt32(structName, fieldName)
	case reflect.Int64:
		return compileInt64(structName, fieldName)
	case reflect.Uint8:
		return compileUint8(structName, fieldName)
	case reflect.Uint16:
		return compileUint16(structName, fieldName)
	case reflect.Uint32:
		return compileUint32(structName, fieldName)
	case reflect.Uint64:
		return compileUint64(structName, fieldName)
	case reflect.String:
		return compileString(structName, fieldName)
	case reflect.Bool:
		return compileBool(structName, fieldName)
	case reflect.Float32:
		return compileFloat32(structName, fieldName)
	case reflect.Float64:
		return compileFloat64(structName, fieldName)
	}
	return nil, &errors.UnmarshalTypeError{Type: runtime.RType2Type(typ)}
}

//func compileInt(structName, fieldName string) (Codec, error) {
//    return newInt32Decoder(structName, fieldName, func(p unsafe.Pointer, v int32) {
//        *(*int)(p) = int(v)
//    }), nil
//}

func compileInt8(structName, fieldName string) (Codec, error) {
	return newInt8Decoder(structName, fieldName), nil
}

func compileInt16(structName, fieldName string) (Codec, error) {
	return newInt16Decoder(structName, fieldName), nil
}

func compileInt32(structName, fieldName string) (Codec, error) {
	return newInt32Decoder(structName, fieldName), nil
}

func compileInt64(structName, fieldName string) (Codec, error) {
	return newInt64Decoder(structName, fieldName), nil
}

//func compileUint(structName, fieldName string) (Codec, error) {
//    return newUint32Decoder(structName, fieldName, func(p unsafe.Pointer, v uint32) {
//        *(*uint)(p) = uint(v)
//    }), nil
//}

func compileUint8(structName, fieldName string) (Codec, error) {
	return newUint8Decoder(structName, fieldName), nil
}

func compileUint16(structName, fieldName string) (Codec, error) {
	return newUint16Decoder(structName, fieldName), nil
}

func compileUint32(structName, fieldName string) (Codec, error) {
	return newUint32Decoder(structName, fieldName), nil
}

func compileUint64(structName, fieldName string) (Codec, error) {
	return newUint64Decoder(structName, fieldName), nil
}

func compileBool(structName, fieldName string) (Codec, error) {
	return newBoolDecoder(structName, fieldName), nil
}

func compileString(structName, fieldName string) (Codec, error) {
	return newStringDecoder(structName, fieldName), nil
}

func compileFloat32(structName, fieldName string) (Codec, error) {
	return newFloat32Decoder(structName, fieldName), nil
}

func compileFloat64(structName, fieldName string) (Codec, error) {
	return newFloat64Decoder(structName, fieldName), nil
}

func compileInterface(typ *runtime.Type, structName, fieldName string, structTypeToDecoder map[uintptr]Codec) (Codec, error) {
	return newInterfaceDecoder(typ, structName, fieldName), nil
}

func compileBytes(structName, fieldName string) (Codec, error) {
	return newBytesDecoder(structName, fieldName), nil
}

func compileSlice(typ *runtime.Type, structName, fieldName string, structTypeToDecoder map[uintptr]Codec) (Codec, error) {
	elem := typ.Elem()
	decoder, err := compile(elem, structName, fieldName, structTypeToDecoder)
	if err != nil {
		return nil, err
	}
	return newSliceDecoder(decoder, elem, elem.Size(), structName, fieldName), nil
}

func compilePtr(typ *runtime.Type, structName, fieldName string, structTypeToDecoder map[uintptr]Codec) (Codec, error) {
	dec, err := compile(typ.Elem(), structName, fieldName, structTypeToDecoder)
	if err != nil {
		return nil, err
	}
	return newPtrDecoder(dec, typ.Elem(), structName, fieldName), nil
}

func compileArray(typ *runtime.Type, structName, fieldName string, structTypeToDecoder map[uintptr]Codec) (Codec, error) {
	elem := typ.Elem()
	decoder, err := compile(elem, structName, fieldName, structTypeToDecoder)
	if err != nil {
		return nil, err
	}
	return newArrayDecoder(decoder, elem, typ.Len(), structName, fieldName), nil
}

func compileMap(typ *runtime.Type, structName, fieldName string, structTypeToDecoder map[uintptr]Codec) (Codec, error) {
	keyDec, err := compile(typ.Key(), structName, fieldName, structTypeToDecoder)
	if err != nil {
		return nil, err
	}
	valueDec, err := compile(typ.Elem(), structName, fieldName, structTypeToDecoder)
	if err != nil {
		return nil, err
	}
	return newMapDecoder(typ, typ.Key(), keyDec, typ.Elem(), valueDec, structName, fieldName), nil
}

func compileStruct(typ *runtime.Type, structName, fieldName string, structTypeToDecoder map[uintptr]Codec) (Codec, error) {
	fieldNum := typ.NumField()
	var fieldVec []*structFieldSet
	typeptr := uintptr(unsafe.Pointer(typ))
	if dec, exists := structTypeToDecoder[typeptr]; exists {
		return dec, nil
	}
	structDec := newStructDecoder(structName, fieldName, fieldVec)
	structTypeToDecoder[typeptr] = structDec
	structName = typ.Name()
	var allFields []*structFieldSet
	for i := 0; i < fieldNum; i++ {
		field := typ.Field(i)
		if runtime.IsIgnoredStructField(field) {
			continue
		}
		isUnexportedField := unicode.IsLower([]rune(field.Name)[0])
		tag := runtime.StructTagFromField(field)
		dec, err := compile(runtime.Type2RType(field.Type), structName, field.Name, structTypeToDecoder)
		if err != nil {
			return nil, err
		}
		if field.Anonymous && !tag.IsTaggedKey {
			if stDec, ok := dec.(*structType); ok {
				if runtime.Type2RType(field.Type) == typ {
					// recursive definition
					continue
				}
				for _, v := range stDec.fieldVec {
					fieldSet := &structFieldSet{
						dec:         v.dec,
						offset:      field.Offset + v.offset,
						isTaggedKey: v.isTaggedKey,
					}
					allFields = append(allFields, fieldSet)
				}
			} else if pdec, ok := dec.(*ptrType); ok {
				contentDec := pdec.contentDecoder()
				if pdec.typ == typ {
					// recursive definition
					continue
				}
				var fieldSetErr error
				if isUnexportedField {
					fieldSetErr = fmt.Errorf(
						"json: cannot set embedded pointer to unexported struct: %v",
						field.Type.Elem(),
					)
				}
				if dec, ok := contentDec.(*structType); ok {
					for _, v := range dec.fieldVec {
						fieldSet := &structFieldSet{
							dec:         newAnonymousFieldDecoder(pdec.typ, v.offset, v.dec),
							offset:      field.Offset,
							isTaggedKey: v.isTaggedKey,
							err:         fieldSetErr,
						}
						allFields = append(allFields, fieldSet)
					}
				} else {
					fieldSet := &structFieldSet{
						dec:         pdec,
						offset:      field.Offset,
						isTaggedKey: tag.IsTaggedKey,
					}
					allFields = append(allFields, fieldSet)
				}
			} else {
				fieldSet := &structFieldSet{
					dec:         dec,
					offset:      field.Offset,
					isTaggedKey: tag.IsTaggedKey,
				}
				allFields = append(allFields, fieldSet)
			}
		} else {
			fieldSet := &structFieldSet{
				dec:         dec,
				offset:      field.Offset,
				isTaggedKey: tag.IsTaggedKey,
			}
			allFields = append(allFields, fieldSet)
		}
	}
	delete(structTypeToDecoder, typeptr)
	structDec.fieldVec = allFields
	return structDec, nil
}

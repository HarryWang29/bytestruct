package errors

import (
	"fmt"
	"reflect"
)

type UnmarshalTypeError struct {
	Value  string
	Type   reflect.Type
	Offset int64
	Struct string
	Field  string
}

func (e *UnmarshalTypeError) Error() string {
	if e.Struct != "" || e.Field != "" {
		return fmt.Sprintf("bytestruct: cannot unmarshal %s into Go struct field %s.%s of type %s",
			e.Value, e.Struct, e.Field, e.Type,
		)
	}
	return fmt.Sprintf("bytestruct: cannot unmarshal %s into Go value of type %s", e.Value, e.Type)
}

package bytestruct

import (
	"bytes"
	"encoding/binary"
	"github.com/HarryWang29/bytestruct/types"
	"io"
)

func MarshalLE(v interface{}, optionFunc ...OptionFunc) ([]byte, error) {
	return Marshal(binary.LittleEndian, v, optionFunc...)
}

func UnmarshalLE(data []byte, v interface{}, optionFunc ...OptionFunc) error {
	return Unmarshal(bytes.NewReader(data), binary.LittleEndian, v, optionFunc...)
}

func MarshalBE(v interface{}, optionFunc ...OptionFunc) ([]byte, error) {
	return Marshal(binary.BigEndian, v, optionFunc...)
}

func UnmarshalBE(data []byte, v interface{}, optionFunc ...OptionFunc) error {
	return Unmarshal(bytes.NewReader(data), binary.BigEndian, v, optionFunc...)
}

func Marshal(order binary.ByteOrder, v interface{}, optionFunc ...OptionFunc) ([]byte, error) {
	return NewWriter(order, optionFunc...).Marshal(v)
}

func Unmarshal(data io.ReadSeeker, order binary.ByteOrder, v interface{}, optionFunc ...OptionFunc) error {
	return NewReader(order, optionFunc...).Unmarshal(data, v)
}

func SetParseUint8(f func(*types.Option, io.Reader) (uint64, error)) {
	types.ParseUint8 = f
}
func SetParseUint16(f func(*types.Option, io.Reader) (uint64, error)) {
	types.ParseUint16 = f
}
func SetParseUint32(f func(*types.Option, io.Reader) (uint64, error)) {
	types.ParseUint32 = f
}
func SetParseUint64(f func(*types.Option, io.Reader) (uint64, error)) {
	types.ParseUint64 = f
}
func SetParseInt8(f func(*types.Option, io.Reader) (uint64, error)) {
	types.ParseInt8 = f
}
func SetParseInt16(f func(*types.Option, io.Reader) (uint64, error)) {
	types.ParseInt16 = f
}
func SetParseInt32(f func(*types.Option, io.Reader) (uint64, error)) {
	types.ParseInt32 = f
}
func SetParseInt64(f func(*types.Option, io.Reader) (uint64, error)) {
	types.ParseInt64 = f
}
func SetParseBool(f func(*types.Option, io.Reader) (uint64, error)) {
	types.ParseBool = f
}
func SetParseString(f func(*types.Option, io.Reader) (string, error)) {
	types.ParseString = f
}

func SetParseStringLen(f func(*types.Option, io.Reader) (uint64, error)) {
	types.ParseStringLen = f
}
func SetParseMapLen(f func(*types.Option, io.Reader) (uint64, error)) {
	types.ParseMapLen = f
}
func SetParseSliceLen(f func(*types.Option, io.Reader) (uint64, error)) {
	types.ParseSliceLen = f
}
func SetParseBytes(f func(*types.Option, io.Reader) ([]byte, error)) {
	types.ParseBytes = f
}

func SetPutUint8(f func(*types.Option, *bytes.Buffer, interface{}) error) {
	types.PutUint8 = f
}
func SetPutUint16(f func(*types.Option, *bytes.Buffer, interface{}) error) {
	types.PutUint16 = f
}
func SetPutUint32(f func(*types.Option, *bytes.Buffer, interface{}) error) {
	types.PutUint32 = f
}
func SetPutUint64(f func(*types.Option, *bytes.Buffer, interface{}) error) {
	types.PutUint64 = f
}
func SetPutInt8(f func(*types.Option, *bytes.Buffer, interface{}) error) {
	types.PutInt8 = f
}
func SetPutInt16(f func(*types.Option, *bytes.Buffer, interface{}) error) {
	types.PutInt16 = f
}
func SetPutInt32(f func(*types.Option, *bytes.Buffer, interface{}) error) {
	types.PutInt32 = f
}
func SetPutInt64(f func(*types.Option, *bytes.Buffer, interface{}) error) {
	types.PutInt64 = f
}
func SetPutBool(f func(*types.Option, *bytes.Buffer, interface{}) error) {
	types.PutBool = f
}
func SetPutString(f func(*types.Option, *bytes.Buffer, string) error) {
	types.PutString = f
}
func SetPutLen(f func(*types.Option, *bytes.Buffer, interface{}) error) {
	types.PutLen = f
}
func SetPutBytes(f func(*types.Option, *bytes.Buffer, []byte) error) {
	types.PutBytes = f
}

type OptionFunc func(option *types.Option)

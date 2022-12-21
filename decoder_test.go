package bytestruct

import (
    "encoding/binary"
    "encoding/hex"
    "github.com/HarryWang29/bytestruct/types"
    "testing"
)

func TestReader_Unmarshal(t *testing.T) {
    //s := struct {
    //	ID        uint32
    //	Timestamp uint32
    //	UserID    uint64
    //}{}
    var r uint8
    s := "b68b01003e19a0634ab9b93b00000000"
    reader, err := NewFromHexString(s, binary.LittleEndian, false)
    if err != nil {
        t.Fatal(err)
    }
    err = reader.Unmarshal(&r)
    if err != nil {
        t.Fatal(err)
    }
    t.Logf("%+v", r)
    err = reader.Unmarshal(&r)
    if err != nil {
        t.Fatal(err)
    }
    t.Logf("%+v", r)
}

func TestReader_Unmarshal_bool(t *testing.T) {
    var r bool
    reader := NewFromBytes([]byte{0xa1}, binary.LittleEndian, false)
    err := reader.Unmarshal(&r)
    if err != nil {
        t.Fatal(err)
    }
    t.Logf("%+v", r)
}

func TestReader_Unmarshal_string_2_zero(t *testing.T) {
    var r string
    reader := NewFromBytes([]byte{0x61, 0x62, 0x63, 0x00, 0x64}, binary.LittleEndian, false)
    reader.s.Option.Flags = types.OptionFlagsString2Zero
    err := reader.Unmarshal(&r)
    if err != nil {
        t.Fatal(err)
    }
    if r != "abc" {
        t.Fatal("decode string error: ", r)
    }
}

func TestReader_Unmarshal_len_string(t *testing.T) {
    var r string
    reader := NewFromBytes([]byte{0x3, 0x61, 0x62, 0x63, 0x64}, binary.LittleEndian, false)
    err := reader.Unmarshal(&r)
    if err != nil {
        t.Fatal(err)
    }
    if r != "abc" {
        t.Fatal("decode string error: ", r)
    }
}

func TestReader_Unmarshal_slice(t *testing.T) {
    var v []int32
    s := NewFromBytes([]byte{0x02, 0xa1, 0xa2, 0x00, 0x00, 0xa3, 0xa4, 0, 0}, binary.LittleEndian, false)
    err := s.Unmarshal(&v)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(v)

    var vv []int64
    s = NewFromBytes([]byte{0x02, 0xa1, 0xa2, 0x00, 0x00, 0xa3, 0xa4, 0, 0}, binary.LittleEndian, false)
    err = s.Unmarshal(&vv)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(vv)

    var strs []string
    b, _ := hex.DecodeString("020361626303646566")
    s = NewFromBytes(b, binary.LittleEndian, false)
    err = s.Unmarshal(&strs)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(strs)
}

func TestReader_Unmarshal_bytes(t *testing.T) {
    var bytes []byte
    b, _ := hex.DecodeString("020361626303646566")
    s := NewFromBytes(b, binary.LittleEndian, false)
    err := s.Unmarshal(&bytes)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(bytes)
}

func TestReader_Unmarshal_ptr(t *testing.T) {
    var v *int32
    b, _ := hex.DecodeString("02036162")
    s := NewFromBytes(b, binary.LittleEndian, false)
    err := s.Unmarshal(&v)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(*v)
}

func TestReader_Unmarshal_array(t *testing.T) {
    var v [3]int32
    b, _ := hex.DecodeString("0203616202036162")
    s := NewFromBytes(b, binary.LittleEndian, false)
    err := s.Unmarshal(&v)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(v)
}

func TestReader_Unmarshal_map(t *testing.T) {
    //var v map[string]int32
    v := make(map[string]int32)
    b, _ := hex.DecodeString("0203616263020361620362636402036162")
    s := NewFromBytes(b, binary.LittleEndian, false)
    err := s.Unmarshal(&v)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(v)
}

func TestReader_Unmarshal_struct(t *testing.T) {
    type A struct {
        A int8
    }
    type B struct {
        B int16
    }
    type C struct {
        A  A
        AA A
        B  B
        t  int32
        tt int32 `bytestruct:"-"`
        C  string
    }

    var v C
    b, _ := hex.DecodeString("02036162026162")
    s := NewFromBytes(b, binary.LittleEndian, false)
    err := s.Unmarshal(&v)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(v)
}

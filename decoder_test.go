package bytestruct

import (
	"encoding/binary"
	"github.com/HarryWang29/bytestruct/types"
	"testing"
)

func TestReader_Unmarshal(t *testing.T) {
	s := struct {
		ID        uint32
		Timestamp uint32
		UserID    uint64
	}{}
	h := "b68b01003e19a0634ab9b93b00000000"
	err := NewReader(binary.LittleEndian).UnmarshalFromHexString(h, &s)
	if err != nil {
		t.Fatal(err)
	}
	if s.ID != 101302 || s.Timestamp != 1671436606 || s.UserID != 1002027338 {
		t.Fatal("decode error")
	}
}

func TestReader_Unmarshal_bool(t *testing.T) {
	var r bool
	err := UnmarshalLE([]byte{0xa1}, &r)
	if err != nil {
		t.Fatal(err)
	}
	if r != true {
		t.Fatal("decode bool error: ", r)
	}
}

func TestReader_Unmarshal_string_2_zero(t *testing.T) {
	var r string
	reader := NewReader(binary.LittleEndian, types.SetString2Zero)
	err := reader.UnmarshalFromBytes([]byte{0x61, 0x62, 0x63, 0x00, 0x64}, &r)
	if err != nil {
		t.Fatal(err)
	}
	if r != "abc" {
		t.Fatal("decode string error: ", r)
	}
}

func TestReader_Unmarshal_len_string(t *testing.T) {
	var r string
	reader := NewReader(binary.LittleEndian)
	err := reader.UnmarshalFromBytes([]byte{0x3, 0, 0, 0, 0x61, 0x62, 0x63, 0x64}, &r)
	if err != nil {
		t.Fatal(err)
	}
	if r != "abc" {
		t.Fatal("decode string error: ", r)
	}
}

//
func TestReader_Unmarshal_slice(t *testing.T) {
	t.Run("slice_int32", func(t *testing.T) {
		var v []int32
		s := NewReader(binary.LittleEndian)
		err := s.UnmarshalFromBytes([]byte{0x02, 0, 0, 0, 0xa1, 0xa2, 0x00, 0x00, 0xa3, 0xa4, 0, 0}, &v)
		if err != nil {
			t.Fatal(err)
		}
		if len(v) != 2 || v[0] != 41633 || v[1] != 42147 {
			t.Fatal("decode slice error: ", v)
		}
	})

	t.Run("slice_int64", func(t *testing.T) {
		var v []int64
		s := NewReader(binary.LittleEndian)
		err := s.UnmarshalFromBytes([]byte{0x01, 0, 0, 0, 0xa1, 0xa2, 0x00, 0x00, 0xa3, 0xa4, 0, 0}, &v)
		if err != nil {
			t.Fatal(err)
		}
		if len(v) != 1 || v[0] != 181019986666145 {
			t.Fatal("decode slice error: ", v)
		}
	})

	t.Run("slice_string", func(t *testing.T) {
		var strs []string
		s := NewReader(binary.LittleEndian)
		err := s.UnmarshalFromHexString("020000000300000061626303000000646566", &strs)
		if err != nil {
			t.Fatal(err)
		}
		if len(strs) != 2 || strs[0] != "abc" || strs[1] != "def" {
			t.Fatal("decode slice error: ", strs)
		}
	})

	t.Run("slice_bytes", func(t *testing.T) {
		var bytes []byte
		s := NewReader(binary.LittleEndian)
		err := s.UnmarshalFromHexString("0200000061626303000000646566", &bytes)
		if err != nil {
			t.Fatal(err)
		}
		if len(bytes) != 2 || bytes[0] != 0x61 || bytes[1] != 0x62 {
			t.Fatal("decode slice error: ", bytes)
		}
	})
}

func TestReader_Unmarshal_ptr(t *testing.T) {
	var v *int32
	s := NewReader(binary.LittleEndian)
	err := s.UnmarshalFromHexString("02036162", &v)
	if err != nil {
		t.Fatal(err)
	}
	if v == nil || *v != 1650524930 {
		t.Fatal("decode ptr error: ", v)
	}
}

func TestReader_Unmarshal_array(t *testing.T) {
	var v [3]int32
	s := NewReader(binary.LittleEndian)
	err := s.UnmarshalFromHexString("0203616202036162", &v)
	if err != nil {
		t.Fatal(err)
	}
	if v[0] != 1650524930 || v[1] != 1650524930 || v[2] != 0 {
		t.Fatal("decode array error: ", v)
	}
}

func TestReader_Unmarshal_map(t *testing.T) {
	//must make map before decode
	//var v map[string]int32
	v := make(map[string]int32)
	s := NewReader(binary.LittleEndian)
	err := s.UnmarshalFromHexString("02000000030000006162630200000002000000616264020361", &v)
	if err != nil {
		t.Fatal(err)
	}
	if len(v) != 2 || v["ab"] != 1627587172 || v["abc"] != 2 {
		t.Fatal("decode map error: ", v)
	}
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
		T  int32
		Tt int32 `bytestruct:"-"`
		C  string
	}
	t.Run("struct", func(t *testing.T) {
		var v C
		s := NewReader(binary.LittleEndian)
		err := s.UnmarshalFromHexString("0203616201000000020000006162", &v)
		if err != nil {
			t.Fatal(err)
		}
		if v.A.A != 2 || v.AA.A != 3 || v.B.B != 25185 || v.T != 1 || v.Tt != 0 || v.C != "ab" {
			t.Fatal("decode struct error: ", v)
		}
	})

	t.Run("struct_varint", func(t *testing.T) {
		SetParseStringLen(ParseUvarintZigzag)
		var v C
		s := NewReader(binary.LittleEndian)
		err := s.UnmarshalFromHexString("0203616201000000046162", &v)
		if err != nil {
			t.Fatal(err)
		}
		if v.A.A != 2 || v.AA.A != 3 || v.B.B != 25185 || v.T != 1 || v.Tt != 0 || v.C != "ab" {
			t.Fatal("decode struct error: ", v)
		}
	})
}

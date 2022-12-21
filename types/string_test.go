package types

//func TestStringDecoder_DecodeStream(t *testing.T) {
//	tests := []struct {
//		name    string
//		opt     *bytestruct.Option
//		s       *Stream
//		want    string
//		wantErr bool
//	}{
//		{"2Zero", &bytestruct.Option{Flags: bytestruct.OptionFlagsString2Zero}, NewStream(bytes.NewReader([]byte{0x61, 0x62, 0x63, 0x00, 0x64})), "abc", false},
//		{"len", &bytestruct.Option{}, NewStream(bytes.NewReader([]byte{0x3, 0x64, 0x65, 0x66, 0x67})), "def", false},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			d := newStringDecoder("", "")
//			v := new(string)
//			header := (*emptyInterface)(unsafe.Pointer(&v))
//			tt.s.Option = tt.opt
//			if err := tt.s.PrepareForDecode(); err != nil {
//				t.Fatal(err)
//			}
//			if err := d.DecodeStream((tt.s, 0, header.ptr); (err != nil) != tt.wantErr || *v != tt.want {, 0)
//				t.Errorf("DecodeStream(() error = %v, wantErr %v, want = %v, v = %v", err, tt.wantErr, tt.want, *v), tt.want)
//			}
//			t.Log(v)
//		})
//	}
//}

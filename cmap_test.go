package pdf

import "testing"

func TestMapBFRangeStringIdentityGreek(t *testing.T) {
	// Identity ToUnicode: <0000> <FFFF> <0000>
	lo := "\x00\x00"
	dst := "\x00\x00"
	got := mapBFRangeString("\x03\xb2", lo, dst) // β
	want := "\x03\xb2"
	if got != want {
		t.Fatalf("mapped=%q want=%q", got, want)
	}
	if utf16Decode(got) != "β" {
		t.Fatalf("decoded=%q", utf16Decode(got))
	}
}

func TestCMapIdentityDecodeGreekAndASCII(t *testing.T) {
	m := &cmap{}
	m.space[1] = []byteRange{{"\x00\x00", "\xff\xff"}}
	m.bfrange = []bfrange{{
		lo:  "\x00\x00",
		hi:  "\xff\xff",
		dst: Value{data: "\x00\x00"},
	}}

	// "14 " + "βιβλιον "
	raw := "" +
		"\x00" + "1" + "\x00" + "4" + "\x00" + " " +
		"\x03\xb2\x03\xb9\x03\xb2\x03\xbb\x03\xb9\x03\xbf\x03\xbd\x00 "
	got := m.Decode(raw)
	want := "14 βιβλιον "
	if got != want {
		t.Fatalf("got=%q want=%q", got, want)
	}
}

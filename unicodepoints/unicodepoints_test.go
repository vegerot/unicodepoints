package unicodepoints_test

import (
	"testing"
	"unicodepoints/unicodepoints"
)

func TestA(t *testing.T) {
	got := unicodepoints.CodePointToBytes(0x41)
	want := []byte{0x41}

	if len(got) != len(want) {
		t.Errorf("got %q, want %q", got, want)
	}
	for idx := range got {
		if got[idx] != want[idx] {
			t.Errorf("got %q, want %q", got, want)
		}
	}
}

func TestB(t *testing.T) {
	got := unicodepoints.CodePointToBytes(0x42)
	want := []byte{0x42}

	if len(got) != len(want) {
		t.Errorf("got %q, want %q", got, want)
	}
	for idx := range got {
		if got[idx] != want[idx] {
			t.Errorf("got %q, want %q", got, want)
		}
	}
}

func TestAlef(t *testing.T) {
	var alefCodePoint uint32 = 0x5D0
	got := unicodepoints.CodePointToBytes(alefCodePoint)
	want := []byte{0xD7, 0x90}

	if len(got) != len(want) {
		t.Errorf("got %q, want %q", got, want)
		return
	}
	for idx := range got {
		if got[idx] != want[idx] {
			t.Errorf("got %q, want %q", got, want)
		}
	}
}

func TestEuro(t *testing.T) {
	var euroCodePoint uint32 = 0x20AC // €
	got := unicodepoints.CodePointToBytes(euroCodePoint)
	want := []byte{0xE2, 0x82, 0xAC}

	if len(got) != len(want) {
		t.Errorf("got %q, want %q", got, want)
		return
	}
	for idx := range got {
		if got[idx] != want[idx] {
			t.Errorf("got %q, want %q from %q and %q", got[idx], want[idx], got, want)
			return
		}
	}
}

func TestParty(t *testing.T) {
	var euroCodePoint uint32 = 0x1F973 // "🥳"
	got := unicodepoints.CodePointToBytes(euroCodePoint)
	want := []byte{0xF0, 0x9F, 0xA5, 0xB3}

	if len(got) != len(want) {
		t.Errorf("got %q, want %q", got, want)
		return
	}
	for idx := range got {
		if got[idx] != want[idx] {
			t.Errorf("got 0x%x, want 0x%x from 0x%x and 0x%x", got[idx], want[idx], got, want)
			return
		}
	}
}

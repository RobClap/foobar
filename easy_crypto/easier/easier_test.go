package easier

import (
	"testing"
)

var rotatingModTests = []struct {
	a        int //input
	b        int //input
	expected int //out
}{
	{0, 5, 0},
	{8, 10, 8},
	{-8, -10, 2},
	{10, 10, 0},
	{-10, 10, 0},
	{10, 7, 3},
	{-10, 7, 4},
}

func TestRotatingMod(t *testing.T) {
	for _, tt := range rotatingModTests {
		if actual := rotatingMod(tt.a, tt.b); actual != tt.expected {
			t.Errorf("RotatingMod(%d,%d): expected %d,actual %d", tt.a, tt.b, tt.expected, actual)
		}
	}
}

var RotTests = []struct {
	str_in, charset string
	shift           int
	expected        string //out
}{
	{"Prova", Alphabet, 0, "Prova"},
	{"Prova", Alphabet, 3, "Suryd"},
	{"Suryd", Alphabet, -3, "Prova"},
	{"Suryd", Alphabet, -len(Alphabet) - 3, "Prova"},
	{"Prova", Alphabet, len(Alphabet) + 3, "Suryd"},
	{"Prova", Kb_alphabet, 0, "Prova"},
	{"Prova", Kb_alphabet, len(Kb_alphabet), "Prova"},
	{"Prova", Kb_alphabet, 3, "Dusmf"},
	{"Prova", "!", 3, "Prova"},
	{"Prova!", Alphabet, 3, "Suryd!"},
}

func TestRot(t *testing.T) {
	for _, tt := range RotTests {
		if actual := Rot(tt.str_in, tt.charset, tt.shift); actual != tt.expected {
			t.Errorf("Rot(%s,%s,%d): expected %s,actual %s", tt.str_in, tt.charset, tt.shift, tt.expected, actual)
		}
	}
}

var SubstituteTests = []struct {
	str_in, alpha1, alpha2 string
	expected               string //out
}{
	{"Prova", "a", "b", "Provb"},
	{"Prova", "", "", "Prova"},
	{"Prova", "arv", "cde", "Pdoec"},
	{"Pippo", "Pp", "pP", "piPPo"},
}

func TestSubstitute(t *testing.T) {
	for _, tt := range SubstituteTests {
		if actual := Substitute(tt.str_in, tt.alpha1, tt.alpha2); actual != tt.expected {
			t.Errorf("Substitute(%s,%s,%s): expected %s,actual %s", tt.str_in, tt.alpha1, tt.alpha2, tt.expected, actual)
		}
	}
}

var ReverseTests = []struct {
	str_in   string
	expected string //out
}{
	{"Prova", "avorP"},
}

func TestReverse(t *testing.T) {
	for _, tt := range ReverseTests {
		if actual := Reverse(tt.str_in); actual != tt.expected {
			t.Errorf("Reverse(%s): expected %s,actual %s", tt.str_in, tt.expected, actual)
		}
	}
}

var ShiftTests = []struct {
	str_in   string
	split    int
	expected string //out
}{
	{"abcde", 2, "cdeab"},
	{"abcde", 0, "abcde"},
	{"abcde", 5, "abcde"},
}

func TestShift(t *testing.T) {
	for _, tt := range ShiftTests {
		if actual := Shift(tt.str_in, tt.split); actual != tt.expected {
			t.Errorf("Shift(%s,%d): expected %s,actual %s", tt.str_in, tt.split, tt.expected, actual)
		}
	}
}

var PreviousAsKeyTests = []struct {
	str_in, charset string
	start           int
	expected        string //out
}{
	{"abcde", Alphabet, 0, "abdfh"},
	{"abcde", Alphabet, 1, "bbdfh"},
}

func TestPreviousAsKey(t *testing.T) {
	for _, tt := range PreviousAsKeyTests {
		if actual := PreviousAsKey(tt.str_in, tt.charset, tt.start); actual != tt.expected {
			t.Errorf("PreviousAsKey(%s,%s,%d): expected %s,actual %s", tt.str_in, tt.charset, tt.start, tt.expected, actual)
		}
	}
}

var PreviousAsKeyDecodeTests = []struct {
	str_in, charset string
	start           int
	expected        string //out
}{
	{"abdfh", Alphabet, 0, "abcde"},
	{"bbdfh", Alphabet, 1, "abcde"},
}

func TestPreviousAsKeyDecode(t *testing.T) {
	for _, tt := range PreviousAsKeyDecodeTests {
		if actual := PreviousAsKeyDecode(tt.str_in, tt.charset, tt.start); actual != tt.expected {
			t.Errorf("PreviousAsKeyDecode(%s,%s,%d): expected %s,actual %s", tt.str_in, tt.charset, tt.start, tt.expected, actual)
		}
	}
}

var CascadeTests = []struct {
	str_in, charset string
	start           int
	expected        string //out
}{
	{"abcde", Alphabet, 0, "abdgk"},
	{"abcde", Alphabet, 1, "bcehl"},
}

func TestCascade(t *testing.T) {
	for _, tt := range CascadeTests {
		if actual := Cascade(tt.str_in, tt.charset, tt.start); actual != tt.expected {
			t.Errorf("Cascade(%s,%s,%d): expected %s,actual %s", tt.str_in, tt.charset, tt.start, tt.expected, actual)
		}
	}
}

var CascadeDecodeTests = []struct {
	str_in, charset string
	start           int
	expected        string //out
}{
	{"abdgk", Alphabet, 0, "abcde"},
	{"bcehl", Alphabet, 1, "abcde"},
}

func TestCascadeDecode(t *testing.T) {
	for _, tt := range CascadeDecodeTests {
		if actual := CascadeDecode(tt.str_in, tt.charset, tt.start); actual != tt.expected {
			t.Errorf("CascadeDecode(%s,%s,%d): expected %s,actual %s", tt.str_in, tt.charset, tt.start, tt.expected, actual)
		}
	}
}

var VigenereTests = []struct {
	str_in, key, charset string
	expected             string //out
}{
	{"aaaaa", "abcde", Alphabet, "abcde"},
	{"abcde", "bbbbb", Alphabet, "bcdef"},
	{"abcdefghijklmn", "bbbbb", Alphabet, "bcdefghijklmno"},
}

func TestVigenere(t *testing.T) {
	for _, tt := range VigenereTests {
		if actual := Vigenere(tt.str_in, tt.key, tt.charset); actual != tt.expected {
			t.Errorf("Vigenere(%s,%s,%s): expected %s,actual %s", tt.str_in, tt.key, tt.charset, tt.expected, actual)
		}
	}
}

var VigenereDecodeTests = []struct {
	str_in, key, charset string
	expected             string //out
}{
	{"abcde", "abcde", Alphabet, "aaaaa"},
	{"bcdef", "bbbbb", Alphabet, "abcde"},
	{"bcdefghijklmno", "bbbbb", Alphabet, "abcdefghijklmn"},
}

func TestVigenereDecode(t *testing.T) {
	for _, tt := range VigenereDecodeTests {
		if actual := VigenereDecode(tt.str_in, tt.key, tt.charset); actual != tt.expected {
			t.Errorf("VigenereDecode(%s,%s,%s): expected %s,actual %s", tt.str_in, tt.key, tt.charset, tt.expected, actual)
		}
	}
}

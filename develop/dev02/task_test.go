package main

import (
	"testing"
)

func TestUnpackStr(t *testing.T) {
	var tests = []struct {
		strIn  string
		strOut string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"", ""},
	}

	for _, v := range tests {
		testStrOut, err := unpackStr(v.strIn)
		if err != nil {
			t.Error("error unpacking string: ", err)
		}
		if testStrOut != v.strOut {
			t.Errorf("expected %s, got %s", v.strOut, testStrOut)
		}
	}
}

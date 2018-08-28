package smlrpc

import (
	"testing"
)

func TestLowerIdent(t *testing.T) {
	for _, test := range []struct {
		in, out string
	}{
		{"A", "a"},
		{"Base", "base"},
		{"Something", "something"},
		{"HTTP", "http"},
		{"HTTPConnection", "httpConnection"},
		{"_A", "_A"},
		{"lowerAlready", "lowerAlready"},
	} {
		if got := lowerIdent(test.in); got != test.out {
			t.Errorf(
				"lowerIdent(%q), want %q, got %q",
				test.in, test.out, got,
			)
		}
	}
}

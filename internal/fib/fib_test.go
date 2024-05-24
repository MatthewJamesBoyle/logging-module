package fib_test

import (
	"github.com/matthewjamesboyle/logging-module/internal/fib"
	"testing"
)

func TestFib(t *testing.T) {
	cases := []struct {
		name  string
		input int
		want  int
	}{
		{"Fib 0", 0, 0},
		{"Fib 1", 1, 1},
		{"Fib 2", 2, 1},
		{"Fib 3", 3, 2},
		{"Fib 4", 4, 3},
		{"Fib 5", 5, 5},
		{"Fib 6", 6, 6},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := fib.Fib(tc.input)
			if got != tc.want {
				t.Errorf(
					"Fib(%d) = %d; want %d",
					tc.input,
					got,
					tc.want,
				)
			}
		})
	}
}

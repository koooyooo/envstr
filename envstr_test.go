package envstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApply(t *testing.T) {
	tests := []struct {
		name  string
		setUp func(*testing.T)
		in    string
		out   string
	}{
		{
			name: "Basic",
			setUp: func(t *testing.T) {
				t.Setenv("Hello", "World")
			},
			in:  "I like this ${Hello} !",
			out: "I like this World !",
		},
		{
			name: "No Env Found",
			setUp: func(t *testing.T) {
				t.Setenv("Stop", "World")
			},
			in:  "I like this ${Hello} !",
			out: "I like this  !",
		},
		{
			name: "No Replacer Found",
			setUp: func(t *testing.T) {
				t.Setenv("Hello", "World")
			},
			in:  "I like this Hello !",
			out: "I like this Hello !",
		},
		{
			name: "Reversed Replacer",
			setUp: func(t *testing.T) {
				t.Setenv("Hello", "World")
			},
			in:  "I like this }Hello${ !",
			out: "I like this }Hello${ !",
		},
		{
			name: "Multi",
			setUp: func(t *testing.T) {
				t.Setenv("Hello", "World")
				t.Setenv("Foo", "Bar")
			},
			in:  "I like this ${Hello} of ${Foo} !",
			out: "I like this World of Bar !",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setUp(t)
			out := Apply(test.in)
			assert.Equal(t, test.out, out)
		})
	}
}

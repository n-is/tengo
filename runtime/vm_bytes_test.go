package runtime_test

import (
	"testing"

	"github.com/n-is/tengo/objects"
)

func TestBytes(t *testing.T) {
	expect(t, `out = bytes("Hello World!")`, nil, []byte("Hello World!"))
	expect(t, `out = bytes("Hello") + bytes(" ") + bytes("World!")`, nil, []byte("Hello World!"))

	// bytes[] -> int
	expect(t, `out = bytes("abcde")[0]`, nil, 97)
	expect(t, `out = bytes("abcde")[1]`, nil, 98)
	expect(t, `out = bytes("abcde")[4]`, nil, 101)
	expect(t, `out = bytes("abcde")[10]`, nil, objects.UndefinedValue)
}

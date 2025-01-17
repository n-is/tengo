package script_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/n-is/tengo/assert"
	"github.com/n-is/tengo/compiler/token"
	"github.com/n-is/tengo/objects"
	"github.com/n-is/tengo/script"
)

type Counter struct {
	value int64
}

func (o *Counter) TypeName() string {
	return "counter"
}

func (o *Counter) String() string {
	return fmt.Sprintf("Counter(%d)", o.value)
}

func (o *Counter) BinaryOp(op token.Token, rhs objects.Object) (objects.Object, error) {
	switch rhs := rhs.(type) {
	case *Counter:
		switch op {
		case token.Add:
			return &Counter{value: o.value + rhs.value}, nil
		case token.Sub:
			return &Counter{value: o.value - rhs.value}, nil
		}
	case *objects.Int:
		switch op {
		case token.Add:
			return &Counter{value: o.value + rhs.Value}, nil
		case token.Sub:
			return &Counter{value: o.value - rhs.Value}, nil
		}
	}

	return nil, errors.New("invalid operator")
}

func (o *Counter) IsFalsy() bool {
	return o.value == 0
}

func (o *Counter) Equals(t objects.Object) bool {
	if tc, ok := t.(*Counter); ok {
		return o.value == tc.value
	}

	return false
}

func (o *Counter) Copy() objects.Object {
	return &Counter{value: o.value}
}

func (o *Counter) Call(args ...objects.Object) (objects.Object, error) {
	return &objects.Int{Value: o.value}, nil
}

func TestScript_CustomObjects(t *testing.T) {
	c := compile(t, `a := c1(); s := string(c1); c2 := c1; c2++`, M{
		"c1": &Counter{value: 5},
	})
	compiledRun(t, c)
	compiledGet(t, c, "a", int64(5))
	compiledGet(t, c, "s", "Counter(5)")
	compiledGetCounter(t, c, "c2", &Counter{value: 6})

	c = compile(t, `
arr := [1, 2, 3, 4]
for x in arr {
	c1 += x
}
out := c1()
`, M{
		"c1": &Counter{value: 5},
	})
	compiledRun(t, c)
	compiledGet(t, c, "out", int64(15))
}

func compiledGetCounter(t *testing.T, c *script.Compiled, name string, expected *Counter) bool {
	v := c.Get(name)
	if !assert.NotNil(t, v) {
		return false
	}

	actual := v.Value().(*Counter)
	if !assert.NotNil(t, actual) {
		return false
	}

	return assert.Equal(t, expected.value, actual.value)
}

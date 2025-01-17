package objects_test

import (
	"testing"

	"github.com/n-is/tengo/assert"
	"github.com/n-is/tengo/objects"
)

func TestMap_Index(t *testing.T) {
	m := &objects.Map{Value: make(map[string]objects.Object)}
	k := &objects.Int{Value: 1}
	v := &objects.String{Value: "abcdef"}
	err := m.IndexSet(k, v)

	assert.NoError(t, err)

	res, err := m.IndexGet(k)
	assert.NoError(t, err)
	assert.Equal(t, v, res)
}

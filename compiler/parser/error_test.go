package parser_test

import (
	"testing"

	"github.com/n-is/tengo/assert"
	"github.com/n-is/tengo/compiler/parser"
	"github.com/n-is/tengo/compiler/source"
)

func TestError_Error(t *testing.T) {
	err := &parser.Error{Pos: source.FilePos{Offset: 10, Line: 1, Column: 10}, Msg: "test"}
	assert.Equal(t, "Parse Error: test\n\tat 1:10", err.Error())
}

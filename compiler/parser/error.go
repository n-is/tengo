package parser

import (
	"fmt"

	"github.com/n-is/tengo/compiler/source"
)

// Error represents a parser error.
type Error struct {
	Pos source.FilePos
	Msg string
}

func (e Error) Error() string {
	if e.Pos.Filename != "" || e.Pos.IsValid() {
		return fmt.Sprintf("Parse Error: %s\n\tat %s", e.Msg, e.Pos)
	}

	return fmt.Sprintf("Parse Error: %s", e.Msg)
}

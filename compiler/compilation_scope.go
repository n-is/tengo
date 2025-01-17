package compiler

import "github.com/n-is/tengo/compiler/source"

// CompilationScope represents a compiled instructions
// and the last two instructions that were emitted.
type CompilationScope struct {
	instructions []byte
	symbolInit   map[string]bool
	sourceMap    map[int]source.Pos
}

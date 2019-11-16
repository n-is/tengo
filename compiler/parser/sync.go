package parser

import "github.com/n-is/tengo/compiler/token"

var stmtStart = map[token.Token]bool{
	token.Break:    true,
	token.Continue: true,
	token.For:      true,
	token.If:       true,
	token.Return:   true,
	token.Export:   true,
}

package ast_test

import (
	"testing"

	"github.com/n-is/tengo/compiler/ast"
)

func TestIdentListString(t *testing.T) {
	identListVar := &ast.IdentList{
		List: []*ast.Ident{
			{Name: "a"},
			{Name: "b"},
			{Name: "c"},
		},
		VarArgs: true,
	}

	expectedVar := "(a, b, ...c)"
	if str := identListVar.String(); str != expectedVar {
		t.Fatalf("expected string of %#v to be %s, got %s", identListVar, expectedVar, str)
	}

	identList := &ast.IdentList{
		List: []*ast.Ident{
			{Name: "a"},
			{Name: "b"},
			{Name: "c"},
		},
		VarArgs: false,
	}

	expected := "(a, b, c)"
	if str := identList.String(); str != expected {
		t.Fatalf("expected string of %#v to be %s, got %s", identList, expected, str)
	}
}

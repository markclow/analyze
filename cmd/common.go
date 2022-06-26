package cmd

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func parseFile(filename string) (*ast.File, error) {
	fset := token.NewFileSet()
	root, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return root, nil
}

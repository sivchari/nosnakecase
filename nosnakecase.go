package nosnakecase

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "nosnakecase is a linter that detects snake case of variable naming and function name."

// Analyzer is a nosnakecase linter.
var Analyzer = &analysis.Analyzer{
	Name: "nosnakecase",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			report(pass, n.Pos(), n.Name)
		}
	})

	return nil, nil
}

func report(pass *analysis.Pass, pos token.Pos, name string) {
	// skip import _ "xxx"
	if name == "_" {
		return
	}
	// skip package xxx_test
	if strings.Contains(name, "_test") {
		return
	}
	if strings.Contains(name, "_") {
		pass.Reportf(pos, "%s is used under score. You should use mixedCap or MixedCap.", name)
		return
	}
}

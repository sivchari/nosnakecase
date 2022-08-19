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
	result := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.SelectorExpr)(nil),
		(*ast.Ident)(nil),
	}

	skip := map[token.Pos]struct{}{}

	result.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.SelectorExpr:
			// Skip field selectors to avoid throwing false positives on stdlib constants etc.
			skip[n.Sel.Pos()] = struct{}{}
		case *ast.Ident:
			pos := n.Pos()
			if _, ok := skip[pos]; !ok {
				report(pass, pos, n.Name)
			} else {
				delete(skip, pos)
			}
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

	// If prefix is Test or Benchmark, Fuzz, skip
	// FYI https://go.dev/blog/examples
	if strings.HasPrefix(name, "Test") || strings.HasPrefix(name, "Benchmark") || strings.HasPrefix(name, "Fuzz") {
		return
	}

	if strings.Contains(name, "_") {
		pass.Reportf(pos, "%s contains underscore. You should use mixedCap or MixedCap.", name)
		return
	}
}

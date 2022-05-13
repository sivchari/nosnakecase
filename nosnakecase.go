package nosnakecase

import (
	"go/ast"
	"go/token"
	"strings"
	"sync"

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
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			var wg sync.WaitGroup
			wg.Add(4)

			go func() {
				defer wg.Done()
				name := n.Name
				report(pass, name.Pos(), name.Name)
			}()

			go func() {
				defer wg.Done()
				params := n.Type.Params
				for _, param := range params.List {
					for _, ident := range param.Names {
						report(pass, ident.Pos(), ident.Name)
					}
				}
			}()

			go func() {
				defer wg.Done()
				results := n.Type.Results
				if results == nil {
					return
				}
				for _, list := range results.List {
					for _, ident := range list.Names {
						report(pass, ident.Pos(), ident.Name)
					}
				}
			}()

			go func() {
				defer wg.Done()
				if n.Body.List == nil {
					return
				}
				for _, list := range n.Body.List {
					switch l := list.(type) {
					case *ast.AssignStmt:
						for _, lh := range l.Lhs {
							ident, ok := lh.(*ast.Ident)
							if !ok {
								continue
							}
							report(pass, ident.Pos(), ident.Name)
						}
					case *ast.DeclStmt:
						gendecl, ok := l.Decl.(*ast.GenDecl)
						if !ok {
							continue
						}
						for _, spec := range gendecl.Specs {
							valspec, ok := spec.(*ast.ValueSpec)
							if !ok {
								continue
							}
							for _, ident := range valspec.Names {
								report(pass, ident.Pos(), ident.Name)
							}
						}
					}
				}
			}()

			wg.Wait()
		}
	})

	return nil, nil
}

func report(pass *analysis.Pass, pos token.Pos, name string) {
	if name == "_" {
		return
	}
	if strings.Contains(name, "_") {
		pass.Reportf(pos, "%s is used under score. You should use mixedCap or MixedCap.", name)
		return
	}
}

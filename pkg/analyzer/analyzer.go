package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "gochecklog",
	Doc:      "Checks that log is correct.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		call := node.(*ast.CallExpr)

		if len(call.Args) == 0 {
			return
		}

		if !isLogFunctionCall(pass, call) {
			return
		}

		literals := extractStringLiterals(call.Args[0])
		for _, lit := range literals {
			if lit.Value != "" {
				checkRules(pass, lit.Pos, lit.Value)
			}
		}

	})

	return nil, nil
}

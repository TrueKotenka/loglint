package analyzer

import (
	"go/ast"
	"go/token"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "loglint",
	Doc:      "Checks log messages for lowercase start, English only, no special chars/emojis, and no sensitive data",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

var (
	sensitiveRegex = regexp.MustCompile(`(?i)(password|token|api_key|apikey|secret)`)
	logMethods     = map[string]bool{"Info": true, "Error": true, "Warn": true, "Debug": true, "Fatal": true}
)

func run(pass *analysis.Pass) (interface{}, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{(*ast.CallExpr)(nil)}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return
		}

		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}

		if !logMethods[sel.Sel.Name] {
			return
		}

		isLogCall := false
		if ident, ok := sel.X.(*ast.Ident); ok {
			if ident.Name == "slog" || ident.Name == "log" {
				isLogCall = true
			}
		} else {
			tv, ok := pass.TypesInfo.Types[sel.X]
			if ok && (strings.Contains(tv.Type.String(), "zap.Logger") || strings.Contains(tv.Type.String(), "zap.SugaredLogger")) {
				isLogCall = true
			}
		}

		if !isLogCall || len(call.Args) == 0 {
			return
		}

		checkLogArgument(pass, call.Args[0])
	})

	return nil, nil
}

func checkLogArgument(pass *analysis.Pass, arg ast.Expr) {
	var message string
	var strNode *ast.BasicLit

	ast.Inspect(arg, func(n ast.Node) bool {
		if lit, ok := n.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			message = strings.Trim(lit.Value, `"`)
			strNode = lit
			return false
		}
		if ident, ok := n.(*ast.Ident); ok {
			if sensitiveRegex.MatchString(ident.Name) {
				pass.Reportf(ident.Pos(), "log message contains potentially sensitive variable: %s", ident.Name)
			}
		}
		return true
	})

	if strNode == nil || message == "" {
		return
	}

	runes := []rune(message)
	if len(runes) > 0 && unicode.IsUpper(runes[0]) {
		fixedMsg := string(unicode.ToLower(runes[0])) + string(runes[1:])
		pass.Report(analysis.Diagnostic{
			Pos:     strNode.Pos(),
			Message: "log messages must start with a lowercase letter",
			SuggestedFixes: []analysis.SuggestedFix{{
				Message: "Make lowercase",
				TextEdits: []analysis.TextEdit{{
					Pos:     strNode.Pos(),
					End:     strNode.End(),
					NewText: []byte(`"` + fixedMsg + `"`),
				}},
			}},
		})
	}

	for _, r := range runes {
		if r > unicode.MaxASCII || (!unicode.IsLetter(r) && !unicode.IsSpace(r) && !unicode.IsDigit(r)) {
			pass.Reportf(strNode.Pos(), "log messages must be in English and contain no special characters or emojis")
			break
		}
	}
}

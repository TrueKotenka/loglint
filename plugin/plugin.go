package plugin

import (
	"loglint/pkg/analyzer"

	"golang.org/x/tools/go/analysis"
)

// Экспортируемая переменная, которую ищет golangci-lint
var AnalyzerPlugin analyzerPlugin

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{analyzer.Analyzer}
}

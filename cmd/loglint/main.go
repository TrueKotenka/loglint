package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"loglint/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}

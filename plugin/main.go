package main

import (
	"github.com/resessh/sqlpassctxcheck"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

func (a analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		sqlpassctxcheck.Analyzer,
	}
}

// nolint
var AnalyzerPlugin analyzerPlugin

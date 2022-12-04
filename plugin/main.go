package main

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
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

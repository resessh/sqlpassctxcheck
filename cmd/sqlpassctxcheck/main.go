package main

import (
	"github.com/resessh/sqlpassctxcheck"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(sqlpassctxcheck.Analyzer)
}

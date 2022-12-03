package sqlpassctxcheck_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/resessh/sqlpassctxcheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, sqlpassctxcheck.Analyzer, "a")
}

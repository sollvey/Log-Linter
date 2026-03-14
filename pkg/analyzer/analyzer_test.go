package analyzer_test

import (
	"fmt"
	"linter/pkg/analyzer"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

// func TestAnalyzerSlog(t *testing.T) {
// 	cwd, err := os.Getwd()
// 	if err == nil {
// 		fmt.Println("error")
// 	}
// 	rp_dir := filepath.Dir(filepath.Dir(cwd))
// 	testdata_dir := filepath.Join(rp_dir, "testdata")
// 	analysistest.Run(t, testdata_dir, analyzer.Analyzer, "slog")
// }

func TestSlogAnalyzer(t *testing.T) {
	cwd, err := os.Getwd()
	if err == nil {
		fmt.Println("error")
	}
	rp_dir := filepath.Dir(filepath.Dir(cwd))
	testdata := filepath.Join(rp_dir, "testdata")

	// testdata := filepath.Join("..", "..", "testdata")
	analysistest.Run(t, testdata, analyzer.Analyzer, "slog")
}

func TestZapAnalyzer(t *testing.T) {
	cwd, err := os.Getwd()
	if err == nil {
		fmt.Println("error")
	}
	rp_dir := filepath.Dir(filepath.Dir(cwd))
	testdata := filepath.Join(rp_dir, "testdata")

	// testdata := filepath.Join("..", "..", "testdata")
	analysistest.Run(t, testdata, analyzer.Analyzer, "zap")
}

package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/sivchari/nosnakecase"
)

func main() { unitchecker.Main(nosnakecase.Analyzer) }

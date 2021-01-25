package main

import (
	"flag"
	conf "github.com/zero-yy/export-config/test/go/gen"
)

var (
	inputExcelPath       = flag.String("input", "", "path of input excel")
	outputDataGoPath     = flag.String("outg", "", "path of output data go")
	outputDataCsharpPath = flag.String("outcs", "", "path of output data csharp")
)

func main() {
	flag.Parse()

	conf.MustSave(*inputExcelPath, *outputDataGoPath)
}

package sheet

import (
	"bytes"
	"github.com/zero-yy/export-config/internal/config"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)

//func toTempSheets(sheets []*Sheet) []TempSheet {
//	ret := make([]TempSheet, 0, len(sheets))
//	for _, s := range sheets {
//		ret = append(ret, TempSheet{Name: s.Name})
//	}
//	return ret
//}

func generateLoader(sheets []*Sheet) {
	generateLoadForX(sheets,
		goLoadTemplate,
		path.Join(config.C.OutputCodeGoPath, "conf_loader_func.go"))

	generateLoadForX(sheets,
		csharpLoadTemplate,
		path.Join(config.C.OutputCodeCSharpPath, "ConfLoaderFunc.cs"))

}

func generateLoadForX(sheets []*Sheet, templateStr string, outputFileName string) {
	var buff bytes.Buffer
	temp, err := template.New("conf").
		//Funcs(template.FuncMap{"ToLower": strings.ToLower}).
		Parse(templateStr)

	if err != nil {
		panic(err)
	}

	err = temp.Execute(&buff, &TempData{Sheets: sheets})
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(outputFileName, buff.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func generateSaver(sheets []*Sheet) {
	var buff bytes.Buffer
	temp, err := template.New("conf").
		//Funcs(template.FuncMap{"ToLower": strings.ToLower}).
		Parse(goSaveTemplate)

	if err != nil {
		panic(err)
	}

	err = temp.Execute(&buff, &TempData{Sheets: sheets})
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(
		path.Join(config.C.OutputCodeGoPath, "conf_saver_func.go"),
		buff.Bytes(),
		os.ModePerm)

	if err != nil {
		panic(err)
	}
}

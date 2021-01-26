package sheet

import (
	"bytes"
	"github.com/zero-yy/export-config/internal/config"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)

func generateCode(sheets []*Sheet) {
	generateFrame(sheets)

	// loader
	generateLoader(sheets)

	// saver func
	generateSaver(sheets)
}

func generateFrame(sheets []*Sheet) {
	// go
	generateFrameForX(goGenFrame, config.C.OutputCodeGoPath)
	generateFrameForX(csGenFrame, config.C.OutputCodeCSharpPath)
}

func generateFrameForX(frame map[string]string, outputPath string) {
	for k, v := range frame {
		relativePath, _ := path.Split(k)
		err := os.MkdirAll(path.Join(outputPath, relativePath), os.ModePerm)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(path.Join(outputPath, k), []byte(v), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

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

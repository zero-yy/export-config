package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	InputExcelPath  string
	OutputProtoPath string

	OutputDataCSharpPath string
	OutputDataGoPath     string
	OutputCodeCSharpPath string
	OutputCodeGoPath     string

	ProtoPackage      string
	CSharpNamespace   string
	GoPackage         string
	GoPackageFullPath string
}

var C Config

func MustInit(fn string) {
	_, err := toml.DecodeFile(fn, &C)
	if err != nil {
		panic(err)
	}
	//fmt.Println(C)

	_ = os.MkdirAll(C.OutputProtoPath, os.ModePerm)
	_ = os.MkdirAll(C.OutputDataCSharpPath, os.ModePerm)
	_ = os.MkdirAll(C.OutputDataGoPath, os.ModePerm)
	_ = os.MkdirAll(C.OutputCodeCSharpPath, os.ModePerm)
	_ = os.MkdirAll(C.OutputCodeGoPath, os.ModePerm)
}

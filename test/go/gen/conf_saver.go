package conf

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/zero-yy/export-config/sheet"
	"io/ioutil"
	"os"
)

// sheet name -> saver
var DataSaveFunc = make(map[string]func(s *sheet.Sheet, outputPath string))

// excel data -> pb bytes
func MustSave(inputPath, outputPath string) {
	initSaveFunc()
	save(inputPath, outputPath)
}

func save(inputPath, outputPath string) {
	sheets := sheet.CreateSheets(inputPath)
	for _, s := range sheets {
		DataSaveFunc[s.Name](s, outputPath)
	}
}

func saveJson(m proto.Message, fullName string) {
	b, err := json.MarshalIndent(m, "", "  ")
	panicIfErr(err)

	err = ioutil.WriteFile(fullName, b, os.ModePerm)
	panicIfErr(err)
}

func saveBytes(m proto.Message, fullName string) {
	buf := proto.NewBuffer(nil)
	buf.SetDeterministic(true)
	err := buf.Marshal(m)
	panicIfErr(err)

	// TODO...
	//m.Crc32 = crc32.ChecksumIEEE(buf.Bytes())

	// recalculate with right crc32!
	buf = proto.NewBuffer(nil)
	buf.SetDeterministic(true)
	err = buf.Marshal(m)
	panicIfErr(err)

	err = ioutil.WriteFile(fullName, buf.Bytes(), os.ModePerm)
	panicIfErr(err)
}

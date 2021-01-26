package sheet

import (
	"github.com/zero-yy/export-config/internal/config"
)

func GenCode() {
	sheets := CreateSheets(config.C.InputExcelPath)

	// proto file
	generateProto(sheets)

	// pb code
	callProtoC()

	generateCode(sheets)
}

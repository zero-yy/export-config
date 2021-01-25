package sheet

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

var (
	// file name => utc
	gFileLastChange = make(map[string]time.Time)
)

func hasSuffix(n string, lowerSuffixMap map[string]bool) bool {
	n = strings.ToLower(n)
	ext := filepath.Ext(n)
	// kill dot.
	ext = ext[1:]
	//fmt.Println(n, ext, lowerSuffixMap)
	if has, ok := lowerSuffixMap[ext]; ok {
		return has
	}
	return false
}

func getFiles(dirPath string, lowerSuffixMap map[string]bool) (files []string, err error) {
	files = make([]string, 0)

	err = filepath.Walk(dirPath, func(filename string, fi os.FileInfo, err error) error {
		if fi == nil {
			return nil
		}

		if fi.IsDir() {
			// no recursive
			return nil
		}
		if hasSuffix(fi.Name(), lowerSuffixMap) {
			if strings.Contains(fi.Name(), "~") {
				return nil
			}
			if strings.HasPrefix(fi.Name(), SkipPrefix) {
				return nil
			}

			files = append(files, filename)

			gFileLastChange[filename] = fi.ModTime().UTC()
		}
		return nil
	})
	return files, err
}

func CreateSheets(inputPath string) []*Sheet {
	// suffixAry should be lower
	fileNames, err := getFiles(inputPath,
		map[string]bool{
			"xlsm": true,
			"xlsx": true,
			"xls":  true})
	fmt.Println(inputPath, fileNames)
	if err != nil {
		panic(err.Error())
	}

	sheets := make([]*Sheet, 0)
	sheetMap := make(map[string]*Sheet)

	for _, fileName := range fileNames {
		xf, err := xlsx.OpenFile(fileName)
		if err != nil {
			panic(err.Error())
		}

		for _, xs := range xf.Sheets {
			if strings.Contains(xs.Name, SkipPrefix) {
				continue
			}

			if _, ok := sheetMap[xs.Name]; ok {
				panic(fmt.Errorf("sheet(%s) name repeated", xs.Name))
			}

			s := NewSheet(xs)
			sheets = append(sheets, s)
			sheetMap[xs.Name] = s
		}
	}
	return sheets
}

//---------------------------------------------------------------------------------------------
var (
	cfgMgrCS = "using System.Collections.Generic;\nusing Config;\nusing System.IO;\nusing UnityEngine;\nusing System;\npublic static partial class CfgMgr\n{" +
		`
	static CfgMgr()
	{
        if (!Directory.Exists(path))
        {
            Directory.CreateDirectory(path);
        }

		__TableNames__
        foreach (string name in TableNames)
        {
            Stream sm = null;
            if (File.Exists(path + name))
                sm = new FileStream(path + name, FileMode.Open);
            else
            {
                var ft = Resources.Load(cfgPath + name) as TextAsset;
                sm = new MemoryStream(ft.bytes);
            }
            Type t = Type.GetType("Config." + name);
            cacheTable(t, sm);
			sm.Close();
        }
    }
`

	csName = map[string]string{
		"int32":  "int",
		"int64":  "long",
		"uint32": "uint",
		"uint64": "ulong",
		"string": "string",
	}
)

func parseMgrCS(xlsxFile *xlsx.File) {
	for _, sheet := range xlsxFile.Sheets {
		if strings.Contains(sheet.Name, "__") {
			fmt.Printf("\t\tSKIP %v in parseMgrCS", sheet.Name)
			continue
		}
		fmt.Printf("\tparseMgrCS %v", sheet.Name)

		fpn := removeSpace(sheet.Rows[kNameRowIndex].Cells[0].Value) // 首列变量名
		fpt := sheet.Rows[kTypeRowIndex].Cells[0].Value              // 首列变量类型

		cfgMgrCS += "\tstatic Dictionary<" + csName[fpt] + ", " + sheet.Name + ".Types.Record> " + sheet.Name + "_map = new Dictionary<" + csName[fpt] + ", " + sheet.Name + ".Types.Record>();\n"
		cfgMgrCS += "\tpublic static " + sheet.Name + ".Types.Record GetRecordBy" + titleName(fpn) + "(this " + sheet.Name + " obj, " + csName[fpt] + " " + fpn + ")\n"
		cfgMgrCS += "\t{\n\t\tif (" + sheet.Name + "_map.Count != obj.Records.Count)\n\t\t{\n\t\t\t" + sheet.Name + "_map.Clear();\n\t\t\tforeach (" + sheet.Name + ".Types.Record r in obj.Records)\n\t\t\t{\n\t\t\t\t" + sheet.Name + "_map[r." + titleName(fpn) + "] = r;\n\t\t\t}\n\t\t}\n\t\t\n\t"
		cfgMgrCS += "\tif (!" + sheet.Name + "_map" + ".ContainsKey(" + fpn + ")) {Debug.LogErrorFormat(\"not find {0}\", " + fpn + ");\treturn null;}"
		cfgMgrCS += "\n\t\treturn " + sheet.Name + "_map[" + fpn + "];\n\t}\n"

		cfgMgrCS = strings.Replace(cfgMgrCS, "__TableNames__", "TableNames.Add(\""+sheet.Name+"\");\n\t\t__TableNames__", -1)
	}
}

func saveMgrCS(outdir string) {
	cfgMgrCS = strings.Replace(cfgMgrCS, "__TableNames__", "", -1)
	cfgMgrCS += "}"
	bmsg := []byte(cfgMgrCS)
	if err := ioutil.WriteFile(outdir+"/CfgMgr2.cs", bmsg, os.ModeExclusive); err != nil {
		panic(err.Error())
	}
}

//-------------------------------------------------------------------------------------------
func removeSpace(str string) string {
	str = strings.TrimSpace(str)
	return str
}

func titleName(str string) string {
	strs := strings.Split(str, "_")
	ostr := ""
	for i := 0; i < len(strs); i++ {
		_, err := strconv.Atoi(strs[i])
		if err != nil {
			ostr += strings.Title(strs[i])
		} else {
			ostr += "_" + strings.Title(strs[i])
			if i < len(strs)+1 {
				ostr += "_"
			}
		}
	}
	return ostr
}

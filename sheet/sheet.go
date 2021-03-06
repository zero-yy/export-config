package sheet

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"reflect"
	"strings"
)

const (
	kDecorateId  = "id"
	kDecorateAry = "ary"
)

type Col struct {
	Name      string
	CamelName string

	ProtoType string

	PureName      string
	CamelPureName string

	Decorate string
}
type Sheet struct {
	Id      string
	CamelId string

	ColAry []*Col
	ColMap map[int]*Col

	// default, (go and so on...)
	DefaultIdColType string
	CSharpIdColType  string
	IdColIndex       int

	CamelName string
	Name      string
	Xs        *xlsx.Sheet
}

// xid@id --> xid, id
// xx/0 xx/1 --> xxs, ""
func getPureName(colName string) (pureColName, decorate string) {
	pureColName = colName

	if strings.Contains(colName, "@") {
		a := strings.Split(colName, "@")
		pureColName = a[0]
		if len(a) > 1 {
			decorate = a[1]
		}
	} else if strings.Contains(colName, "/") {
		a := strings.Split(colName, "/")
		pureColName = a[0] + "s"
		decorate = kDecorateAry
	}

	decorate = strings.ToLower(decorate)

	return
}

func NewSheet(xs *xlsx.Sheet) (s *Sheet) {
	s = &Sheet{}
	s.Xs = xs

	s.Name = xs.Name
	s.CamelName = toCamelCase(xs.Name)

	s.ColMap = make(map[int]*Col)

	for i := KColStart; i < len(xs.Rows[kNameRowIndex].Cells); i++ {
		colName := xs.Rows[kNameRowIndex].Cells[i].Value
		originalColType := xs.Rows[kTypeRowIndex].Cells[i].Value

		defaultColType, ok := typeParse[originalColType]
		if !ok {
			panic(fmt.Errorf("sheet(%s) typeParse error: unknown type(%s) cell(%s)", xs.Name, originalColType, CellName(kTypeRowIndex, int32(i))))
		}

		csharpColType := defaultColType
		if _, ok2 := typeConverToCSharpt[defaultColType]; ok2 {
			csharpColType = typeConverToCSharpt[defaultColType]
		}

		pureColName, decorate := getPureName(colName)
		setSheetInfo := false
		if decorate == kDecorateId {
			setSheetInfo = true
		} else if pureColName == kDecorateId {
			if s.Id != pureColName {
				panic(fmt.Errorf("sheet(%s) repeat id %s vs %s", xs.Name, s.Id, pureColName))
			}
			setSheetInfo = true
		}

		if setSheetInfo {
			s.Id = pureColName
			s.CamelId = toCamelCase(s.Id)
			s.DefaultIdColType = defaultColType
			s.CSharpIdColType = csharpColType
			s.IdColIndex = i
		}

		c := &Col{
			Name:          colName,
			CamelName:     toCamelCase(colName),
			ProtoType:     defaultColType,
			PureName:      pureColName,
			CamelPureName: toCamelCase(pureColName),
			Decorate:      decorate,
		}

		s.ColAry = append(s.ColAry, c)
		s.ColMap[i] = c
	}

	if len(s.Id) == 0 {
		panic(fmt.Errorf("sheet(%s) not found id or @id in all column", xs.Name))
	}
	return
}

func (s *Sheet) getProtoCode() string {
	dealedPureNames := make(map[string]bool)

	var body string
	for index, v := range s.ColAry {
		if _, ok := dealedPureNames[v.PureName]; ok {
			continue
		}

		colType := v.ProtoType
		if v.Decorate == kDecorateAry {
			colType = "repeated " + colType
		}

		decorateComment := ""
		if len(v.Decorate) > 0 {
			decorateComment = fmt.Sprintf("// decorate:%s", v.Decorate)
		}
		line := fmt.Sprintf("\t\t%s %s = %d; %s\n", colType, v.PureName, index+1, decorateComment)

		body += line

		dealedPureNames[v.PureName] = true
	}

	// kill last \n
	body = strings.TrimRight(body, "\n")

	protoStr := fmt.Sprintf(kMessageFmt, s.Xs.Name, body, s.DefaultIdColType)
	return protoStr
}

func (s *Sheet) GenerateGoCode() string {
	// TODO..
	return ""
}

func (s *Sheet) GenerateCSharpCode() string {
	return ""
}

func panicIfErr(err error, sheetName string, cellName string) {
	if err != nil {
		panic(fmt.Errorf("sheet(%s) cell(%s) err(%v)", sheetName, cellName, err))
	}
}

func getCellValue(protoType string, cell *xlsx.Cell) (v reflect.Value, err error) {
	tFunc, ok := typeReflect[protoType]
	if !ok {
		err = fmt.Errorf("not found typeReflect for %s", protoType)
		return
	}

	t := tFunc()
	v = reflect.New(t).Elem()
	//fmt.Print(t, v)

	switch protoType {
	case "int32", "int", "int64", "long":
		iValue, _err := cell.Int()
		err = _err
		if err != nil {
			break
		}
		v.SetInt(int64(iValue))

	case "uint32", "uint", "uint64", "ulong":
		_iValue, _err := cell.Int()
		err = _err
		if err != nil {
			break
		}
		iValue := uint64(_iValue)
		v.SetUint(uint64(iValue))

	case "string":
		s := cell.String()
		v.SetString(s)

	case "float", "float32", "float64", "double":
		iValue, _err := cell.Float()
		err = _err
		if err != nil {
			break
		}
		v.SetFloat(iValue)

	default:
		err = fmt.Errorf("unsupport %s", protoType)
	}

	return
}

// unmarshal row data to record
func (s *Sheet) ParseRecordData(rowIndex int, rowData *xlsx.Row, record interface{}) {
	t := reflect.TypeOf(record)
	_ = t

	v := reflect.ValueOf(record).Elem()

	//fmt.Println("ParseRecordData", t, v)

	for colIndex := KColStart; colIndex < len(rowData.Cells); colIndex++ {
		c := s.ColMap[colIndex]
		//fmt.Println(colIndex, c)

		cell := rowData.Cells[colIndex]

		cellValue, err := getCellValue(c.ProtoType, cell)
		panicIfErr(err, s.Name, CellName(int32(rowIndex), int32(colIndex)))

		fv := v.FieldByName(c.CamelPureName)

		if c.Decorate != kDecorateAry {
			fv.Set(cellValue)
		} else {
			switch fv.Kind() {
			case reflect.Slice:
				fv.Set(reflect.Append(fv, cellValue))
			default:
				fmt.Println(fv)
				panicIfErr(fmt.Errorf("unknown kind:%v, cellInfo:%+v, pls check toCamelName, maybe un compare", fv.Kind(), c), s.Name, CellName(int32(rowIndex), int32(colIndex)))
			}
		}
	}
}

func CellName(row, col int32) string {
	colStr := make([]byte, 0)
	col += 1
	for col > 0 {
		x := 'A' + byte((col-1)%26)
		colStr = append([]byte{x}, colStr...)
		col /= 26
	}
	return fmt.Sprintf("%s%d", string(colStr), row+1)
}

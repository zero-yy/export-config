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

	IdColType  string
	IdColIndex int

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

		colType, ok := typeParse[originalColType]
		if !ok {
			panic(fmt.Errorf("sheet(%s) typeParse error: unknown type(%s) cell(%s)", xs.Name, originalColType, CellName(kTypeRowIndex, int32(i))))
		}

		pureColName, decorate := getPureName(colName)
		if decorate == kDecorateId {
			s.Id = pureColName
			s.CamelId = toCamelCase(s.Id)
			s.IdColType = colType
			s.IdColIndex = i
		} else if pureColName == kDecorateId {
			if s.Id != pureColName {
				panic(fmt.Errorf("sheet(%s) repeat id %s vs %s", xs.Name, s.Id, pureColName))
			}
			s.Id = pureColName
			s.CamelId = toCamelCase(s.Id)
			s.IdColType = colType
			s.IdColIndex = i
		}

		c := &Col{
			Name:          colName,
			CamelName:     toCamelCase(colName),
			ProtoType:     colType,
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

	protoStr := fmt.Sprintf(kMessageFmt, s.Xs.Name, body, s.IdColType)
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

// unmarshal row data to record
func (s *Sheet) ParseRecordData(rowIndex int, rowData *xlsx.Row, record interface{}) {
	t := reflect.TypeOf(record)
	v := reflect.ValueOf(record).Elem()

	fmt.Println("ParseRecordData", t, v)

	for colIndex := KColStart; colIndex < len(rowData.Cells); colIndex++ {
		c := s.ColMap[colIndex]
		fmt.Println(colIndex, c)

		cell := rowData.Cells[colIndex]

		var _i int64
		var _u uint64
		var _s string
		var _f float64
		//var _v reflect.Value

		switch c.ProtoType {
		case "int32", "int", "int64", "long":
			iValue, err := cell.Int()
			panicIfErr(err, s.Name, CellName(int32(rowIndex), int32(colIndex)))
			_i = int64(iValue)
			//_v = reflect.ValueOf(_i)

		case "uint32", "uint", "uint64", "ulong":
			_iValue, err := cell.Int()
			iValue := uint64(_iValue)
			panicIfErr(err, s.Name, CellName(int32(rowIndex), int32(colIndex)))
			_u = iValue
			//_v = reflect.ValueOf(_u)

		case "string":
			_s = cell.String()
			//_v = reflect.ValueOf(_s)

		case "float", "float32", "float64", "double":
			iValue, err := cell.Float()
			panicIfErr(err, s.Name, CellName(int32(rowIndex), int32(colIndex)))
			_f = iValue
			//_v = reflect.ValueOf(_f)
		}

		fv := v.FieldByName(c.CamelPureName)

		if c.Decorate != kDecorateAry {
			switch fv.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				fv.SetInt(_i)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				fv.SetUint(_u)
			case reflect.String:
				fv.SetString(_s)
			case reflect.Float32, reflect.Float64:
				fv.SetFloat(_f)

			default:
				panicIfErr(fmt.Errorf("unknown kind "), s.Name, CellName(int32(rowIndex), int32(colIndex)))
			}
		} else {
			switch fv.Kind() {
			case reflect.Slice:
				//TODO
				fmt.Println("todo slice")
				//fv.Set(reflect.Append(fv, _v))
			default:
				panicIfErr(fmt.Errorf("unknown kind "), s.Name, CellName(int32(rowIndex), int32(colIndex)))
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

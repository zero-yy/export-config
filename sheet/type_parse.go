package sheet

import (
	"reflect"
	"strconv"
	"strings"
)

var (
	typeParse = map[string]string{
		"int32": "int32",
		"int":   "int32",
		"int64": "int64",
		"long":  "int64",

		"uint32": "uint32",
		"uint":   "uint32",
		"uint64": "uint64",
		"ulong":  "uint64",

		"string": "string",

		"float":   "float",
		"float32": "float",
		"double":  "double",
		"float64": "double",

		// not support now!
		"list<int32>": "repeated int32",
		"list<int>":   "repeated int32",
		"list<int64>": "repeated int64",
		"list<long>":  "repeated int64",

		"list<uint32>": "repeated uint32",
		"list<uint>":   "repeated uint32",
		"list<uint64>": "repeated uint64",
		"list<ulong>":  "repeated uint64",

		"list<string>": "repeated string",

		"list<float>":   "repeated float",
		"list<float32>": "repeated float",
		"list<double>":  "repeated double",
		"list<float64>": "repeated double",
	}

	typeConverToCSharpt = map[string]string{
		"int32":  "int",
		"uint32": "uint",

		"int64":  "long",
		"uint64": "long",
	}

	typeReflect = map[string]func() reflect.Type{
		"int32": func() reflect.Type { return reflect.TypeOf(int32(0)) },
		"int":   func() reflect.Type { return reflect.TypeOf(int32(0)) },
		"int64": func() reflect.Type { return reflect.TypeOf(int64(0)) },
		"long":  func() reflect.Type { return reflect.TypeOf(int64(0)) },

		"uint32": func() reflect.Type { return reflect.TypeOf(uint32(0)) },
		"uint":   func() reflect.Type { return reflect.TypeOf(uint32(0)) },
		"uint64": func() reflect.Type { return reflect.TypeOf(uint64(0)) },
		"ulong":  func() reflect.Type { return reflect.TypeOf(uint64(0)) },

		"string": func() reflect.Type { return reflect.TypeOf("") },

		"float":   func() reflect.Type { return reflect.TypeOf(float32(0)) },
		"float32": func() reflect.Type { return reflect.TypeOf(float32(0)) },
		"double":  func() reflect.Type { return reflect.TypeOf(float64(0)) },
		"float64": func() reflect.Type { return reflect.TypeOf(float64(0)) },

		// not support others!
	}
)

var (
	typeCheck = map[string]func(string) error{
		"int32": checkInt,
		"int":   checkInt,
		"int64": checkInt,
		"long":  checkInt,

		"uint32": checkUint,
		"uint":   checkUint,
		"uint64": checkUint,
		"ulong":  checkUint,

		"string": checkString,

		"float":   checkFloat,
		"float32": checkFloat,
		"double":  checkFloat,
		"float64": checkFloat,

		"list<int32>": checkListWith(checkInt),
		"list<int>":   checkListWith(checkInt),
		"list<int64>": checkListWith(checkInt),
		"list<long>":  checkListWith(checkInt),

		"list<uint32>": checkListWith(checkUint),
		"list<uint>":   checkListWith(checkUint),
		"list<uint64>": checkListWith(checkUint),
		"list<ulong>":  checkListWith(checkUint),

		"list<string>": checkListWith(checkString),

		"list<float>":   checkListWith(checkFloat),
		"list<float32>": checkListWith(checkFloat),
		"list<double>":  checkListWith(checkFloat),
		"list<float64>": checkListWith(checkFloat),
	}
)

type tCheckFunc func(str string) error

func checkInt(str string) error {
	_, err := strconv.ParseInt(str, 10, 0)
	return err
}

func checkUint(str string) error {
	_, err := strconv.ParseUint(str, 10, 0)
	return err
}

func checkString(str string) error {
	return nil
}

func checkFloat(str string) error {
	_, err := strconv.ParseFloat(str, 32)
	return err
}

func checkListWith(check tCheckFunc) tCheckFunc {
	return func(str string) error {
		for _, s := range strings.Split(str, ListSeparator) {
			if len(s) > 0 {
				if err := check(s); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

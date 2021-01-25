package conf

import (
	"sort"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

// sheet name -> loader
// Could use for hot load outside!
var DataLoadFunc = make(map[string]func(dataPath string))

func MustLoad(dataPath string) {
	initLoadFunc()

	// Load by order!
	names := make([]string, 0)
	for n, _ := range DataLoadFunc {
		names = append(names, n)
	}
	sort.Strings(names)

	for _, n := range names {
		DataLoadFunc[n](dataPath)
	}
}

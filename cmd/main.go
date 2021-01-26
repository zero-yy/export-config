package main

import (
	"flag"
	"fmt"
	config "github.com/zero-yy/export-config/internal/config"
	"github.com/zero-yy/export-config/sheet"
)

var (
	configFileName = flag.String("config", "./default.toml", "config file")
)

func main() {
	flag.Parse()
	fmt.Printf("start run export-config --config=%s\n", *configFileName)

	config.MustInit(*configFileName)

	sheet.GenCode()
}

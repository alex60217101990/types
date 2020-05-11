package main

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/alex60217101990/types/configs"
	"github.com/alex60217101990/types/enums"
)

func main() {
	var c configs.Configs
	var tests = map[enums.FormatType]string{
		enums.Json: "./configs/configs.json",
		enums.Yaml: "./configs/configs.yaml",
	}
	for tk, tv := range tests {
		err := configs.ReadConfigFile(&c, tk, tv)
		if err != nil {
			fmt.Println(err)
			debug.PrintStack()
			continue
		}
		log.Printf("%#v\n", c.Firewall)
	}
}

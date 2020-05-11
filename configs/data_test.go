package configs

import (
	"fmt"
	"testing"

	"github.com/alex60217101990/types/enums"
)

func TestPrintTestConfigs(t *testing.T) {
	var c Configs
	var err error
	var tests = map[enums.FormatType]string{
		enums.Json: "./configs.json",
		enums.Yaml: "./configs.yaml",
	}
	for tk, tv := range tests {
		testname := fmt.Sprintf("%s", tk)
		t.Run(testname, func(t *testing.T) {
			err = c.PrintTestConfigs(tk, tv)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

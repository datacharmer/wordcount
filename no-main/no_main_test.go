package no_main

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestScriptGeneric(t *testing.T) {
	for i := 1; i <= 4; i++ {
		directory := fmt.Sprintf("testdata/%d", i)
		files, err := filepath.Glob(directory + "/*.tx*")
		if err == nil && len(files) > 0 {
			t.Run(directory, func(t *testing.T) {
				testscript.Run(t, testscript.Params{
					Dir: directory,
				})
			})
		}
	}
}

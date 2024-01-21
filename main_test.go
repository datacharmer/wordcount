package main

import (
	"os"
	"testing"

	"github.com/datacharmer/wordcount/cmd"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestWordCount(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir:                 "testdata/wordcount",
		RequireExplicitExec: true,
	})
}

func TestMain(m *testing.M) {
	exitCode := testscript.RunMain(m, map[string]func() int{
		"wordcount": cmd.RunMain,
	})
	os.Exit(exitCode)
}

func TestWordCountAdvanced(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir:                 "testdata/advanced",
		Condition:           customConditions,
		Cmds:                customCommands(),
		RequireExplicitExec: true,
		Setup: func(env *testscript.Env) error {
			env.Setenv("WORDCOUNT_VERSION", cmd.Version)
			return nil
		},
	})
}

package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	monkey "github.com/agiledragon/gomonkey/v2"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestRootCommand(t *testing.T) {
	viper.AddConfigPath("../test/")

	t.Run("when command ends with an error", func(t *testing.T) {
		exitCalled := false
		osExit := func(code int) {
			if code == 1 {
				exitCalled = true
			}
		}
		patch := monkey.ApplyFunc(os.Exit, osExit)
		defer patch.Reset()

		commandToRun = func(_ io.Writer) bool {
			return false
		}

		rootCmd.Execute()
		assert.True(t, exitCalled, "Expects a call to exit")
	})

	t.Run("when command ends without an error", func(t *testing.T) {
		exitCalled := false
		osExit := func(code int) {
			if code == 1 {
				exitCalled = true
			}
		}
		patch := monkey.ApplyFunc(os.Exit, osExit)
		defer patch.Reset()
		commandToRun = func(_ io.Writer) bool {
			return true
		}

		rootCmd.Execute()
		assert.False(t, exitCalled, "Expects to finish")
	})

	t.Run("checks configuration file", func(t *testing.T) {
		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.Execute()

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		cmdOutput := string(out)
		if !strings.Contains(cmdOutput, "Running arch-go command") {
			t.Fatal("Expects a log containing the running command.")
		}
		if !strings.Contains(cmdOutput, "Using config file:") || !strings.Contains(cmdOutput, "/test/arch-go.yml") {
			t.Fatal("Expects a log containing the configuration file.")
		}
	})

	t.Run("Force an error trying to get current directory", func(t *testing.T) {
		exitCalled := false
		osExit := func(code int) {
			if code == 1 {
				exitCalled = true
			}
		}

		patch := monkey.ApplyFuncReturn(os.Getwd, nil, fmt.Errorf("foobar"))
		defer patch.Reset()
		patchExit := monkey.ApplyFunc(os.Exit, osExit)
		defer patchExit.Reset()

		rootCmd.Execute()

		if !exitCalled {
			t.Fatal("Expects to call os.Exit when arch-go is not able to get current directory.")
		}
	})
}

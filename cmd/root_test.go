package cmd

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	monkey "github.com/agiledragon/gomonkey/v2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/v2/api"
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
		patch := monkey.ApplyFunc(os.Exit, func(code int) {
			if code == 1 {
				exitCalled = true
			}
		})
		commandToRun = func(_ io.Writer) bool {
			return true
		}

		defer patch.Reset()

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

		if !strings.Contains(cmdOutput, "Using configuration file:") || !strings.Contains(cmdOutput, "/test/arch-go.yml") {
			t.Fatal("Expects a log containing the configuration file.")
		}
	})

	t.Run("checks if successful run returns zero exit code", func(t *testing.T) {
		nonZeroExitCode := false
		passingAPIResult := &api.Result{
			Time: time.Now(),
			Pass: true,
		}
		patchCheck := monkey.ApplyFuncReturn(api.CheckArchitecture, passingAPIResult)
		patchExit := monkey.ApplyFunc(os.Exit, func(code int) {
			if code != 0 {
				nonZeroExitCode = true
			}
		})

		rootCmd.SetArgs([]string{
			"",
		})

		defer patchCheck.Reset()
		defer patchExit.Reset()

		rootCmd.Execute()

		if nonZeroExitCode {
			t.Fatal("Expects to call os.Exit with exit code 0 when arch-go validates successfully.")
		}
	})

	t.Run("checks if not successful run returns non-zero exit code", func(t *testing.T) {
		viper.Reset()

		nonZeroExitCode := false
		nonPassingAPIResult := &api.Result{
			Time: time.Now(),
			Pass: false,
		}
		patchCheck := monkey.ApplyFuncReturn(api.CheckArchitecture, nonPassingAPIResult)
		patchExit := monkey.ApplyFunc(os.Exit, func(code int) {
			if code != 0 {
				nonZeroExitCode = true
			}
		})

		defer patchCheck.Reset()
		defer patchExit.Reset()

		rootCmd.SetArgs([]string{
			"",
		})

		err := rootCmd.Execute()

		if !nonZeroExitCode && err == nil {
			t.Fatal("Expects to call os.Exit with exit code 1 when arch-go does find errors when validating.")
		}
	})

	t.Run("Force an error trying to get current directory", func(t *testing.T) {
		exitCalled := false
		patch := monkey.ApplyFuncReturn(os.Getwd, nil, errors.New("foobar"))
		patchExit := monkey.ApplyFunc(os.Exit, func(code int) {
			if code == 1 {
				exitCalled = true
			}
		})

		defer patch.Reset()
		defer patchExit.Reset()

		rootCmd.Execute()

		if !exitCalled {
			t.Fatal("Expects to call os.Exit when arch-go is not able to get current directory.")
		}
	})
}

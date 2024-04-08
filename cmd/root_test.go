package cmd

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestRootCommand(t *testing.T) {
	viper.AddConfigPath("../test/")

	t.Run("when command ends with an error", func(t *testing.T) {
		exitCalled := false
		osExit = func(code int) {
			if code != 1 {
				t.Fatalf("Expects an error")
			}
			exitCalled = true
		}
		commandToRun = func() bool {
			return false
		}

		Execute()
		assert.True(t, exitCalled, "Expects a call to exit")
	})

	t.Run("when command ends without an error", func(t *testing.T) {
		exitCalled := false
		osExit = func(code int) {
			if code != 1 {
				t.Fatalf("Expects an error")
			}
			exitCalled = true
		}
		commandToRun = func() bool {
			return true
		}

		Execute()
		assert.False(t, exitCalled, "Expects to finish")
	})

	t.Run("checks configuration file", func(t *testing.T) {
		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		Execute()

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		cmdOutput := string(out)
		//		fmt.Printf(cmdOutput)
		if !strings.Contains(cmdOutput, "Running arch-go command") {
			t.Fatal("Expects a log containing the running command.")
		}
		if !strings.Contains(cmdOutput, "Using config file:") || !strings.Contains(cmdOutput, "/test/arch-go.yml") {
			t.Fatal("Expects a log containing the configuration file.")
		}
	})
}

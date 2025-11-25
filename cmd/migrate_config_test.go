package cmd

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"

	monkey "github.com/agiledragon/gomonkey/v2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/arch-go/arch-go/v2/api/configuration"
	"github.com/arch-go/arch-go/v2/internal/utils/values"
)

func TestMigrateConfigCommand(t *testing.T) {
	viper.AddConfigPath("../test/")

	t.Run("when arch-go.yaml contains current schema but has no rules", func(t *testing.T) {
		var exitCode int

		cmd := NewMigrateConfigCommand()
		patch := monkey.ApplyFuncReturn(configuration.LoadConfig, &configuration.Config{}, nil)
		patchExit := monkey.ApplyFunc(os.Exit, func(c int) { exitCode = c })

		defer patch.Reset()
		defer patchExit.Reset()

		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.Execute()

		out, _ := io.ReadAll(b)
		expected := `Invalid Configuration: configuration file should have at least one rule
`
		assert.Equal(t, expected, string(out))
		assert.Equal(t, 1, exitCode)
	})

	t.Run("when arch-go.yaml contains current schema and has rules", func(t *testing.T) {
		var exitCode int

		cmd := NewMigrateConfigCommand()
		currentConfiguration := configuration.Config{
			FunctionsRules: []*configuration.FunctionsRule{
				{
					Package:  "foobar",
					MaxLines: values.GetIntRef(10),
				},
			},
		}
		patch := monkey.ApplyFuncReturn(configuration.LoadConfig, &currentConfiguration, nil)
		patchExit := monkey.ApplyFunc(os.Exit, func(c int) { exitCode = c })

		defer patch.Reset()
		defer patchExit.Reset()

		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.Execute()

		out, _ := io.ReadAll(b)

		expected := `File is already compatible with version 1
`
		assert.Equal(t, expected, string(out))
		assert.Equal(t, 0, exitCode)
	})

	t.Run("when arch-go.yaml contains old schema", func(t *testing.T) {
		var exitCode int

		cmd := NewMigrateConfigCommand()
		patch := monkey.ApplyFuncReturn(configuration.LoadConfig, nil, errors.New("not found"))
		patchOldSchema := monkey.ApplyFuncReturn(configuration.LoadDeprecatedConfig, &configuration.DeprecatedConfig{}, nil)
		patchExit := monkey.ApplyFunc(os.Exit, func(c int) { exitCode = c })

		defer patch.Reset()
		defer patchOldSchema.Reset()
		defer patchExit.Reset()

		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.Execute()

		out, _ := io.ReadAll(b)
		expected := `Migrating deprecated configuration to current schema.
Deprecated configuration backup at: old-arch-go.yml
Configuration saved at: arch-go.yml
`
		_, err1 := os.Stat("old-arch-go.yml")
		os.Remove("old-arch-go.yml")

		_, err2 := os.Stat("arch-go.yml")
		os.Remove("arch-go.yml")

		require.NoError(t, err1, string(out))
		require.NoError(t, err2, string(out))
		assert.Equal(t, expected, string(out))
		assert.Equal(t, 0, exitCode)
	})

	t.Run("when arch-go.yaml cannot be loaded", func(t *testing.T) {
		var exitCode int

		cmd := NewMigrateConfigCommand()
		patch := monkey.ApplyFuncReturn(configuration.LoadConfig, nil, errors.New("not found"))
		patchOldSchema := monkey.ApplyFuncReturn(configuration.LoadDeprecatedConfig, nil, errors.New("not loaded"))
		patchExit := monkey.ApplyFunc(os.Exit, func(c int) { exitCode = c })

		defer patch.Reset()
		defer patchOldSchema.Reset()
		defer patchExit.Reset()

		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.Execute()

		out, _ := io.ReadAll(b)
		expected := `Error: not loaded
`
		assert.Equal(t, expected, string(out))
		assert.Equal(t, 1, exitCode)
	})
}

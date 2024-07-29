package packages

import (
	"fmt"
	"os"

	"golang.org/x/mod/modfile"
)

const goModFile = "go.mod"

func GetMainPackage() (string, error) {
	if _, err := os.Stat(goModFile); err != nil {
		return "", fmt.Errorf("could not load %s file. %w", goModFile, err)
	}

	content, _ := os.ReadFile(goModFile)
	modulePath := modfile.ModulePath(content)

	return modulePath, nil
}

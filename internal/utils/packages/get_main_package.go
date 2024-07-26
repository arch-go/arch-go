package packages

import (
	"fmt"
	"os"

	"golang.org/x/mod/modfile"
)

const goModFile = "go.mod"

func GetMainPackage() (string, error) {
	if _, err := os.Stat(goModFile); err != nil {
		return "", fmt.Errorf("Could not load %s file. %s\n", goModFile, err.Error())
	}

	content, _ := os.ReadFile(goModFile)
	modulePath := modfile.ModulePath(content)

	return modulePath, nil
}

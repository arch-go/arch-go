package packages

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"os"
)

const goModFile = "go.mod"

func GetMainPackage() (string, error) {
	if _, err := os.Stat(goModFile); err == nil {
		content, _ := os.ReadFile(goModFile)
		modulePath := modfile.ModulePath(content)
		return modulePath, nil
	} else {
		return "", fmt.Errorf("Could not load %s file. %s\n", goModFile, err.Error())
	}
}

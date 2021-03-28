package packages

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"io/ioutil"
	"os"
)

const goModFile = "go.mod"

func GetMainPackage() (string, error) {
	if _, err := os.Stat(goModFile); err == nil {
		content, _ := ioutil.ReadFile(goModFile)
		modulePath := modfile.ModulePath(content)
		return modulePath, nil
	} else {
		return "", fmt.Errorf("Could not load %s file. %s\n", goModFile, err.Error())
	}
}

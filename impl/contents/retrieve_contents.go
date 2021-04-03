package contents

import (
	"github.com/fdaines/arch-go/model"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func retrieveContents(pkg *model.PackageInfo, mainPackage string) (*PackageContents, error) {
	var methods, functions, interfaces, types int
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	packageDir := strings.Replace(pkg.PackageData.ImportPath, mainPackage, path, 1)

	for _, srcFile := range pkg.PackageData.GoFiles {
		data, err := ioutil.ReadFile(filepath.Join(packageDir, srcFile))
		if err != nil {
			return nil, err
		}
		fileset := token.NewFileSet()
		node, err := parser.ParseFile(fileset, srcFile, data, 0)
		if err != nil {
			return nil, err
		}
		ast.Inspect(node, func(n ast.Node) bool {
			switch t := n.(type) {
			case *ast.FuncDecl:
				if t.Recv != nil {
					methods++
				} else {
					functions++
				}
			case *ast.InterfaceType:
				interfaces++
			case *ast.StructType:
				types++
			}
			return true
		})
	}

	return &PackageContents{
		Methods:    methods,
		Functions:  functions,
		Interfaces: interfaces,
		Types:      types,
	}, nil
}

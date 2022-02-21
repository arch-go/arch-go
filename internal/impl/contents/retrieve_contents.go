package contents

import (
	"github.com/fdaines/arch-go/internal/model"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func retrieveContents(pkg *model.PackageInfo, mainPackage string) (*PackageContents, error) {
	var methods, functions, interfaces, structs int
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	packageDir := strings.Replace(pkg.PackageData.ImportPath, mainPackage, path, 1)

	packageContents := &PackageContents{}
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
		packageContents = inspectFile(node, packageContents)
	}

	return &PackageContents{
		Methods:    methods,
		Functions:  functions,
		Interfaces: interfaces,
		Structs:    structs,
	}, nil
}

func inspectFile(node *ast.File, contents *PackageContents) *PackageContents {
	ast.Inspect(node, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.FuncDecl:
			if t.Recv != nil {
				contents.Methods++
			} else {
				contents.Functions++
			}
		case *ast.InterfaceType:
			if t.Methods != nil && len(t.Methods.List) > 0 {
				contents.Interfaces++
			}
		case *ast.StructType:
			contents.Structs++
		}
		return true
	})
	return contents
}

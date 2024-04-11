package contents

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"github.com/fdaines/arch-go/internal/model"
)

func retrieveContents(pkg *model.PackageInfo) (*PackageContents, error) {
	packageContents := &PackageContents{}
	for _, srcFile := range pkg.PackageData.GoFiles {
		data, err := os.ReadFile(filepath.Join(pkg.PackageData.Dir, srcFile))
		if err != nil {
			return nil, err
		}
		fileset := token.NewFileSet()
		node, err := parser.ParseFile(fileset, srcFile, data, 0)
		if err != nil {
			return nil, err
		}
		inspectFile(node, packageContents)
	}

	return packageContents, nil
}

func inspectFile(node *ast.File, contents *PackageContents) {
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
}

package functions

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/arch-go/arch-go/internal/model"
	"github.com/arch-go/arch-go/internal/utils/packages"
)

func RetrieveFunctions(pkg *model.PackageInfo, mainPackage string) ([]*FunctionDetails, error) {
	var functionDetailsCollection []*FunctionDetails

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	packageDir := strings.Replace(pkg.PackageData.ImportPath, mainPackage, path, 1)

	for _, srcFile := range pkg.PackageData.GoFiles {
		srcFilePath := filepath.Join(pkg.Path, srcFile)

		data, err := os.ReadFile(filepath.Join(packageDir, srcFile))
		if err != nil {
			return nil, err
		}

		fileset := token.NewFileSet()

		node, err := parser.ParseFile(fileset, srcFile, data, 0)
		if err != nil {
			return nil, err
		}

		functionDetailsCollection = resolveFunctionDetails(node, srcFile, srcFilePath, fileset, functionDetailsCollection)
	}

	return functionDetailsCollection, nil
}

func resolveFunctionDetails(
	node *ast.File,
	srcFile string,
	srcFilePath string,
	fileset *token.FileSet,
	functionDetailsCollection []*FunctionDetails,
) []*FunctionDetails {
	ast.Inspect(node, func(n ast.Node) bool {
		if typ, ok := n.(*ast.FuncDecl); ok {
			functionDetails := &FunctionDetails{
				File:     srcFile,
				FilePath: srcFilePath,
				Name:     typ.Name.Name,
				IsPublic: packages.IsPublic(typ.Name.Name),
				NumLines: fileset.Position(typ.End()).Line - fileset.Position(typ.Pos()).Line,
			}

			if typ.Type.Params != nil {
				functionDetails.NumParams = typ.Type.Params.NumFields()
			}

			if typ.Type.Results != nil {
				functionDetails.NumReturns = typ.Type.Results.NumFields()
			}

			functionDetailsCollection = append(functionDetailsCollection, functionDetails)
		}

		return true
	})

	return functionDetailsCollection
}

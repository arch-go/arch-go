package naming

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"github.com/arch-go/arch-go/internal/model"
)

func getInterfacesMatching(pkg *model.PackageInfo, pattern string) ([]InterfaceDescription, error) {
	var interfaces []InterfaceDescription

	comparator, patternValue := getPatternComparator(pattern)

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

		interfaces = retrieveMatchingInterfaces(node, comparator, patternValue, data, interfaces)
	}

	return interfaces, nil
}

func retrieveMatchingInterfaces(
	node *ast.File,
	comparator func(s string, prefix string) bool,
	patternValue string,
	data []byte,
	interfaces []InterfaceDescription,
) []InterfaceDescription {
	ast.Inspect(node, func(n ast.Node) bool {
		if typ, ok := n.(*ast.GenDecl); ok {
			if typ.Tok.String() == "type" {
				// ok if panics
				ts := typ.Specs[0].(*ast.TypeSpec) //nolint: forcetypeassert

				it, ok := ts.Type.(*ast.InterfaceType)
				if ok && comparator(ts.Name.String(), patternValue) {
					currentInterface := InterfaceDescription{
						Name: ts.Name.String(),
					}

					for _, m := range it.Methods.List {
						// ok if panics
						method := m.Type.(*ast.FuncType) //nolint: forcetypeassert

						currentInterface.Methods = append(currentInterface.Methods, MethodDescription{
							Name:         m.Names[0].String(),
							Parameters:   getParameters(string(data), method.Params),
							ReturnValues: getReturnValues(string(data), method.Results),
						})
					}

					interfaces = append(interfaces, currentInterface)
				}
			}
		}

		return true
	})

	return interfaces
}

func getStructsWithMethods(pkg *model.PackageInfo) ([]StructDescription, error) {
	var structs []StructDescription

	structsMap := make(map[string][]*ast.FuncDecl)
	fileStructsMap := make(map[*ast.FuncDecl]string)

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

		ast.Inspect(node, func(n ast.Node) bool {
			ft, ok := n.(*ast.FuncDecl)
			if ok && ft.Recv != nil {
				structName := resolveStructName(ft)
				structsMap[structName] = append(structsMap[structName], ft)
				fileStructsMap[ft] = string(data)
			}

			return true
		})
	}

	structs = resolveStructsDescription(structsMap, fileStructsMap, structs)

	return structs, nil
}

func resolveStructsDescription(
	structsMap map[string][]*ast.FuncDecl, fileStructsMap map[*ast.FuncDecl]string, structs []StructDescription,
) []StructDescription {
	if len(structsMap) > 0 {
		for k, v := range structsMap {
			currentStruct := StructDescription{
				Name: k,
			}

			for _, vx := range v {
				md := MethodDescription{
					Name:         vx.Name.String(),
					Parameters:   getParameters(fileStructsMap[vx], vx.Type.Params),
					ReturnValues: getReturnValues(fileStructsMap[vx], vx.Type.Results),
				}

				currentStruct.Methods = append(currentStruct.Methods, md)
			}

			structs = append(structs, currentStruct)
		}
	}

	return structs
}

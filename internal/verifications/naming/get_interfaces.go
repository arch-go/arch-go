package naming

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/arch-go/arch-go/internal/model"
	"github.com/arch-go/arch-go/internal/utils/output"
	"github.com/arch-go/arch-go/internal/utils/packages"
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
					interfaceDescription := InterfaceDescription{
						Name: ts.Name.String(),
					}
					retrieveMethods(&interfaceDescription, it, data, node)
					interfaces = append(interfaces, interfaceDescription)
				}
			}
		}

		return true
	})

	return interfaces
}

func retrieveMethods(currentInterface *InterfaceDescription, it *ast.InterfaceType, data []byte, node *ast.File) {
	for _, field := range it.Methods.List {
		switch tp := field.Type.(type) {
		case *ast.FuncType:
			retrieveNormalMethods(currentInterface, field, tp, data)
		case *ast.Ident:
			retrieveMethodsFromEmbeddedInterface(currentInterface, tp, data, node)
		case *ast.IndexExpr:
			retrieveGenericsMethods(currentInterface, tp, data, node)
		case *ast.SelectorExpr:
			retrieveOtherPackageMethods(currentInterface, tp, node)
		default:
			panic(fmt.Sprintf("unknown type %T to retrieve method", tp))
		}
	}
}

func retrieveNormalMethods(currentInterface *InterfaceDescription, field *ast.Field, ft *ast.FuncType, data []byte) {
	addMethods(currentInterface, field.Names[0].String(), ft, data)
}

func addMethods(currentInterface *InterfaceDescription, name string, method *ast.FuncType, data []byte) {
	currentInterface.Methods = append(currentInterface.Methods, MethodDescription{
		Name:         name,
		Parameters:   getParameters(string(data), method.Params),
		ReturnValues: getReturnValues(string(data), method.Results),
	})
}

func retrieveMethodsFromEmbeddedInterface(currentInterface *InterfaceDescription, ident *ast.Ident,
	data []byte, node *ast.File,
) {
	it := ident.Obj.Decl.(*ast.TypeSpec).Type.(*ast.InterfaceType) //nolint: forcetypeassert
	retrieveMethods(currentInterface, it, data, node)
}

func retrieveGenericsMethods(currentInterface *InterfaceDescription, ie *ast.IndexExpr, data []byte,
	node *ast.File,
) {
	ident, ok := ie.X.(*ast.Ident)
	if !ok {
		se := ie.X.(*ast.SelectorExpr) //nolint: forcetypeassert
		retrieveOtherPackageMethods(currentInterface, se, node)

		return
	}

	it := ident.Obj.Decl.(*ast.TypeSpec).Type.(*ast.InterfaceType) //nolint: forcetypeassert
	retrieveMethods(currentInterface, it, data, node)
}

func retrieveOtherPackageMethods(currentInterface *InterfaceDescription, se *ast.SelectorExpr, node *ast.File) {
	ident := se.X.(*ast.Ident) //nolint: forcetypeassert

	var importPath string

	for _, imp := range node.Imports {
		// checks if import has an alias name
		if imp.Name != nil {
			if ident.Name == imp.Name.Name {
				importPath = strings.Trim(imp.Path.Value, "\"")

				break
			}
		} else {
			path := strings.Trim(imp.Path.Value, "\"")
			if strings.HasSuffix(path, ident.Name) {
				importPath = path

				break
			}
		}
	}

	if importPath == "" {
		panic("import path should not be empty")
	}

	pkgs, _ := packages.GetBasicPackagesInfo(importPath, output.CreateNilWriter(), false)
	i := slices.IndexFunc(pkgs, func(pkg *model.PackageInfo) bool {
		return pkg.Path == importPath
	})

	if i == -1 {
		panic("import should have been found")
	}

	interfaces, _ := getInterfacesMatching(pkgs[i], se.Sel.Name)
	for _, inf := range interfaces {
		for _, met := range inf.Methods {
			currentInterface.Methods = append(currentInterface.Methods, MethodDescription{
				Name:         met.Name,
				Parameters:   met.Parameters,
				ReturnValues: met.ReturnValues,
			})
		}
	}
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

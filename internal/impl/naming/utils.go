package naming

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/utils/packages"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func getInterfacesMatching(module string, pattern string) ([]InterfaceDescription, error) {
	var interfaces []InterfaceDescription

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pkgs, _ := packages.GetBasicPackagesInfo(false)
	comparator, patternValue := getPatternComparator(pattern)
	for _, pkg := range pkgs {
		packageDir := strings.Replace(pkg.PackageData.ImportPath, module, path, 1)
		for _, srcFile := range pkg.PackageData.GoFiles {
			data, err := os.ReadFile(filepath.Join(packageDir, srcFile))
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
	}

	return interfaces, nil
}

func retrieveMatchingInterfaces(node *ast.File, comparator func(s string, prefix string) bool, patternValue string, data []byte, interfaces []InterfaceDescription) []InterfaceDescription {
	ast.Inspect(node, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.GenDecl:
			if t.Tok.String() == "type" {
				ts := t.Specs[0].(*ast.TypeSpec)
				it, ok := ts.Type.(*ast.InterfaceType)
				if ok && comparator(ts.Name.String(), patternValue) {
					currentInterface := InterfaceDescription{
						Name: ts.Name.String(),
					}
					for _, m := range it.Methods.List {
						method := m.Type.(*ast.FuncType)
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

func getStructsWithMethods(module string, pkg model.PackageVerification) ([]StructDescription, error) {
	var structs []StructDescription
	structsMap := make(map[string][]*ast.FuncDecl)
	fileStructsMap := make(map[*ast.FuncDecl]string)

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	packageDir := strings.Replace(pkg.Package.PackageData.ImportPath, module, path, 1)
	for _, srcFile := range pkg.Package.PackageData.GoFiles {
		data, err := os.ReadFile(filepath.Join(packageDir, srcFile))
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

func resolveStructsDescription(structsMap map[string][]*ast.FuncDecl, fileStructsMap map[*ast.FuncDecl]string, structs []StructDescription) []StructDescription {
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

func resolveStructName(ft *ast.FuncDecl) string {
	se, ok := ft.Recv.List[0].Type.(*ast.StarExpr)
	if ok {
		return fmt.Sprintf("*%v", se.X)
	}
	ie, ok := ft.Recv.List[0].Type.(*ast.Ident)
	if ok {
		return fmt.Sprintf("%v", ie.Name)
	}
	return ""
}

func getPatternComparator(pattern string) (func(s string, prefix string) bool, string) {
	if strings.HasPrefix(pattern, "*") {
		return strings.HasSuffix, pattern[1:]
	}
	if strings.HasSuffix(pattern, "*") {
		return strings.HasPrefix, pattern[:len(pattern)-1]
	}
	return strings.EqualFold, pattern
}

func getReturnValues(fileContent string, results *ast.FieldList) []string {
	if results == nil || results.List == nil {
		return []string{}
	}
	returnValues := make([]string, results.NumFields(), results.NumFields())
	for index, p := range results.List {
		returnValues[index] = fmt.Sprintf("%s", fileContent[p.Type.Pos()-1:p.Type.End()-1])
	}

	return returnValues
}

func getParameters(fileContent string, params *ast.FieldList) []string {
	if params == nil || params.List == nil {
		return []string{}
	}
	parameters := make([]string, params.NumFields(), params.NumFields())
	for index, p := range params.List {
		parameters[index] = fmt.Sprintf("%s", fileContent[p.Type.Pos()-1:p.Type.End()-1])
	}
	return parameters
}

func implementsInterface(s StructDescription, i InterfaceDescription) bool {
	methodsLeft := len(i.Methods)
	if len(s.Methods) < methodsLeft {
		return false
	}

	for _, sm := range s.Methods {
		for _, im := range i.Methods {
			if sm.Name == im.Name {
				paramsComplies := areEquals(sm.Parameters, im.Parameters)
				returnValuesComplies := areEquals(sm.ReturnValues, im.ReturnValues)

				if paramsComplies && returnValuesComplies {
					methodsLeft--
				}
			}
		}
	}

	return methodsLeft == 0
}

func areEquals(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

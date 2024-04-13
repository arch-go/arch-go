package naming

import (
	"fmt"
	"go/ast"
	"strings"
)

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
	if results == nil {
		return []string{}
	}
	returnValues := make([]string, results.NumFields(), results.NumFields())
	for index, p := range results.List {
		returnValues[index] = fmt.Sprintf("%s", fileContent[p.Type.Pos()-1:p.Type.End()-1])
	}

	return returnValues
}

func getParameters(fileContent string, params *ast.FieldList) []string {
	if params == nil {
		return []string{}
	}
	parameters := make([]string, params.NumFields(), params.NumFields())
	for index, p := range params.List {
		parameters[index] = fmt.Sprintf("%s", fileContent[p.Type.Pos()-1:p.Type.End()-1])
	}
	return parameters
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

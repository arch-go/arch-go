package naming

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveStructsDescription(t *testing.T) {
	t.Run("test empty inputs", func(t *testing.T) {
		var (
			inputStructs         []StructDescription
			expectedDescriptions []StructDescription
		)

		inputMap := make(map[string][]*ast.FuncDecl)
		inputFileStructMap := make(map[*ast.FuncDecl]string)

		result := resolveStructsDescription(inputMap, inputFileStructMap, inputStructs)

		assert.Equal(t, expectedDescriptions, result)
	})

	t.Run("test filled inputs", func(t *testing.T) {
		fn1 := &ast.FuncDecl{
			Name: &ast.Ident{Name: "funcName1"},
			Type: &ast.FuncType{
				Func: token.Pos(1),
				Params: &ast.FieldList{
					Closing: token.Pos(10),
					List: []*ast.Field{
						{
							Type: &ast.Ident{
								Name:    "foo1",
								NamePos: token.Pos(5),
							},
						},
					},
				},
			},
		}
		fn2 := &ast.FuncDecl{
			Name: &ast.Ident{Name: "funcName2"},
			Type: &ast.FuncType{
				Func: token.Pos(1),
				Params: &ast.FieldList{
					Closing: token.Pos(10),
					List: []*ast.Field{
						{
							Type: &ast.Ident{
								Name:    "bar1",
								NamePos: token.Pos(5),
							},
						},
						{
							Type: &ast.Ident{
								Name:    "bar2",
								NamePos: token.Pos(25),
							},
						},
					},
				},
			},
		}
		inputMap := make(map[string][]*ast.FuncDecl)
		inputMap["foobar"] = []*ast.FuncDecl{fn1, fn2}

		inputFileStructMap := make(map[*ast.FuncDecl]string)
		inputFileStructMap[fn1] = "jglfsjgskdjgaiendkcnoizcmoewacfoeijoaiewcmjoidvndfvndijfnvienofickewdcmewkprpotrtyea"
		inputFileStructMap[fn2] = "turyeproivpotirevogiv49gprtwig9045wrwifm4givbg94tnugb09rif3490cj09di29wexmqj9ef09q34"
		inputStructs := []StructDescription{}

		expectedDescriptions := []StructDescription{
			{
				Name: "foobar",
				Methods: []MethodDescription{
					{
						Name:         "funcName1",
						Parameters:   []string{"sjgs"},
						ReturnValues: []string{},
					},
					{
						Name:         "funcName2",
						Parameters:   []string{"epro", "prtw"},
						ReturnValues: []string{},
					},
				},
			},
		}

		result := resolveStructsDescription(inputMap, inputFileStructMap, inputStructs)

		assert.Equal(t, expectedDescriptions, result)
	})
}

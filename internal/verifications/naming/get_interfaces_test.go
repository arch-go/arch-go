package naming

import (
	"go/ast"
	"go/token"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/arch-go/arch-go/v2/internal/model"
	"github.com/arch-go/arch-go/v2/internal/utils/output"
	"github.com/arch-go/arch-go/v2/internal/utils/packages"
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

func TestGetInterfacesMatching(t *testing.T) {
	pkgPath := "github.com/arch-go/arch-go/v2/internal/verifications/naming/test/pkgone"
	pkgs, err := packages.GetBasicPackagesInfo(pkgPath, output.CreateNilWriter(), false)
	require.NoError(t, err)

	i := slices.IndexFunc(pkgs, func(pkg *model.PackageInfo) bool {
		return pkg.Path == pkgPath
	})
	pkg := pkgs[i]

	t.Run("Get interfaces from normal interface", func(t *testing.T) {
		interfaces, err := getInterfacesMatching(pkg, "Service")
		require.NoError(t, err)

		assert.Len(t, interfaces, 1)
		assert.Equal(t, "Service", interfaces[0].Name)

		assert.Equal(t, "NormalFunc", interfaces[0].Methods[0].Name)
		assert.Equal(t, []string{"[]string"}, interfaces[0].Methods[0].Parameters)
		assert.Equal(t, []string{"int", "error"}, interfaces[0].Methods[0].ReturnValues)

		assert.Equal(t, "EmbeddedGenericsFunc", interfaces[0].Methods[1].Name)
		assert.Equal(t, []string{"T"}, interfaces[0].Methods[1].Parameters)
		assert.Equal(t, []string{"T", "error"}, interfaces[0].Methods[1].ReturnValues)

		assert.Equal(t, "OtherPackageEmbeddedGenericsFunc", interfaces[0].Methods[2].Name)
		assert.Equal(t, []string{"T"}, interfaces[0].Methods[2].Parameters)
		assert.Equal(t, []string{}, interfaces[0].Methods[2].ReturnValues)

		assert.Equal(t, "EmbeddedNormalFunc", interfaces[0].Methods[3].Name)

		assert.Equal(t, "OtherPackageEmbeddedNormalFunc", interfaces[0].Methods[4].Name)

		assert.Equal(t, "Write", interfaces[0].Methods[5].Name)
		assert.Equal(t, []string{"[]byte"}, interfaces[0].Methods[5].Parameters)
		assert.Equal(t, []string{"int", "error"}, interfaces[0].Methods[5].ReturnValues)

		assert.Equal(t, "MarshalJSON", interfaces[0].Methods[6].Name)
		assert.Equal(t, []string{}, interfaces[0].Methods[6].Parameters)
		assert.Equal(t, []string{"[]byte", "error"}, interfaces[0].Methods[6].ReturnValues)
	})

	t.Run("Get interfaces from generics interface", func(t *testing.T) {
		interfaces, err := getInterfacesMatching(pkg, "ServiceWithGenerics")
		require.NoError(t, err)

		assert.Len(t, interfaces, 1)
		assert.Equal(t, "ServiceWithGenerics", interfaces[0].Name)

		assert.Equal(t, "NormalFunc", interfaces[0].Methods[0].Name)
		assert.Equal(t, []string{"[]string"}, interfaces[0].Methods[0].Parameters)
		assert.Equal(t, []string{"int", "error"}, interfaces[0].Methods[0].ReturnValues)

		assert.Equal(t, "EmbeddedGenericsFunc", interfaces[0].Methods[1].Name)
		assert.Equal(t, []string{"T"}, interfaces[0].Methods[1].Parameters)
		assert.Equal(t, []string{"T", "error"}, interfaces[0].Methods[1].ReturnValues)

		assert.Equal(t, "OtherPackageEmbeddedGenericsFunc", interfaces[0].Methods[2].Name)
		assert.Equal(t, []string{"T"}, interfaces[0].Methods[2].Parameters)
		assert.Equal(t, []string{}, interfaces[0].Methods[2].ReturnValues)

		assert.Equal(t, "EmbeddedNormalFunc", interfaces[0].Methods[3].Name)

		assert.Equal(t, "OtherPackageEmbeddedNormalFunc", interfaces[0].Methods[4].Name)

		assert.Equal(t, "Write", interfaces[0].Methods[5].Name)
		assert.Equal(t, []string{"[]byte"}, interfaces[0].Methods[5].Parameters)
		assert.Equal(t, []string{"int", "error"}, interfaces[0].Methods[5].ReturnValues)

		assert.Equal(t, "MarshalJSON", interfaces[0].Methods[6].Name)
		assert.Equal(t, []string{}, interfaces[0].Methods[6].Parameters)
		assert.Equal(t, []string{"[]byte", "error"}, interfaces[0].Methods[6].ReturnValues)
	})
}

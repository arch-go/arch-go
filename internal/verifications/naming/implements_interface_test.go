package naming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamingRuleImplementsInterface(t *testing.T) {
	t.Run("implements interface case 1", func(t *testing.T) {
		structDescription := StructDescription{}
		interfaceDescription := InterfaceDescription{}

		result := implementsInterface(structDescription, interfaceDescription)

		assert.True(t, result)
	})

	t.Run("implements interface case 2", func(t *testing.T) {
		structDescription := StructDescription{
			Name: "myStruct",
			Methods: []MethodDescription{
				{
					Name:         "method1",
					Parameters:   []string{},
					ReturnValues: []string{},
				},
			},
		}
		interfaceDescription := InterfaceDescription{
			Name: "foobar",
			Methods: []MethodDescription{
				{
					Name:         "method1",
					Parameters:   []string{},
					ReturnValues: []string{},
				},
				{
					Name:         "method2",
					Parameters:   []string{},
					ReturnValues: []string{},
				},
			},
		}

		result := implementsInterface(structDescription, interfaceDescription)

		assert.False(t, result)
	})

	t.Run("implements interface case 3", func(t *testing.T) {
		structDescription := StructDescription{
			Name: "myStruct",
			Methods: []MethodDescription{
				{
					Name:         "method1",
					Parameters:   []string{"int"},
					ReturnValues: []string{},
				},
				{
					Name:         "method5",
					Parameters:   []string{},
					ReturnValues: []string{"int"},
				},
				{
					Name:         "method2",
					Parameters:   []string{"string"},
					ReturnValues: []string{"bool"},
				},
			},
		}
		interfaceDescription := InterfaceDescription{
			Name: "foobar",
			Methods: []MethodDescription{
				{
					Name:         "method1",
					Parameters:   []string{"int"},
					ReturnValues: []string{},
				},
				{
					Name:         "method2",
					Parameters:   []string{"string"},
					ReturnValues: []string{"bool"},
				},
			},
		}

		result := implementsInterface(structDescription, interfaceDescription)

		assert.True(t, result)
	})

	t.Run("implements interface case 4", func(t *testing.T) {
		structDescription := StructDescription{
			Name: "myStruct",
			Methods: []MethodDescription{
				{
					Name:         "method1x",
					Parameters:   []string{"int"},
					ReturnValues: []string{},
				},
				{
					Name:         "method5",
					Parameters:   []string{},
					ReturnValues: []string{"int"},
				},
				{
					Name:         "method2",
					Parameters:   []string{"string"},
					ReturnValues: []string{"bool"},
				},
			},
		}
		interfaceDescription := InterfaceDescription{
			Name: "foobar",
			Methods: []MethodDescription{
				{
					Name:         "method1",
					Parameters:   []string{"int"},
					ReturnValues: []string{},
				},
				{
					Name:         "method2",
					Parameters:   []string{"string"},
					ReturnValues: []string{"bool"},
				},
			},
		}

		result := implementsInterface(structDescription, interfaceDescription)

		assert.False(t, result)
	})

	t.Run("implements interface case 5", func(t *testing.T) {
		structDescription := StructDescription{
			Name: "myStruct",
			Methods: []MethodDescription{
				{
					Name:         "method1",
					Parameters:   []string{"*int"},
					ReturnValues: []string{},
				},
				{
					Name:         "method5",
					Parameters:   []string{},
					ReturnValues: []string{"int"},
				},
				{
					Name:         "method2",
					Parameters:   []string{"string"},
					ReturnValues: []string{"bool"},
				},
			},
		}
		interfaceDescription := InterfaceDescription{
			Name: "foobar",
			Methods: []MethodDescription{
				{
					Name:         "method1",
					Parameters:   []string{"int"},
					ReturnValues: []string{},
				},
				{
					Name:         "method2",
					Parameters:   []string{"string"},
					ReturnValues: []string{"bool"},
				},
			},
		}

		result := implementsInterface(structDescription, interfaceDescription)

		assert.False(t, result)
	})
}

package naming

import "fmt"

func implementsInterface(s StructDescription, i InterfaceDescription) bool {
	fmt.Printf("Struct: %s\n", s.Name)
	methodsLeft := len(i.Methods)
	if len(s.Methods) < methodsLeft {
		fmt.Printf("\t[%s] -> No!\n", s.Name)
		return false
	}

	fmt.Printf("\t[%s] -> %v\n", s.Name, methodsLeft)
	for _, sm := range s.Methods {
		for _, im := range i.Methods {
			if sm.Name == im.Name {
				equalParams := areEquals(sm.Parameters, im.Parameters)
				equalReturnValues := areEquals(sm.ReturnValues, im.ReturnValues)

				fmt.Printf("\t\t[%s][%s][%s] -> (%v)(%v)\n", s.Name, sm.Name, im.Name, equalParams, equalReturnValues)
				fmt.Printf("\t\t\t[%v](%v)\n", sm.Parameters, im.Parameters)
				fmt.Printf("\t\t\t[%v](%v)\n", sm.ReturnValues, im.ReturnValues)
				if equalParams && equalReturnValues {
					methodsLeft--
					fmt.Printf("\t\t[%s] -> %v\n", s.Name, methodsLeft)
				}
			}
		}
	}

	fmt.Printf("\t[%s] -> %v\n", s.Name, methodsLeft)
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

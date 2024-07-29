package naming

import "slices"

func implementsInterface(structDesc StructDescription, interfaceDesc InterfaceDescription) bool {
	methodsLeft := len(interfaceDesc.Methods)
	if len(structDesc.Methods) < methodsLeft {
		return false
	}

	for _, sm := range structDesc.Methods {
		for _, im := range interfaceDesc.Methods {
			if sm.Name == im.Name {
				equalParams := slices.Equal(sm.Parameters, im.Parameters)
				equalReturnValues := slices.Equal(sm.ReturnValues, im.ReturnValues)

				if equalParams && equalReturnValues {
					methodsLeft--
				}
			}
		}
	}

	return methodsLeft == 0
}

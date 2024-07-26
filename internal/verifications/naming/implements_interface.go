package naming

func implementsInterface(structDesc StructDescription, interfaceDesc InterfaceDescription) bool {
	methodsLeft := len(interfaceDesc.Methods)
	if len(structDesc.Methods) < methodsLeft {
		return false
	}

	for _, sm := range structDesc.Methods {
		for _, im := range interfaceDesc.Methods {
			if sm.Name == im.Name {
				equalParams := areEquals(sm.Parameters, im.Parameters)
				equalReturnValues := areEquals(sm.ReturnValues, im.ReturnValues)

				if equalParams && equalReturnValues {
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

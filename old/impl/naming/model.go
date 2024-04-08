package naming

type InterfaceDescription struct {
	Name    string
	Methods []MethodDescription
}

type MethodDescription struct {
	Name 		 string
	Parameters   []string
	ReturnValues []string
}

type StructDescription struct {
	Name    string
	Methods []MethodDescription
}

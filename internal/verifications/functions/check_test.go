package functions

var functionTestDetails = []*FunctionDetails{ //nolint:gochecknoglobals
	{
		FilePath:   "/foo/bar/myfile.go",
		File:       "myfile.go",
		Name:       "myfunction1",
		NumLines:   10,
		IsPublic:   true,
		NumParams:  2,
		NumReturns: 2,
	},
	{
		FilePath:   "/foo/bar/myfile.go",
		File:       "myfile.go",
		Name:       "myfunction2",
		NumLines:   100,
		IsPublic:   true,
		NumParams:  2,
		NumReturns: 1,
	},
	{
		FilePath:   "/foo/bar/myfile2.go",
		File:       "myfile2.go",
		Name:       "myfunction21",
		NumLines:   15,
		IsPublic:   true,
		NumParams:  1,
		NumReturns: 2,
	},
	{
		FilePath:   "/foo/bar/myfile2.go",
		File:       "myfile2.go",
		Name:       "myfunction22",
		NumLines:   100,
		IsPublic:   true,
		NumParams:  2,
		NumReturns: 5,
	},
	{
		FilePath:   "/foo/bar/myfile2.go",
		File:       "myfile2.go",
		Name:       "myfunction23",
		NumLines:   200,
		IsPublic:   false,
		NumParams:  20,
		NumReturns: 21,
	},
}

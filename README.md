
[![Codecov](https://codecov.io/gh/arch-go/arch-go/branch/master/graph/badge.svg)](https://codecov.io/gh/arch-go/arch-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/arch-go/arch-go.svg)](https://pkg.go.dev/github.com/fdaines/arch-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/arch-go/arch-go)](https://goreportcard.com/report/github.com/fdaines/arch-go)
<img align="center" src="logo.png" alt="Arch-Go" title="Arch-Go" width="600px"/>

# Arch-Go
Architecture checks for Go projects

# Supported rules

## Dependencies Checks
Supports defining import rules
- Allowed dependencies
  - internal dependencies (same module)
  - standard library dependencies
  - external dependencies (3rd party packages)
- Not allowed dependencies
    - internal dependencies (same module)
    - standard library dependencies
    - external dependencies (3rd party packages)
  
## Package Content Checks
Allows you to define the contents of a set of packages, e.g. you can define that a desired package should only contain interfaces definitions.
The supported checks are:
* shouldNotContainInterfaces
* shouldNotContainStructs
* shouldNotContainFunctions
* shouldNotContainMethods
* shouldOnlyContainInterfaces
* shouldOnlyContainStructs
* shouldOnlyContainFunctions
* shouldOnlyContainMethods

## Function checks
Checks some functions properties, like the following:
- Maximum number of parameters
- Maximum number of return values
- Maximum number of public functions per file
- Maximum number of lines in the function body

## Naming rules checks
Checks some naming rules, like the following:
- If a struct implements an interface that match some name pattern, then it's name should starts or ends with a specific pattern. For example, all structs that implements 'Verificator' interface, should have a name that ends with 'Verificator'
  
# Configuration

## File arch-go.yml
```yaml
version: 1
threshold:
  compliance: 100
  coverage: 100
dependenciesRules:
  - package: "**.impl.*"
    shouldOnlyDependsOn:
      internal:
          - "**.foo.*"
          - "*.bar.*"
    shouldNotDependsOn: 
      internal: ["**.model.**"]
  - package: "**.utils.**"
    shouldOnlyDependsOn:
      - "**.model.**"
  - package: "**.foobar.**"
    shouldOnlyDependsOn:
      external:
        - "gopkg.in/yaml.v3"
  - package: "**.example.**"
    shouldNotDependsOn:
      external:
        - "github.com/foobar/example-module"
contentsRules:
  - package: "**.impl.model"
    shouldNotContainInterfaces: true
  - package: "**.impl.configuration"
    shouldOnlyContainFunctions: true
  - package: "**.impl.dependencies"
    shouldNotContainStructs: true
    shouldNotContainInterfaces: true
    shouldNotContainMethods: true
    shouldNotContainFunctions: true
functionsRules:
  - package: "**.impl.**"
    maxParameters: 3
    maxReturnValues: 2
    maxPublicFunctionPerFile: 1
    maxLines: 50
namingRules:
  - package: "**.arch-go.**"
    interfaceImplementationNamingRule:
      structsThatImplement: "*Connection"
      shouldHaveSimpleNameEndingWith: "Connection"
```

## Package name patterns
The package name can be defined as a fixed value or using _*_ special character, to create a simple pattern.

| Example         | Description                                                                                                                                                           |
| --------------- |:----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| *.name          | Package should end with _name_ and anything before, supporting only one level (for example _foo/name_, but no _foo/bar/name_)                                         |
| **.name         | Package should end with _name_ and anything before, supporting multiple levels (for example either _foo/name_ and _foo/bar/name_)                                     |
| name.*          | Package should start with _name_ and anything before, supporting only one level (for example _name/foo_, but no _name/foo/bar_)                                       |
| name.**         | Package should start with _name_ and anything before, supporting multiple levels (for example either _name/foo_ and _name/foo/bar_)                                   |
| \*\*.name.**    | Package should contain _name_, supporting multiple levels before and after (for example both _foo/name/x/y/z_, _foo/bar/name_ and _foo/bar/name/x_)                   |
| \*\*.foo/bar.** | Package should contain _foo/bar_, supporting multiple levels before and after (for example both _x/y/foo/bar/w/z_, _foo/bar/name_ and _x/y/foo/bar_)                  |
| foo.**.bar      | Package should start with _foo_, and ends with _bar_, and can have anything between them. (for example _foo/bar_, _foo/test/blah/bar_ and _foo/ok/bar_)               |
| foo.*.bar       | Package should start with _foo_, and ends with _bar_, and can have only one level between them. (for example _foo/bar_ and _foo/ok/bar_, but no _foo/test/blah/bar_)  |

## Threshold configuration
Current version supports threshold configuration for compliance and coverage rate. By default both limits are set to 100%.

### Compliance rate threshold
Represents how much the compliance level of the module is considering all the rules defined in the `arch-go.yml` file. For example, if there are 4 rules and the module meets 3 of them, then its compliance level will be 75%.

Arch-Go will check that the compliance level of your module must be equals or greater than the compliance threshold defined in your `arch-go.yml` file, if not then the verification will fail.

### Coverage rate threshold
Represents how many packages in this module were evaluated by at least one rule.
* At this moment Arch-Go doesn't analize internal components as structs, interfaces, functions or methods, just verifies if a certain package complies with a rule package pattern, so then this rule is evaluated against this package

Arch-Go will check that the coverage level of your module must be equals or greater than the threshold defined in your `arch-go.yml` file, if not then the verification will fail.

# Usage

## Using Arch-Go in command line

To install Arch-Go, run
```bash
$ go install -v github.com/fdaines/arch-go@latest
```

To execute this tool you have to be in the module path
```bash
$ cd [path-to-your-module]
```

Now you can execute Arch-Go tool
```bash
$ arch-go [flags]
```

### Describing your architecture guidelines
Arch-Go includes a command to describe the architecture rules from `arch-go.yml` file.

```bash
$ arch-go describe
```
The output of the `describe` command is similar to:
```
$ arch-go describe
Dependency Rules
        * Packages that match pattern '**.cmd.*',
                * Should only depends on packages that matches:
                        - '**.arch-go.**'
Function Rules
        * Packages that match pattern '**.arch-go.**' should comply with the following rules:
                * Functions should not have more than 50 lines
                * Functions should not have more than 4 parameters
                * Functions should not have more than 2 return values
                * Files should not have more than 5 public functions
Content Rules
        * Packages that match pattern '**.impl.model' should not contain functions or methods
        * Packages that match pattern '**.impl.config' should only contain functions
Naming Rules
        * Packages that match pattern '**.arch-go.**' should comply with:
                * Structs that implement interfaces matching name '*Verification' should have simple name ending with 'Verification'
Threshold Rules
        * The module must comply with at least 100% of the rules described above.
        * The rules described above must cover at least 100% of the packages in this module.

Time: 0.000 seconds
```

## Supported flags

| Flag      | Description                                                                                                                              |
|-----------|:-----------------------------------------------------------------------------------------------------------------------------------------|
| --color   | If not set (default: auto) in a tty the colors are printed. If set to yes or no the default is overridden. This can be useful in the CI. |
| --verbose | Includes detailed information while the command is running. The shorthand is _-v_                                                        |
| --html    | Generates a simple HTL report with the evaluation result.                                                                                |


## Examples
```bash
$ arch-go 
$ arch-go -v
$ arch-go --verbose
$ arch-go --html
$ arch-go --color
$ arch-go describe
```

## Using Arch-Go programmatically
The current version of Arch-Go allows us to include architecture checks as part of the tests run by the go test tool.

You need to include Arch-Go as a dependency in your project, using
```
go get github.com/fdaines/arch-go@latest
```

Then you need to create Architecture Tests, there is an example of a simple test case:
```go
package architecture_test

import (
	"testing"

	archgo "github.com/fdaines/arch-go/api"
	config "github.com/fdaines/arch-go/api/configuration"
)

func TestArchitecture(t *testing.T) {
	configuration := config.Config{
		DependenciesRules: []*config.DependenciesRule{
			{
				Package: "**.cmd.**",
				ShouldOnlyDependsOn: &config.Dependencies{
					Internal: []string{
						"**.cmd.**",
						"**.internal.**",
					},
				},
			},
		},
	}
	moduleInfo := config.Load("github.com/fdaines/my-go-project")

	result := archgo.CheckArchitecture(moduleInfo, configuration)

	if !result.Passes {
		t.Fatal("Project doesn't pass architecture tests")
	}
}
```
The `result` variable will store more than the verification result, 
including details for each rule type and analyzed packages, 
so then you can access all this data to create assertions as you need.



# Contributions
Feel free to contribute.

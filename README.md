# Arch-Go
Architecture checks for Go projects

**Supported rules:**
- Dependencies checks
    - Allowed dependencies
    - Not allowed dependencies
- Package content checks
    - Only interfaces
    - Only structs
    - Only functions
    - Only methods
    - Anything but interfaces
    - Anything but structs
    - Anything but functions
    - Anything but methods
- Cyclic dependencies

# Description of supported rules

## Dependencies Checks
Supports defining import rules

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

## Cyclic Dependencies checks
Checks if a set of packages contains cyclic dependencies.


# Configuration

## File arch-go.yml
```yaml
dependenciesRules:
  - package: "**.impl.*"
    shouldOnlyDependsOn:
      - "**.foo.*"
      - "*.bar.*"
    shouldNotDependsOn: ["**.model.**"]
  - package: "**.utils.**"
    shouldOnlyDependsOn:
      - "**.model.**"

contentsRules:
  - package: "**.impl.model"
    shouldNotContainInterfaces: true
  - package: "**.impl.config"
    shouldOnlyContainFunctions: true
  - package: "**.impl.dependencies"
    shouldNotContainStructs: true
    shouldNotContainInterfaces: true
    shouldNotContainMethods: true
    shouldNotContainFunctions: true

cyclesRules:
  - package: "**.cmd"
    shouldNotContainCycles: true

```

## Package name patterns
The package name can be defined as a fixed value or using _*_ special character, to create a simple pattern.

| Example    | Description                                                                                                                                         |
| ---------- |:---------------------------------------------------------------------------------------------------------------------------------------------------:|
| *.name     | Package should end with _name_ and anything before, supporting multiple levels (for example either _foo/name_ and _foo/bar/name_)                   |
| **.name    | Package should end with _name_ and anything before, supporting only one level (for example _foo/name_, but no _foo/bar/name_)                       |
| name.*     | Package should start with _name_ and anything before, supporting multiple levels (for example either _name/foo_ and _name/foo/bar_)                 |
| name.**    | Package should start with _name_ and anything before, supporting only one level (for example _name/foo_, but no _name/foo/bar_)                     |
| **.name.** | Package should contain _name_, supporting multiple levels before and after (for example both _foo/name/x/y/z_, _foo/bar/name_ and _foo/bar/name/x_) |


# Usage
To install Arch-Go, run
```bash
$ go get -u github.com/fdaines/arch-go
```

To execute this tool you have to be in the module path
```bash
$ cd [path-to-your-module]
```

Now you can execute Arch-Go tool
```bash
$ arch-go [flags]
```

## Supported flags

| Flag      | Description                                                                                     |
| --------- |:-----------------------------------------------------------------------------------------------:|
| --verbose | Includes detailed information while the command is running                                      |


## Examples
```bash
$ arch-go 
$ arch-go -v
```

# Contributions
Feel free to contribute.

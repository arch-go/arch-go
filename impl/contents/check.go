package contents

import (
	"github.com/fdaines/arch-go/config"
	"github.com/fdaines/arch-go/model"
	"github.com/fdaines/arch-go/model/result"
	"github.com/fdaines/arch-go/utils/text"
	"regexp"
)

type ContentsRule struct {
	results []*result.ContentsRuleResult
	rule    *config.ContentsRule
	module  *model.ModuleInfo
}

func NewContentsRule(results []*result.ContentsRuleResult, rule *config.ContentsRule, module *model.ModuleInfo) *ContentsRule {
	return &ContentsRule{
		rule:    rule,
		results: results,
		module:  module,
	}
}

func (c *ContentsRule) CheckRule() []*result.ContentsRuleResult {
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(c.rule.Package))
	for _, p := range c.module.Packages {
		if packageRegExp.MatchString(p.Path) {
			contents, _ := retrieveContents(p, c.module.MainPackage)
			c.checkInterfaces(contents)
			c.checkTypes(contents)
			c.checkMethods(contents)
			c.checkFunctions(contents)
		}
	}
	return c.results
}

func (c *ContentsRule) checkInterfaces(contents *PackageContents) {
	if contents.Interfaces > 0 {
		if c.rule.ShouldNotContainInterfaces {
			c.results = appendError(c.results, c.rule.Package, "should not contain interfaces")
		}
		if c.rule.ShouldOnlyContainTypes {
			c.results = appendError(c.results, c.rule.Package, "should only contain types")
		}
		if c.rule.ShouldOnlyContainMethods {
			c.results = appendError(c.results, c.rule.Package, "should only contain methods")
		}
		if c.rule.ShouldOnlyContainFunctions {
			c.results = appendError(c.results, c.rule.Package, "should only contain functions")
		}
		if c.rule.ShouldOnlyContainInterfaces {
			c.results = appendSuccess(c.results, c.rule.Package, "should only contain interfaces")
		}
	}
	if c.rule.ShouldNotContainInterfaces {
		c.results = appendSuccess(c.results, c.rule.Package, "should not contain interfaces")
	}
}

func (c *ContentsRule) checkTypes(contents *PackageContents) {
	if contents.Types > 0 {
		if c.rule.ShouldNotContainTypes {
			c.results = appendError(c.results, c.rule.Package, "should not contain types")
		}
		if c.rule.ShouldOnlyContainInterfaces {
			c.results = appendError(c.results, c.rule.Package, "should only contain interfaces")
		}
		if c.rule.ShouldOnlyContainMethods {
			c.results = appendError(c.results, c.rule.Package, "should only contain methods")
		}
		if c.rule.ShouldOnlyContainFunctions {
			c.results = appendError(c.results, c.rule.Package, "should only contain functions")
		}
		if c.rule.ShouldOnlyContainTypes {
			c.results = appendSuccess(c.results, c.rule.Package, "should only contain types")
		}
	}
	if c.rule.ShouldNotContainTypes {
		c.results = appendSuccess(c.results, c.rule.Package, "should not contain types")
	}
}

func (c *ContentsRule) checkMethods(contents *PackageContents) {
	if contents.Methods > 0 {
		if c.rule.ShouldNotContainMethods {
			c.results = appendError(c.results, c.rule.Package, "should not contain methods")
		}
		if c.rule.ShouldOnlyContainTypes {
			c.results = appendError(c.results, c.rule.Package, "should only contain types")
		}
		if c.rule.ShouldOnlyContainInterfaces {
			c.results = appendError(c.results, c.rule.Package, "should only contain interfaces")
		}
		if c.rule.ShouldOnlyContainFunctions {
			c.results = appendError(c.results, c.rule.Package, "should only contain functions")
		}
		if c.rule.ShouldOnlyContainMethods {
			c.results = appendSuccess(c.results, c.rule.Package, "should only contain methods")
		}
	}
	if c.rule.ShouldNotContainMethods {
		c.results = appendSuccess(c.results, c.rule.Package, "should not contain methods")
	}
}

func (c *ContentsRule) checkFunctions(contents *PackageContents) {
	if contents.Functions > 0 {
		if c.rule.ShouldNotContainFunctions {
			c.results = appendError(c.results, c.rule.Package, "should not contain functions")
		}
		if c.rule.ShouldOnlyContainTypes {
			c.results = appendError(c.results, c.rule.Package, "should only contain types")
		}
		if c.rule.ShouldOnlyContainMethods {
			c.results = appendError(c.results, c.rule.Package, "should only contain methods")
		}
		if c.rule.ShouldOnlyContainInterfaces {
			c.results = appendError(c.results, c.rule.Package, "should only contain interfaces")
		}
		if c.rule.ShouldOnlyContainFunctions {
			c.results = appendSuccess(c.results, c.rule.Package, "should only contain functions")
		}
	}
	if c.rule.ShouldNotContainFunctions {
		c.results = appendSuccess(c.results, c.rule.Package, "should not contain functions")
	}
}

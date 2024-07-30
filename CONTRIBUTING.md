# Welcome to arch-go's contributing guide

Thank you for investing your time in contributing to this project!
Any contribution you make, will make arch-go better for everyone :sparkles:!

Before you continue, here are some important resources:

* To keep the community around arch-go approachable and respectable, please read our [Code of Conduct](./CODE_OF_CONDUCT.md).
* [arch-go's Discussions](https://github.com/arch-go/arch-go/discussions) can help you in getting answers to your questions


## How can I contribute?

There are many ways you can contribute to arch-go. Here are some ideas:

* **Give this project a star**: It may not seem like much, but it really makes a difference. This is something that everyone can do to help. GitHub stars help the project gaining visibility and stand out.
* **Join the community**: Helping people can be as easy as by just sharing your own experience. You can also help by listening to issues and ideas of other people and offering a different perspective, or providing some related information that might help. Take a look at [arch-go's Discussions](https://github.com/arch-go/arch-go/discussions). Bonus: You get GitHub achievements for answered discussions :wink:.
* **Help with open issues**: There may be many [open issues](https://github.com/arch-go/arch-go/issues). Some of them may lack necessary information, some may be duplicates of older issues. Most are waiting for being implemented.
* **You spot a problem**: Search if an [issue already exists](https://github.com/arch-go/arch-go/issues). If a related issue doesn't exist, please open a new issue using a relevant [issue form](https://github.com/arch-go/arch-go/issues/new/choose). You have no obligation to offer a solution or code to fix an issue you open. 

## Disclosing vulnerabilities

Please disclose vulnerabilities by making use of [Security Advisories](https://github.com/arch-go/arch-go/security/advisories). Do not use GitHub issues for that!

## Contribute content

Unless you are fixing a known bug, we strongly recommend discussing it with the core team via a [GitHub Issue](https://github.com/arch-go/arch-go/issues) or in [arch-go's Discussions](https://github.com/arch-go/arch-go/discussions) before getting started.

**Important:** Only PRs with [signed commits](https://docs.github.com/en/authentication/managing-commit-signature-verification/signing-commits) will be accepted.

The general process is as follows:

Set up your local development environment to contribute to arch-go:

1. [Fork](https://github.com/arch-go/arch-go/fork), then clone the repository.
  
   ```bash
   > git clone https://github.com/your_github_username/arch-go.git
   > cd arch-go
   > git remote add upstream https://github.com/arch-go/arch-go.git
   > git fetch -p upstream
   ```

2. Install required tools:
  * [Golang](https://go.dev/dl/) - latest version.
  * [golangci-lint](https://golangci-lint.run/usage/install/#local-installation) to lint the code.
  * [go-licenses](https://github.com/google/go-licenses) to ensure all dependencies have allowed licenses.

3. Verify that tests and other checks pass locally.
   ```bash
   > git pull
   > git checkout main
   > golangci-lint run
   > go-licenses check --disallowed_types=forbidden,restricted,reciprocal,permissive,unknown  --ignore=github.com/hashicorp/hcl .
   > go test -v -gcflags=all=-l ./...
   ```
   
4. When creating your PR, please follow the guide you'll see in the PR template to streamline the review process.

5. At this point, you're waiting on us to review your changes. We *try* to respond to issues and pull requests within a few days, and we may suggest some improvements or alternatives. Once your changes are approved, one of the project maintainers will merge them.

### Contribute Code

After having installed the tools listed above:

1. Create a new feature branch.
   ```bash
   > git checkout -b cool_new_feature
   ```

2. Make your changes, and verify that all tests and lints still pass.
   ```bash
   > golangci-lint run
   > go-licenses check --disallowed_types=forbidden,restricted,reciprocal,permissive,unknown  --ignore=github.com/hashicorp/hcl .
   > go test -v -gcflags=all=-l ./...
   ```

3. When you're satisfied with the change, push it to your fork and make a pull request.
   ```bash
   > git push origin cool_new_feature
   # Open a PR at https://github.com/arch-go/arch-go/compare
   ```
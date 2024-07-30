<!--
Please note: This project expects semantic PRs. That means each PR name must comply to the pattern "tag: description",
with "tag" being one from the following list and "description" being the name/description of your PR:

* build - The PR is related to the build system of the project
* chore - The PR is about housekeeping, like fixing typos, and alike
* ci - The PR updates the CI
* docs - The PR updates the documentation
* deps - The PR updates dependencies
* feat - The PR implements a new feature
* fix - The PR fixes for a bug
* perf - The PR implements new performance tests
* refactor - The PR is about code refactoring
* revert - The PR reverts one of the previous PRs
* style - The PR updates the UI style
* test - The PR implements new tests
* wip - The PR is not yet ready for review; it is Work in Progress.

E.g.: "feat: Some cool feature".

If this PR introduces a breaking change, please put a "!" after the "tag" and before the colon.

E.g.: "feat!: Some cool feature that breaks current behaviour"

When you create PRs which are not yet complete in sense of the implementation/changes, please use the "wip" tag
to inform the maintainers that you're not ready with your work yet, and you're not expecting a review.

As of today only PRs tagged with "feat", "fix", "deps" and "docs" as well as all breaking change PRs (those with
a "!" after the tag) are included into the change list of a release. This may change in the future. All contributors
are referenced in the release, despite whether the actual PR is listed or not.

If your PR addresses multiple changes, you can represent these by adding additional semantic commit messages at the
end of the Description section (See below)
-->

## Related issue(s)

<!--
If this pull request

1. is a fix for a known bug, link the issue where the bug was reported in the format of `fixes #1234`;
2. is a fix for a previously unknown bug, please create a new issue first, which describes the bug;
3. implements a new feature, link the issue describing the idea in the format of `closes #1234`;
4. improves documentation, updates dependencies, implements new tests, etc, no issue reference is required. Please delete this section in such case.

You can discuss changes with maintainers in the [GitHub Discussions](https://github.com/arch-go/arch-go/discussions) in this repository.
-->

## Checklist

<!--
Remove the boxes, which are not applicable and put an `x` in the boxes that apply.
You can also fill these out after creating the PR.
-->

- [ ] I agree to follow this project's [Code of Conduct](../CODE_OF_CONDUCT.md).
- [ ] I have read, and I am following this repository's [Contributing Guidelines](../CONTRIBUTING.md).
- [ ] I have read the [Security Policy](../SECURITY.md).
- [ ] I have referenced an issue describing the bug/feature request.
- [ ] I have added tests that prove the correctness of my implementation.
- [ ] I have updated the documentation.

## Description
<!--
Describe your changes here to communicate 

1. why your PR should be accepted, why you chose the solution you did and what alternatives you considered, etc...
2. which changes/updates it introduces. If your change includes breaking changes please add a code block documenting the breaking change
-->

## Changelist
<!--
If your PR addresses multiple changes, you can represent these by adding additional semantic commit messages at the
end of this section. Otherwise, if the name of this PR is enough, delete this section.

E.g.
Other changes done by this PR:

feat: Other cool feature

fix: Fixes this and that issue

refactor: This and that refactored
-->

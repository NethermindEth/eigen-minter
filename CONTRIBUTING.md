# Contributing to eigen-minter

The Nethermind team maintains guidelines for contributing to the Nethermind repos. Check out our [docs page](https://docs.nethermind.io/nethermind/) for more info about us.

### Code of Conduct

Have you read the [code of conduct](<repo>/blob/main/CODE_OF_CONDUCT.md)?

## Bugs and Feature Request

Before you make your changes, check to see if an [issue](<repo>/issues) exists already for the change you want to make.

### Don't see your issue? Open one

If you spot something new, open an issue using a [template](<repo>/issues/new/choose). We'll use the issue to have a conversation about the problem you want to fix.

### Open a Pull Request

When you're done making changes and you'd like to propose them for review, use the pull request template to open your PR (pull request).

If your PR is not ready for review and merge because you are still working on it, please convert it to draft and add to it the label `wip` (work in progress). This label allows to filter correctly the rest of PR not `wip`.

### Do you intend to add a new feature or change an existing one?

Suggest your change by opening an issue and starting a discussion.

### Improving Issues and PR

Please add, if possible, a reviewer, assignees and labels to your issue and PR.

## DOs and DON'Ts

Please do:

* **DO** give priority to the current style of the project or file you're changing even if it diverges from the general guidelines.
* **DO** include tests when adding new features. When fixing bugs, start with adding a test that highlights how the current behavior is broken.
* **DO** especially follow our rules in the [Contributing](https://github.com/NethermindEth/nethermind/blob/master/CODE_OF_CONDUCT.md#contributing) section of our code of conduct.
* **DO** write idiomatic golang code

Please do not:

* **DON'T** fill the issues and PR descriptions vaguely. The elements in the templates are there for a good reason. Help the team.
* **DON'T** surprise us with big pull requests. Instead, file an issue and start a discussion so we can agree on a direction before you invest a large amount of time.

## Branch Naming

Branch names must follow `kebab-case` pattern. Follow the pattern `feature/<name>` or `fix/<name>` `(folder/<name>)` when it is possible and add issue reference if applicable.

## Commit Naming

Commits must follow the `<type>(<scope>): <subject>` pattern, as stated in the [Conventional Commits specification](https://www.conventionalcommits.org/en/v1.0.0/). The `scope` should be the chart name.

The following is the list of types that can be used:

- `feat:` for new features
- `fix:` for bugfixes
- `improvement:` for enhancements
- `docs:` for documentation and examples
- `refactor:` for code refactoring
- `test:` for tests
- `ci:` for CI purpose
- `chore:` for chores stuff

## Pre-commit hooks

Before making any commits, we strongly recommend running our **pre-commit hooks** over your code. This will run linters and static checks to ensure your code is formatted correctly and you have no syntax errors. To install pre-commit hooks, run the following command:

```
make install-pre-commit
```

To run pre-commit hooks after you have installed them, and without making any commits, run the following command:

```
make pre-commit
```

The same pre-commit hooks will be run over your code after you push your commits; it is not mandatory to run them locally. However, relying on the CI for the pre-commit check could be an ineffective process, requiring you to push more commits than intended for any likely fix.

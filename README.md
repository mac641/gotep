# gotep
gotep is a terminal-based REST client designed to execute HTTP tests based on the Jetbrains HTTP-Client.

## Contribution guidelines
Please read this and follow these guidelines, if you wish to contribute.

### Git hooks
Before you can start contributing, enable git commit hooks to ensure unified commit messages and properly formatted source code.
Therefore, clone the repository and execute `git config core.hooksPath .hooks; chmod +x .hooks/*` in the project's root directory.

* Commit messages have to follow the [Angular commit message guidelines](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#-commit-message-format).
* All go files need to be properly formatted before they can be committed. Use `make fmt`.

### Code comments
Communicating throughout the source code can be done using one of the following comment types.

```golang
// FIXME - for a bug that has to be approached later on
// TODO - for a feature / architecture / design change which depends on something else and needs to be implemented on a future point in time
// NOTE - for an important info
```

### Git workflow
This repo follows the **rebase-merge** workflow. Therefore, feature branches need to be rebased onto the *main* branch before they can be merged.

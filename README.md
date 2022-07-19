# gotep
gotep is a terminal-based REST client designed to execute HTTP tests based on the Jetbrains HTTP-Client.

## Contributing
Please read and follow the guidelines, if you wish to contribute.

### Dev environment setup
1. Clone the repository.
2. Enable git hooks by executing `git config core.hooksPath .hooks; chmod +x .hooks/*`.
3. Spin up a docker dev container and run `go install` inside the `src` directory to install all necessary go
dependencies.
  * ```bash
    docker run --rm -it --name gotep-dev -v $PWD:/go/gotep golang:1.18
    ```

### Miscellaneous
* This repo follows the **rebase-merge** git strategy.
* Commit messages have to follow an extended version of
[Angular commit message guidelines](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#-commit-message-format).
* All go files need to be properly formatted before they can be committed. Use `make fmt`.
* Comments should begin with the following annotations:
  * ```golang
    // FIXME - for a bug that has to be approached later on
    // TODO - for a feature / architecture / design change which depends on something else and needs to be implemented on a future point in time
    // NOTE - for an important info
    ```


![Lint, test and build gotep](https://github.com/mac641/gotep/actions/workflows/main.yml/badge.svg?branch=main)

# gotep
gotep is a terminal-based REST client designed to execute HTTP tests based on the Jetbrains HTTP-Client.

## Usage 💡
1. Download the OS-specific binary from the release page.
2. Copy the binary to a location which is in your `$PATH` variable.
3. Execute your tests by typing `gotep run <testfile with .http or .rest ending>` in your terminal.
  Find out more about supported command flags by executing the help commands described below.
4. If you need help, type `gotep -h` or `gotep help`. For subcommand-specific help, type `gotep <subcommand> -h`.

## Bugs 🐞
If you encounter bugs during use, feel free to open a PR containing a fix. 😉

## Contributing 🧑‍💻
Please read and follow the guidelines, if you wish to contribute.

### Dev environment setup
1. Ensure you have [go-task](https://taskfile.dev/installation/) installed.
2. Clone the repository.
3. Enable git hooks by executing `git config core.hooksPath .hooks; chmod +x .hooks/*`.
4. Run `go install` inside the `src` directory to install all necessary go dependencies. You can either spin up a
docker dev container using the command below or develop on your local machine.

```bash
  docker run --rm -it --name gotep-dev -v $PWD:/go/gotep golang:1.18
```

### Miscellaneous
* This repo follows the **rebase-merge** git strategy.
* Commit messages have to follow an extended version of
[Angular commit message guidelines](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#-commit-message-format).
* All go files need to be properly formatted before they can be committed. Use `task fmt`.
* Comments should begin with the following annotations:
  * ```golang
    // FIXME - for a bug that has to be approached later on
    // TODO - for a feature / architecture / design change which depends on something else and needs to be implemented on a future point in time
    // NOTE - for an important info
    ```

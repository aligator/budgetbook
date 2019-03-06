# <img src="https://sternentstehung.de/budgetbook-md.png" alt="BudgetBook"> BudgetBook

> A simple CLI tool for managing your personal fincances.

## Installation and Usage

### Build the binary

Building the executable from `%GOPATH%\src\budgetBook` may look like this:

```sh
go build -o budgetbook.exe main.go
```

It's recommended to choose `budgetbook` as output binary name since this is the name of the root command.

### Execute commands

Example for adding a new financial transaction category:

```sh
budgetbook add-cat --name Coffee --budget 120 --is-capped
```

To assign flag values with whitespaces, use the attribute-style syntax:

```sh
budgetbook add-cat --name="Coffee To Go" --budget 120 --is-capped
```

### Retrieve the available commands

All sub-commands are lodged in `app\cmds.go`. However, the easier way is a simple help flag.

```sh
budgetbook -h
```

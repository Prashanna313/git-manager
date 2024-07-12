# Git Manager

A simple Cobra CLI application to manage multiple local Git repositories. This tool allows you to:

- Fetch the latest changes from the remote repository.
- Delete merged branches.
- Switch Git user configurations.

## Features

1. **Fetch Latest Changes**: Fetches the latest changes from the remote for all repositories within a specified directory.
2. **Delete Merged Branches**: Deletes local branches that have been merged into the main or master branch.
3. **Switch Git User**: Switches the Git user configuration for all repositories within a specified directory.

## Installation

Ensure you have Go installed. If not, download and install it from [golang.org](https://golang.org/dl/).

Clone the repository:

```bash
git clone https://github.com/Prashanna313/git-manager.git
cd git-manager
```

## Build
To build the CLI application, run:


```bash
go build -o git-cli
```
This will generate an executable named git-cli in the current directory.

## Usage
- Fetch Latest Changes for All Repositories
```bash
./git-cli fetch -p /path/to/your/repos
```
- Delete Merged Branches for All Repositories
```bash
./git-cli delete-merged -p /path/to/your/repos
```
- Switch Git User for All Repositories
```bash
./git-cli switch-user "Your Name" "your.email@example.com" -p /path/to/your/repos
```
## Flags
-p, --path : Path to the directory containing repositories (default is current directory).

## Testing
To run the tests, use the following command:

```bash
go test -v
```
This will execute all the tests and provide detailed output for each test case.

## Contributing
Feel free to submit issues, fork the repository and send pull requests!

## License
This project is licensed under the MIT License. See the LICENSE file for details.
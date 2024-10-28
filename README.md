<img alt="Static Badge" src="https://img.shields.io/badge/Go-1.23.3 +-blue">


# IP Finder CLI

**IP Finder CLI** is a command-line interface (CLI) application developed in Golang for searching IP addresses and retrieving detailed server information. This application uses the `github.com/urfave/cli` package to streamline command creation and organize the CLI interface.

## Prerequisites

- **Golang**: Ensure that Go is installed, version 1.23 or higher. [Install Go](https://go.dev/doc/install)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/ip-finder-cli.git
    cd ip-finder-cli
    ```

2. Run the project:

    ```bash
    go run main.go ip --host "name of website"
    ```
    or
   ```bash
   go run main.go server --host "name of website"

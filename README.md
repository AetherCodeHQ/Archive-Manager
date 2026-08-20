# Archive Manager

![CI](https://github.com/Qyroxen/Archive-Manager/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Archive-Manager/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Archive-Manager?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Archive-Manager)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Archive-Manager)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Archive-Manager?style=social)](https://github.com/Qyroxen/Archive-Manager/stargazers)

## What is it?

Archive Manager is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Archive-Manager.git
cd Archive-Manager
go build -o archivemanager .

# Run
./archivemanager --help
```

## CLI Usage

```bash
# Basic usage
./archivemanager

# With flags
./archivemanager --verbose --output json

# Get help
./archivemanager --help
```

## Examples

```bash
# Example 1
./archivemanager example1

# Example 2
./archivemanager example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o archivemanager .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Archive-Manager/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Archive-Manager?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Archive-Manager/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Archive-Manager?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Archive-Manager/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Archive-Manager" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Archive-Manager/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Archive-Manager" alt="Pull Requests">
  </a>
</p>

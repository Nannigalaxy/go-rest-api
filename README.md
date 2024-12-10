# Go REST API Boilerplate

A simple yet scalable, and maintainable RESTful API built using **Go (Golang)** to serve as boilerplate for new projects. The goal is to design for low-latency and high-volume applications, while also offering developers greater control through a balanced level of abstraction.

Note: This project is still in early stage. Feel free to raise PRs and contribute. 

External Go packages used:  
- Framework: [gin](https://github.com/gin-gonic/gin)
- Database (postgresql): [jackc/pgx](https://github.com/jackc/pgx)
- Logging (rollbar): [rollbar](https://github.com/rollbar/rollbar-go)

## Table of Contents

- [Installation](#installation)
- [Development](#development)
- [Contributing](#contributing)


## Getting started

### Prerequisites

- Go version 1.23.2

### Clone the repository

```bash
git clone https://github.com/nannigalaxy/go-rest-api
cd go-rest-api
```
### Install dependencies
Run `go mod download` to install the Go dependencies.

## Development
Running services locally without Docker
```bash
go run main.go
```

Additionally to live-reload you can use [`air`](https://github.com/air-verse/air)

## Contributing
We welcome contributions to this project! If you find any bugs, have ideas for new features, or would like to improve the codebase, feel free to fork the repository and submit a pull request.

## TODO
- Add test framework
- Add swagger docs
- Fix nginx configuration
- Add database migration tool
- Add observability framework (e.g, OpenTelemetry)
# Tiny Server

A simple HTTP server written in Golang.

## Installation

```
go install github.com/qwinsi/tinyserver@latest
```

Then you can run the server with:

```shell
tinyserver <directory>
```

## Usage

Run `./tinyserver --help` to see the usage.

```shell
$ ./tinyserver --help
 Usage:   tinyserver [OPTIONS] <directory>
 Options: 
    -p, --port The port number which you want to bind on.
                         (If omitted, use 80 by default)
```

## Development

1. Clone the repository.

2. Build the binary.

```shell
go build -o tinyserver ./main.go
```

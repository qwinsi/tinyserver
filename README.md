# Tiny Server

A simple HTTP server written in Golang.

# Building

```shell
go build -o tinyserver  ./src/main.go
```

# Running

Run `./tinyserver --help` to see the usage.

```shell
$ ./tinyserver --help
 Usage:   tinyserver [OPTIONS] <directory>
 Options: 
    -p, --port The port number which you want to bind on.
                         (If omitted, use 80 by default)
```

/**
  Tiny Server
	@Version: 0.3.0
	@Last Modified: 2021-03-24
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
)

type Config struct {
	Port         int
	DocumentRoot string
}

var config = Config{
	Port:         80,
	DocumentRoot: "/var/www/html",
}

const USAGE = `
 Usage:   %s [OPTIONS] <directory>
 Options: 
    -p, --port The port number which you want to bind on.
                         (If omitted, use 80 by default)
`

func showUsageOf(programName string) {
	fmt.Printf(USAGE, programName)
}
func readConfig(config *Config) {
	for idx := 1; idx < len(os.Args); idx++ {

		switch os.Args[idx] {
		case "-h":
			fallthrough
		case "--help":
			showUsageOf("tinyserver")
			os.Exit(0)
		case "-p":
			fallthrough
		case "--port":
			if len(os.Args)-1 == idx {
				_, _ = fmt.Fprintln(os.Stderr, "Bad command: got -p/--port but no port number specified.")
				os.Exit(1)
			}
			idx++
			port, err := strconv.Atoi(os.Args[idx])
			if err != nil || port <= 0 || port >= 65535 {
				_, _ = fmt.Fprintf(os.Stderr, "Bad command: Not a proper port number followed with -p/--port.")
				os.Exit(1)
			}
			config.Port = port
		default:
			// assume this parameter is <directory>
			root := os.Args[idx]
			if fileInfo, err := os.Stat(root); err != nil || !fileInfo.IsDir() {
				_, _ = fmt.Fprintf(os.Stderr, "Error: the directory %s is not accessible.\n", config.DocumentRoot)
				os.Exit(1)
			}
			if length := len(root); '/' == root[length-1] || '\\' == root[length-1] {
				config.DocumentRoot = root[0:(length - 1)]
			} else {
				config.DocumentRoot = root
			}
		}

	}

}

func handler(respWriter http.ResponseWriter, req *http.Request) {
	var (
		filePath    string
		statusCode  int
		contentType string
		body        []byte
	)
	// Only support GET and POST at present.
	if req.Method != "GET" && req.Method != "POST" {
		respWriter.WriteHeader(http.StatusForbidden)
		return
	}

	filePath = config.DocumentRoot + req.URL.Path

	// try as a file
	body, err := ioutil.ReadFile(filePath)
	// try as a directory
	if err != nil {
		if '/' != filePath[len(filePath)-1] {
			filePath += "/"
		}
		filePath += "index.html"
		body, err = ioutil.ReadFile(filePath)
	}


	if err == nil {
		statusCode = http.StatusOK
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/MIME_types/Common_types
		switch path.Ext(filePath) {
		case ".html":
			contentType = "text/html; charset=utf-8"
		case ".css":
			contentType = "text/css; charset=utf-8"
		case ".js", ".mjs":
			contentType = "application/javascript; charset=utf-8"
		case ".svg":
			contentType = "image/svg+xml; charset=utf-8"
		case ".json":
			contentType = "application/json; charset=utf-8"
		default:
			// No assignment here,
			// just use the default Content-Type set by the `http` library in golang.
		}
	} else {
		statusCode = http.StatusNotFound
		body = []byte("<h1>404 Not Found</h1>")
	}

	if contentType != "" {
		header := respWriter.Header()
		header.Set("Content-Type", contentType)
	}

	respWriter.WriteHeader(statusCode)
	_, _ = respWriter.Write(body)

}

func main() {
	readConfig(&config)
	http.HandleFunc("/", handler)

	fmt.Printf("Server started at Port %d and Directory %s ... \n", config.Port, config.DocumentRoot)
	addr := fmt.Sprintf(":%d", config.Port)
	err := http.ListenAndServe(addr, nil)
	_, _ = fmt.Fprintln(os.Stderr, err)
}

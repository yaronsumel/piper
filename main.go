package main

import (
	"flag"
	"github.com/yaronsumel/pipe"
)

// usage
// server -> $ tail -f log.txt | piper 1.2.3.4:8080
// client -> $ piper 1.2.3.4:8080 > log.txt

var address = flag.String("a", "", "address to bind TLS server or connect to running server")
var verbose = flag.Bool("v", false, "verbose mode")

func main() {

	flag.Parse()

	// server mode
	if pipe.IsNamedPipe() {
		newServer(*address, *verbose).listen()
		return
	}
	// client mode
	newClient(*address, *verbose).connect()
}

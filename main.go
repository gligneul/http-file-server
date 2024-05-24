// Run a HTTP that serves files in the current directory

package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Get port
	port := flag.String("port", "7777", "port to serve files on")
	flag.Parse()

	// Define the file server root
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// Start the server and log any errors
	log.Printf("Serving files on HTTP port %s", *port)
	addr := "localhost:" + *port
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

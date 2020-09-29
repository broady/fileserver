package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr,
			"usage: fileserver [host]:port /path/to/dir\n\n"+
				"For example:\n"+
				"\tfileserver :9999 .\n\n")
		os.Exit(2)
	}
	http.Handle("/", http.FileServer(http.Dir(os.Args[2])))
	log.Fatal(http.ListenAndServe(os.Args[1], nil))
}

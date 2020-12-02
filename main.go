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
	http.Handle("/", http.FileServer(custom404fs{http.Dir(os.Args[2])}))
	log.Fatal(http.ListenAndServe(os.Args[1], nil))
}

type fileSystem struct {
	base http.FileSystem
}

func (fs fileSystem) Open(name string) (http.File, error) {
	f, err := fs.base.Open(name)
	if os.IsNotExist(err) {
		f, err = fs.base.Open("404.html")
	}
	return f, err
}

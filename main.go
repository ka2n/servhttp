package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "3000", "port(default: 3000)")
	flag.Parse()

	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}

	path, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("Listen: http://localhost:%v, base: %v", port, path)

	fs := http.FileServer(http.Dir(path))
	http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// do some configuration stuff
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	// make a new mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/item", showItem)
	mux.HandleFunc("/item/add", addItem)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}

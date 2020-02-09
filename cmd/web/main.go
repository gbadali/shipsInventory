package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// do some configuration stuff
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	// leveled loging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// make a new mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/item", showItem)
	mux.HandleFunc("/item/add", addItem)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

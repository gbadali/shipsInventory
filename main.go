package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Ships Inventory"))
}

func showItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific item in the inventory..."))
}

func addItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Add an item to the inventory(test)..."))
}

//test

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/item", showItem)
	mux.HandleFunc("/item/add", addItem)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

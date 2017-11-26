package main

import (
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "approved.jpg")
}
func ProofHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "proof.jpg")
}

func main() {
	http.HandleFunc("/proof", ProofHandler)
	http.HandleFunc("/", RootHandler)
	http.ListenAndServe(":8181", nil)
}

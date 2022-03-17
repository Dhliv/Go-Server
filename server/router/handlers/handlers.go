package handlers

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is te API endpoint")
}

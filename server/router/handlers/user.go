package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dhliv/Go-Server/server/types"
)

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user types.User
	err := decoder.Decode(&user)

	if err != nil {
		fmt.Fprintf(w, "error: %v\n", err)
		return
	}

	response, err := user.ToJson()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

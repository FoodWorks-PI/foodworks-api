package api

import (
	"log"
	"net/http"

	"foodworks.ml/m/internal/auth"
)

func EchoRequest(w http.ResponseWriter, r *http.Request) {
	user := auth.ForContext(r.Context())
	log.Println(user)
}

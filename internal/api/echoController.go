package api

import (
	"log"
	"net/http"

	"github.com/go-chi/jwtauth"
)

func EchoRequest(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		log.Println(cookie.Name)
		log.Println(cookie.Value)
	}
	_, claims, err := jwtauth.FromContext(r.Context())
	// Allow unauthenticated users in
	if claims == nil || err != nil || len(claims) == 0 {
		log.Println(err)
		return
	}
	identity := claims["session"].(map[string]interface{})["identity"]
	log.Println(identity)
}

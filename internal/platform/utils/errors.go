package utils

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

//TODO: Add Graphql error
func SimpleFail(w http.ResponseWriter, err error, code int) {
	log.Err(err).Msg("")
	w.WriteHeader(code)
}

package auth

import (
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/lestrrat/go-jwx/jwk"
	"github.com/rs/zerolog/log"
)

var TokenAuth *jwtauth.JWTAuth

func InitAuth() {
	set, err := jwk.Fetch("http://127.0.0.1:4456/.well-known/jwks.json")
	if err != nil {
		log.Panic().Err(err).Msg("Error gettint jwks")
	}
	public, private := set.Keys[0].Materialize()
	TokenAuth = jwtauth.New("RS256", public, private)
}

// Middleware decodes the jwt and packs the session into context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, claims, err := jwtauth.FromContext(r.Context())
			if claims == nil || err != nil {
				log.Err(err).Msg("Error getting jwt")
				next.ServeHTTP(w, r)
				return
			}
			identity := claims["session"].(map[string]interface{})["identity"]
			log.Debug().Interface("session", identity).Msg("")
			// parsed, err := gabs.ParseJSON(claims)
			// log.Println(parsed)
			// log.Println(err)

			next.ServeHTTP(w, r)
		})
	}
}

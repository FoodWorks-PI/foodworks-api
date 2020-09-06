package auth

import (
	"log"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/lestrrat/go-jwx/jwk"
)

var TokenAuth *jwtauth.JWTAuth

func InitAuth() {
	set, err := jwk.Fetch("http://127.0.0.1:4456/.well-known/jwks.json")
	if err != nil {
		log.Panic(err)
	}
	public, private := set.Keys[0].Materialize()
	TokenAuth = jwtauth.New("RS256", public, private)
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// for _, cookie := range r.Cookies() {
			// 	fmt.Fprint(w, cookie.Name)
			// 	log.Println(cookie.Name)
			// }
			// log.Println(formatRequest(r))
			// log.Println(r.Cookies())
			// c, err := r.Cookie("auth-cookie")

			// Allow unauthenticated users in
			// if err != nil || c == nil {
			// 	next.ServeHTTP(w, r)
			// 	return
			// }

			_, claims, err := jwtauth.FromContext(r.Context())
			// Allow unauthenticated users in
			if claims == nil || err != nil {
				log.Println(err)
				next.ServeHTTP(w, r)
				return
			}
			identity := claims["session"].(map[string]interface{})["identity"]
			log.Println(identity)
			// parsed, err := gabs.ParseJSON(claims)
			// log.Println(parsed)
			// log.Println(err)

			next.ServeHTTP(w, r)
		})
	}
}

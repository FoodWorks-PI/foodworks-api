package auth

import (
	"context"
	"net/http"
	"strings"

	"foodworks.ml/m/internal/platform/utils"
	"github.com/Jeffail/gabs/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/lestrrat/go-jwx/jwk"
	"github.com/rs/zerolog/log"
)

var TokenAuth *jwtauth.JWTAuth

func InitAuth(jwksUrl string) {
	set, err := jwk.Fetch(jwksUrl)
	if err != nil {
		log.Panic().Err(err).Msg("Error gettint jwks")
	}
	public, private := set.Keys[0].Materialize()
	TokenAuth = jwtauth.New("RS256", public, private)
}

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}
type UserCredentials struct {
	Id    string
	Email string
}

func ForContext(ctx context.Context) *UserCredentials {
	raw, _ := ctx.Value(userCtxKey).(*UserCredentials)
	return raw
}

// Middleware decodes the jwt and packs the session into context
// TODO: Might make errors specific for graphql
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, _, err := jwtauth.FromContext(r.Context())
			if err != nil {
				utils.SimpleFail(w, err, http.StatusBadRequest)
				return
			}
			// Copied from jwt.Parser.ParseUnverified
			parts := strings.Split(token.Raw, ".")

			claimBytes, err := jwt.DecodeSegment(parts[1])
			if err != nil {
				utils.SimpleFail(w, err, http.StatusBadRequest)
				return
			}

			log.Info().Msg(token.Raw)
			parsed, err := gabs.ParseJSON(claimBytes)
			if err != nil {
				utils.SimpleFail(w, err, http.StatusBadRequest)
				return
			}
			log.Info().Interface("", parsed).Msg("")
			id, ok := parsed.Path("session.identity.id").Data().(string)
			if !ok {
				utils.SimpleFail(w, err, http.StatusBadRequest)
				return
			}
			email, ok := parsed.Path("session.identity.traits.email").Data().(string)
			if !ok {
				utils.SimpleFail(w, err, http.StatusBadRequest)
				return
			}

			userCredentials := &UserCredentials{Id: id, Email: email}

			ctx := context.WithValue(r.Context(), userCtxKey, userCredentials)

			// and call the next with our new context
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

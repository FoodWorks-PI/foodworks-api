package api

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"foodworks.ml/m/internal/auth"
	"github.com/go-chi/jwtauth"

	"foodworks.ml/m/internal/generated/ent"
	generated "foodworks.ml/m/internal/generated/graphql"
	"foodworks.ml/m/internal/resolver"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
)

type API struct {
	Router *chi.Mux
}

func (a *API) SetupRoutes(entClient *ent.Client, redisClient *redis.Client) {
	// init server
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.DefaultCompress)
	// Add CORS middleware around every request
	router.Use(cors.New(cors.Options{
		// AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		// Debug:            true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}).Handler)
	// router.Use(auth.Middleware())
	router.Use(render.SetContentType(render.ContentTypeJSON))

	srv := newGraphQLServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{Client: entClient, Redis: redisClient}}))

	router.Route("/graphql", func(router chi.Router) {
		router.Handle("/", srv)
		router.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	})

	router.Route("/", func(router chi.Router) {
		router.HandleFunc("/echo", EchoRequest)
	})

	router.Group(func(router chi.Router) {
		//TODO: Verify works with unit tests
		if flag.Lookup("test.v") == nil {
			auth.InitAuth()
			router.Use(jwtauth.Verifier(auth.TokenAuth))
			router.Use(jwtauth.Authenticator)
			fmt.Println("normal run")
		} else {
			fmt.Println("run under go test")
		}
		router.HandleFunc("/echoAuth", EchoRequest)
	})
	a.Router = router
}

func (a *API) StartServer() {
	appPort := os.Getenv("APPLICATION_PORT")
	log.Printf("connect to http://localhost:%s/graphql/playground for GraphQL playground", appPort)
	err := http.ListenAndServe(appPort, a.Router)
	if err != nil {
		log.Panic().Err(err)
	}
}

// This is needed for using CORS with websockets
func newGraphQLServer(es graphql.ExecutableSchema) *handler.Server {
	srv := handler.New(es)
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		KeepAlivePingInterval: 15 * time.Second,
	})

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return srv
}

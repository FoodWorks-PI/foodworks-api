package api

import (
	"context"
	"flag"
	"math/rand"
	"net/http"
	"os"
	"time"

	"foodworks.ml/m/internal/generated/ent/order"
	"foodworks.ml/m/internal/generated/graphql/model"
	"foodworks.ml/m/internal/recommendations"

	"foodworks.ml/m/internal/platform/filehandler"
	"github.com/elastic/go-elasticsearch/v6"

	"github.com/jmoiron/sqlx"

	"foodworks.ml/m/internal/platform"

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

func (a *API) SetupRoutes(entClient *ent.Client, dbClient *sqlx.DB, redisClient *redis.Client, dataStoreConfig platform.DataStoreConfig, elasticClient *elasticsearch.Client, fileHandler *filehandler.FileHandler, recommender recommendations.Reccomender) {
	localFileHandler := *fileHandler
	// init server
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(5))
	router.Use(middleware.Heartbeat("/health"))
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	}).Handler)
	// router.Use(auth.Middleware())

	router.Group(func(router chi.Router) {
		//TODO: Verify works with unit tests
		router.Use(render.SetContentType(render.ContentTypeJSON))
		if flag.Lookup("test.v") == nil {
			auth.InitAuth(dataStoreConfig.JWKSURL)
			router.Use(jwtauth.Verifier(auth.TokenAuth))
			router.Use(jwtauth.Authenticator)
			router.Use(auth.Middleware())
			log.Info().Msg("Auth Enabled")
		} else {
			log.Info().Msg("Auth Disabled")
		}
		router.Get("/media/{path}", func(w http.ResponseWriter, r *http.Request) {
			path := chi.URLParam(r, "path")
			localFileHandler.Download(w, r, path)
		})
		router.HandleFunc("/echoAuth", EchoRequest)
		router.Route("/graphql", func(router chi.Router) {
			config := generated.Config{
				Resolvers: &resolver.Resolver{EntClient: entClient, Redis: redisClient, DBClient: dbClient,
					ElasticClient: elasticClient, FileHandler: fileHandler, Recommender: recommender}}
			config.Directives.HasRole = resolver.HasRole(entClient)
			srv := newGraphQLServer(generated.NewExecutableSchema(config))
			router.Handle("/", srv)
			router.Handle("/playground", playground.Handler("GraphQL playground", "/api/graphql"))
		})
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

func (a *API) StartAutoPayment(client *ent.Client, db *sqlx.DB) {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for range ticker.C {
			ctx := context.Background()
			pendingOrders, err := client.Order.Query().Where(order.OrderStateEQ(model.OrderStatePendingPayment.String())).All(ctx)
			paidOrders, err := client.Order.Query().Where(order.OrderStateEQ(model.OrderStatePaid.String())).All(ctx)
			if err != nil {
				log.Err(err).Msg("")
			}
			for _, order := range pendingOrders {
				ind := rand.Float64()
				var orderState model.OrderState
				if ind < 0.80 {
					orderState = model.OrderStatePaid
				} else if ind < 0.95 {
					orderState = model.OrderStateCancelled
				} else {
					orderState = model.OrderStateError
				}
				_, _ = order.Update().SetOrderState(orderState.String()).SetUpdatedAt(time.Now().UTC()).Save(ctx)
			}
			for _, order := range paidOrders {
				ind := rand.Float64()
				var orderState model.OrderState
				if ind < 0.90 {
					orderState = model.OrderStateCompleted
				} else if ind < 0.95 {
					orderState = model.OrderStateCancelled
				} else {
					orderState = model.OrderStateError
				}
				_, _ = order.Update().SetOrderState(orderState.String()).SetUpdatedAt(time.Now().UTC()).Save(ctx)
			}
		}
	}()
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

package test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"testing"

	"foodworks.ml/m/internal/api"
	"foodworks.ml/m/internal/platform"
	"foodworks.ml/m/internal/platform/filehandler"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
)

type IntegrationTestSuite struct {
	suite.Suite
	TestSupport   *TestSupport
	DockerSupport *DockerSupport
}

type DockerSupport struct {
	Pool             *dockertest.Pool
	PostgresResource *dockertest.Resource
	RedisResource    *dockertest.Resource
}

func (suite *IntegrationTestSuite) TearDownSuite() {
	// When you're done, kill and remove the container
	if err := suite.DockerSupport.Pool.Purge(suite.DockerSupport.PostgresResource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	if err := suite.DockerSupport.Pool.Purge(suite.DockerSupport.RedisResource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func (suite *IntegrationTestSuite) SetupSuite() {
	// Setup Config
	var config platform.DataStoreConfig
	var testSupport TestSupport
	testSupport.Config = &config
	var dockerSupport DockerSupport

	// Setup Docker
	var err error
	var pool *dockertest.Pool
	// Windows workaround: see https://github.com/ory/dockertest/issues/142
	if runtime.GOOS == "windows" {
		fmt.Println("Hello from Windows")
		pool, err = dockertest.NewPool("npipe:////./pipe/docker_engine")
	} else {
		pool, err = dockertest.NewPool("")
	}

	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	suite.TestSupport = &testSupport
	suite.DockerSupport = &dockerSupport
	dockerSupport.Pool = pool

	// Setup Postgres
	var db *sql.DB
	database := "foodworks-test"
	postgresResource, err := pool.Run("postgres", "latest", []string{"POSTGRES_PASSWORD=foodworks", "POSTGRES_DB=" + database})
	dockerSupport.PostgresResource = postgresResource
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	_ = postgresResource.Expire(60)

	if err = pool.Retry(func() error {
		var err error
		config.DatabaseURL = fmt.Sprintf("postgres://postgres:foodworks@localhost:%s/%s?sslmode=disable", postgresResource.GetPort("5432/tcp"), database)
		db, err = sql.Open("pgx", config.DatabaseURL)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	// Setup Redis
	var rdb *redis.Client
	redisResource, err := pool.Run("redis", "latest", nil)
	dockerSupport.RedisResource = redisResource
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	_ = redisResource.Expire(60)

	config.RedisPass = ""
	var ctx = context.Background()
	if err = pool.Retry(func() error {
		config.RedisAddr = fmt.Sprintf("localhost:%s", redisResource.GetPort("6379/tcp"))
		rdb = redis.NewClient(&redis.Options{
			Addr: config.RedisAddr,
		})

		return rdb.Ping(ctx).Err()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Setup API
	SetupAPI(&testSupport)
}

type TestSupport struct {
	Handler http.Handler
	Config  *platform.DataStoreConfig
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

//TODO: Dummy test, need to seed db

func (suite *IntegrationTestSuite) Test_Auth_Disabled() {
	apitest.New().
		Handler(suite.TestSupport.Handler).
		Post("/echoAuth").
		Expect(suite.T()).
		Status(http.StatusOK).
		End()
}

// func (suite *IntegrationTestSuite) Test_Create_Fetch() {
// 	apitest.New().
// 		Handler(suite.TestSupport.Handler).
// 		Post("/graphql").
// 		GraphQLQuery(`mutation { createUser(input:{name:"John", age:20})}`).
// 		Expect(suite.T()).
// 		Status(http.StatusOK).
// 		End()
// 	apitest.New().
// 		Handler(suite.TestSupport.Handler).
// 		Post("/graphql").
// 		GraphQLQuery(`query { users(where:{name:"John"}) {name} }`).
// 		Expect(suite.T()).
// 		Status(http.StatusOK).
// 		Assert(jsonpath.Len(`$.data.users`, 1)).
// 		End()
// }

// TODO: Abstract
func SetupAPI(support *TestSupport) {

	db, client := platform.NewEntClient(*support.Config)

	rdb := platform.NewRedisClient(*support.Config)

	elasticClient := platform.NewElasticSearchClient(*support.Config)

	api := api.API{}
	fileHandler := filehandler.NewDiskUploader()
	api.SetupRoutes(client, db, rdb, *support.Config, elasticClient, &fileHandler)
	support.Handler = api.Router
}

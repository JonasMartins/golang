package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"src/graph"
	"src/graph/generated"
	"src/infra/orm/gorm/config/db"
	"src/main/auth"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const defaultPort string = "4000"

type Application struct {
	srv    *handler.Server
	logger *log.Logger
	port   uint16
	host   string
	db     *gorm.DB
}

func Run() error {
	app, err := Build()
	if err != nil {
		return err
	}
	router := chi.NewRouter()

	router.Use(auth.Middleware(app.db))

	router.Handle("/graphql", playground.Handler("Project", "/query"))
	router.Handle("/query", app.srv)

	app.logger.Printf("server running at: http://%s:%d/graphql", app.host, app.port)
	app.logger.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", app.port), router))

	return nil
}

func Build() (*Application, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Application need a .env file.")
	}

	s_port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	if s_port == "" {
		s_port = defaultPort
	}

	if host == "" {
		host = "localhost"
	}

	port, err := strconv.ParseUint(s_port, 10, 16)
	if err != nil {
		log.Fatal("Error defiing the server port.")
	}

	db := db.Init()

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}})
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	srv := handler.NewDefaultServer(schema)

	app := Application{
		srv:    srv,
		logger: logger,
		port:   uint16(port),
		host:   host,
		db:     db,
	}

	if err != nil {
		return nil, err
	}

	return &app, nil
}

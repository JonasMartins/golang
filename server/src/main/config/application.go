package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"server/src/graph"
	"server/src/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort string = "4000"

type application struct {
	srv    *handler.Server
	logger *log.Logger
	port   uint16
	host   string
}

func Run() {
	app := build()

	http.Handle("/", playground.Handler("Project", "/query"))
	http.Handle("/query", app.srv)

	app.logger.Printf("server running at: http://%s:%d", app.host, app.port)
	app.logger.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", app.port), nil))
}

func build() *application {

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

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	srv := handler.NewDefaultServer(schema)

	app := application{
		srv:    srv,
		logger: logger,
		port:   uint16(port),
		host:   host,
	}

	return &app
}

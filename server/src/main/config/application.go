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
	dbHandler "src/infra/orm/gorm/repositories"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort string = "4000"

type Application struct {
	srv    *handler.Server
	logger *log.Logger
	port   uint16
	host   string
}

func Run() error {
	app, err := Build()
	if err != nil {
		return err
	}

	DB := db.Init()
	orm := dbHandler.New(DB)

	fmt.Println("orm ", orm)

	http.Handle("/", playground.Handler("Project", "/query"))
	http.Handle("/query", app.srv)

	app.logger.Printf("server running at: http://%s:%d", app.host, app.port)
	app.logger.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", app.port), nil))

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

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	srv := handler.NewDefaultServer(schema)

	app := Application{
		srv:    srv,
		logger: logger,
		port:   uint16(port),
		host:   host,
	}

	if err != nil {
		return nil, err
	}

	return &app, nil
}

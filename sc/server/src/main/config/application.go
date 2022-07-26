package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"src/graph/generated"
	resolver "src/graph/resolvers"
	"src/infra/orm/gorm/config/db"
	"src/main/auth"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

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
	origin := os.Getenv("ALLOWED_ORIGIN")
	if origin == "" {
		log.Fatal("need a origin address")
		os.Exit(1)
	}

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{origin},
		AllowCredentials: true,
	}).Handler)

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
		os.Exit(1)
	}

	s_port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	origin := os.Getenv("ALLOWED_ORIGIN")
	if origin == "" || s_port == "" || host == "" {
		log.Fatal("please add a valid env")
		os.Exit(1)
	}

	port, err := strconv.ParseUint(s_port, 10, 16)
	if err != nil {
		log.Fatal("Error defiing the server port.")
	}

	db := db.Init()

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		DB: db,
	},
		Directives: generated.DirectiveRoot{},
	})
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	srv := handler.NewDefaultServer(schema)
	srv.Use(extension.FixedComplexityLimit(10))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return r.Host == origin
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

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

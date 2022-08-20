package config

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{origin},
		AllowCredentials: true,
	})

	filesDir := http.Dir("/Users/jonasmartinssouza/Documents/Dev/golang/golang/sc/server/public/images")

	router.Handle("/graphql", playground.Handler("Project", "/query"))
	router.Handle("/query", c.Handler(app.srv))

	// palliative solution
	FileServer(router, "/files", filesDir)

	app.logger.Printf("server running at: http://%s:%d/graphql", app.host, app.port)
	app.logger.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", app.port), router))

	return nil
}

func Build() (*Application, error) {

	var mb int64 = 1 << 20

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Application need a .env file.")
		os.Exit(1)
	}

	s_port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	origin := os.Getenv("ALLOWED_ORIGIN")
	redisAddr := os.Getenv("REDIS_ADDRESS")
	storage := os.Getenv("UPLOAD_STORAGE")
	if origin == "" || s_port == "" || host == "" || redisAddr == "" || storage == "" {
		log.Fatal("please add a valid env")
		os.Exit(1)
	}

	cache, err := NewCache(context.Background(), redisAddr, 24*time.Hour)
	if err != nil {
		log.Fatalf("cannot create APQ redis cache: %v", err)
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
	srv := handler.New(schema)
	srv.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return r.Header["Origin"][0] == origin
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		// InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
		// 	return auth.WebSocketInit(ctx, initPayload)
		// },
	})
	srv.AddTransport(transport.MultipartForm{
		MaxMemory:     32 * mb,
		MaxUploadSize: 15 * mb,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(cache)
	srv.Use(extension.Introspection{})
	srv.Use(extension.FixedComplexityLimit(30))
	srv.Use(extension.AutomaticPersistedQuery{Cache: cache})

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

// solution for now, allowing the server to display the images
// but in the future, those images will be stored on a aws bucket
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

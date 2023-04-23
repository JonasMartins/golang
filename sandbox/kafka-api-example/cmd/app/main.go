package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	apachekafka "project/internal/infra/apacheKafka"
	"project/internal/infra/repository"
	"project/internal/infra/web"
	"project/internal/usecase"

	"github.com/go-chi/chi/v5"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.NewProductRepositoryMysql(db)
	createProductUsecase := usecase.NewCreateProductUsecase(repo)
	listProductsUsecase := usecase.NewListProductsUsecase(repo)

	productHandlers := web.NewProductHandlers(createProductUsecase, listProductsUsecase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductsHandler)

	go http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)
	go apachekafka.Consume([]string{"producs"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDTO{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			log.Printf("%v", err)
		}
		_, err = createProductUsecase.Execute(&dto)
	}
}

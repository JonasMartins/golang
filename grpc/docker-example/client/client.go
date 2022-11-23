package main

import (
	"context"
	"fmt"
	"log"
	"os"

	poetry "poetry-example/proto"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	log := log.New(os.Stdout, "GRPC CLIENT : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	if err := run(ctx, log); err != nil {
		log.Println("main: error:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *log.Logger) error {
	log.Println("main: Initializing GRPC client")
	defer log.Println("main: Completed")
	conn, err := grpc.Dial("localhost:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		errors.Wrap(err, "dialing")
	}
	client := poetry.NewProtobufServiceClient(conn)
	res, err := client.RandomPoetries(ctx, &poetry.RandomPoetriesRequest{NumberOfPoetries: int32(2)})
	if err != nil {
		return errors.Wrap(err, "calling 'client.RandomPoetries()'")
	}
	for _, poetry := range res.List {
		fmt.Println("\nTitle: ", poetry.Title)
		fmt.Println("Author: ", poetry.Author)
		fmt.Print("\n")
		for _, line := range poetry.Lines {
			fmt.Printf("\t%s\n", line)
		}
		fmt.Print("\n")
	}
	return nil
}

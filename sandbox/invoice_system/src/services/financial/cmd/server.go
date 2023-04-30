package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"project/src/pkg/discovery"
	"project/src/pkg/discovery/consul"
	"project/src/services/financial/internal/repository/mysql"
	"time"

	"project/src/gen"
	general "project/src/pkg/utils"
	configs "project/src/services/financial/configs"
	financialController "project/src/services/financial/internal/controller/financial"
	gRPCHandler "project/src/services/financial/internal/handler/grpc"

	"google.golang.org/grpc"

	// "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const serviceName = "financial"

func Run_gRPCServer() {
	cfg := GetConfig()
	port := cfg.API.Port
	registry, err := consul.NewRegistry(cfg.API.LoadBalancerAddr)
	if err != nil {
		log.Fatalf("Could not create a new register %v", err)
	}
	tlsCredentials, err := general.LoadTLSCredentials()
	if err != nil {
		log.Fatalf("could not load TLS credentials %v", err)
	}
	fmt.Println(tlsCredentials.Info())

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("%s:%d", cfg.API.Domain, cfg.API.Port)); err != nil {
		log.Fatalf("Could not register a new instance %v", err)
	}
	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)
	financialRepo, err := mysql.New()
	if err != nil {
		log.Fatalf("Could not start a db connection %v", err)
	}
	defer financialRepo.Db.Close()

	ctrl := financialController.New(financialRepo)
	h := gRPCHandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.API.Domain, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Started %s gRPC server at %s", serviceName, lis.Addr().String())
	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterFinancialServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Could not start gRPC server: %v", err)
	}
}

// * Get a pointer to config object with all
// * needed variables from yaml file
func GetConfig() *configs.Config {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config %v", err)
	}
	return cfg
}

func RunLoop() {
	for {
		fmt.Println("Server Running for ever")
		time.Sleep(1 * time.Minute)
	}
}

package main

import (
	"fmt"
	"kakao/platform/internal/proto/pb"
	"kakao/platform/services/account/server"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	os.Setenv("ACCOUNT_SERVICE_PORT", "8082")
	os.Setenv("ACCOUNT_SERVICE_HOST", "account")
	os.Setenv("ACCOUNT_DB", "account_kakao_db")
	os.Setenv("POSTGRES_HOST", "172.17.0.2")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_USER", "kakao_dba")
	os.Setenv("POSTGRES_PASSWORD", "kakao@123")

	port, exists := os.LookupEnv(("ACCOUNT_SERVICE_PORT"))
	if !exists {
		log.Fatalln("No ACCOUNT_SERVICE_PORT .env variable specified")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalln("TCP server start error", err.Error())
	}

	accountServer, err := server.NewAccountSerivce()
	if err != nil {
		log.Fatalln("could not create server instance", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterAccountServiceServer(grpcServer, accountServer)

	log.Println("starting account service with config")

	// if err := grpcServer.Serve(lis); err != nil {
	// 	fmt.Println(lis )
	// 	log.Fatalln("account service error", err)
	// }

	fmt.Println("Before Serve call")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("Error during Serve:", err)
		log.Fatalln("account service error", err)
	}
	fmt.Println("After Serve call")

}

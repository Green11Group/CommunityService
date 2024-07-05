package main

import (
	pb "GreenThumb/Community-Service/genproto/community"
	"GreenThumb/Community-Service/service"
	"GreenThumb/Community-Service/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("Error connection to Postgres?", err)
	}

	community := postgres.NewCommunity(db)

	communityserver := servicecommunity.NewCommunityServer(db, community)

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatal("Failed to listen", err)
	}

	server := grpc.NewServer()
	pb.RegisterCommunityServiceServer(server, communityserver)

	log.Println("Server is running on port :50053")
	if err := server.Serve(listener); err != nil {
		log.Fatal("Error starting the server", err)
	}
}

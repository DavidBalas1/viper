package main

import (
	"flag"
	"log"

	"viper/agent/modules/shell"
	pb "viper/protos/cmds"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "The address to connect to.")
)

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAgentClient(conn)
	go shell.RunModule(client)
	select {}
}
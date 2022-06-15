package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"viper/controller/agents"
	"viper/controller/commands"
	pb "viper/protos/cmds"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

var (
	agentServerPort        = flag.Int("port", 50051, "Agent server port")
	agentManagerServerPort = flag.Int("management-port", 50052, "Agent management server port")
)

func initAgentInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		peer, ok := peer.FromContext(stream.Context())
		if !ok {
			log.Fatal("Failed to get peer from context")
		}
		agents.InitAgent(peer.Addr.String())
		handler(srv, stream)
		return nil
	}
}

func runAgentServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *agentServerPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer(grpc.StreamInterceptor(initAgentInterceptor()))
	pb.RegisterAgentServer(server, &commands.AgentServer{})
	log.Printf("Agent server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func runAgentManagerServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *agentManagerServerPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterAgentManagerServer(server, &commands.AgentManagerServer{})
	log.Printf("Agent manager server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func main() {
	flag.Parse()
	go runAgentServer()
	go runAgentManagerServer()
	select {}
}

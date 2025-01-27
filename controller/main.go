package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"net"
	"viper"

	_ "embed"
	"viper/controller/agents"
	"viper/controller/commands"
	pb "viper/protos/cmds"

	"google.golang.org/grpc"
)

func runAgentServer() {
	certBuffers := viper.Conf.Controller.Cert
	cert, err := tls.X509KeyPair([]byte(certBuffers.Cert), []byte(certBuffers.Key))
	if err != nil {
		log.Fatalf("Failed to load certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	ok := caCertPool.AppendCertsFromPEM([]byte(viper.Conf.Agent.Cert.Cert))
	if !ok {
		log.Fatalf("Failed to parse agent's client certificate.")
	}
	tlsCfg := &tls.Config{Certificates: []tls.Certificate{cert}, ClientAuth: tls.RequireAndVerifyClientCert, ClientCAs: caCertPool}
	lis, err := tls.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", viper.Conf.Controller.AgentServerPort), tlsCfg)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Agent server listening at %v", lis.Addr())
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("Failed to accept connection: %v", err)
		}
		go agents.InitAgent(conn)
	}
}

func runAgentManagerServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", viper.Conf.Controller.AgentManagerServerPort))
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

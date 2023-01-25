package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/rhythmicsoul/nginx-mgmt/proto/controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	controller.UnimplementedAddServiceServer
}

func loadTLSCreds() (credentials.TransportCredentials, error) {
	caCert, err := ioutil.ReadFile("/home/rhythmic/certs/ca/rootCA.crt")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to add custom ca cert")
	}

	serverCert, err := tls.LoadX509KeyPair("/home/rhythmic/certs/mbishnu.com/fullchain.pem",
		"/home/rhythmic/certs/mbishnu.com/key.pem")
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

func (s *server) NewAgentToken(ctx context.Context, req *controller.Empty) (*controller.AgentToken, error) {
	log.Printf("est")
	return &controller.AgentToken{Token: "an agent token"}, nil
}

func (s *server) BiT(stream controller.AddService_BiTServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if in.Name == "version" {
			dat, _ := os.ReadFile("../tmp/version")
			ver, _ := strconv.Atoi(string(dat))
			stream.Send(
				&controller.Test{
					Name: string(dat),
					Age:  int32(ver),
					Add:  "nginx",
				})
			continue
		}

		log.Println(in)
		stream.Send(&controller.Test{
			Name: "bishnu",
			Age:  20,
			Add:  "Baikuntha",
		})
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:9091")
	if err != nil {
		log.Fatalf("Starting of server failed: %v", err)
	}

	tlsCreds, err := loadTLSCreds()
	if err != nil {
		log.Fatalf("Couldn't load the certs. Error: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(tlsCreds))
	// s := grpc.NewServer()
	controller.RegisterAddServiceServer(s, &server{})
	log.Printf("server listening at: %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

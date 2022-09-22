package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/rhythmicsoul/nginx-mgmt/proto/controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func loadTLSCreds() (credentials.TransportCredentials, error) {
	caCert, err := ioutil.ReadFile("/home/rhythmic/certs/ca/rootCA.crt")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to add custom ca cert")
	}

	serverCert, err := tls.LoadX509KeyPair("/home/rhythmic/certs/node1/fullchain.pem",
		"/home/rhythmic/certs/node1/key.pem")
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	tlsCreds, err := loadTLSCreds()
	if err != nil {
		log.Fatalf("Could not load tls config/ca certs. Error: %v", err)
	}
	dctx, dcancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer dcancel()

	conn, err := grpc.DialContext(dctx, "mbishnu.com.np:9091", grpc.WithTransportCredentials(tlsCreds), grpc.WithBlock(), grpc.WithReturnConnectionError())

	if err != nil {
		log.Fatalf("Cannot Connect to Server: %v", err)
	}
	defer conn.Close()

	c := controller.NewAddServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, _ := c.NewAgentToken(ctx, &controller.Empty{})
	log.Printf("the token is: %s", r.GetToken())
}

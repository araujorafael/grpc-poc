package rpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

// Client Client abstraction
type Client struct {
	address        string
	connection     *grpc.ClientConn
	ctx            context.Context
	DiscountClient DiscountServiceClient
}

// NewClient create a new implememntation RPC client
func NewClient(address string) Client {
	return Client{
		address: address,
		ctx:     context.Background(),
	}
}

// Connect create grpc connection with specified configurarion
func (c Client) Connect() *grpc.ClientConn {
	k := keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}
	opts := grpc.WaitForReady(false)
	conn, err := grpc.Dial(
		c.address,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(k),
		grpc.WithDefaultCallOptions(opts),
	)
	if err != nil {
		log.Println("Cannot connect on "+c.address, err.Error())
	} else {
		log.Println("Connected on ", c.address)
	}
	c.connection = conn
	return conn
}

// Disconnect disconnects client
func (c Client) Disconnect() {
	if c.connection != nil {
		log.Println("Desconectei de " + c.address)
		c.connection.Close()
	}
}

package main

import (
	"context"
	"log"
	"time"

	config "github.com/davidmontoyago/go-grpc-gossiping-cluster/api"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	conn, configClient := newClient(":9000")
	defer conn.Close()
	configReq := &config.PutConfigRequest{
		Key:   "cluster.config.test.prop",
		Value: "distributed!",
	}
	resp, err := configClient.Put(ctx, configReq)
	if err != nil {
		log.Fatal("error calling config.Put", err)
	}
	log.Println("Successfully put config", resp)

	log.Println("checking other nodes...")
	conn, configClient = newClient(":9001")
	defer conn.Close()
	getReq := &config.GetConfigRequest{
		Key: "cluster.config.test.prop",
	}
	for {
		resp, err = configClient.Get(ctx, getReq)
		if resp.GetValue() == "distributed!" {
			log.Println("config found! done.")
			break
		} else {
			log.Println("config not available yet; trying again in 3 seconds...")
			time.Sleep(3 * time.Second)
		}
	}
}

func newClient(addr string) (*grpc.ClientConn, config.ConfigServiceClient) {
	var err error
	var conn *grpc.ClientConn

	conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to connect to api: %s", err)
	}

	return conn, config.NewConfigServiceClient(conn)
}

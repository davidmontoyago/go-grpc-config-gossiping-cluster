package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	config "github.com/davidmontoyago/go-grpc-gossiping-cluster/api"
	"google.golang.org/grpc"

	"github.com/hashicorp/memberlist"
)

// Node is a gRPC serving node of the cluster
type Node struct {
	config.UnimplementedConfigServiceServer
	// name to identify the node by
	Name string

	// host and node ports for gossiping and api
	addr       string
	apiPort    int
	gossipPort int

	// addr:port of any node in the cluster; empty if it's the first node
	clusterNodeAddr string

	// node internal state
	config     map[string]string
	grpcServer *grpc.Server
	memberlist *memberlist.Memberlist
}

// NewNode creates new gRPC serving node but does not start serving
func NewNode(name string, addr string, apiPort, gossipPort int, clusterNodeAddr string) *Node {
	return &Node{
		Name:            name,
		addr:            addr,
		apiPort:         apiPort,
		gossipPort:      gossipPort,
		clusterNodeAddr: clusterNodeAddr,
		config:          make(map[string]string, 10),
	}
}

// Put adds config to the local store
func (n *Node) Put(ctx context.Context, req *config.PutConfigRequest) (*config.Config, error) {
	key := req.GetKey()
	value := req.GetValue()

	// update local state
	n.config[key] = value
	log.Println("succesfully put config", key, value)

	go n.distributeConfig(*req)

	return &config.Config{Key: key, Value: value}, nil
}

// Get fetches config from the local store
func (n *Node) Get(ctx context.Context, req *config.GetConfigRequest) (*config.Config, error) {
	key := req.GetKey()
	value := n.config[key]

	return &config.Config{Key: key, Value: value}, nil
}

// Start async runs gRPC server and joins cluster
func (n *Node) Start() chan error {
	errChan := make(chan error)
	go n.serve(errChan)
	go n.joinCluster(errChan)
	return errChan
}

// Shutdown stops gRPC server and leaves cluster
func (n *Node) Shutdown() {
	n.grpcServer.GracefulStop()
	n.memberlist.Leave(15 * time.Second)
	n.memberlist.Shutdown()
}

func (n *Node) serve(errChan chan error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", n.addr, n.apiPort))
	if err != nil {
		log.Printf("failed to listen on %s: %v", n.addr, err)
		errChan <- err
	}
	log.Println("grpc api serving on", n.addr, n.apiPort)

	n.grpcServer = grpc.NewServer()
	config.RegisterConfigServiceServer(n.grpcServer, n)
	if err := n.grpcServer.Serve(lis); err != nil {
		log.Println("failed to serve", err)
		errChan <- err
	}
}

func (n *Node) joinCluster(errChan chan error) {
	config := memberlist.DefaultLocalConfig()
	config.Name = n.Name
	config.BindAddr = n.addr
	config.BindPort = n.gossipPort
	config.AdvertisePort = config.BindPort

	var err error
	n.memberlist, err = memberlist.Create(config)
	if err != nil {
		log.Println("failed to init memberlist", err)
		errChan <- err
	}

	var nodeAddr string
	if n.clusterNodeAddr != "" {
		log.Printf("not the first node, joining %s...", n.clusterNodeAddr)
		nodeAddr = n.clusterNodeAddr
	} else {
		log.Println("first node of the cluster...")
		nodeAddr = fmt.Sprintf("%s:%d", n.addr, n.gossipPort)
	}
	_, err = n.memberlist.Join([]string{nodeAddr})
	if err != nil {
		log.Println("failed to join cluster", err)
		errChan <- err
	}
	log.Println("succesfully joined cluster via", nodeAddr)
}

func (n *Node) distributeConfig(req config.PutConfigRequest) {
	for _, member := range n.memberlist.Members() {
		if member == n.memberlist.LocalNode() {
			continue
		}

		log.Println("distributing config to", member.Addr)
	}
}

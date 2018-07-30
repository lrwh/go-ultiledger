package api

import (
	"log"
	"net"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	p "github.com/ultiledger/go-ultiledger/peer"
	"github.com/ultiledger/go-ultiledger/ultpb/rpc"
)

// ultNode is the central controller for ultiledger
type ultNode struct {
	// config viper
	config *ultNodeConfig
	// zap logger
	logger *zap.SugaredLogger
	// ULTNode server
	server *ULTNodeServer

	// IP address of this node
	IP string
	// start time of the node
	StartTime int64
	// list of connected peers
	Peers []*p.Peer

	// stop channel
	stopChan chan struct{}
}

// NewULTNode creates a ultNode which controls all the sub tasks
func NewULTNode(conf *ultNodeConfig) *ultNode {
	// initialize logger
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	// get outbound IP
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	addr := conn.LocalAddr().(*net.UDPAddr)

	node := &ultNode{
		config:    conf,
		logger:    l.Sugar(),
		server:    &ULTNodeServer{},
		IP:        addr.String(),
		StartTime: time.Now().Unix(),
		stopChan:  make(chan struct{}),
	}

	return node
}

// Start checks the provided configurations, if the config is valid,
// it will trigger sub goroutines to do the sub tasks.
func (u *ultNode) Start() error {
	// TODO(bobonovski) check the validity of config in viper

	go u.serveULTNode()
	select {}
	return nil
}

// Restart checks the provided configurations, if the config is valid,
// it will trigger sub goroutines to do the sub tasks.
func (u *ultNode) Restart() error {
	log.Println("restart called")
	return nil
}

// serve starts a listener on the port and starts to accept request
func (u *ultNode) serveULTNode() {
	// register rpc service and start the ULTNode server
	listener, err := net.Listen("tcp", u.config.Port)
	if err != nil {
		u.logger.Fatal(err)
	}

	s := grpc.NewServer()
	rpc.RegisterULTNodeServer(s, u.server)

	u.logger.Infof("start to serve gRPC requests on %s", u.config.Port)
	go s.Serve(listener)

	for {
		select {
		case <-u.stopChan:
			u.logger.Infof("gracefully shutdown gRPC server")
			s.GracefulStop()
			return
		}
	}
}

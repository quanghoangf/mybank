package gapi

import (
	"fmt"

	db "github.com/quanghoangf/mybank/db/sqlc"
	"github.com/quanghoangf/mybank/pb"
	"github.com/quanghoangf/mybank/token"
	"github.com/quanghoangf/mybank/util"
	"github.com/quanghoangf/mybank/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedMybankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}

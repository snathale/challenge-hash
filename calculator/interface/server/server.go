package server

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
	"github.com/snathale/challenge-hash/calculator/application/controller"
	"github.com/snathale/challenge-hash/calculator/interface/proto"
	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
)

var GrpcNewServerError = errors.New("impossible create a grpc server")

type Server struct {
	server *grpc.Server
	lis    net.Listener
}

func NewServer(config Config, ctrl controller.Controller) (*Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		fmt.Println("aquiii")
		logrus.WithError(err).Warning(GrpcNewServerError)
		return nil, GrpcNewServerError
	}
	grpcServer := grpc.NewServer()
	calculateServer := NewCalculatorServer(ctrl)
	proto.RegisterCalculatorServer(grpcServer, calculateServer)
	return &Server{
		server: grpcServer,
		lis:    lis,
	}, nil
}

func (s *Server) Run() <-chan error {
	errChan := make(chan error)
	go func() {
		defer close(errChan)
		if err := s.server.Serve(s.lis); err != nil {
			errChan <- err
		}
	}()
	return errChan
}

func (s *Server) Close() {
	s.lis.Close()
}

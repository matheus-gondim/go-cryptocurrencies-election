package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/matheus-gondim/go-cryptocurrencies-election/src/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CryptocurrencyElectionServer struct {
	db *gorm.DB
	pb.UnimplementedCryptocurrencyElectionServer
}

func NewCryptocurrencyElectionServer(db *gorm.DB) *CryptocurrencyElectionServer {
	return &CryptocurrencyElectionServer{db: db}
}

func (s *CryptocurrencyElectionServer) Run(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to listen: %v", err))
	}

	server := grpc.NewServer()
	pb.RegisterCryptocurrencyElectionServer(server, s)
	log.Printf("server listening at %v", lis.Addr())

	return server.Serve(lis)
}

func (s *CryptocurrencyElectionServer) CreateNew(ctx context.Context, in *pb.CreateCryptocurrency) (*pb.Cryptocurrency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNew not implemented")

}

func (s *CryptocurrencyElectionServer) FindCryptocurrencies(ctx context.Context, in *pb.GetCryptocurrencyParams) (*pb.CryptocurrencyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCryptocurrencies not implemented")
}

func (s *CryptocurrencyElectionServer) FindById(context.Context, *pb.CryptocurrencyId) (*pb.Cryptocurrency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (s *CryptocurrencyElectionServer) DeleteById(context.Context, *pb.CryptocurrencyId) (*pb.CryptocurrencyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (s *CryptocurrencyElectionServer) UpdateById(context.Context, *pb.UpdateCryptocurrency) (*pb.CryptocurrencyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateById not implemented")
}
func (s *CryptocurrencyElectionServer) UpvoteById(context.Context, *pb.CryptocurrencyId) (*pb.Cryptocurrency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpvoteById not implemented")
}
func (s *CryptocurrencyElectionServer) DownvoteById(context.Context, *pb.CryptocurrencyId) (*pb.Cryptocurrency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownvoteById not implemented")
}

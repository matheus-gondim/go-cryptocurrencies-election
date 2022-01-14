package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/matheus-gondim/go-cryptocurrencies-election/src/entity"
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
	if err := in.IsValid(); err != nil {
		return nil, err
	}

	c := new(entity.Cryptocurrency).FromCreate(in)

	if err := s.db.Create(&c).Error; err != nil {
		return nil, err
	}

	return c.ToOutput(), nil
}

func (s *CryptocurrencyElectionServer) FindCryptocurrencies(ctx context.Context, in *pb.GetCryptocurrencyParams) (*pb.CryptocurrencyList, error) {
	cryptocurrencylist := []entity.Cryptocurrency{}
	if err := s.db.Find(&cryptocurrencylist).Error; err != nil {
		return nil, err
	}

	cryptocurrencies := []*pb.Cryptocurrency{}
	for _, c := range cryptocurrencylist {
		cryptocurrencies = append(cryptocurrencies, c.ToOutput())
	}

	return &pb.CryptocurrencyList{Cryptocurrencies: cryptocurrencies}, nil
}

func (s *CryptocurrencyElectionServer) FindById(ctx context.Context, in *pb.CryptocurrencyId) (*pb.Cryptocurrency, error) {
	c, err := s.findCryptocurrencyById(in.GetId())
	if err != nil {
		return nil, err
	}

	return c.ToOutput(), nil
}

func (s *CryptocurrencyElectionServer) DeleteById(ctx context.Context, in *pb.CryptocurrencyId) (*pb.CryptocurrencyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}

func (s *CryptocurrencyElectionServer) UpdateById(context.Context, *pb.UpdateCryptocurrency) (*pb.CryptocurrencyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateById not implemented")
}

func (s *CryptocurrencyElectionServer) UpvoteById(ctx context.Context, in *pb.CryptocurrencyId) (*pb.Cryptocurrency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpvoteById not implemented")
}

func (s *CryptocurrencyElectionServer) DownvoteById(ctx context.Context, in *pb.CryptocurrencyId) (*pb.Cryptocurrency, error) {
	c, err := s.findCryptocurrencyById(in.GetId())
	if err != nil {
		return nil, err
	}

	c.Dislike++

	if err := s.db.Save(c).Error; err != nil {
		return nil, err
	}

	return c.ToOutput(), nil
}

func (s *CryptocurrencyElectionServer) findCryptocurrencyById(id int64) (*entity.Cryptocurrency, error) {
	if id <= 0 {
		return nil, errors.New("Id is a required parameter")
	}
	c := &entity.Cryptocurrency{}

	if err := s.db.First(c, id).Error; err != nil {
		return nil, err
	}

	return c, nil
}

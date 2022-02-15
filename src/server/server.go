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
	"gorm.io/gorm"
)

type CryptocurrencyElectionServer struct {
	db *gorm.DB
	pb.UnimplementedCryptocurrencyElectionServer
	steams map[int64]func()
}

func NewCryptocurrencyElectionServer(db *gorm.DB) *CryptocurrencyElectionServer {
	return &CryptocurrencyElectionServer{
		db:     db,
		steams: make(map[int64]func()),
	}
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
	c, err := s.findCryptocurrencyById(in.GetId())
	if err != nil {
		return nil, err
	}

	if err := s.db.Delete(c).Error; err != nil {
		return nil, err
	}

	message := fmt.Sprintf("%s cryptocurrency has been deleted", c.Name)
	return &pb.CryptocurrencyMessage{Message: message}, nil
}

func (s *CryptocurrencyElectionServer) UpdateById(ctx context.Context, in *pb.UpdateCryptocurrency) (*pb.CryptocurrencyMessage, error) {
	if err := in.IsValid(); err != nil {
		return nil, err
	}

	c, err := s.findCryptocurrencyById(in.GetId())
	if err != nil {
		return nil, err
	}

	c.FromUpdate(in)

	s.update(c)
	message := fmt.Sprintf("%d cryptocurrency has been updated", c.ID)
	return &pb.CryptocurrencyMessage{Message: message}, nil
}

func (s *CryptocurrencyElectionServer) UpvoteById(ctx context.Context, in *pb.CryptocurrencyId) (*pb.Cryptocurrency, error) {
	c, err := s.findCryptocurrencyById(in.GetId())
	if err != nil {
		return nil, err
	}

	c.Like++

	s.update(c)
	return c.ToOutput(), nil
}

func (s *CryptocurrencyElectionServer) DownvoteById(ctx context.Context, in *pb.CryptocurrencyId) (*pb.Cryptocurrency, error) {
	c, err := s.findCryptocurrencyById(in.GetId())
	if err != nil {
		return nil, err
	}

	c.Dislike++

	s.update(c)
	return c.ToOutput(), nil
}

func (s *CryptocurrencyElectionServer) FetchResponse(in *pb.CryptocurrencyId, srv pb.CryptocurrencyElection_FetchResponseServer) error {
	id := in.GetId()

	if id <= 0 {
		return errors.New("Id is a required parameter")
	}

	s.steams[id] = func() {
		c, _ := s.findCryptocurrencyById(id)

		if err := srv.Send(c.ToOutput()); err != nil {
			log.Printf("send error %v", err)
		}
	}

	s.steams[id]()

	select {}
}

func (s *CryptocurrencyElectionServer) update(c *entity.Cryptocurrency) (*entity.Cryptocurrency, error) {
	if err := s.db.Save(c).Error; err != nil {
		return nil, err
	}

	if s.steams[c.ID] != nil {
		s.steams[c.ID]()
	}

	return c, nil
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

package server

import (
	"context"
	"log"
	"reflect"
	"testing"

	"github.com/matheus-gondim/go-cryptocurrencies-election/src/entity"
	"github.com/matheus-gondim/go-cryptocurrencies-election/src/pb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDBTest() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatal()
	}
	db.AutoMigrate(&entity.Cryptocurrency{})
	return db
}

func closeDBTest(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()
}

type Server struct {
	db                                        *gorm.DB
	UnimplementedCryptocurrencyElectionServer pb.UnimplementedCryptocurrencyElectionServer
}

func Test_CreateNew_WhenSuccess(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	type args struct {
		ctx context.Context
		in  *pb.CreateCryptocurrency
	}
	tests := []struct {
		name    string
		fields  Server
		args    args
		want    *pb.Cryptocurrency
		wantErr bool
	}{
		{
			name:   "",
			fields: Server{db: db},
			args: args{
				context.Background(),
				&pb.CreateCryptocurrency{Name: "Cryptocurrency Test", Symbol: "CCT"},
			},
			want:    &pb.Cryptocurrency{Id: 1, Name: "Cryptocurrency Test", Symbol: "CCT", Like: 0, Dislike: 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptocurrencyElectionServer{
				db: tt.fields.db,
				UnimplementedCryptocurrencyElectionServer: tt.fields.UnimplementedCryptocurrencyElectionServer,
			}
			got, err := s.CreateNew(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptocurrencyElectionServer.CreateNew() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CryptocurrencyElectionServer.CreateNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CreateNew_BadRequest(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	type args struct {
		ctx context.Context
		in  *pb.CreateCryptocurrency
	}
	tests := []struct {
		name    string
		fields  Server
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "",
			fields: Server{db: db},
			args: args{
				context.Background(),
				&pb.CreateCryptocurrency{Name: "", Symbol: ""},
			},
			want:    "name: cannot be blank; symbol: cannot be blank.",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptocurrencyElectionServer{
				db: tt.fields.db,
				UnimplementedCryptocurrencyElectionServer: tt.fields.UnimplementedCryptocurrencyElectionServer,
			}
			_, err := s.CreateNew(tt.args.ctx, tt.args.in)
			if !reflect.DeepEqual(err.Error(), tt.want) {
				t.Errorf("CryptocurrencyElectionServer.CreateNew() = %v, want %v", err, tt.want)
			}
		})
	}
}

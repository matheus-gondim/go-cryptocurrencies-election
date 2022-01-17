package server

import (
	"context"
	"fmt"
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

var (
	cryptocurrencyMock       = &pb.Cryptocurrency{Id: 1, Name: "Cryptocurrency Test", Symbol: "CCT", Like: 0, Dislike: 0}
	createCryptocurrencyMock = &pb.CreateCryptocurrency{Name: "Cryptocurrency Test", Symbol: "CCT"}
	contextMock              = context.Background()
	cryptocurrencyListMock   = &pb.CryptocurrencyList{Cryptocurrencies: []*pb.Cryptocurrency{cryptocurrencyMock}}
	cryptocurrencyIdMock     = &pb.CryptocurrencyId{Id: 1}
)

type BasicTestStruct struct {
	name    string
	fields  Server
	wantErr bool
}

func Test_CreateNew_WhenSuccess(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	type args struct {
		ctx context.Context
		in  *pb.CreateCryptocurrency
	}
	tests := []struct {
		BasicTestStruct
		args args
		want *pb.Cryptocurrency
	}{
		{
			BasicTestStruct: BasicTestStruct{
				name:    "should create a cryptocurrency object when successful",
				fields:  Server{db: db},
				wantErr: false,
			},
			args: args{contextMock, createCryptocurrencyMock},
			want: cryptocurrencyMock,
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

func Test_CreateNew_WhenBadRequest(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	type args struct {
		ctx context.Context
		in  *pb.CreateCryptocurrency
	}
	tests := []struct {
		BasicTestStruct
		args args
		want string
	}{
		{
			BasicTestStruct: BasicTestStruct{
				name:    "should return an error when the dto is invalid",
				fields:  Server{db: db},
				wantErr: true,
			},
			args: args{contextMock, &pb.CreateCryptocurrency{}},
			want: "name: cannot be blank; symbol: cannot be blank.",
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

func Test_FindCryptocurrencies_WhenSuccess(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	db.Create(cryptocurrencyMock)

	type args struct {
		ctx context.Context
		in  *pb.GetCryptocurrencyParams
	}

	tests := []struct {
		BasicTestStruct
		args args
		want *pb.CryptocurrencyList
	}{
		{
			BasicTestStruct: BasicTestStruct{
				name:    "should return a list of cryptocurrencies when successful",
				fields:  Server{db: db},
				wantErr: false,
			},
			args: args{ctx: contextMock, in: &pb.GetCryptocurrencyParams{}},
			want: cryptocurrencyListMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptocurrencyElectionServer{
				db: tt.fields.db,
				UnimplementedCryptocurrencyElectionServer: tt.fields.UnimplementedCryptocurrencyElectionServer,
			}
			got, err := s.FindCryptocurrencies(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptocurrencyElectionServer.FindCryptocurrencies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CryptocurrencyElectionServer.FindCryptocurrencies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_FindById_WhenSuccess(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	db.Create(cryptocurrencyMock)

	type args struct {
		ctx context.Context
		in  *pb.CryptocurrencyId
	}

	tests := []struct {
		BasicTestStruct
		args args
		want *pb.Cryptocurrency
	}{
		{
			BasicTestStruct: BasicTestStruct{
				name:    "should return a cryptocurrency by id when successful",
				fields:  Server{db: db},
				wantErr: false,
			},
			args: args{contextMock, cryptocurrencyIdMock},
			want: cryptocurrencyMock,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptocurrencyElectionServer{
				db: tt.fields.db,
				UnimplementedCryptocurrencyElectionServer: tt.fields.UnimplementedCryptocurrencyElectionServer,
			}
			got, err := s.FindById(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptocurrencyElectionServer.FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CryptocurrencyElectionServer.FindById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_FindById_WhenBadRequest(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	type args struct {
		ctx context.Context
		in  *pb.CryptocurrencyId
	}

	tests := []struct {
		BasicTestStruct
		args args
		want string
	}{
		{
			BasicTestStruct: BasicTestStruct{
				name:    "should return an error when the dto is invalid",
				fields:  Server{db: db},
				wantErr: false,
			},
			args: args{contextMock, &pb.CryptocurrencyId{}},
			want: "Id is a required parameter",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptocurrencyElectionServer{
				db: tt.fields.db,
				UnimplementedCryptocurrencyElectionServer: tt.fields.UnimplementedCryptocurrencyElectionServer,
			}
			_, err := s.FindById(tt.args.ctx, tt.args.in)

			if !reflect.DeepEqual(err.Error(), tt.want) {
				t.Errorf("CryptocurrencyElectionServer.FindById() = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_FindById_WhenNotFoundEntity(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	db.Create(cryptocurrencyMock)

	type args struct {
		ctx context.Context
		in  *pb.CryptocurrencyId
	}

	tests := []struct {
		BasicTestStruct
		args args
		want string
	}{
		{
			BasicTestStruct: BasicTestStruct{
				name:    "should return an error when entity not found",
				fields:  Server{db: db},
				wantErr: false,
			},
			args: args{contextMock, &pb.CryptocurrencyId{Id: 2}},
			want: "record not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptocurrencyElectionServer{
				db: tt.fields.db,
				UnimplementedCryptocurrencyElectionServer: tt.fields.UnimplementedCryptocurrencyElectionServer,
			}
			_, err := s.FindById(tt.args.ctx, tt.args.in)

			if !reflect.DeepEqual(err.Error(), tt.want) {
				t.Errorf("CryptocurrencyElectionServer.FindById() = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_UpvoteById_WhenSuccess(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	db.Create(cryptocurrencyMock)

	type args struct {
		ctx context.Context
		in  *pb.CryptocurrencyId
	}
	tests := []struct {
		BasicTestStruct
		args args
		want int64
	}{
		{
			BasicTestStruct: BasicTestStruct{
				name:    "should add a like for cryptocurrency",
				fields:  Server{db: db},
				wantErr: false,
			},
			args: args{contextMock, &pb.CryptocurrencyId{Id: 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptocurrencyElectionServer{
				db: tt.fields.db,
				UnimplementedCryptocurrencyElectionServer: tt.fields.UnimplementedCryptocurrencyElectionServer,
			}
			got, err := s.UpvoteById(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptocurrencyElectionServer.UpvoteById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Like, tt.want) {
				t.Errorf("CryptocurrencyElectionServer.UpvoteById() = %v, want %v", got.Like, tt.want)
			}
		})
	}
}

func Test_UpvoteById_WhenNotFoundEntity(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	db.Create(cryptocurrencyMock)

	type args struct {
		ctx context.Context
		in  *pb.CryptocurrencyId
	}

	tests := []struct {
		BasicTestStruct
		args args
		want string
	}{
		{
			BasicTestStruct: BasicTestStruct{
				name:    "should return an error when entity not found",
				fields:  Server{db: db},
				wantErr: false,
			},
			args: args{contextMock, &pb.CryptocurrencyId{Id: 2}},
			want: "record not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptocurrencyElectionServer{
				db: tt.fields.db,
				UnimplementedCryptocurrencyElectionServer: tt.fields.UnimplementedCryptocurrencyElectionServer,
			}
			_, err := s.UpvoteById(tt.args.ctx, tt.args.in)

			if !reflect.DeepEqual(err.Error(), tt.want) {
				t.Errorf("CryptocurrencyElectionServer.UpvoteById() = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_DownvoteById_WhenSuccess(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	db.Create(cryptocurrencyMock)

	type args struct {
		ctx context.Context
		in  *pb.CryptocurrencyId
	}
	tests := []struct {
		BasicTestStruct
		args args
		want int64
	}{
		{
			BasicTestStruct: BasicTestStruct{
				name:    "should add a dislike for cryptocurrency",
				fields:  Server{db: db},
				wantErr: false,
			},
			args: args{contextMock, &pb.CryptocurrencyId{Id: 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptocurrencyElectionServer{
				db: tt.fields.db,
				UnimplementedCryptocurrencyElectionServer: tt.fields.UnimplementedCryptocurrencyElectionServer,
			}
			got, err := s.DownvoteById(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptocurrencyElectionServer.DownvoteById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Dislike, tt.want) {
				t.Errorf("CryptocurrencyElectionServer.DownvoteById() = %v, want %v", got.Dislike, tt.want)
			}
		})
	}
}

func Test_DownvoteById_WhenNotFoundEntity(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	db.Create(cryptocurrencyMock)

	type args struct {
		ctx context.Context
		in  *pb.CryptocurrencyId
	}

	tests := []struct {
		BasicTestStruct
		args args
		want string
	}{
		{
			BasicTestStruct: BasicTestStruct{
				name:    "should return an error when entity not found",
				fields:  Server{db: db},
				wantErr: false,
			},
			args: args{contextMock, &pb.CryptocurrencyId{Id: 2}},
			want: "record not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptocurrencyElectionServer{
				db: tt.fields.db,
				UnimplementedCryptocurrencyElectionServer: tt.fields.UnimplementedCryptocurrencyElectionServer,
			}
			_, err := s.DownvoteById(tt.args.ctx, tt.args.in)

			if !reflect.DeepEqual(err.Error(), tt.want) {
				t.Errorf("CryptocurrencyElectionServer.DownvoteById() = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_DeleteById_WhenSuccess(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	db.Create(cryptocurrencyMock)

	type args struct {
		ctx context.Context
		in  *pb.CryptocurrencyId
	}
	tests := []struct {
		BasicTestStruct
		args args
		want *pb.CryptocurrencyMessage
	}{
		{
			BasicTestStruct: BasicTestStruct{
				name:    "should delete a cryptocurrency by id",
				fields:  Server{db: db},
				wantErr: false,
			},
			args: args{contextMock, &pb.CryptocurrencyId{Id: 1}},
			want: &pb.CryptocurrencyMessage{
				Message: fmt.Sprintf("%s cryptocurrency has been deleted", cryptocurrencyMock.Name),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptocurrencyElectionServer{
				db: tt.fields.db,
				UnimplementedCryptocurrencyElectionServer: tt.fields.UnimplementedCryptocurrencyElectionServer,
			}
			got, err := s.DeleteById(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptocurrencyElectionServer.DeleteById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CryptocurrencyElectionServer.DeleteById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DeleteById_WhenNotFoundEntity(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	db.Create(cryptocurrencyMock)

	type args struct {
		ctx context.Context
		in  *pb.CryptocurrencyId
	}

	tests := []struct {
		BasicTestStruct
		args args
		want string
	}{
		{
			BasicTestStruct: BasicTestStruct{
				name:    "should return an error when entity not found",
				fields:  Server{db: db},
				wantErr: false,
			},
			args: args{contextMock, &pb.CryptocurrencyId{Id: 2}},
			want: "record not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptocurrencyElectionServer{
				db: tt.fields.db,
				UnimplementedCryptocurrencyElectionServer: tt.fields.UnimplementedCryptocurrencyElectionServer,
			}
			_, err := s.DeleteById(tt.args.ctx, tt.args.in)

			if !reflect.DeepEqual(err.Error(), tt.want) {
				t.Errorf("CryptocurrencyElectionServer.DeleteById() = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_UpdateById_WhenSuccess(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	db.Create(cryptocurrencyMock)

	type args struct {
		ctx context.Context
		in  *pb.UpdateCryptocurrency
	}
	tests := []struct {
		BasicTestStruct
		args args
		want *pb.CryptocurrencyMessage
	}{
		{
			BasicTestStruct: BasicTestStruct{
				name:    "should update a cryptocurrency by id",
				fields:  Server{db: db},
				wantErr: false,
			},
			args: args{contextMock, &pb.UpdateCryptocurrency{Id: 1, Name: "name updated", Symbol: "sybup"}},
			want: &pb.CryptocurrencyMessage{
				Message: fmt.Sprintf("%d cryptocurrency has been updated", cryptocurrencyMock.Id),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptocurrencyElectionServer{
				db: tt.fields.db,
				UnimplementedCryptocurrencyElectionServer: tt.fields.UnimplementedCryptocurrencyElectionServer,
			}
			got, err := s.UpdateById(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptocurrencyElectionServer.UpdateById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CryptocurrencyElectionServer.UpdateById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_UpdateById_WhenBadRequest(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	db.Create(cryptocurrencyMock)

	type args struct {
		ctx context.Context
		in  *pb.UpdateCryptocurrency
	}
	tests := []struct {
		BasicTestStruct
		args args
		want string
	}{
		{
			BasicTestStruct: BasicTestStruct{
				name:    "should return bad request error when id is less than or equal to zero",
				fields:  Server{db: db},
				wantErr: false,
			},
			args: args{contextMock, &pb.UpdateCryptocurrency{Id: 0}},
			want: "id: cannot be blank.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptocurrencyElectionServer{
				db: tt.fields.db,
				UnimplementedCryptocurrencyElectionServer: tt.fields.UnimplementedCryptocurrencyElectionServer,
			}
			_, err := s.UpdateById(tt.args.ctx, tt.args.in)
			if !reflect.DeepEqual(err.Error(), tt.want) {
				t.Errorf("CryptocurrencyElectionServer.UpdateById() = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_UpdateById_WhenNotFoundEntity(t *testing.T) {
	db := initDBTest()
	defer closeDBTest(db)

	db.Create(cryptocurrencyMock)

	type args struct {
		ctx context.Context
		in  *pb.UpdateCryptocurrency
	}
	tests := []struct {
		BasicTestStruct
		args args
		want string
	}{
		{
			BasicTestStruct: BasicTestStruct{
				name:    "should return an error when entity not found",
				fields:  Server{db: db},
				wantErr: false,
			},
			args: args{contextMock, &pb.UpdateCryptocurrency{Id: 2}},
			want: "record not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptocurrencyElectionServer{
				db: tt.fields.db,
				UnimplementedCryptocurrencyElectionServer: tt.fields.UnimplementedCryptocurrencyElectionServer,
			}
			_, err := s.UpdateById(tt.args.ctx, tt.args.in)

			if !reflect.DeepEqual(err.Error(), tt.want) {
				t.Errorf("CryptocurrencyElectionServer.UpdateById() = %v, want %v", err, tt.want)
			}
		})
	}
}

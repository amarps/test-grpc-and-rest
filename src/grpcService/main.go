package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	pb "testGRPCandREST/src/grpcService/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedLoggerServer
}

var (
	logs *mongo.Collection
	port = flag.Int("p", 50051, "Listening port")
)

func quitOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}

func quitOnErrorf(err error, msg string, a ...interface{}) {
	if err != nil {
		log.Fatalf("%s: %s", fmt.Sprintf(msg, a...), err.Error())
	}
}

func init() {
	flag.Parse()

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	quitOnError(err, "Failed to connect to mongodb")
	err = client.Ping(context.TODO(), readpref.Primary())
	quitOnError(err, "Failed to connect to mongodb")

	logs = client.Database(os.Getenv("MONGO_DATABASE")).Collection("AmayTestLog")
	log.Println("Connected to mongoDB")
}

func main() {
	hostAddr := fmt.Sprintf("localhost:%d", *port)
	lis, err := net.Listen("tcp", hostAddr)
	quitOnErrorf(err, "Failed to establish connection at %s", hostAddr)

	s := grpc.NewServer()
	pb.RegisterLoggerServer(s, &server{})
	reflection.Register(s)
	log.Printf("server listneing at %v", lis.Addr())
	err = s.Serve(lis)
	quitOnErrorf(err, "Failed to serve on %v", lis.Addr())
}

func (s *server) Read(ctx context.Context, in *pb.ID) (*pb.Log, error) {
	if in.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "Argument ID can't be empty/null")
	}

	var res *pb.Log
	cur := logs.FindOne(ctx, bson.M{"id": in.Id})
	if cur != nil {
		cur.Decode(&res)
	}

	return res, nil
}

func (s *server) Write(ctx context.Context, in *pb.Log) (*pb.ID, error) {
	if in.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "Argument ID can't be empty/null")
	}

	_, err := logs.InsertOne(ctx, in)
	return &pb.ID{Id: in.Id}, err
}

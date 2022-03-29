package server

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCServer struct {
	pb_gen.UnimplementedGraduateDesignApiServer
}

func (s *GRPCServer) Ping(ctx context.Context, req *pb_gen.PingRequest) (*pb_gen.PingResponse, error) {
	return &pb_gen.PingResponse{Msg: "PONG"}, nil
}

func (s *GRPCServer) UserLogin(ctx context.Context, req *pb_gen.UserLoginRequest) (*pb_gen.UserLoginResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GRPCServer) CreateDir(ctx context.Context, req *pb_gen.CreateDirRequest) (*pb_gen.CreateDirResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GRPCServer) UploadFile(ctx context.Context, req *pb_gen.UploadFileRequest) (*pb_gen.UploadFileReponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GRPCServer) DownloadFile(ctx context.Context, req *pb_gen.DownloadFileRequest) (*pb_gen.DownloadFileResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GRPCServer) GetNode(ctx context.Context, req *pb_gen.GetNodeRequest) (*pb_gen.GetNodeResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GRPCServer) SearchFile(ctx context.Context, req *pb_gen.SearchFileRequest) (*pb_gen.SearchFileResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GRPCServer) mustEmbedUnimplementedGraduateDesignApiServer() {
	panic("not implemented") // TODO: Implement
}

func GRPCMain() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	grpc.WithTransportCredentials(insecure.NewCredentials())
	pb_gen.RegisterGraduateDesignApiServer(s, &GRPCServer{})
	httpServer := provideHTTP(":8888", s)
	if err := httpServer.Serve(lis); err != nil {
		panic(err)
	}
}

func provideHTTP(endpoint string, s *grpc.Server) *http.Server {
	ctx := context.Background()
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := pb_gen.RegisterGraduateDesignApiHandlerFromEndpoint(ctx, gwmux, endpoint, opts)
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	return &http.Server{
		Addr:    endpoint,
		Handler: grpcHandleFunc(s, mux),
	}
}

func grpcHandleFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

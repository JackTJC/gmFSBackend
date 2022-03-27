package server

import (
	"context"
	"net"

	"github.com/JackTJC/gmFS_backend/pb_gen"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb_gen.UnimplementedGraduateDesignApiServer
}

func (s *GRPCServer) UserLogin(_ context.Context, _ *pb_gen.UserLoginRequest) (*pb_gen.UserLoginResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GRPCServer) CreateDir(_ context.Context, _ *pb_gen.CreateDirRequest) (*pb_gen.CreateDirResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GRPCServer) UploadFile(_ context.Context, _ *pb_gen.UploadFileRequest) (*pb_gen.UploadFileReponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GRPCServer) DownloadFile(_ context.Context, _ *pb_gen.DownloadFileRequest) (*pb_gen.DownloadFileResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GRPCServer) GetNode(_ context.Context, _ *pb_gen.GetNodeRequest) (*pb_gen.GetNodeResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GRPCServer) SearchFile(_ context.Context, _ *pb_gen.SearchFileRequest) (*pb_gen.SearchFileResponse, error) {
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
	pb_gen.RegisterGraduateDesignApiServer(s, &GRPCServer{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

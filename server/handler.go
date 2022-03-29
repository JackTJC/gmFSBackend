package server

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/JackTJC/gmFS_backend/util"
)

type server struct {
	pb_gen.UnimplementedGraduateDesignApiServer
}

func (s *server) Ping(ctx context.Context, req *pb_gen.PingRequest) (*pb_gen.PingResponse, error) {
	defer util.RecoverFromPanic()
	return &pb_gen.PingResponse{Msg: "PONG" + req.GetName()}, nil
}

func (s *server) UserLogin(ctx context.Context, req *pb_gen.UserLoginRequest) (*pb_gen.UserLoginResponse, error) {
	defer util.RecoverFromPanic()
	panic("not implemented") // TODO: Implement
}

func (s *server) CreateDir(ctx context.Context, req *pb_gen.CreateDirRequest) (*pb_gen.CreateDirResponse, error) {
	defer util.RecoverFromPanic()
	panic("not implemented") // TODO: Implement
}

func (s *server) UploadFile(ctx context.Context, req *pb_gen.UploadFileRequest) (*pb_gen.UploadFileReponse, error) {
	defer util.RecoverFromPanic()
	panic("not implemented") // TODO: Implement
}

func (s *server) DownloadFile(ctx context.Context, req *pb_gen.DownloadFileRequest) (*pb_gen.DownloadFileResponse, error) {
	defer util.RecoverFromPanic()
	panic("not implemented") // TODO: Implement
}

func (s *server) GetNode(ctx context.Context, req *pb_gen.GetNodeRequest) (*pb_gen.GetNodeResponse, error) {
	defer util.RecoverFromPanic()
	panic("not implemented") // TODO: Implement
}

func (s *server) SearchFile(ctx context.Context, req *pb_gen.SearchFileRequest) (*pb_gen.SearchFileResponse, error) {
	defer util.RecoverFromPanic()
	panic("not implemented") // TODO: Implement
}

func (s *server) mustEmbedUnimplementedGraduateDesignApiServer() {
	panic("not implemented") // TODO: Implement
}

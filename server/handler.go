package server

import (
	"context"

	"github.com/JackTJC/gmFS_backend/logs"
	"github.com/JackTJC/gmFS_backend/method"
	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/JackTJC/gmFS_backend/util"
)

type server struct {
	pb_gen.UnimplementedGraduateDesignApiServer
}

func (s *server) UserRegister(ctx context.Context, req *pb_gen.UserRegisterRequest) (resp *pb_gen.UserRegisterResponse, err error) {
	defer util.RecoverFromPanic()
	defer func() {
		logs.Sugar.Infof("request = %+v", req)
		logs.Sugar.Infof("response = %+v", resp)
	}()
	h := method.NewUserRegisterHandler(ctx, req)
	h.Run()
	return h.Resp, nil

}

func (s *server) Ping(ctx context.Context, req *pb_gen.PingRequest) (resp *pb_gen.PingResponse, err error) {
	defer util.RecoverFromPanic()
	defer func() {
		logs.Sugar.Infof("request = %+v", req)
		logs.Sugar.Infof("response = %+v", resp)
	}()
	return &pb_gen.PingResponse{Msg: "PONG" + req.GetName()}, nil
}

func (s *server) UserLogin(ctx context.Context, req *pb_gen.UserLoginRequest) (resp *pb_gen.UserLoginResponse, err error) {
	defer util.RecoverFromPanic()
	defer func() {
		logs.Sugar.Infof("request = %+v", req)
		logs.Sugar.Infof("response = %+v", resp)
	}()
	h := method.NewUserLoginHandler(ctx, req)
	h.Run()
	return h.Resp, nil
}

func (s *server) CreateDir(ctx context.Context, req *pb_gen.CreateDirRequest) (resp *pb_gen.CreateDirResponse, err error) {
	defer util.RecoverFromPanic()
	defer func() {
		logs.Sugar.Infof("request = %+v", req)
		logs.Sugar.Infof("response = %+v", resp)
	}()
	h := method.NewCreateDirHandler(ctx, req)
	h.Run()
	return h.Resp, nil
}

func (s *server) UploadFile(ctx context.Context, req *pb_gen.UploadFileRequest) (resp *pb_gen.UploadFileReponse, err error) {
	defer util.RecoverFromPanic()
	defer func() {
		logs.Sugar.Infof("request = %+v", req)
		logs.Sugar.Infof("response = %+v", resp)
	}()
	h := method.NewUploadFileHandler(ctx, req)
	h.Run()
	return h.Resp, nil
}

func (s *server) DownloadFile(ctx context.Context, req *pb_gen.DownloadFileRequest) (resp *pb_gen.DownloadFileResponse, err error) {
	defer util.RecoverFromPanic()
	defer func() {
		logs.Sugar.Infof("request = %+v", req)
		logs.Sugar.Infof("response = %+v", resp)
	}()
	h := method.NewDownloadFileHandler(ctx, req)
	h.Run()
	return h.Resp, nil
}

func (s *server) GetNode(ctx context.Context, req *pb_gen.GetNodeRequest) (resp *pb_gen.GetNodeResponse, err error) {
	defer util.RecoverFromPanic()
	defer func() {
		logs.Sugar.Infof("request = %+v", req)
		logs.Sugar.Infof("response = %+v", resp)
	}()
	h := method.NewGetNodeHandler(ctx, req)
	h.Run()
	return h.Resp, nil
}

func (s *server) SearchFile(ctx context.Context, req *pb_gen.SearchFileRequest) (resp *pb_gen.SearchFileResponse, err error) {
	defer util.RecoverFromPanic()
	defer func() {
		logs.Sugar.Infof("request = %+v", req)
		logs.Sugar.Infof("response = %+v", resp)
	}()
	h := method.NewSearchFileHandler(ctx, req)
	h.Run()
	return h.Resp, nil
}

func (s *server) mustEmbedUnimplementedGraduateDesignApiServer() {
	panic("not implemented") // TODO: Implement
}

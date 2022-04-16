package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/JackTJC/gmFS_backend/util"
)

type GetUserInfoHandler struct {
	ctx context.Context
	Req *pb_gen.GetUserInfoRequest
}

func NewGetUserInfoHandler(ctx context.Context, req *pb_gen.GetUserInfoRequest) *GetUserInfoHandler {
	return &GetUserInfoHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *GetUserInfoHandler) Run() (resp *pb_gen.GetUserInfoResponse) {
	resp = &pb_gen.GetUserInfoResponse{
		BaseResp: util.BuildBaseResp(pb_gen.StatusCode_Success),
	}
	return
}

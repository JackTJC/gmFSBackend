package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/JackTJC/gmFS_backend/util"
)

type CreateDirHandler struct {
	ctx context.Context
	Req *pb_gen.CreateDirRequest
}

func NewCreateDirHandler(ctx context.Context, req *pb_gen.CreateDirRequest) *CreateDirHandler {
	return &CreateDirHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *CreateDirHandler) Run() (resp *pb_gen.CreateDirResponse) {
	resp = &pb_gen.CreateDirResponse{
		BaseResp: util.BuildBaseResp(pb_gen.StatusCode_Success),
	}
	return
}

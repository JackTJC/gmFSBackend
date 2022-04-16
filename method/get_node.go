package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/JackTJC/gmFS_backend/util"
)

type GetNodeHandler struct {
	ctx context.Context
	Req *pb_gen.GetNodeRequest
}

func NewGetNodeHandler(ctx context.Context, req *pb_gen.GetNodeRequest) *GetNodeHandler {
	return &GetNodeHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *GetNodeHandler) Run() (resp *pb_gen.GetNodeResponse) {
	resp = &pb_gen.GetNodeResponse{
		BaseResp: util.BuildBaseResp(pb_gen.StatusCode_Success),
	}
	return resp
}

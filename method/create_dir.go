package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
)

type CreateDirHandler struct {
	ctx context.Context
	Req *pb_gen.CreateDirRequest

	Resp *pb_gen.CreateDirResponse
}

func NewCreateDirHandler(ctx context.Context, req *pb_gen.CreateDirRequest) *CreateDirHandler {
	return &CreateDirHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *CreateDirHandler) Run() (resp *pb_gen.CreateDirResponse) {
	resp = &pb_gen.CreateDirResponse{}
	return
}

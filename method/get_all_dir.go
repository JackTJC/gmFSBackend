package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
)

type GetAllDirHandler struct {
	ctx context.Context
	Req *pb_gen.GetAllDirRequest
}

func NewGetAllDirHandler(ctx context.Context, req *pb_gen.GetAllDirRequest) *GetAllDirHandler {
	return &GetAllDirHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *GetAllDirHandler) Run() (resp *pb_gen.GetAllDirResponse) {
	return nil
}

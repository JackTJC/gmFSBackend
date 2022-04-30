package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
)

type RegisterFileHandler struct {
	ctx context.Context
	Req *pb_gen.RegisterFileRequest
}

func NewRegisterFileHandler(ctx context.Context, req *pb_gen.RegisterFileRequest) *RegisterFileHandler {
	return &RegisterFileHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *RegisterFileHandler) Run() (resp *pb_gen.RegisterFileResponse) {
	return nil
}

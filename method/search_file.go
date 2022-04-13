package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
)

type SearchFileHandler struct {
	ctx context.Context
	Req *pb_gen.SearchFileRequest
}

func NewSearchFileHandler(ctx context.Context, req *pb_gen.SearchFileRequest) *SearchFileHandler {
	return &SearchFileHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *SearchFileHandler) Run() (resp *pb_gen.SearchFileResponse) {
	resp = &pb_gen.SearchFileResponse{}
	return resp
}

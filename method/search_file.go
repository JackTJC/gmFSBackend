package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/JackTJC/gmFS_backend/util"
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
	resp = &pb_gen.SearchFileResponse{
		BaseResp: util.BuildBaseResp(pb_gen.StatusCode_Success),
	}
	return resp
}

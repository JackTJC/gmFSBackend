package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
)

type DownloadFileHandler struct {
	ctx  context.Context
	Req  *pb_gen.DownloadFileRequest
	Resp *pb_gen.DownloadFileResponse
}

func NewDownloadFileHandler(ctx context.Context, req *pb_gen.DownloadFileRequest) *DownloadFileHandler {
	return &DownloadFileHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *DownloadFileHandler) Run() {

}

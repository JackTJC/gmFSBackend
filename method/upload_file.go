package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
)

type UploadFileHandler struct {
	ctx  context.Context
	Req  *pb_gen.UploadFileRequest
	Resp *pb_gen.UploadFileReponse
}

func NewUploadFileHandler(ctx context.Context, req *pb_gen.UploadFileRequest) *UploadFileHandler {
	return &UploadFileHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *UploadFileHandler) Run() {

}

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

func (h *UploadFileHandler) Run() {

}

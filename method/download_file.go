package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
)

type DownloadFileHandler struct {
	ctx  context.Context
	Req  *pb_gen.DownloadFileRequest
	Resp *pb_gen.DownloadFileRequest
}

func (h *DownloadFileHandler) Run() {

}

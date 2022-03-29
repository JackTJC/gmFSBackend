package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
)

type SearchFileHandler struct {
	ctx  context.Context
	Req  *pb_gen.SearchFileRequest
	Resp *pb_gen.SearchFileResponse
}

func (h *SearchFileHandler) Run() {

}

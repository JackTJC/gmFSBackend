package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
)

type GetNodeHandler struct {
	ctx  context.Context
	Req  *pb_gen.GetNodeRequest
	Resp *pb_gen.GetNodeResponse
}

func (h *GetNodeHandler) Run() {

}

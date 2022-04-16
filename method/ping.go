package method

import (
	"context"
	"fmt"

	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/JackTJC/gmFS_backend/util"
)

type PingHandler struct {
	ctx context.Context
	Req *pb_gen.PingRequest
}

func NewPingHandler(ctx context.Context, req *pb_gen.PingRequest) *PingHandler {
	return &PingHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *PingHandler) Run() (resp *pb_gen.PingResponse) {
	resp = &pb_gen.PingResponse{
		BaseResp: util.BuildBaseResp(pb_gen.StatusCode_Success),
	}
	resp.Msg = fmt.Sprintf("PONG %v", h.Req.GetName())
	return
}

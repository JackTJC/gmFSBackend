package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
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
	resp = &pb_gen.PingResponse{}
	return
}

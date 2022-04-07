package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
)

type UserRegisterHandler struct {
	ctx context.Context

	Req  *pb_gen.UserRegisterRequest
	Resp *pb_gen.UserRegisterResponse
}

func NewUserRegisterHandler(ctx context.Context, req *pb_gen.UserRegisterRequest) *UserRegisterHandler {
	return &UserRegisterHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *UserRegisterHandler) Run() {

}

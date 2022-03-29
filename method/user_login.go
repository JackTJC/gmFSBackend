package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
)

type UserLoginHandler struct {
	ctx  context.Context
	Req  *pb_gen.UserLoginRequest
	Resp *pb_gen.UserLoginResponse
}

func NewUserLoginHandler(ctx context.Context, req *pb_gen.UserLoginRequest) *UserLoginHandler {
	return &UserLoginHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *UserLoginHandler) Run() {

}

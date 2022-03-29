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

func (h *UserLoginHandler) Run() {

}

package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/pb_gen"
)

type CreateDirHandler struct {
	ctx context.Context
	Req *pb_gen.CreateDirRequest

	Resp *pb_gen.CreateDirResponse
}

func (h *CreateDirHandler) Run() {

}

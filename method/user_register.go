package method

import (
	"context"
	"errors"

	"github.com/JackTJC/gmFS_backend/dal/db"
	"github.com/JackTJC/gmFS_backend/model"
	"github.com/JackTJC/gmFS_backend/pb_gen"
)

type UserRegisterHandler struct {
	ctx context.Context

	Req  *pb_gen.UserRegisterRequest
	Resp *pb_gen.UserRegisterResponse
}

func NewUserRegisterHandler(ctx context.Context, req *pb_gen.UserRegisterRequest) *UserRegisterHandler {
	return &UserRegisterHandler{
		ctx:  ctx,
		Req:  req,
		Resp: &pb_gen.UserRegisterResponse{},
	}
}

func (h *UserRegisterHandler) Run() {
	if err := h.checkParams(); err != nil {
		return
	}
	user := &model.UserInfo{
		UserName: h.Req.GetUserName(),
		Password: h.Req.GetPassword(),
	}
	err := db.User.CreateUser(user)
	if err != nil {
		h.Resp.BaseResp.StatusCode = -1
		h.Resp.BaseResp.Message = "db失败"
		return
	}
	h.Resp.BaseResp = &pb_gen.BaseResp{
		StatusCode: 0,
		Message:    "success",
	}
}

func (h *UserRegisterHandler) checkParams() error {
	if h.Req.GetUserName() == "" || h.Req.GetPassword() == "" {
		h.Resp.BaseResp.Message = "参数非法"
		return errors.New("参数非法")
	}
	return nil
}

package method

import (
	"context"
	"errors"

	"github.com/JackTJC/gmFS_backend/dal/db"
	"github.com/JackTJC/gmFS_backend/logs"
	"github.com/JackTJC/gmFS_backend/model"
	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/JackTJC/gmFS_backend/util"
)

type UserRegisterHandler struct {
	ctx context.Context
	Req *pb_gen.UserRegisterRequest
}

func NewUserRegisterHandler(ctx context.Context, req *pb_gen.UserRegisterRequest) *UserRegisterHandler {
	return &UserRegisterHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *UserRegisterHandler) Run() (resp *pb_gen.UserRegisterResponse) {
	resp = &pb_gen.UserRegisterResponse{}
	if err := h.checkParams(); err != nil {
		return
	}
	user := &model.UserInfo{
		UserName: h.Req.GetUserName(),
		Password: h.Req.GetPassword(),
	}
	err := db.User.CreateUser(user)
	if err != nil {
		if err == db.ErrUserExist {
			logs.Sugar.Errorf("user:%v have exist", h.Req.GetUserName())
			resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_UserExist)
			return
		}
		logs.Sugar.Errorf("create user error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	// TODO 为用户创建根文件夹
	resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_Success)
	return
}

func (h *UserRegisterHandler) checkParams() error {
	if h.Req.GetUserName() == "" || h.Req.GetPassword() == "" {
		return errors.New("参数非法")
	}
	return nil
}

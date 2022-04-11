package method

import (
	"context"
	"errors"

	"github.com/JackTJC/gmFS_backend/dal/db"
	"github.com/JackTJC/gmFS_backend/logs"
	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/JackTJC/gmFS_backend/util"
)

type UserLoginHandler struct {
	ctx  context.Context
	Req  *pb_gen.UserLoginRequest
	Resp *pb_gen.UserLoginResponse
}

func NewUserLoginHandler(ctx context.Context, req *pb_gen.UserLoginRequest) *UserLoginHandler {
	return &UserLoginHandler{
		ctx:  ctx,
		Req:  req,
		Resp: &pb_gen.UserLoginResponse{},
	}
}

func (h *UserLoginHandler) Run() {
	if err := h.checkParams(); err != nil {
		logs.Sugar.Errorf("checkParams error:%v", err)
		h.Resp.BaseResp = &pb_gen.BaseResp{
			StatusCode: pb_gen.StatusCode_CommonErr,
		}
		return
	}
	user, err := db.User.GetUserByName(h.Req.GetUserName())
	if err != nil {
		logs.Sugar.Errorf("GetUserByName error:%v", err)
		if err == db.ErrUserNotFound {
			h.Resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_UserNotFound)
			return
		}
		h.Resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	if user.Password != h.Req.GetPassword() {
		logs.Sugar.Errorf("auth error, wrong passwd")
		h.Resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_WrongPasswd)
		return
	}
	h.Resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_Success)
}

func (h *UserLoginHandler) checkParams() error {
	if len(h.Req.GetUserName()) == 0 {
		return errors.New("empty user name")
	}
	if len(h.Req.GetPassword()) == 0 {
		return errors.New("empty passwd")
	}
	return nil
}

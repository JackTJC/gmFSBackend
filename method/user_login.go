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
	ctx context.Context
	Req *pb_gen.UserLoginRequest
}

func NewUserLoginHandler(ctx context.Context, req *pb_gen.UserLoginRequest) *UserLoginHandler {
	return &UserLoginHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *UserLoginHandler) Run() (resp *pb_gen.UserLoginResponse) {
	resp = &pb_gen.UserLoginResponse{
		BaseResp: util.BuildBaseResp(pb_gen.StatusCode_Success),
	}
	if err := h.checkParams(); err != nil {
		logs.Sugar.Errorf("checkParams error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	user, err := db.User.GetByName(h.ctx, h.Req.GetUserName())
	if err != nil {
		logs.Sugar.Errorf("GetUserByName error:%v", err)
		if err == db.ErrUserNotFound {
			resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_UserNotFound)
			return
		}
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	if user.Password != h.Req.GetPassword() {
		logs.Sugar.Errorf("auth error, wrong passwd")
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_WrongPasswd)
		return
	}
	resp.UserInfo = &pb_gen.UserInfo{
		RootId:   int64(user.RootNodeID),
		UserName: user.UserName,
		Email:    user.Email,
	}
	token := util.GenUUID() // 暂时这样
	resp.Token = token
	return
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

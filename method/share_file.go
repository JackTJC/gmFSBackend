package method

import (
	"context"
	"errors"

	"github.com/JackTJC/gmFS_backend/dal/cache"
	"github.com/JackTJC/gmFS_backend/dal/db"
	"github.com/JackTJC/gmFS_backend/logs"
	"github.com/JackTJC/gmFS_backend/model"
	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/JackTJC/gmFS_backend/util"
)

type ShareFileHandler struct {
	ctx context.Context
	Req *pb_gen.ShareFileRequest
	uid uint64
}

func NewShareFileHandler(ctx context.Context, req *pb_gen.ShareFileRequest) *ShareFileHandler {
	return &ShareFileHandler{
		ctx: ctx,
		Req: req,
	}
}
func (h *ShareFileHandler) Run() (resp *pb_gen.ShareFileResponse) {
	defer func() {
		logs.Sugar.Infof("req = %+v, resp = %+v", h.Req, resp)
	}()
	resp = &pb_gen.ShareFileResponse{}
	if err := h.checkParams(); err != nil {
		logs.Sugar.Errorf("check params error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return nil
	}
	user, err := db.User.GetByName(h.ctx, h.Req.GetUserName())
	if err != nil {
		if err == db.ErrUserNotFound {
			logs.Sugar.Errorf("user not found")
			resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_UserNotFound)
			return
		}
		logs.Sugar.Errorf("get user by name error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return nil
	}
	shareFile := &model.ShareFile{
		FileID: uint64(h.Req.GetFileId()),
		From:   h.uid,
		To:     user.UserID,
		Key:    string(h.Req.GetKey()),
		Status: int(db.NotProcess),
	}
	if err := db.ShareFile.Create(h.ctx, shareFile); err != nil {
		logs.Sugar.Errorf("create share file in db error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	return
}

func (h *ShareFileHandler) checkParams() error {
	if h.Req.GetFileId() <= 0 {
		return errors.New("illegal file id")
	}
	if h.Req.GetUserName() == "" {
		return errors.New("illegal user name")
	}
	uid, err := cache.Token.GetUID(h.Req.GetBaseReq().GetToken())
	if err != nil {
		return err
	}
	h.uid = uid
	return nil
}

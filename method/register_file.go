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

type RegisterFileHandler struct {
	ctx context.Context
	uid uint64
	Req *pb_gen.RegisterFileRequest
}

func NewRegisterFileHandler(ctx context.Context, req *pb_gen.RegisterFileRequest) *RegisterFileHandler {
	return &RegisterFileHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *RegisterFileHandler) Run() (resp *pb_gen.RegisterFileResponse) {
	resp = &pb_gen.RegisterFileResponse{}
	if err := h.checkParams(); err != nil {
		logs.Sugar.Errorf("register file check param error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	err := db.Transaction(h.ctx, func(ctx context.Context) error {
		// 保存文件夹到文件的引用关系
		relation := &model.NodeRel{
			ParentID: uint64(h.Req.GetDirId()),
			ChildID:  uint64(h.Req.GetFileId()),
		}
		if err := db.NodeRel.Create(h.ctx, relation); err != nil {
			logs.Sugar.Errorf("register file save db error:%v", err)
			resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
			return err
		}
		// 保存密钥
		sk := &model.SecretKey{
			UserID: h.uid,
			FileID: uint64(h.Req.GetFileId()),
			Key:    string(h.Req.GetSecretKey()),
		}
		if err := db.SecretKey.Create(ctx, sk); err != nil {
			logs.Sugar.Errorf("register file save key error:%v", err)
			resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
			return err
		}
		return nil
	})
	if err != nil {
		logs.Sugar.Errorf("transaction error:%v", err)
		if resp.GetBaseResp().GetStatusCode() == pb_gen.StatusCode_Success {
			resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		}
		return
	}

	return
}

func (h *RegisterFileHandler) checkParams() error {
	if h.Req.GetDirId() <= 0 {
		return errors.New("illegal dir id")
	}
	if h.Req.GetFileId() <= 0 {
		return errors.New("illegal file id")
	}
	uid, err := cache.Token.GetUID(h.Req.GetBaseReq().GetToken())
	if err != nil {
		return err
	}
	h.uid = uid
	return nil
}

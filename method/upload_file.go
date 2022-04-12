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

type UploadFileHandler struct {
	ctx  context.Context
	Req  *pb_gen.UploadFileRequest
	Resp *pb_gen.UploadFileReponse
}

func NewUploadFileHandler(ctx context.Context, req *pb_gen.UploadFileRequest) *UploadFileHandler {
	return &UploadFileHandler{
		ctx:  ctx,
		Req:  req,
		Resp: &pb_gen.UploadFileReponse{},
	}
}

func (h *UploadFileHandler) Run() {
	if err := h.checkParams(); err != nil {
		logs.Sugar.Errorf("checkParams error:%v", err)
		h.Resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	node := &model.Node{
		NodeType: uint(pb_gen.NodeType_File),
		Name:     h.Req.GetFileName(),
		Content:  string(h.Req.GetContent()),
	}
	if err := db.Node.Create(node); err != nil {
		logs.Sugar.Errorf("CreateFile error:%v", err)
		h.Resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
}

func (h *UploadFileHandler) checkParams() error {
	if len(h.Req.GetFileName()) == 0 {
		return errors.New("empty file name")
	}
	if len(h.Req.GetContent()) == 0 {
		return errors.New("empty file content")
	}
	return nil
}

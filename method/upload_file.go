package method

import (
	"context"
	"errors"
	"fmt"

	"github.com/JackTJC/gmFS_backend/dal/db"
	objstore "github.com/JackTJC/gmFS_backend/dal/obj_store"
	"github.com/JackTJC/gmFS_backend/logs"
	"github.com/JackTJC/gmFS_backend/model"
	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/JackTJC/gmFS_backend/util"
)

type UploadFileHandler struct {
	ctx context.Context
	Req *pb_gen.UploadFileRequest
}

func NewUploadFileHandler(ctx context.Context, req *pb_gen.UploadFileRequest) *UploadFileHandler {
	return &UploadFileHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *UploadFileHandler) Run() (resp *pb_gen.UploadFileReponse) {
	resp = &pb_gen.UploadFileReponse{
		BaseResp: util.BuildBaseResp(pb_gen.StatusCode_Success),
	}
	if err := h.checkParams(); err != nil {
		logs.Sugar.Errorf("checkParams error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	nodeID := uint64(util.GenId())
	if err := objstore.UploadFile(h.ctx, GenCosFileKey(int64(nodeID)), h.Req.Content); err != nil {
		logs.Sugar.Errorf("upload file cos upload error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	err := db.Transaction(h.ctx, func(ctx context.Context) error {
		nodeRel := &model.NodeRel{
			ParentID: uint64(h.Req.GetParentId()),
			ChildID:  nodeID,
		}
		if err := db.NodeRel.Create(ctx, nodeRel); err != nil {
			logs.Sugar.Errorf("CreateNodeRel error:%v", err)
			resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
			return err
		}
		node := &model.Node{
			NodeID:   nodeID,
			NodeType: uint(pb_gen.NodeType_File),
			Name:     h.Req.GetFileName(),
			Content:  string(h.Req.GetContent()),
		}
		if err := db.Node.Create(ctx, node); err != nil {
			logs.Sugar.Errorf("Create Node error:%v", err)
			resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
			return err
		}
		return nil
	})
	if err != nil {
		logs.Sugar.Errorf("Transaction error:%v", err)
		if resp.GetBaseResp().GetStatusCode() == pb_gen.StatusCode_Success {
			resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		}
		return
	}
	return
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

func GenCosFileKey(fileId int64) string {
	return fmt.Sprintf("file/%v", fileId)
}

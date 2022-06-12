package method

import (
	"context"

	"github.com/JackTJC/gmFS_backend/dal/cache"
	"github.com/JackTJC/gmFS_backend/dal/db"
	"github.com/JackTJC/gmFS_backend/logs"
	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/JackTJC/gmFS_backend/util"
)

type GetRecvFileHandler struct {
	uid uint64
	ctx context.Context
	Req *pb_gen.GetRecvFileRequest
}

func NewGetRecvFileHandler(ctx context.Context, req *pb_gen.GetRecvFileRequest) *GetRecvFileHandler {
	return &GetRecvFileHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *GetRecvFileHandler) Run() (resp *pb_gen.GetRecvFileResponse) {
	defer func() {
		if resp.GetBaseResp().GetStatusCode() == pb_gen.StatusCode_Success {
			resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_Success)
		}
		logs.Sugar.Infof("req = %+v, resp = %+v", h.Req, resp)
	}()
	resp = &pb_gen.GetRecvFileResponse{}
	if err := h.checkParams(); err != nil {
		logs.Sugar.Errorf("check params error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	// 查询未处理的分享
	shareFileList, err := db.ShareFile.GetByDstUID(h.ctx, h.uid)
	if err != nil {
		if err == db.ErrEmptyShareFile {
			logs.Sugar.Errorf("user:%v have no share file", h.uid)
			resp.ShareFileList = []*pb_gen.SharedFile{}
			return
		}
		logs.Sugar.Errorf("get share file list, error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	// 查文件
	var fileIDList []uint64
	for _, f := range shareFileList {
		fileIDList = append(fileIDList, f.FileID)
	}
	fileMap, err := db.Node.MGetByNodeId(h.ctx, fileIDList)
	if err != nil {
		logs.Sugar.Errorf("get file error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	// 查user
	var userIDList []uint64
	for _, f := range shareFileList {
		userIDList = append(userIDList, f.From)
	}
	userMap, err := db.User.MGetByID(h.ctx, userIDList)
	if err != nil {
		logs.Sugar.Errorf("get user error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}

	// 写resp
	for _, shareFile := range shareFileList {
		resp.ShareFileList = append(resp.ShareFileList, &pb_gen.SharedFile{
			ShareId:  int64(shareFile.ID),
			FileId:   int64(shareFile.FileID),
			Key:      []byte(shareFile.Key),
			From:     userMap[shareFile.From].UserName,
			FileName: fileMap[shareFile.FileID].Name,
		})
	}
	return
}

func (h *GetRecvFileHandler) checkParams() error {
	uid, err := cache.Token.GetUID(h.Req.GetBaseReq().GetToken())
	if err != nil {
		return err
	}
	h.uid = uid
	return nil
}

package method

import (
	"context"
	"errors"

	"github.com/JackTJC/gmFS_backend/dal/cache"
	"github.com/JackTJC/gmFS_backend/dal/db"
	"github.com/JackTJC/gmFS_backend/logs"
	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/JackTJC/gmFS_backend/util"
)

type SearchFileHandler struct {
	ctx context.Context
	uid uint64
	Req *pb_gen.SearchFileRequest
}

func NewSearchFileHandler(ctx context.Context, req *pb_gen.SearchFileRequest) *SearchFileHandler {
	return &SearchFileHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *SearchFileHandler) Run() (resp *pb_gen.SearchFileResponse) {
	resp = &pb_gen.SearchFileResponse{
		BaseResp: util.BuildBaseResp(pb_gen.StatusCode_Success),
	}
	if err := h.checkParams(); err != nil {
		logs.Sugar.Errorf("search file check params error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	// 获取该用户的所有file id
	sks, err := db.SecretKey.GetByUID(h.ctx, h.uid)
	if err != nil {
		logs.Sugar.Errorf("search file get sks error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return nil
	}
	var fileIDList []uint64
	for _, sk := range sks {
		fileIDList = append(fileIDList, sk.FileID)
	}
	if len(fileIDList) == 0 {
		return
	}
	// 搜索该关键词
	searchIdxs, err := db.SearchIndex.SearchByKwAndFileID(h.ctx, h.Req.GetKeyword(), fileIDList)
	if err != nil {
		logs.Sugar.Errorf("search file search error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return nil
	}
	// 拉取node
	var searchResultIDList []uint64
	for _, idx := range searchIdxs {
		searchResultIDList = append(searchResultIDList, idx.FileID)
	}
	nodeMap, err := db.Node.MGetByNodeId(h.ctx, searchResultIDList)
	// 下发数据
	for _, node := range nodeMap {
		resp.NodeList = append(resp.NodeList, &pb_gen.Node{
			NodeType:   pb_gen.NodeType_File,
			NodeId:     int64(node.NodeID),
			NodeName:   node.Name,
			CreateTime: node.CreateTime.Unix(),
			UpdateTime: node.UpdateTime.Unix(),
		})
	}
	return resp
}

func (h *SearchFileHandler) checkParams() error {
	if h.Req.GetKeyword() == "" {
		return errors.New("empty keyword")
	}

	uid, err := cache.Token.GetUID(h.Req.GetBaseReq().GetToken())
	if err != nil {
		return err
	}
	h.uid = uid
	return nil
}

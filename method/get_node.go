package method

import (
	"context"
	"errors"

	"github.com/JackTJC/gmFS_backend/dal/db"
	"github.com/JackTJC/gmFS_backend/logs"
	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/JackTJC/gmFS_backend/util"
)

type GetNodeHandler struct {
	ctx context.Context
	Req *pb_gen.GetNodeRequest
}

func NewGetNodeHandler(ctx context.Context, req *pb_gen.GetNodeRequest) *GetNodeHandler {
	return &GetNodeHandler{
		ctx: ctx,
		Req: req,
	}
}

func (h *GetNodeHandler) Run() (resp *pb_gen.GetNodeResponse) {
	resp = &pb_gen.GetNodeResponse{
		BaseResp: util.BuildBaseResp(pb_gen.StatusCode_Success),
	}
	if err := h.checkParams(); err != nil {
		logs.Sugar.Errorf("check params error:%v", err)
		util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	// 获取当前节点
	nodeMap, err := db.Node.MGetByNodeId(h.ctx, []int64{h.Req.GetNodeId()})
	if err != nil {
		logs.Sugar.Errorf("MGetByNodeId error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	node, ok := nodeMap[h.Req.GetNodeId()]
	if !ok {
		logs.Sugar.Warnf("node id:%v not exist", h.Req.GetNodeId())
		return
	}
	// 打包节点元数据
	resp.Node = &pb_gen.Node{
		NodeId:      int64(node.NodeID),
		NodeName:    node.Name,
		NodeContent: []byte(node.Content),
		CreateTime:  node.CreateTime.Unix(),
		UpdateTime:  node.UpdateTime.Unix(),
	}
	if node.NodeType == uint(pb_gen.NodeType_File) {
		return
	}
	// 获取子节点
	subNodeList, err := db.NodeRel.GetByParent(h.ctx, h.Req.GetNodeId())
	if err != nil {
		logs.Sugar.Errorf("GetByParent error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	subNodeIDs := make([]int64, 0, len(subNodeList))
	for _, nodeRel := range subNodeList {
		subNodeIDs = append(subNodeIDs, int64(nodeRel.ChildID))
	}
	nodeMap, err = db.Node.MGetByNodeId(h.ctx, subNodeIDs)
	if err != nil {
		logs.Sugar.Errorf("MGetByNodeId error:%v", err)
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_CommonErr)
		return
	}
	// 打包子节点数据
	pbSubNodes := make([]*pb_gen.Node, 0, len(nodeMap))
	for _, node := range nodeMap {
		pbSubNodes = append(pbSubNodes, &pb_gen.Node{
			NodeId:     int64(node.NodeID),
			NodeName:   node.Name,
			NodeType:   pb_gen.NodeType(node.NodeType),
			CreateTime: node.CreateTime.Unix(),
			UpdateTime: node.UpdateTime.Unix(),
		})
	}
	resp.SubNodes = pbSubNodes
	resp.Node.SubNodeList = pbSubNodes
	return
}

func (h *GetNodeHandler) checkParams() error {
	if h.Req.GetNodeId() <= 0 {
		return errors.New("illegal node id")
	}
	return nil
}

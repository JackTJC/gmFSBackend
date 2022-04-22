package http_server

import (
	"net/http"

	"github.com/JackTJC/gmFS_backend/logs"
	"github.com/JackTJC/gmFS_backend/method"
	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/JackTJC/gmFS_backend/util"
	"github.com/gin-gonic/gin"
)

// 设置路由
func setRoute(r *gin.Engine) {
	r.Use(recognizeContentType)
	r.Use(writeContentType)
	r.POST("/ping", ping)
	r.POST("/user/login", userLogin)
	r.POST("/user/register", userRegister)
	r.POST("/dir/create", createDir)
	r.POST("/file/upload", uploadFile)
	r.POST("/file/download", downloadFile)
	r.POST("/file/search", searchFile)
	r.POST("/node/get", getNode)
	r.POST("/user/info/get", getUserInfo)
}

func recognizeContentType(c *gin.Context) {
	c.Keys = make(map[string]interface{})
	var format reqFormat
	if c.ContentType() == pbContentType {
		format = pbReqFormat
	} else {
		format = jsonReqFormat
	}
	c.Keys[formatKey] = format
	c.Next()
}

func writeContentType(c *gin.Context) {
	format, ok := c.Keys[formatKey].(reqFormat)
	if !ok {
		c.Next()
		return
	}
	switch format {
	case jsonReqFormat:
		c.Header("Content-Type", jsonContentType)
	case pbReqFormat:
		c.Header("Content-Type", pbContentType)
	default:
		c.Header("Content-Type", jsonContentType)
	}
	c.Next()
}

// template
func ping(c *gin.Context) {
	req := &pb_gen.PingRequest{}
	if err := readBody(c, req); err != nil { // 从http request body解析初请求
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewPingHandler(c.Request.Context(), req).Run() // 业务逻辑
	if err := writeBody(c, resp); err != nil {                    // 响应写回http response
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if resp.GetBaseResp().GetStatusCode() == pb_gen.StatusCode_Success {
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_Success)
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}

func userLogin(c *gin.Context) {
	req := &pb_gen.UserLoginRequest{}
	if err := readBody(c, req); err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewUserLoginHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if resp.GetBaseResp().GetStatusCode() == pb_gen.StatusCode_Success {
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_Success)
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}
func userRegister(c *gin.Context) {
	req := &pb_gen.UserRegisterRequest{}
	if err := readBody(c, req); err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewUserRegisterHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if resp.GetBaseResp().GetStatusCode() == pb_gen.StatusCode_Success {
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_Success)
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}
func createDir(c *gin.Context) {
	req := &pb_gen.CreateDirRequest{}
	if err := readBody(c, req); err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewCreateDirHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if resp.GetBaseResp().GetStatusCode() == pb_gen.StatusCode_Success {
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_Success)
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}
func getUserInfo(c *gin.Context) {
	req := &pb_gen.GetUserInfoRequest{}
	if err := readBody(c, req); err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewGetUserInfoHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if resp.GetBaseResp().GetStatusCode() == pb_gen.StatusCode_Success {
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_Success)
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}
func uploadFile(c *gin.Context) {
	req := &pb_gen.UploadFileRequest{}
	if err := readBody(c, req); err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewUploadFileHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if resp.GetBaseResp().GetStatusCode() == pb_gen.StatusCode_Success {
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_Success)
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}
func downloadFile(c *gin.Context) {
	req := &pb_gen.DownloadFileRequest{}
	if err := readBody(c, req); err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewDownloadFileHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if resp.GetBaseResp().GetStatusCode() == pb_gen.StatusCode_Success {
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_Success)
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}

func searchFile(c *gin.Context) {
	req := &pb_gen.SearchFileRequest{}
	if err := readBody(c, req); err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewSearchFileHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if resp.GetBaseResp().GetStatusCode() == pb_gen.StatusCode_Success {
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_Success)
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}
func getNode(c *gin.Context) {
	req := &pb_gen.GetNodeRequest{}
	if err := readBody(c, req); err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewGetNodeHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if resp.GetBaseResp().GetStatusCode() == pb_gen.StatusCode_Success {
		resp.BaseResp = util.BuildBaseResp(pb_gen.StatusCode_Success)
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}

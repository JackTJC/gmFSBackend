package http_server

import (
	"net/http"

	"github.com/JackTJC/gmFS_backend/logs"
	"github.com/JackTJC/gmFS_backend/method"
	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
)

var jsonPbMarshaler *jsonpb.Marshaler

func HTTPMain() {
	r := gin.Default()
	setRoute(r)
	r.Run(":9000")
}

// 设置路由
func setRoute(r *gin.Engine) {
	r.POST("/ping", ping)
	r.POST("/user/login", userLogin)
	r.POST("/user/register", userRegister)
	r.POST("/dir/create", createDir)
	r.POST("/file/upload", uploadFile)
	r.POST("/file/download", downloadFile)
	r.POST("/file/search")
	r.POST("/node/get", getNode)
	r.POST("/user/info/get", getUserInfo)
}

// template
func ping(c *gin.Context) {
	req := &pb_gen.PingRequest{}
	err := jsonpb.Unmarshal(c.Request.Body, req) // 从req body解析出request
	if err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewPingHandler(c.Request.Context(), req).Run() // 执行逻辑
	err = jsonPbMarshaler.Marshal(c.Writer, resp)                 // 写回到http response
	if err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}

func userLogin(c *gin.Context) {
	req := &pb_gen.UserLoginRequest{}
	err := jsonpb.Unmarshal(c.Request.Body, req)
	if err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewUserLoginHandler(c.Request.Context(), req).Run()
	err = jsonPbMarshaler.Marshal(c.Writer, resp)
	if err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}

func userRegister(c *gin.Context) {
	req := &pb_gen.UserRegisterRequest{}
	err := jsonpb.Unmarshal(c.Request.Body, req)
	if err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewUserRegisterHandler(c.Request.Context(), req).Run()
	err = jsonPbMarshaler.Marshal(c.Writer, resp)
	if err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)

}

func createDir(c *gin.Context) {
	req := &pb_gen.CreateDirRequest{}
	err := jsonpb.Unmarshal(c.Request.Body, req)
	if err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewCreateDirHandler(c.Request.Context(), req).Run()
	err = jsonPbMarshaler.Marshal(c.Writer, resp)
	if err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)

}

func downloadFile(c *gin.Context) {
	req := &pb_gen.DownloadFileRequest{}
	err := jsonpb.Unmarshal(c.Request.Body, req)
	if err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewDownloadFileHandler(c.Request.Context(), req).Run()
	err = jsonPbMarshaler.Marshal(c.Writer, resp)
	if err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}

func getNode(c *gin.Context) {
	req := &pb_gen.GetNodeRequest{}
	err := jsonpb.Unmarshal(c.Request.Body, req)
	if err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewGetNodeHandler(c.Request.Context(), req).Run()
	err = jsonPbMarshaler.Marshal(c.Writer, resp)
	if err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}

func getUserInfo(c *gin.Context) {
	req := &pb_gen.GetUserInfoRequest{}
	err := jsonpb.Unmarshal(c.Request.Body, req)
	if err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewGetUserInfoHandler(c.Request.Context(), req).Run()
	err = jsonPbMarshaler.Marshal(c.Writer, resp)
	if err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}

func searchFile(c *gin.Context) {
	req := &pb_gen.SearchFileRequest{}
	err := jsonpb.Unmarshal(c.Request.Body, req)
	if err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewSearchFileHandler(c.Request.Context(), req).Run()
	err = jsonPbMarshaler.Marshal(c.Writer, resp)
	if err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}

func uploadFile(c *gin.Context) {
	req := &pb_gen.UploadFileRequest{}
	err := jsonpb.Unmarshal(c.Request.Body, req)
	if err != nil {
		logs.Sugar.Errorf("unmarshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewUploadFileHandler(c.Request.Context(), req).Run()
	err = jsonPbMarshaler.Marshal(c.Writer, resp)
	if err != nil {
		logs.Sugar.Errorf("marshal error:%v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	logs.Sugar.Infof("req = %+v, resp = %+v", req, resp)
	c.Status(http.StatusOK)
}

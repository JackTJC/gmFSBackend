package http_server

import (
	"net/http"

	"github.com/JackTJC/gmFS_backend/method"
	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/gin-gonic/gin"
)

// 设置路由
func setRoute(r *gin.Engine) {
	r.POST("/ping", ping)
	r.POST("/user/login", userLogin)
	r.POST("/user/register", userRegister)
	r.POST("/dir/create", createDir)
	r.POST("/file/upload", uploadFile)
	r.POST("/file/search", searchFile)
	r.POST("/node/get", getNode)
	r.POST("/file/register", registerFile)
	r.POST("/file/share", shareFile)
	r.POST("/file/recv/get", getRecvFile)
}

// template
func ping(c *gin.Context) {
	req := &pb_gen.PingRequest{}
	if err := readBody(c, req); err != nil { // 从http request body解析初请求
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewPingHandler(c.Request.Context(), req).Run() // 业务逻辑
	if err := writeBody(c, resp); err != nil {                    // 响应写回http response
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func userLogin(c *gin.Context) {
	req := &pb_gen.UserLoginRequest{}
	if err := readBody(c, req); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewUserLoginHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
func userRegister(c *gin.Context) {
	req := &pb_gen.UserRegisterRequest{}
	if err := readBody(c, req); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewUserRegisterHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
func createDir(c *gin.Context) {
	req := &pb_gen.CreateDirRequest{}
	if err := readBody(c, req); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewCreateDirHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func uploadFile(c *gin.Context) {
	req := &pb_gen.UploadFileRequest{}
	if err := readBody(c, req); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewUploadFileHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func searchFile(c *gin.Context) {
	req := &pb_gen.SearchFileRequest{}
	if err := readBody(c, req); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewSearchFileHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
func getNode(c *gin.Context) {
	req := &pb_gen.GetNodeRequest{}
	if err := readBody(c, req); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewGetNodeHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func registerFile(c *gin.Context) {
	req := &pb_gen.RegisterFileRequest{}
	if err := readBody(c, req); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewRegisterFileHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func shareFile(c *gin.Context) {
	req := &pb_gen.ShareFileRequest{}
	if err := readBody(c, req); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewShareFileHandler(c.Request.Context(), req).Run()
	if err := writeBody(c, resp); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func getRecvFile(c *gin.Context) {
	req := &pb_gen.GetRecvFileRequest{}
	if err := readBody(c, req); err != nil { // 从http request body解析初请求
		c.Status(http.StatusInternalServerError)
		return
	}
	resp := method.NewGetRecvFileHandler(c.Request.Context(), req).Run() // 业务逻辑
	if err := writeBody(c, resp); err != nil {                           // 响应写回http response
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

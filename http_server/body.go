package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

var (
	marshaler = jsonpb.Marshaler{EmitDefaults: false}
	unmarshal = jsonpb.Unmarshaler{AllowUnknownFields: true}
)

func readBodyFromReq(c *gin.Context, msg proto.Message) error {
	return unmarshal.Unmarshal(c.Request.Body, msg)
}

func writeBodyToResp(c *gin.Context, msg proto.Message) error {
	return marshaler.Marshal(c.Writer, msg)
}

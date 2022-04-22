package http_server

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

type reqFormat int8

const (
	jsonReqFormat   reqFormat = 1
	pbReqFormat     reqFormat = 2
	formatKey                 = "req_format"
	jsonContentType           = "application/json"
	pbContentType             = "application/x-protobuf"
)

var (
	marshaler = jsonpb.Marshaler{EmitDefaults: false}
	unmarshal = jsonpb.Unmarshaler{AllowUnknownFields: true}
)

func readBody(c *gin.Context, msg proto.Message) error {
	format, ok := c.Keys[formatKey].(reqFormat)
	if !ok {
		return errors.New("type assert error")
	}
	switch format {
	case jsonReqFormat:
		c.Header("Content-Type", jsonContentType)
		return unmarshal.Unmarshal(c.Request.Body, msg)
	case pbReqFormat:
		bodyData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return err
		}
		return proto.Unmarshal(bodyData, msg)
	default:
		return errors.New("unknown req format")
	}
}

func writeBody(c *gin.Context, msg proto.Message) error {
	format, ok := c.Keys[formatKey].(reqFormat)
	if !ok {
		return errors.New("type assert error")
	}
	switch format {
	case jsonReqFormat:
		return marshaler.Marshal(c.Writer, msg)
	case pbReqFormat:
		bodyData, err := proto.Marshal(msg)
		if err != nil {
			return err
		}
		c.Data(http.StatusOK, pbContentType, bodyData)
		return nil
	default:
		return errors.New("unknown req format")
	}
}

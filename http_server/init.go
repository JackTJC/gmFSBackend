package http_server

import "github.com/golang/protobuf/jsonpb"

func InitHTTPServer() {
	jsonPbMarshaler = &jsonpb.Marshaler{}
}

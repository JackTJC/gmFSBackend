package test

import (
	"context"
	"testing"

	"github.com/JackTJC/gmFS_backend/method"
	"github.com/JackTJC/gmFS_backend/pb_gen"
	"github.com/kr/pretty"
)

func TestUplooadFile(t *testing.T) {
	ctx := context.Background()
	req := &pb_gen.UploadFileRequest{
		FileName: "tianjincai's test file",
		Content:  []byte("this is test file content"),
	}
	h := method.NewUploadFileHandler(ctx, req)
	pretty.Println(h.Run())
}

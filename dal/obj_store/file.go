package objstore

import (
	"bytes"
	"context"
)

func UploadFile(ctx context.Context, key string, content []byte) error {
	reader := bytes.NewBuffer(content)
	_, err := client.Object.Put(ctx, key, reader, nil)
	return err
}

package files

import (
	"bytes"
	"context"
	"github.com/JPratama7/whahep/types"
	"go.mau.fi/whatsmeow"
	"io"
)

func (f *Files) UploadIO(ctx context.Context, data io.Reader, extension string) (res whatsmeow.UploadResponse, err error) {

	payload := bytes.Buffer{}

	_, err = payload.ReadFrom(data)
	if err != nil {
		return
	}

	res, err = f.cli.Upload(ctx, payload.Bytes(), types.GetMediaTypeByExtension(extension))

	return
}

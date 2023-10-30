package files

import (
	"context"
	"go.mau.fi/whatsmeow"
	"io"
)

func (f *Files) DownloadIO(ctx context.Context, msg whatsmeow.DownloadableMessage, write io.Writer) (err error) {
	payload, err := f.cli.Download(msg)
	if err != nil {
		_ = ctx.Err()
		return
	}

	_, err = write.Write(payload)
	ctx.Done()
	return
}

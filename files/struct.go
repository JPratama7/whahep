package files

import "go.mau.fi/whatsmeow"

type Files struct {
	cli *whatsmeow.Client
}

func NewFiles(client *whatsmeow.Client) *Files {
	return &Files{cli: client}
}

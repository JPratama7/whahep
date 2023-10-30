package sender

import (
	"context"
	"github.com/JPratama7/whahep"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
	"net/http"
	"time"
)

func (s *Sender) Message(ctx context.Context, msg string, toJID types.JID) (resp whatsmeow.SendResponse, err error) {
	var wamsg waProto.Message
	wamsg.Conversation = proto.String(msg)
	resp, err = s.client.SendMessage(ctx, toJID, &wamsg)
	return
}

func (s *Sender) ReadMark(info *types.MessageInfo) (resp whatsmeow.SendResponse, err error) {
	err = s.client.MarkRead([]types.MessageID{info.ID}, time.Now(), info.Chat, info.Sender)
	return
}

func (s *Sender) ListMessage(ctx context.Context, lstmsg whahep.ListMessage, toJID types.JID) (resp whatsmeow.SendResponse, err error) {
	var lms []*waProto.ListMessage_Section
	for _, sec := range lstmsg.Sections {

		var lmr []*waProto.ListMessage_Row
		for _, lst := range sec.Rows {
			tmplst := waProto.ListMessage_Row{
				Title:       proto.String(lst.Title),
				Description: proto.String(lst.Description),
				RowId:       proto.String(lst.RowId),
			}
			lmr = append(lmr, &tmplst)
		}

		tmpsec := waProto.ListMessage_Section{
			Title: proto.String(sec.Title),
			Rows:  lmr,
		}
		lms = append(lms, &tmpsec)
	}

	message := &waProto.Message{
		ListMessage: &waProto.ListMessage{
			Title:       proto.String(lstmsg.Title),
			Description: proto.String(lstmsg.Description),
			FooterText:  proto.String(lstmsg.FooterText),
			ButtonText:  proto.String(lstmsg.ButtonText),
			ListType:    waProto.ListMessage_SINGLE_SELECT.Enum(),
			Sections:    lms,
		},
	}
	viewOnce := &waProto.Message{
		ViewOnceMessage: &waProto.FutureProofMessage{
			Message: message,
		},
	}
	resp, err = s.client.SendMessage(ctx, toJID, viewOnce)
	return

}

func (s *Sender) ButtonMessage(ctx context.Context, btnmsg whahep.ButtonsMessages, toJID types.JID) (resp whatsmeow.SendResponse, err error) {
	var buttons []*waProto.ButtonsMessage_Button
	for _, btn := range btnmsg.Buttons {
		tmpbtn := waProto.ButtonsMessage_Button{
			ButtonId: proto.String(btn.ButtonId),
			ButtonText: &waProto.ButtonsMessage_Button_ButtonText{
				DisplayText: proto.String(btn.DisplayText),
			},
			Type: waProto.ButtonsMessage_Button_RESPONSE.Enum(),
		}
		buttons = append(buttons, &tmpbtn)
	}
	this_message := &waProto.Message{
		ButtonsMessage: &waProto.ButtonsMessage{
			ContentText: proto.String(btnmsg.Message.ContentText),
			FooterText:  proto.String(btnmsg.Message.FooterText),
			Buttons:     buttons,
			HeaderType:  waProto.ButtonsMessage_TEXT.Enum(),
			Header: &waProto.ButtonsMessage_Text{
				Text: btnmsg.Message.HeaderText,
			},
		},
	}
	viewOnce := &waProto.Message{
		ViewOnceMessage: &waProto.FutureProofMessage{
			Message: this_message,
		},
	}
	resp, err = s.client.SendMessage(ctx, toJID, viewOnce)
	return
}

func (s *Sender) DocumentMessage(ctx context.Context, plaintext []byte, filename string, caption string, toJID types.JID) (resp whatsmeow.SendResponse, err error) {
	respupload, err := s.client.Upload(ctx, plaintext, whatsmeow.MediaDocument)
	if err != nil {
		return
	}

	docMsg := &waProto.DocumentMessage{
		Caption:       proto.String(caption),
		Mimetype:      proto.String(http.DetectContentType(plaintext)),
		FileName:      &filename,
		Url:           &respupload.URL,
		DirectPath:    &respupload.DirectPath,
		MediaKey:      respupload.MediaKey,
		FileEncSha256: respupload.FileEncSHA256,
		FileSha256:    respupload.FileSHA256,
		FileLength:    &respupload.FileLength,
	}

	docMessage := &waProto.Message{
		DocumentMessage: docMsg,
	}
	resp, err = s.client.SendMessage(ctx, toJID, docMessage)
	return
}

func (s *Sender) ImageMessage(ctx context.Context, imageUpload whatsmeow.UploadResponse, mimetype, caption string, toJID types.JID) (resp whatsmeow.SendResponse, err error) {

	imgMsg := &waProto.ImageMessage{
		Caption:       proto.String(caption),
		Url:           proto.String(imageUpload.URL),
		DirectPath:    proto.String(imageUpload.DirectPath),
		MediaKey:      imageUpload.MediaKey,
		Mimetype:      proto.String(mimetype),
		FileEncSha256: imageUpload.FileEncSHA256,
		FileSha256:    imageUpload.FileSHA256,
		FileLength:    proto.Uint64(imageUpload.FileLength),
	}

	imgMessage := &waProto.Message{
		ImageMessage: imgMsg,
	}
	resp, err = s.client.SendMessage(ctx, toJID, imgMessage)
	return resp, err
}

func (s *Sender) LiveLocation(ctx context.Context, caption string, latitude, longitude float64, toJID types.JID) (resp whatsmeow.SendResponse, err error) {
	msg := new(waProto.Message)

	msg.LiveLocationMessage = &waProto.LiveLocationMessage{
		DegreesLatitude:                   proto.Float64(latitude),
		DegreesLongitude:                  proto.Float64(longitude),
		AccuracyInMeters:                  nil,
		SpeedInMps:                        nil,
		DegreesClockwiseFromMagneticNorth: nil,
		Caption:                           proto.String(caption),
		SequenceNumber:                    nil,
		TimeOffset:                        nil,
		JpegThumbnail:                     nil,
		ContextInfo:                       nil,
	}

	resp, err = s.client.SendMessage(ctx, toJID, msg)

	return
}

package types

import (
	"go.mau.fi/whatsmeow"
	"strings"
)

func GetMediaTypeByExtension(extension string) whatsmeow.MediaType {

	ext := extension

	if splt := strings.Split(extension, "."); len(splt) > 2 {
		ext = splt[1]
	}

	switch ext {
	case "jpeg", "jpg", "png", "gif":
		return whatsmeow.MediaImage

	case "mp4", "avi":
		return whatsmeow.MediaVideo

	case "mp3", "ogg":
		return whatsmeow.MediaAudio
	default:
		return whatsmeow.MediaDocument
	}

}

package file_picker

import (
	"strings"

	"github.com/pkg/errors"
)

func fileFilter(method string) (string, error) {
	switch method {
	case "ANY":
		return `*.*`, nil
	case "IMAGE":
		return `*.png *.jpg *.jpeg`, nil
	case "AUDIO":
		return `*.mp3`, nil
	case "VIDEO":
		return `*.webm *.mpeg *.mkv *.mp4 *.avi *.mov *.flv`, nil
	default:
		if strings.HasPrefix(method, "__CUSTOM_") {
			resolveType := strings.Split(method, "__CUSTOM_")
			return `*.` + resolveType[1], nil
		}
		return "", errors.New("unknown method")
	}

}

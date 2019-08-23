package file_picker

import (
	"strings"

	"github.com/pkg/errors"
)

func fileFilter(method string) (string, error) {
	switch method {
	case "ANY":
		return `public.item`, nil
	case "IMAGE":
		return `public.image`, nil
	case "AUDIO":
		return `public.audio`, nil
	case "VIDEO":
		return `public.movie`, nil
	default:
		if strings.HasPrefix(method, "__CUSTOM_") {
			resolveType := strings.Split(method, "__CUSTOM_")
			return resolveType[1], nil
		}
		return "", errors.New("unknown method")
	}

}

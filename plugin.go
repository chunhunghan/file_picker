package file_picker

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/pkg/errors"
)

const channelName = "file_picker"

type FilePickerPlugin struct{}

var _ flutter.Plugin = &FilePickerPlugin{} // compile-time type check

func (p *FilePickerPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	dialogProvider := dialogProvider{}

	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	//channel.HandleFunc("openDirectory", p.filePicker(dialogProvider, true))
	channel.HandleFunc("ANY", p.filePicker(dialogProvider, false))
	channel.HandleFunc("IMAGE", p.filePicker(dialogProvider, false))
	channel.HandleFunc("AUDIO", p.filePicker(dialogProvider, false))
	channel.HandleFunc("VIDEO", p.filePicker(dialogProvider, false))

	return nil
}

func (p *FilePickerPlugin) filePicker(dialog dialog, isDirectory bool) func(arguments interface{}) (reply interface{}, err error) {
	return func(arguments interface{}) (reply interface{}, err error) {
		fileDescriptor, _, err := dialog.File("select file", "*", isDirectory)
		if err != nil {
			return nil, errors.Wrap(err, "failed to open dialog picker")
		}

		return fileDescriptor, nil
	}
}

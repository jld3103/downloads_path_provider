package downloads_path_provider

import (
	"errors"
	"path/filepath"

	"github.com/mitchellh/go-homedir"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "downloads_path_provider"

// DownloadsPathProviderPlugin implements flutter.Plugin and handles method.
type DownloadsPathProviderPlugin struct{}

var _ flutter.Plugin = &DownloadsPathProviderPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *DownloadsPathProviderPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getDownloadsDirectory", p.getDownloadsDirectory)
	return nil
}

func (p *DownloadsPathProviderPlugin) getDownloadsDirectory(arguments interface{}) (reply interface{}, err error) {
	dir, err := homedir.Dir()
	if err != nil {
		return nil, errors.New("could not get home directory")
	}
	dir, err = homedir.Expand(dir)
	if err != nil {
		return nil, errors.New("could not expand home directory")
	}
	dir = filepath.Join(dir, "Downloads")
	return dir, nil
}

package downloaders

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/mindscratch/probable-waffle/config"
	"github.com/mindscratch/probable-waffle/downloaders/crate"
)

type DownloaderClient interface {
	// Get a file and return the path to the file or an empty string
	// if it was written to stdout.
	Get() (string, error)
}

var _ DownloaderClient = &crate.Client{}

func New(config config.Config, downloadUri string) (DownloaderClient, error) {
	downloadUrl, err := url.Parse(downloadUri)
	if err != nil {
		return nil, errors.New("Unable to parse download URI: " + err.Error())
	}
	switch downloadUrl.Scheme {
	case crate.Scheme:
		actualUri := fmt.Sprintf("%s/_blobs/%s%s", config.Crate, downloadUrl.Host, downloadUrl.Path)
		crateUrl, err := url.Parse(actualUri)
		if err != nil {
			return nil, errors.New("Problem creating Crate download url: " + err.Error())
		}
		return crate.NewCrateClient(crateUrl)
	}
	return nil, errors.New("Invalid URI scheme")
}

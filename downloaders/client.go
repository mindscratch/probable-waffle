package downloaders

import (
	"errors"
	"fmt"
	"io"
	"net/url"

	"github.com/mindscratch/probable-waffle/config"
	"github.com/mindscratch/probable-waffle/downloaders/crate"
)

type DownloaderClient interface {
	// Get a file and write it using the given writer.
	Get(w io.Writer) error
	Url() *url.URL
}

var _ DownloaderClient = &crate.Client{}

func New(config config.Config, downloadUri string) (DownloaderClient, error) {
	downloadUrl, err := url.Parse(downloadUri)
	if err != nil {
		return nil, errors.New("Unable to parse download URI: " + err.Error())
	}

	var client DownloaderClient = nil

	switch downloadUrl.Scheme {
	case crate.Scheme:
		actualUri := fmt.Sprintf("%s/_blobs/%s%s", config.Crate, downloadUrl.Host, downloadUrl.Path)
		crateUrl, err := url.Parse(actualUri)
		if err != nil {
			return nil, errors.New("Problem creating Crate download url: " + err.Error())
		}

		client, err = crate.NewCrateClient(crateUrl)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Problem creating Crate client: %s", err.Error()))
		}
	default:
		return nil, errors.New("Invalid URI scheme")
	}

	return client, nil
}

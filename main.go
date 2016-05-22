package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/mindscratch/probable-waffle/config"
	"github.com/mindscratch/probable-waffle/downloaders"
)

var (
	appConfig config.Config
)

func getMetadataValue(metadata_file, key string) (string, error) {
	r, re_err := regexp.Compile(fmt.Sprintf("%s: (.*)", key))
	if re_err != nil {
		return "", fmt.Errorf("Unable to compile regex: %s\n", re_err)
	}

	data, err := ioutil.ReadFile(metadata_file)
	if err != nil {
		return "", err
	}
	matches := r.FindStringSubmatch(string(data))
	if len(matches) != 2 {
		return "", fmt.Errorf("Unable to find value where key=%s", key)
	}
	return matches[1], nil
}

func main() {
	flag.Parse()
	appConfig, err := config.New()
	if err != nil {
		fmt.Printf("Configuration error: %s\n", err.Error())
		os.Exit(1)
	}

	metadata_file := "sample.data"
	key := "data_uri"
	retries := 3
	retry_delay := "5s"
	fmt.Println("CONFIG:", metadata_file, key, appConfig.Output, retries, retry_delay)

	downloadUri, err := getMetadataValue(metadata_file, key)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("config: %#v\n", appConfig)
	client, err := downloaders.New(appConfig, downloadUri)
	fmt.Printf("client: %s\n", client)
	fmt.Printf("err: %#v\n", err)

	var outputWriter io.Writer = os.Stdout
	if appConfig.Output != "" {
		//TODO close file "f"
		f, err := os.Create(appConfig.Output)
		if err != nil {
			fmt.Printf("Unable open %s for writing: %s\n", appConfig.Output, err.Error())
			os.Exit(1)
		}
		outputWriter = f
	}

	err = client.Get(outputWriter)
	if err != nil {
		fmt.Printf("get err: %s\n", err.Error())
	}
}

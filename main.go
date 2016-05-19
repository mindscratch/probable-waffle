package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func getValue(metadata_file, key string) (string, error) {
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
		return "", fmt.Errorf("Unable to find value where key=%s\n", key)
	}
	return matches[1], nil
}

func main() {
	metadata_file := "sample.data"
	key := "data_uri"
	output_file := "output.data"
	retries := 3
	retry_delay := "5s"
	fmt.Println("CONFIG:", metadata_file, key, output_file, retries, retry_delay)

	value, err := getValue(metadata_file, key)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("value: %s\n", value)
}

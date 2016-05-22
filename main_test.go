package main

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetMetadataValue(t *testing.T) {
	Convey("Given a metadata file and key", t, func() {
		Convey("When the key is found", func() {
			Convey("The value is returned", func() {
				value, err := getMetadataValue("sample.data", "data_uri")
				So(err, ShouldBeNil)
				So(value, ShouldEqual, "crate://test/22596363b3de40b06f981fb85d82312e8c0ed511")
			})
		})

		Convey("When the key is not found", func() {
			Convey("An error is returned", func() {
				key := "jibberish"
				_, err := getMetadataValue("sample.data", key)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Unable to find value where key="+key)
			})
		})

		Convey("When the file does not exist", func() {
			Convey("An error is returned", func() {
				_, err := getMetadataValue("does-not-exist.data", "something")

				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, &os.PathError{})
			})
		})
	})
}

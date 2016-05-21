package downloaders

import (
	"testing"

	"github.com/mindscratch/probable-waffle/config"
	"github.com/mindscratch/probable-waffle/downloaders/crate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	c, _ := config.New()

	Convey("Given a config and download URI", t, func() {
		Convey("When the download URI is for crate", func() {
			client, err := New(c, "crate://foobar/abc123")
			Convey("The error is nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("A Crate client is returned", func() {
				So(client, ShouldHaveSameTypeAs, &crate.Client{})
			})

			Convey("The Crate download URL is correct", func() {
				So(client.Url().String(), ShouldEqual, c.Crate+"/_blobs/foobar/abc123")
			})
		})

		Convey("When the download URI is malformed", func() {
			Convey("An error is returned", func() {
				_, err := New(c, "://uri")
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "Unable to parse download URI")
			})
		})

		Convey("When the download URI is not supported", func() {
			Convey("An error is returned", func() {
				_, err := New(c, "foo://test")
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Invalid URI scheme")
			})
		})
	})
}

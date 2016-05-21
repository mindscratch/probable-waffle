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
		})
	})
}

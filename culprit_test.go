package sentry

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleCulprit() {
	cl := NewClient(
		// You can set this when creating your client
		Culprit("example"),
	)

	cl.Capture(
		// Or you can set it when sending an event
		Culprit("example"),
	)
}

func TestCulprit(t *testing.T) {
	Convey("Culprit", t, func() {
		Convey("Culprit()", func() {
			Convey("Should return an Option", func() {
				So(Culprit("test"), ShouldImplement, (*Option)(nil))
			})
		})

		Convey("Should use the correct class", func() {
			So(Culprit("test").Class(), ShouldEqual, "culprit")
		})

		Convey("MarshalJSON", func() {
			So(testOptionsSerialize(Culprit("test")), ShouldResemble, "test")
		})
	})
}

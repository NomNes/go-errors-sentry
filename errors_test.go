package errors

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWrap(t *testing.T) {
	Convey("Wrap Error", t, func() {
		err := errors.New("test error")
		wErr := Wrap(err)
		So(wErr, ShouldNotEqual, err)
		So(wErr, ShouldBeError)
		tErr, ok := wErr.(*Error)
		So(ok, ShouldBeTrue)
		So(tErr.err, ShouldEqual, err)
	})
}

func TestUnwrap(t *testing.T) {
	Convey("Wrap Error", t, func() {
		err := errors.New("test error")
		wErr := Wrap(err)
		So(wErr, ShouldNotEqual, err)
		So(Unwrap(err), ShouldEqual, err)
		So(Unwrap(wErr), ShouldEqual, err)
	})
}

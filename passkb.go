package passkb

import (
	"io"
	"time"
)

type Keyboard interface {
	Type(str string, delay time.Duration) error
	io.Closer
}

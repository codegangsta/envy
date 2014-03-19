package autoload

import (
	"github.com/codegangsta/envy/lib"
)

func init() {
	// ignore errors if there are any
	envy.Bootstrap()
}

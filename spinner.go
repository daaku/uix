package uix

import (
	_ "embed"

	"github.com/daaku/cssm"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

//go:embed spinner.css
var kSpinner string

func Spinner(c *cssm.Collector, props ...g.Node) g.Node {
	props = append(props, c.R(kSpinner))
	return h.Div(props...)
}

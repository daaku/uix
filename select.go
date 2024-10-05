package uix

import (
	_ "embed"

	"github.com/daaku/cssm"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

//go:embed select.css
var kSelect string

type SelectOption struct {
	Value string
	Label string
}

func Select(styles *cssm.Collector, currentValue string, options []SelectOption) g.Node {
	return h.Select(styles.R(kSelect),
		g.Map(options, func(o SelectOption) g.Node {
			return h.Option(h.Value(o.Value), g.If(o.Value == currentValue, h.Selected()), g.Text(o.Label))
		}),
	)
}

package uix

import (
	_ "embed"

	"github.com/daaku/cssm"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

//go:embed toggle_switch.css
var kToggleSwitch string

func ToggleSwitch(styles *cssm.Collector, id string, active bool) g.Node {
	return h.Div(styles.R(kToggleSwitch),
		h.Input(h.Type("checkbox"), h.ID(id), g.If(active, h.Checked())),
		h.Label(h.For(id)),
	)
}

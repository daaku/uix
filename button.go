package uix

import (
	_ "embed"

	"github.com/daaku/cssm"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

//go:embed button.css
var kButton string

func PrimaryButton(styles *cssm.Collector, children ...g.Node) g.Node {
	return h.Button(styles.C(kButton, "root", "primary"), g.Group(children))
}

func SecondaryButton(styles *cssm.Collector, children ...g.Node) g.Node {
	return h.Button(styles.C(kButton, "root", "secondary"), g.Group(children))
}

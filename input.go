package uix

import (
	_ "embed"

	"github.com/daaku/cssm"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

//go:embed input.css
var kInput string

func Input(styles *cssm.Collector, attrs ...g.Node) g.Node {
	return h.Input(append([]g.Node{styles.C(kInput, "root")}, attrs...)...)
}

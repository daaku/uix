package uix

import (
	_ "embed"

	"github.com/daaku/cssm"
	g "github.com/maragudk/gomponents"
)

//go:embed reset.css
var kReset string

func Reset(c *cssm.Collector) g.Node {
	_, err := c.Classes(kReset)
	if err != nil {
		panic(err)
	}
	return nil
}

//go:embed system_font.css
var kSystemFont string

func SystemFont(c *cssm.Collector) g.Node {
	_, err := c.Classes(kSystemFont)
	if err != nil {
		panic(err)
	}
	return nil
}

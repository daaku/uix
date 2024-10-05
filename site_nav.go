package uix

import (
	_ "embed"

	"github.com/daaku/cssm"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

//go:embed site_nav.css
var kSiteNav string

func SiteNav(styles *cssm.Collector, logo g.Node, menuItems []g.Node) g.Node {
	const id = "site-nav"
	return h.Nav(styles.C(kSiteNav, "nav"),
		h.Div(logo),
		h.Input(styles.C(kSiteNav, "toggle"),
			h.ID(id), h.Type("checkbox")),
		h.Label(styles.C(kSiteNav, "button-container"), h.For(id),
			h.Div(styles.C(kSiteNav, "button"))),
		h.Ul(styles.C(kSiteNav, "menu"),
			g.Map(menuItems, func(i g.Node) g.Node {
				return h.Li(i)
			})),
	)
}

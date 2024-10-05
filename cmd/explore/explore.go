package main

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/daaku/lands"
	"github.com/daaku/livereload"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	"github.com/sanity-io/litter"

	"github.com/daaku/cssm"
	"github.com/daaku/uix"
	"github.com/daaku/uix/workout"
)

//go:embed icons/*
var icons embed.FS

type example struct {
	path      string
	title     string
	component func(*cssm.Collector) g.Node
}

var examples = []example{
	{
		path:  "/toggle-switch",
		title: "Toggle Switch",
		component: func(styles *cssm.Collector) g.Node {
			return g.Group{
				h.Div(
					g.Text("Toggle Me: "),
					uix.ToggleSwitch(styles, "explore-off", false),
				),
				h.Div(
					g.Text("Toggle Me: "),
					uix.ToggleSwitch(styles, "explore-on", true),
				),
			}
		},
	},
	{
		path:  "/button",
		title: "Button",
		component: func(styles *cssm.Collector) g.Node {
			return g.Group{
				h.Div(
					g.Text("Click something here: "),
					uix.PrimaryButton(styles, g.Text("Submit")),
					g.Text(" "),
					uix.SecondaryButton(styles, g.Text("Cancel")),
				),
				h.Div(
					uix.PrimaryButton(styles,
						uix.SVG(icons, "icons/plus-square-fill"),
						g.Text("Add Item")),
					g.Text(" "),
					uix.SecondaryButton(styles,
						g.Text("Delete Item"),
						uix.SVG(icons, "icons/trash-fill")),
				),
			}
		},
	},
	{
		path:  "/select",
		title: "Select",
		component: func(styles *cssm.Collector) g.Node {
			return h.Div(
				uix.Select(
					styles,
					"",
					[]uix.SelectOption{
						{Value: "year", Label: "Year"},
						{Value: "month", Label: "Month"},
						{Value: "day", Label: "Day"},
					},
				),
			)
		},
	},
	{
		path:  "/workout",
		title: "Single Dumbbell Skull Crusher",
		component: func(c *cssm.Collector) g.Node {
			return g.Group{
				workout.InputWeight(c, 102.5),
				workout.InputReps(c, 5),
			}
		},
	},
	{
		path:  "/edit-exercise",
		title: "Edit Exercise",
		component: func(c *cssm.Collector) g.Node {
			return workout.EditExercise{
				CSS: c,
			}
		},
	},
	{
		path:  "/spinner",
		title: "Spinner",
		component: func(c *cssm.Collector) g.Node {
			const customSize = "10ch"
			const kSpinner = `
				.root {
					display: flex;
					flex-direction: column;
					gap: 1ch;
					align-items: center;
				}
				.root > h3:not(:first-child) {
					margin-top: 3ch;
				}
			`
			return h.Div(c.R(kSpinner),
				h.H3(g.Text("Default")),
				uix.Spinner(c),
				h.H3(g.Text("Custom Size: "+customSize)),
				uix.Spinner(c, h.Style("--spinner-size:"+customSize)),
			)
		},
	},
}

//go:embed explore_main.css
var kExploreMain string

func stdPage(styles *cssm.Collector, title string, body ...g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []g.Node{
			h.Meta(h.Name("viewport"), h.Content("width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0")),
			h.Link(h.Rel("icon"), h.Type("image/png"), h.Href("/assets/favicon.png")),
			uix.Reset(styles),
			uix.SystemFont(styles),
		},
		Body: []g.Node{
			h.Main(styles.R(kExploreMain),
				uix.SiteNav(
					styles,
					h.H3(g.Text(title)),
					[]g.Node{
						h.A(h.Href("/"), g.Text("Home")),
						h.A(h.Href("/about"), g.Text("About Us")),
					},
				),
				h.Div(styles.C(kExploreMain, "content"),
					g.Group(body)),
			),
			h.Script(h.Async(), h.Src("/assets/vendor/htmx.js")),
			h.Script(g.Raw(livereload.JS)),
			styles,
		},
	})
}

type server struct{}

func (s *server) notFound(w http.ResponseWriter, _ *http.Request) error {
	var styles cssm.Collector
	return stdPage(&styles, "Page Not Found",
		h.H1(g.Text("Page Not Found")),
	).Render(w)
}

//go:embed explore_home.css
var kExploreHome string

func (s *server) home(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return s.notFound(w, r)
	}
	var styles cssm.Collector
	return stdPage(&styles, "Home",
		h.Ul(styles.R(kExploreHome),
			g.Map(examples, func(e example) g.Node {
				return h.Li(h.A(h.Href(e.path), g.Text(e.title)))
			}))).Render(w)
}

func (s *server) example(e example) http.HandlerFunc {
	return s.wrap(func(w http.ResponseWriter, r *http.Request) error {
		var styles cssm.Collector
		return stdPage(&styles, e.title, e.component(&styles)).Render(w)
	})
}

func (s *server) wrap(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			fmt.Fprintf(w, "%+v\n", err)
		}
	}
}

//go:embed assets/*
var assets embed.FS

var (
	assetsFS = http.FileServer(http.FS(assets))
	cache    = os.Getenv("LISTEN_FDS") != ""
)

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	if cache || strings.Contains(r.URL.Path, "/vendor/") {
		w.Header().Set("cache-control", "public, max-age=600")
	}
	assetsFS.ServeHTTP(w, r)
}

func run() error {
	s := &server{}
	mux := http.NewServeMux()
	mux.HandleFunc("/livereload", livereload.Handler)
	mux.HandleFunc("/assets/", assetsHandler)
	mux.HandleFunc("/", s.wrap(s.home))
	for _, e := range examples {
		mux.HandleFunc(e.path, s.example(e))
	}
	return lands.ListenAndServe(context.Background(), ":8080", mux)
}

func main() {
	litter.Config.HidePrivateFields = false
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

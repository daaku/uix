package uix

import (
	"embed"
	"io"

	g "github.com/maragudk/gomponents"
)

func SVG(fs embed.FS, name string) g.Node {
	b, err := fs.ReadFile(name + ".svg")
	return g.NodeFunc(func(w io.Writer) error {
		if err != nil {
			return err
		}
		_, err := w.Write(b)
		return err
	})
}

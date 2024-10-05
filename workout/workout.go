package workout

import (
	_ "embed"
	"fmt"
	"io"

	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"

	"github.com/daaku/cssm"
	"github.com/daaku/uix"
)

//go:embed workout.css
var kWorkout string

type InputNumber struct {
	CSS     *cssm.Collector
	Header  string
	Name    string
	Current float64
	Step    float64
}

func (i InputNumber) Render(w io.Writer) error {
	id := "input_number_" + i.Name
	inc := fmt.Sprintf(`document.getElementById("%s").valueAsNumber+=%v`, id, i.Step)
	dec := fmt.Sprintf(`document.getElementById("%s").valueAsNumber-=%v`, id, i.Step)
	return h.Div(
		h.Label(h.For(id),
			h.H3(i.CSS.C(kWorkout, "input-header"), g.Text(i.Header)),
		),
		h.Div(
			i.CSS.C(kWorkout, "input-number-group"),
			uix.SecondaryButton(i.CSS, g.Attr("onclick", dec), g.Text("-")),
			h.Input(
				i.CSS.C(kWorkout, "input-number"),
				h.ID(id),
				h.Type("number"),
				g.Attr("inputmode", "decimal"),
				h.Step(fmt.Sprint(i.Step)),
				h.Value(fmt.Sprint(i.Current)),
			),
			uix.SecondaryButton(i.CSS, g.Attr("onclick", inc), g.Text("+")),
		),
	).Render(w)
}

func InputWeight(c *cssm.Collector, current float64) g.Node {
	return InputNumber{
		CSS:     c,
		Header:  "Weight (kgs)",
		Name:    "weight",
		Current: current,
		Step:    0.5,
	}
}

func InputReps(c *cssm.Collector, current int) g.Node {
	return InputNumber{
		CSS:     c,
		Header:  "Reps",
		Name:    "reps",
		Current: float64(current),
		Step:    1,
	}
}

const UnitDatalistID = "unit-choices"

var UnitDatalist = h.DataList(
	h.ID(UnitDatalistID),
	h.Option(h.Value("kg")),
	h.Option(h.Value("lb")),
	h.Option(h.Value("plate")),
)

//go:embed edit_exercise.css
var kEditExercise string

type EditExercise struct {
	CSS  *cssm.Collector
	Name string
	Unit string
}

func (e EditExercise) Render(w io.Writer) error {
	nameID := "edit_exercise_name_" + e.Name
	unitID := "edit_exercise_unit_" + e.Name
	return h.Div(
		h.H3(g.Text("Edit Exercise")),
		UnitDatalist,
		h.Form(
			e.CSS.C(kEditExercise, "form"),
			h.Label(g.Text("Name"), h.For(nameID)),
			h.Input(g.Attr("size", "1"), h.ID(nameID)),
			h.Label(g.Text("Unit"), h.For(unitID)),
			h.Input(g.Attr("size", "1"), h.ID(unitID), g.Attr("list", UnitDatalistID)),
			h.Div(
				e.CSS.C(kEditExercise, "buttons"),
				uix.PrimaryButton(e.CSS, g.Text("Save")),
				uix.SecondaryButton(e.CSS, g.Text("Cancel")),
			),
		),
	).Render(w)
}

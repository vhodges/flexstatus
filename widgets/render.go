package widgets

import (
	"bytes"
	"html/template"

	"github.com/Masterminds/sprig"
)

func Render(widget Widget) string {

	var line bytes.Buffer

	t := template.Must(template.New("widget_template").
		Funcs(sprig.FuncMap()).
		Parse(widget.GetTemplate()))

	t.Execute(&line, widget)

	return line.String()
}

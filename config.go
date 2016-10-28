package main

import (
	"fmt"
	"time"

	"./widgets"

	"github.com/BurntSushi/toml"
)

type widgetconfig struct {
	Widget []widget
}

// defines a widget.  Must contain a superset of fields a widget can have
type widget struct {
	Kind       string
	Template   string
	Filename   string
	Sleep      int64
	Command    string
	Args       []string
	DateFormat string
	Value      string
}

func get_widgets(filename string) []widgets.Widget {

	active := make([]widgets.Widget, 0)

	var conf widgetconfig

	_, err := toml.DecodeFile(filename, &conf)
	if err != nil {
		fmt.Println(err)
		return default_widgets(active)
	}

	for i := range conf.Widget {
		w := conf.Widget[i]

		switch w.Kind {
		case "Desktop":
			active = append(active, &widgets.Desktop{
				BaseWidget: widgets.BaseWidget{Template: w.Template},
			})
		case "Clock":
			active = append(active, &widgets.Clock{
				BaseWidget: widgets.BaseWidget{Template: w.Template},
				DateFormat: w.DateFormat,
				Millis:     time.Duration(w.Sleep),
			})
		case "PollCommand":
			active = append(active, &widgets.Poller{
				BaseWidget: widgets.BaseWidget{Template: w.Template},
				Command:    w.Command,
				Args:       w.Args,
				Millis:     time.Duration(w.Sleep),
			})
		case "PollFile":
			active = append(active, &widgets.Poller{
				BaseWidget: widgets.BaseWidget{Template: w.Template},
				Filename:   w.Filename,
				Millis:     time.Duration(w.Sleep),
			})
		case "Text":
			active = append(active, &widgets.Text{
				BaseWidget: widgets.BaseWidget{Template: w.Value},
			})
		}
	}

	return active
}

func default_widgets(active []widgets.Widget) []widgets.Widget {

	active = append(active, &widgets.Clock{
		BaseWidget: widgets.BaseWidget{Template: "{{.FormattedDateTime}} :: "},
		DateFormat: "Mon 2 Jan, 15:04",
		Millis:     1000,
	})

	active = append(active, &widgets.Text{
		BaseWidget: widgets.BaseWidget{Template: "flexstatus v0.1.0"},
	})

	return active
}

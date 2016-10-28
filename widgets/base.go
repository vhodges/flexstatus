package widgets

type BaseWidget struct {
	Updates  chan Widget
	Template string
}

func (base *BaseWidget) SetUpdateChannel(updates chan Widget) {
	base.Updates = updates
}

func (base *BaseWidget) GetTemplate() string {
	return base.Template
}

func (base *BaseWidget) Updated(widget Widget) {
	if base.Updates != nil {
		base.Updates <- widget
	}
}

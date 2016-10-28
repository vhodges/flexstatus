package widgets

type Skeleton struct {
	BaseWidget
}

func (widget *Skeleton) Value() string {
	return "Skeleton Widget 1.0"
}

// Don't instantiate or call this!  Will busy wait...
func (widget *Skeleton) Start() {

	go func() {
		for {
			widget.Updated(widget)
		}
	}()
}

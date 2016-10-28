package widgets

import (
	"time"
)

type Clock struct {
	BaseWidget

	DateFormat string
	Millis     time.Duration

	// The time when the ticker fired
	Time time.Time
}

func (widget *Clock) FormattedDateTime() string {
	return widget.Time.Format(widget.DateFormat)
}

func (widget *Clock) Start() {

	if widget.Millis == 0 {
		widget.Millis = 500
	}

	widget.Time = time.Now()

	ticker := time.NewTicker(time.Millisecond * widget.Millis)
	go func() {
		for t := range ticker.C {

			old := widget.Time.Format(widget.DateFormat)

			widget.Time = t

			if old != widget.Time.Format(widget.DateFormat) {
				widget.Updated(widget)
			}
		}
	}()
}

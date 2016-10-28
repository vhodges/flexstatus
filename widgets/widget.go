package widgets

type Widget interface {
	Start()

	// BaseWidget implements these
	GetTemplate() string
	SetUpdateChannel(chan Widget)
	Updated(Widget)
}

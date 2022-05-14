package mediator

type mediator interface {
	canArrive(train) bool
	notifyAboutDeparture()
}

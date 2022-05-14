package chainofresp

type department interface {
	execute(*patient)
	setNext(department)
}

package mediator

type train interface {
	arrive()
	depart()
	permitArrival()
}

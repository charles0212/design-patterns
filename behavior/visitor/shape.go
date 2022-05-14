package visitor

type shape interface {
	getType() string
	accept(visitor)
}

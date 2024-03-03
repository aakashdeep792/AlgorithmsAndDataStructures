package handler

type Handler interface {
	Get()
	Post()
	Update()
	Delete()
}

func NewInventoryHandler() {

}

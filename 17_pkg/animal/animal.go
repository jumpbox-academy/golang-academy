package animal

type animal struct {
	Name string
}

func New() animal {
	return animal{Name: "Cat"}
}

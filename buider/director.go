package buider

type Director struct {
	builder IBuilder
}

func NewDirector (b  IBuilder) *Director{
	return &Director{
		builder : b,
	}
}

func (d *Director) SetBuilder (b IBuilder){
	d.builder = b
}

func (d *Director) 	BuilderHouse() House{
	d.builder.setDoorType()
	d.builder.setNumFloor()
	d.builder.setWindowType()
	return d.builder.getHouse()
}
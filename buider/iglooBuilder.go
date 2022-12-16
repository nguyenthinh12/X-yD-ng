package buider

type iglooBuilder struct{
	windowType string
	doorType string
	floor int
}

func newIglooBuilder() *iglooBuilder{
	return &iglooBuilder{}
}

func (b *iglooBuilder ) setwindowType(){
	b.windowType = "cửa sổ tuyết"
}

func (b *iglooBuilder) setDoorType(){
	b.doorType = "cửa sổ tuyết"
}

func (b  *iglooBuilder) setNumFloor (){
	b.floor = 1
}

func (b *normalBuilder) getHouse() House {
	return House {
		doorType : b.doorType,
		windowType : b.windowType,
		floor     : b.floor
	}
}
package buider

type normalBuilder struc {
	windowType string
	doorType string 
	floor int 
}

func newNormalBuilder() *normalBuilder{
	return &normalBuilder{}
}

func (b *normalBuilder) setWindowType(){
	b.windowType = "cửa sổ bằng gỗ"
}

func (b *normalBuilder) setDoorType(){
	b.doorType = "cửa sổ bằng gỗ"
}

func (b *normalBuider) setNumFloor (){
	b.floor = 2
}

func (b *normalBuilder) getHouse() House {
	return House {
		doorType : b.doorType,
		windowType : b.windowType,
		floor : b.floor
	}
}
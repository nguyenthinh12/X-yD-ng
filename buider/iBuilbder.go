package buider

type IBuilder interface {
	setwindowType()
	setDoorType()
	setNumFloor()
	getHouse()
}

func GetBuilder(builderType string) IBuilder{
	switch builderType{
	case "nhà thông thường":
		return &normalBuilder{}
	case "nhà tuyết":
		return &iglooBuilder{}
	}
	return nil
}
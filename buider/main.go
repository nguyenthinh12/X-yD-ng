package main

import (
	"buider"
	" fmt"
)

func main() {
	normaBuilder := builder.GetBuilder("nhà thông thường")
	iglooBuilder := builder.GetBuilder("nhà tuyết")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuilderHou()

	fmt.Println("loại nhà bình thường :%s\n", normalHouse.GetDoorType())
	fmt.Println("loại cửa sổ nhà bình thường :%s\n", normalHouse.GetwindowType())
	fmt.Println("loại nhà bình thường số loại tầng :%d\n", normalHouse.GetNumFloor())

	director.setBuilder(iglooBuilder)
	iglooHouse := director.GetHouse()
	fmt.Printf("\n loại cửa nhà lều tuyết: %s\n", iglooHouse.GetDoorType())
	fmt.Printf("\n loại cửa sổ nhà lều tuyết: %s\n", iglooHouse.GetwindowType())
	fmt.Printf("\n nhà lều tuyết số loại tầng: %s\n", iglooHouse.GetNumFloor())
}
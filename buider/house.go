package buider

type House struct {
	windowType string
	doorType string
	floor int 
}

func (h *House) GetwindowTepe()string {
	return h.windowType
}

func (h *House) GetDoorType()string{
	return h.doorType
}

func (h *House) GetNumFloor() int {
	return h.floor
}
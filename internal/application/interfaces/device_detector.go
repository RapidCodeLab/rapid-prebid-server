package interfaces

type DeviceData struct {
	DeviceType int
	Platform   int
	OS         int
}

type DeviceDetecor interface {
	Detect(user_agent string) (DeviceData, error)
}

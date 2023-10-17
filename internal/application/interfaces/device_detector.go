package interfaces

type DeviceData struct {
	DeviceType int
	Platform   int
	OS         int
}

type DeviceDetector interface {
	Detect(user_agent string) (DeviceData, error)
}

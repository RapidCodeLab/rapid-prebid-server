package interfaces

type DeviceDetecor interface{
	Detect(user_agent string) ([]byte, error)
}

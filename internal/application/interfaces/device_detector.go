package interfaces

import "github.com/prebid/openrtb/v17/adcom1"

type DeviceData struct {
	Browser    string
	Platform   string
	DeviceType adcom1.DeviceType
}

type DeviceDetector interface {
	Detect(userAgent string) (DeviceData, error)
}

package interfaces

import "github.com/prebid/openrtb/v17/adcom1"

type (
	DeviceDetector interface {
		Detect(userAgent string) DeviceData
	}
	DeviceData struct {
		Browser    string
		Platform   string
		DeviceType adcom1.DeviceType
		UserAgent  string
	}
)

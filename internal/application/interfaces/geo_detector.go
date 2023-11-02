package interfaces

import "net"

type (
	GeoDetector interface {
		Detect(net.IP) (GeoData, error)
	}
	GeoData struct {
		CountryCode string
		Region      string
		City        string
	}
)

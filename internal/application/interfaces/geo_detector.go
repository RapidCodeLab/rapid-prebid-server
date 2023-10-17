package interfaces

import "net"

type GeoData struct {
	CountryCode string
	Region      string
	City        string
}

type GeoDetector interface {
	Detect(net.IP) (GeoData, error)
}

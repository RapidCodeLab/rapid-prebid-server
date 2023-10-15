package interfaces

import "net"

type GeoDetector interface{
	Detect(ip net.IP) ([]byte, error)
}

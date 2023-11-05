package geoip2_detector

import (
	"net"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/oschwald/geoip2-golang"
)

type GeoLocator struct {
	namesLanguage string
	reader        *geoip2.Reader
}

func New(path,
	namesLanguage string,
) (*GeoLocator, error) {
	l := &GeoLocator{}

	r, err := geoip2.Open(path)
	if err != nil {
		return l, err
	}

	l.reader = r
	return l, nil
}

func (l *GeoLocator) Detect(
	ip net.IP,
) (interfaces.GeoData, error) {
	data := interfaces.GeoData{}

	if ip.To4() != nil {
		data.IP = ip.String()
	}
	if ip.To16() != nil {
		data.IPv6 = ip.String()
	}

	d, err := l.reader.City(ip)
	if err != nil {
		return data, err
	}

	data.CountryCode = d.Country.IsoCode
	if len(d.Subdivisions) > 0 {
		data.Region = d.Subdivisions[0].Names[l.namesLanguage]
	}
	cityName, ok := d.City.Names[l.namesLanguage]
	if ok {
		data.City = cityName
	}

	return data, nil
}

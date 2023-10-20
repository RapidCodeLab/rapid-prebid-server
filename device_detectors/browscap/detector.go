package browscap_devicedetector

import (
	"errors"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	bgo "github.com/hamaxx/browscap_go"
	"github.com/prebid/openrtb/v17/adcom1"
)

const (
	unknownDeviceType = 0
)

var notDetected error = errors.New("device not detected")

type DeviceDetector struct {
	reader *bgo.BrowsCap
}

func New(path string) (*DeviceDetector, error) {
	d := &DeviceDetector{}

	r, err := bgo.NewBrowsCapFromFile(path)
	if err != nil {
		return d, err
	}

	d.reader = r
	return d, nil
}

func (d *DeviceDetector) Detect(
	ua string,
) (interfaces.DeviceData, error) {
	// TODO: добавить все доступные данные по юзерагенту,
	// версии браузера, платформы, етс.

	data := interfaces.DeviceData{}

	b, ok := d.reader.GetBrowser(ua)
	if !ok {
		err := notDetected
		return data, err
	}

	data.Browser = b.Browser
	data.Platform = b.Platform

	switch {
	case b.IsIPhone():
		data.DeviceType = adcom1.DevicePhone
	case b.IsMobile():
		data.DeviceType = adcom1.DeviceMobile
	case b.IsDesktop():
		data.DeviceType = adcom1.DevicePC
	case b.IsIPad() || b.IsTablet():
		data.DeviceType = adcom1.DeviceTablet
	case b.IsTv():
		data.DeviceType = adcom1.DeviceTV
	default:
		data.DeviceType = unknownDeviceType
	}

	return data, nil
}

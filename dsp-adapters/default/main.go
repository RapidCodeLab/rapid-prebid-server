package default_dsp_adapter

import (
	"encoding/json"
	"time"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/prebid/openrtb/v17/openrtb2"
	"github.com/valyala/fasthttp"
)

type adapter struct {
	Name          interfaces.DSPName
	httpClient    *fasthttp.Client
	requestTimout time.Duration
	endpintURI    string
}

func (i *adapter) DoRequest(
	bidRequest openrtb2.BidRequest,
) (
	openrtb2.BidResponse,
	error,
) {
	bidResponse := openrtb2.BidResponse{}

	httpReq := fasthttp.AcquireRequest()
	httpRes := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseRequest(httpReq)
		fasthttp.ReleaseResponse(httpRes)
	}()

	reqBody, err := json.Marshal(bidRequest)
	if err != nil {
		return bidResponse, err
	}

	httpReq.SetBody(reqBody)
	httpReq.SetRequestURI(i.endpintURI)

	err = i.httpClient.DoTimeout(
		httpReq,
		httpRes,
		i.requestTimout,
	)
	if err != nil {
		return bidResponse, err
	}

	err = json.Unmarshal(
		httpReq.Body(),
		&bidResponse,
	)
	if err != nil {
		return bidResponse, err
	}

	return bidResponse, nil
}

func NewDSPAdpater(
	c interfaces.DSPAdapterConfig,
) (interfaces.DSPAdapter, error) {
	return &adapter{
		Name:       c.Name,
		httpClient: &fasthttp.Client{},
		endpintURI: c.Endpoint,
	}, nil
}
package default_dsp_adapter

import (
	"encoding/json"
	"fmt"
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

func (i *adapter) GetName() interfaces.DSPName {
	return i.Name
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
	httpReq.Header.SetMethod(fasthttp.MethodPost)

	err = i.httpClient.DoTimeout(
		httpReq,
		httpRes,
		i.requestTimout,
	)
	if err != nil {
		return bidResponse, err
	}
	if httpRes.StatusCode() != fasthttp.StatusOK {
		return bidResponse,
			fmt.Errorf("%s: response status: %d",
				interfaces.DSPResponseErr.Error(),
				httpRes.StatusCode(),
			)
	}

	err = json.Unmarshal(
		httpRes.Body(),
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
		requestTimout: time.Duration(
			time.Duration(c.RequestTimout) *
				time.Millisecond,
		),
	}, nil
}

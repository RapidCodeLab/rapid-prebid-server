package payload_handler

import (
	"encoding/json"
	"strings"

	"github.com/prebid/openrtb/v17/native1/response"
	"github.com/prebid/openrtb/v17/openrtb2"
)

func (h *Handler) buildResponse(
	bids []openrtb2.Bid,
) payloadResponse {
	res := payloadResponse{}

	for _, bid := range bids {
		var (
			adMarkup string
			err      error
		)

		switch bid.MType {
		case openrtb2.MarkupNative:
			adMarkup, err = buildNativeMarkup(bid.AdM)
			if err != nil {
			}
		default:
			adMarkup = bid.AdM
		}
		p := payload{
			EntityID:   bid.ImpID,
			Adm:        adMarkup,
			MarkupType: bid.MType,
		}
		res.Paylads = append(res.Paylads, p)
	}

	return res
}

func buildNativeMarkup(
	data string,
) (string, error) {
	var (
		markup    string
		nativeObj response.Response
	)

	markup = nativeAdMarkupTemplate

	err := json.Unmarshal([]byte(data), &nativeObj)
	if err != nil {
		return markup, err
	}

	for _, asset := range nativeObj.Assets {
		switch true {
		case asset.Title != nil:
			markup = strings.ReplaceAll(markup, "{{TITLE}}", asset.Title.Text)
		case asset.Img != nil:
			markup = strings.ReplaceAll(markup, "{{IMG}}", asset.Img.URL)
		case asset.Data != nil:
			markup = strings.ReplaceAll(markup, "{{DESC}}", asset.Data.Value)
		case asset.Link != nil:
			markup = strings.ReplaceAll(markup, "{{LINK}}", asset.Link.URL)
		}
	}
	return markup, nil
}

var nativeAdMarkupTemplate = `<div>
<a href="{{LINK}}" target="_blank">
<img src="{{IMG}}/>
<br>
<b>{{TITLE}}</b>
<p>{{DESC}}</p>
</a>
</div>`

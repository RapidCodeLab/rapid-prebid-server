package inventoryapi_handler_test

import (
	"errors"
	"testing"

	inventoryapi_handler "github.com/RapidCodeLab/rapid-prebid-server/internal/application/handlers/inventory-api"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	mock_inventory "github.com/RapidCodeLab/rapid-prebid-server/mocks/inventory_api"
	mock_logger "github.com/RapidCodeLab/rapid-prebid-server/mocks/logger"
	"github.com/valyala/fasthttp"
	"go.uber.org/mock/gomock"
)

var validInventoryPyload = "{\"id\":\"someid\", \"name\":\"\", \"inventory_type\": 0, \"iab_categories\":[], \"blocked_advertiser_iab_categories\":[], \"iab_categories_taxonomy\":0}"

func TestHandleReadAllInventorires(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := mock_logger.NewMockLogger(ctrl)
	invAPIStorage := mock_inventory.NewMockInventoryAPI(ctrl)
	handler := inventoryapi_handler.New(l, invAPIStorage)

	testCases := []struct {
		name           string
		expectHTTPCode int
		apiReturnData  []interfaces.Inventory
		apiReturnErr   error
	}{
		{
			name:           "bad gateway",
			expectHTTPCode: fasthttp.StatusBadGateway,
			apiReturnData:  nil,
			apiReturnErr:   errors.New("some error"),
		},
		{
			name:           "no content",
			expectHTTPCode: fasthttp.StatusNoContent,
			apiReturnData:  []interfaces.Inventory{},
			apiReturnErr:   nil,
		},
		{
			name:           "ok",
			expectHTTPCode: fasthttp.StatusOK,
			apiReturnData: []interfaces.Inventory{
				{},
			},
			apiReturnErr: nil,
		},
	}

	for _, tc := range testCases {

		invAPIStorage.EXPECT().
			ReadAllInventories().
			Return(tc.apiReturnData,
				tc.apiReturnErr)

		reqCtx := &fasthttp.RequestCtx{}

		handler.HandleReadAllInventories(reqCtx)

		if reqCtx.Response.StatusCode() !=
			tc.expectHTTPCode {
			t.Fail()
		}
	}
}

func TestHandleCreateInventory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := mock_logger.NewMockLogger(ctrl)

	testCases := []struct {
		name           string
		expectHTTPCode int
		payload        []byte
		apiReturnErr   error
	}{
		{
			name:           "bad gateway",
			payload:        []byte(validInventoryPyload),
			expectHTTPCode: fasthttp.StatusBadGateway,
			apiReturnErr:   errors.New("some error"),
		},
		{
			name:           "bad request",
			payload:        []byte("{"),
			expectHTTPCode: fasthttp.StatusBadRequest,
			apiReturnErr:   nil,
		},
		{
			name:           "ok",
			payload:        []byte(validInventoryPyload),
			expectHTTPCode: fasthttp.StatusOK,
			apiReturnErr:   nil,
		},
	}

	for _, tc := range testCases {

		invAPIStorage := mock_inventory.NewMockInventoryAPI(ctrl)
		handler := inventoryapi_handler.New(l, invAPIStorage)

		invAPIStorage.EXPECT().
			CreateInventory(gomock.Not(nil)).
			AnyTimes().
			Return(tc.apiReturnErr)

		reqCtx := &fasthttp.RequestCtx{}
		reqCtx.Request.SetBody(tc.payload)

		handler.HandleCreateInventory(reqCtx)

		if reqCtx.Response.StatusCode() !=
			tc.expectHTTPCode {
			t.Fail()
		}
	}
}

func TestHandleReadInventory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := mock_logger.NewMockLogger(ctrl)

	testCases := []struct {
		name           string
		expectHTTPCode int
		apiReturnData  interface{}
		apiReturnErr   error
	}{
		{
			name:           "bad gateway",
			expectHTTPCode: fasthttp.StatusBadGateway,
			apiReturnData:  interfaces.Inventory{},
			apiReturnErr:   errors.New("some error"),
		},
		{
			name:           "ok",
			expectHTTPCode: fasthttp.StatusOK,
			apiReturnData:  interfaces.Inventory{},
			apiReturnErr:   nil,
		},
	}

	for _, tc := range testCases {

		invAPIStorage := mock_inventory.NewMockInventoryAPI(ctrl)
		handler := inventoryapi_handler.New(l, invAPIStorage)

		invAPIStorage.EXPECT().
			ReadInventory(gomock.Not(nil)).
			AnyTimes().
			Return(tc.apiReturnData, tc.apiReturnErr)

		reqCtx := &fasthttp.RequestCtx{}
		reqCtx.SetUserValue("id", "someid")

		handler.HandleReadInventory(reqCtx)

		if reqCtx.Response.StatusCode() !=
			tc.expectHTTPCode {
			t.Fail()
		}
	}
}

func TestHandleUpdateInventory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := mock_logger.NewMockLogger(ctrl)

	testCases := []struct {
		name           string
		expectHTTPCode int
		payload        []byte
		apiReturnErr   error
	}{
		{
			name:           "bad gateway",
			payload:        []byte(validInventoryPyload),
			expectHTTPCode: fasthttp.StatusBadGateway,
			apiReturnErr:   errors.New("some error"),
		},
		{
			name:           "bad request",
			payload:        []byte("{"),
			expectHTTPCode: fasthttp.StatusBadRequest,
			apiReturnErr:   nil,
		},
		{
			name:           "ok",
			payload:        []byte(validInventoryPyload),
			expectHTTPCode: fasthttp.StatusOK,
			apiReturnErr:   nil,
		},
	}

	for _, tc := range testCases {

		invAPIStorage := mock_inventory.NewMockInventoryAPI(ctrl)
		handler := inventoryapi_handler.New(l, invAPIStorage)

		invAPIStorage.EXPECT().
			UpdateInventory(gomock.Not(nil)).
			AnyTimes().
			Return(tc.apiReturnErr)

		reqCtx := &fasthttp.RequestCtx{}
		reqCtx.Request.SetBody(tc.payload)

		handler.HandleUpdateInventory(reqCtx)

		if reqCtx.Response.StatusCode() !=
			tc.expectHTTPCode {
			t.Fail()
		}
	}
}

func TestHandleDeleteInventory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := mock_logger.NewMockLogger(ctrl)

	testCases := []struct {
		name           string
		expectHTTPCode int
		apiReturnErr   error
	}{
		{
			name:           "bad gateway",
			expectHTTPCode: fasthttp.StatusBadGateway,
			apiReturnErr:   errors.New("some error"),
		},
		{
			name:           "ok",
			expectHTTPCode: fasthttp.StatusOK,
			apiReturnErr:   nil,
		},
	}

	for _, tc := range testCases {

		invAPIStorage := mock_inventory.NewMockInventoryAPI(ctrl)
		handler := inventoryapi_handler.New(l, invAPIStorage)

		invAPIStorage.EXPECT().
			DeleteInventory(gomock.Not(nil)).
			AnyTimes().
			Return(tc.apiReturnErr)

		reqCtx := &fasthttp.RequestCtx{}
		reqCtx.SetUserValue("id", "someid")

		handler.HandleDeleteInventory(reqCtx)

		if reqCtx.Response.StatusCode() !=
			tc.expectHTTPCode {
			t.Fail()
		}
	}
}

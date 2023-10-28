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

var validEntityPyload = "{\"id\":\"someid\", \"name\":\"\", \"inventory_type\": 0, \"iab_categories\":[], \"blocked_advertiser_iab_categories\":[], \"iab_categories_taxonomy\":0}"

func TestHandleReadAllEntities(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := mock_logger.NewMockLogger(ctrl)
	invAPIStorage := mock_inventory.NewMockInventoryAPI(ctrl)
	handler := inventoryapi_handler.New(l, invAPIStorage)

	testCases := []struct {
		name           string
		expectHTTPCode int
		apiReturnData  []interfaces.Entity
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
			apiReturnData:  []interfaces.Entity{},
			apiReturnErr:   nil,
		},
		{
			name:           "ok",
			expectHTTPCode: fasthttp.StatusOK,
			apiReturnData: []interfaces.Entity{
				{},
			},
			apiReturnErr: nil,
		},
	}

	for _, tc := range testCases {

		invAPIStorage.EXPECT().
			ReadAllEntities().
			Return(tc.apiReturnData,
				tc.apiReturnErr)

		reqCtx := &fasthttp.RequestCtx{}

		handler.HandleReadAllEntities(reqCtx)

		if reqCtx.Response.StatusCode() !=
			tc.expectHTTPCode {
			t.Fail()
		}
	}
}

func TestHandleCreateEntity(t *testing.T) {
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
			payload:        []byte(validEntityPyload),
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
			payload:        []byte(validEntityPyload),
			expectHTTPCode: fasthttp.StatusOK,
			apiReturnErr:   nil,
		},
	}

	for _, tc := range testCases {

		invAPIStorage := mock_inventory.NewMockInventoryAPI(ctrl)
		handler := inventoryapi_handler.New(l, invAPIStorage)

		invAPIStorage.EXPECT().
			CreateEntity(gomock.Not(nil)).
			AnyTimes().
			Return(tc.apiReturnErr)

		reqCtx := &fasthttp.RequestCtx{}
		reqCtx.Request.SetBody(tc.payload)

		handler.HandleCreateEntity(reqCtx)

		if reqCtx.Response.StatusCode() !=
			tc.expectHTTPCode {
			t.Fail()
		}
	}
}

func TestHandleReadEntity(t *testing.T) {
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
			apiReturnData:  interfaces.Entity{},
			apiReturnErr:   errors.New("some error"),
		},
		{
			name:           "ok",
			expectHTTPCode: fasthttp.StatusOK,
			apiReturnData:  interfaces.Entity{},
			apiReturnErr:   nil,
		},
	}

	for _, tc := range testCases {

		invAPIStorage := mock_inventory.NewMockInventoryAPI(ctrl)
		handler := inventoryapi_handler.New(l, invAPIStorage)

		invAPIStorage.EXPECT().
			ReadEntity(gomock.Not(nil)).
			AnyTimes().
			Return(tc.apiReturnData, tc.apiReturnErr)

		reqCtx := &fasthttp.RequestCtx{}
		reqCtx.SetUserValue("id", "someid")

		handler.HandleReadEntity(reqCtx)

		if reqCtx.Response.StatusCode() !=
			tc.expectHTTPCode {
			t.Fail()
		}
	}
}

func TestHandleUpdateEntity(t *testing.T) {
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
			payload:        []byte(validEntityPyload),
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
			payload:        []byte(validEntityPyload),
			expectHTTPCode: fasthttp.StatusOK,
			apiReturnErr:   nil,
		},
	}

	for _, tc := range testCases {

		invAPIStorage := mock_inventory.NewMockInventoryAPI(ctrl)
		handler := inventoryapi_handler.New(l, invAPIStorage)

		invAPIStorage.EXPECT().
			UpdateEntity(gomock.Not(nil)).
			AnyTimes().
			Return(tc.apiReturnErr)

		reqCtx := &fasthttp.RequestCtx{}
		reqCtx.Request.SetBody(tc.payload)

		handler.HandleUpdateEntity(reqCtx)

		if reqCtx.Response.StatusCode() !=
			tc.expectHTTPCode {
			t.Fail()
		}
	}
}

func TestHandleDeleteEntity(t *testing.T) {
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
			DeleteEntity(gomock.Not(nil)).
			AnyTimes().
			Return(tc.apiReturnErr)

		reqCtx := &fasthttp.RequestCtx{}
		reqCtx.SetUserValue("id", "someid")

		handler.HandleDeleteEntity(reqCtx)

		if reqCtx.Response.StatusCode() !=
			tc.expectHTTPCode {
			t.Fail()
		}
	}
}

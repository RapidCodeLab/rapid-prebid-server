package health_handler_test

import (
	"testing"

	health_handler "github.com/RapidCodeLab/rapid-prebid-server/internal/application/handlers/health"
	mock_interfaces "github.com/RapidCodeLab/rapid-prebid-server/mocks/logger"
	"github.com/valyala/fasthttp"
	"go.uber.org/mock/gomock"
)

func TestHandleStatusOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := mock_interfaces.NewMockLogger(ctrl)

	handler := health_handler.New(l)
	reqCtx := &fasthttp.RequestCtx{}

	handler.Handle(reqCtx)

	if reqCtx.Response.StatusCode() != fasthttp.StatusOK{
		t.Fail()
	}

	//TODO: body assert
}

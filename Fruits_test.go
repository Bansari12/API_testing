package example

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestFruits(t *testing.T) {
	// create http.Handler
	handler := FruitsHandler()

	// run server using httptest
	server := httptest.NewServer(handler)
	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	// is it working?
	e.GET("/fruits").
		Expect().
		Status(http.StatusOK).JSON().Array().Empty()
}
package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSetup(t *testing.T) {
	Convey("Given an API instance", t, func() {
		r := http.NewServeMux()
		ctx := context.Background()
		api := Setup(ctx, r)

		// TODO: remove hello world example handler route test case
		Convey("When created the following routes should have been added", func() {
			// Replace the check below with any newly added api endpoints
			So(hasRoute(api.Mux, "/hello", "GET"), ShouldBeTrue)
		})
	})
}

func hasRoute(m *http.ServeMux, path, method string) bool {
	req := httptest.NewRequest(method, path, http.NoBody)

	_, pat := m.Handler(req)
	return pat != ""
}

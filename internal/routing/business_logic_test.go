package routing

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewBLRouter(t *testing.T) {
	r := NewBLRouter()
	srv := httptest.NewServer(r)
	defer srv.Close()

	testCases := []struct {
		route              string
		expectedStatusCode int
	}{
		{"/home", http.StatusOK},
		{"/", http.StatusNotFound},
	}

	for _, c := range testCases {
		resp, err := http.Get(srv.URL + c.route)
		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode != c.expectedStatusCode {
			t.Errorf("Status code is %v, but %v expected", resp.StatusCode, c.expectedStatusCode)
		}
	}

}

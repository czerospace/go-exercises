package testHandler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "http://example.com/add?a=5&b=3", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	Handle(w, req)

	if want, got := "8", w.Body.String(); want != got {
		t.Fatalf("expected a %s, instead got: %s", want, got)
	}

}

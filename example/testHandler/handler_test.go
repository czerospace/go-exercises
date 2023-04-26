package testHandler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// 构建 req
	req, err := http.NewRequest("GET", "http://example.com/add?a=5&b=3", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	// 使用待测函数处理
	Handler(w, req)

	if want, got := "8", w.Body.String(); want != got {
		t.Fatalf("expected a %s, instead got: %s", want, got)
	}

}

package testHandler

import (
	"fmt"
	"net/http"
	"strconv"
)

func Handle(w http.ResponseWriter, r *http.Request) {

	// 获取 URL 参数 a
	a, err := strconv.Atoi(r.URL.Query().Get("a"))
	if err != nil {
		http.Error(w, "Invalid parameter 'a'", http.StatusBadRequest)
		return
	}
	// 获取 URL 参数 b
	b, err := strconv.Atoi(r.URL.Query().Get("b"))
	if err != nil {
		http.Error(w, "Invalid parameter 'b'", http.StatusBadRequest)
		return
	}

	// 计算 a + b
	sum := a + b
	// 将结果发送回客户端
	fmt.Fprintf(w, "%d", sum)
}

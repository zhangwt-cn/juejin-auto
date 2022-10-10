package util

import (
	"net/http"
	"strings"
)

// http 请求简单封装
func Post(url string, headMap map[string] string) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(""))
	if err != nil {
		return nil, err
	}
	for k , v := range headMap {
		req.Header.Add(k, v)
	}
	return http.DefaultClient.Do(req)
}
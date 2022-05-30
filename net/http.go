// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2022-05-27 22:47
// version: 1.0.0
// desc   :

package net

import (
	"github.com/valyala/fasthttp"
)

// Get http GET 请求方式
func Get(url string) (data []byte, err error) {
	req := fasthttp.AcquireRequest()
	// 用完需要释放资源
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	// 用完需要释放资源
	defer fasthttp.ReleaseResponse(resp)

	if err = fasthttp.Do(req, resp); err != nil {
		return
	}
	data = resp.Body()
	return
}

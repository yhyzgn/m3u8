// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2022-05-27 22:55
// version: 1.0.0
// desc   :

package net

import "testing"

func TestGet(t *testing.T) {
	url := "http://devimages.apple.com/iphone/samples/bipbop/bipbopall.m3u8"
	//url := "https://bvujarg.xyz/tv_adult/avid6228f4d21745b/avid6228f4d21745b.m3u8?siteUrl=https://video.awvvvvw.live"

	data, err := Get(url)
	if nil != err {
		panic(err)
	}
	t.Log(string(data))
}

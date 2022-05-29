// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2022-05-27 23:17
// version: 1.0.0
// desc   :

package parser

import "testing"

const (
	//url    = "https://bvujarg.xyz/tv_adult/avid6228f4d21745b/avid6228f4d21745b.m3u8?siteUrl=https://video.awvvvvw.live"
	url = "https://bf1.semaobf1.com/20220520/6A87E01765D2A36C/hls/1500k/index.m3u8"

	master = `#EXTM3U
#EXT-X-STREAM-INF:BANDWIDTH=150000,RESOLUTION=416x234,CODECS="avc1.42e00a,mp4a.40.2"
http://example.com/low/index.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=240000,RESOLUTION=416x234,CODECS="avc1.42e00a,mp4a.40.2"
http://example.com/lo_mid/index.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=440000,RESOLUTION=416x234,CODECS="avc1.42e00a,mp4a.40.2"
http://example.com/hi_mid/index.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=640000,RESOLUTION=640x360,CODECS="avc1.42e00a,mp4a.40.2"
http://example.com/high/index.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=64000,CODECS="mp4a.40.5"
http://example.com/audio/index.m3u8
#EXT-X-ENDLIST
`
)

func TestParseURL(t *testing.T) {
	m3u8, err := FromNetwork(url)
	if nil != err {
		panic(err)
	}
	t.Log(m3u8)
}

func TestParseString(t *testing.T) {
	m3u8, err := FromString(master)
	if nil != err {
		panic(err)
	}
	t.Log(m3u8)
}

// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2022-05-28 00:17
// version: 1.0.0
// desc   :

package parser

import (
	"bytes"
	"errors"
	"github.com/yhyzgn/m3u8/model"
	"github.com/yhyzgn/m3u8/util"
	"regexp"
	"strconv"
	"strings"
)

const (
	extM3u              = "#EXTM3U"
	extVersion          = "#EXT-X-VERSION:"
	extXIFrameStreamInf = "#EXT-X-I-FRAME-STREAM-INF"
	extXStreamInf       = "#EXT-X-STREAM-INF:"
	extInf              = "#EXTINF:"
	extXKey             = "#EXT-X-KEY:"
	extXEnd             = "#EXT-X-ENDLIST:"

	fieldProgramID  = "PROGRAM-ID"
	fieldBandWidth  = "BANDWIDTH"
	fieldResolution = "RESOLUTION"
	fieldCodeCS     = "CODECS"
	fieldURI        = "URI"

	cryptMethodAES  model.CryptMethod = "AES-128"
	cryptMethodNONE model.CryptMethod = "NONE"
)

// regex pattern for extracting `key=value` parameters from a line
var linePattern = regexp.MustCompile(`([a-zA-Z-]+)=("[^"]+"|[^",]+)`)

// Doc 文档信息
type Doc struct {
	m3u8      *model.M3U8
	Lines     [][]byte
	LineIndex int
}

// NewDoc 创建一个文档对象
func NewDoc(lines [][]byte, m3u8 *model.M3U8) *Doc {
	return &Doc{
		m3u8:      m3u8,
		Lines:     lines,
		LineIndex: -1,
	}
}

// Parse 开始解析
func (d *Doc) Parse() (res *model.M3U8, err error) {
	res = d.m3u8

	// 函数内全局变量，因为每个 key 都将应用于其下面的所有 ts 片段
	var key *model.Key

	for d.HasNextLine() {
		line := d.Line()

		if 0 == d.LineIndex {
			// 校验第一行，是否符合 m3u8 文件规范
			// #EXTM3U
			if string(line) != extM3u {
				err = errors.New("illegal m3u8 content")
				return
			}
			continue
		}

		// 其他行
		switch true {
		case bytes.HasPrefix(line, []byte(extVersion)):
			// 版本信息
			vsnBys := bytes.Split(line, []byte(":"))[1]
			if version, e := strconv.Atoi(string(vsnBys)); nil != e {
				err = e
			} else {
				res.Version = version
			}
			break
		case bytes.HasPrefix(line, []byte(extInf)):
			// 分片TS的信息，如时长，带宽等
			// #EXTINF:duration,<title>
			extInf := bytes.Split(line, []byte(":"))[1]
			spt := bytes.Split(extInf, []byte(","))
			// 直接获取下一行的 ts 文件信息
			if d.HasNextLine() {
				tsURL := util.BuildRealURL(res.URL, string(d.Line()))
				// ts 片段将引用往上查找到最近的 key，如果为 nil 则说明不需要 key
				res.TsList = append(res.TsList, model.TS{
					Duration: string(spt[0]),
					Title:    string(spt[1]),
					URL:      tsURL,
					Key:      key,
				})
			}
			break
		case bytes.HasPrefix(line, []byte(extXKey)):
			// 加密解密信息
			keyInfMap := parseLineKVMap(string(line))
			method := model.CryptMethod(keyInfMap["METHOD"])
			if method != "" && method != cryptMethodAES && method != cryptMethodNONE {
				err = errors.New("invalid EXT-X-KEY method: " + string(method))
				return
			}
			keyURI := util.BuildRealURL(res.URL, keyInfMap["URI"])
			key = &model.Key{
				Method: method,
				URI:    keyURI,
				IV:     keyInfMap["IV"],
			}
			break
		case bytes.HasPrefix(line, []byte(extXIFrameStreamInf)):
			// 指定一个包含多媒体信息的 media URI 作为PlayList
			infMap := parseLineKVMap(string(line))

			// 此时，下一行就是改播放源的 URL
			var playURL string
			if uri, ok := infMap[fieldURI]; ok {
				playURL = util.BuildRealURL(res.URL, uri)
			}

			res.PlayList = append(res.PlayList, model.PlayItem{
				ProgramID:  infMap[fieldProgramID],
				BandWidth:  infMap[fieldBandWidth],
				Resolution: infMap[fieldResolution],
				CodeCS:     infMap[fieldCodeCS],
				URL:        playURL,
			})
			break
		case bytes.HasPrefix(line, []byte(extXStreamInf)):
			// 指定一个包含多媒体信息的 media URI 作为PlayList
			infMap := parseLineKVMap(string(line))

			// 此时，下一行就是改播放源的 URL
			var playURL string
			if d.HasNextLine() {
				playURL = util.BuildRealURL(res.URL, string(d.Line()))
			}

			res.PlayList = append(res.PlayList, model.PlayItem{
				ProgramID:  infMap[fieldProgramID],
				BandWidth:  infMap[fieldBandWidth],
				Resolution: infMap[fieldResolution],
				CodeCS:     infMap[fieldCodeCS],
				URL:        playURL,
			})
			break
		case bytes.HasPrefix(line, []byte(extXEnd)):
			// 文档末尾了
			return
		}
	}
	return
}

// Line 读取当前行
func (d *Doc) Line() (line []byte) {
	line = d.Lines[d.LineIndex]
	return
}

// HasNextLine 判断是否有下一行，如果有，自动将游标移动到下一行
func (d *Doc) HasNextLine() bool {
	if d.LineIndex >= len(d.Lines)-1 {
		return false
	}
	d.LineIndex++
	return true
}

func parseLineKVMap(line string) map[string]string {
	r := linePattern.FindAllStringSubmatch(line, -1)
	params := make(map[string]string)
	for _, arr := range r {
		params[arr[1]] = strings.Trim(arr[2], "\"")
	}
	return params
}

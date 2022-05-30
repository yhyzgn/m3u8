// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2022-05-27 23:02
// version: 1.0.0
// desc   :

package parser

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/yhyzgn/goat/file"
	"github.com/yhyzgn/m3u8/model"
	"github.com/yhyzgn/m3u8/net"
	"io"
)

// FromFile 从文件一些
func FromFile(filename string) (m3u8 *model.M3U8, err error) {
	if !file.Exists(filename) {
		err = errors.New("no such file + '" + filename + "'")
	} else {
		bys, e := file.Read(filename)
		if nil != e {
			err = e
			return
		}
		m3u8, err = formBytes(bys, &model.M3U8{
			FileName: filename,
		})
	}
	return
}

// FromNetwork 从网络解析
func FromNetwork(url string) (m3u8 *model.M3U8, err error) {
	data, err := net.Get(url)
	if nil != err {
		return
	}
	m3u8, err = formBytes(data, &model.M3U8{
		URL: url,
	})
	return
}

// FromString 从文本字符串解析
func FromString(src string) (m3u8 *model.M3U8, err error) {
	return formBytes([]byte(src), &model.M3U8{})
}

// FromBytes 从二进制数据解析
func FromBytes(bys []byte) (m3u8 *model.M3U8, err error) {
	return formBytes(bys, &model.M3U8{})
}

func formBytes(bys []byte, m3u8 *model.M3U8) (res *model.M3U8, err error) {
	buff := bufio.NewReader(bytes.NewBuffer(bys))

	lines := make([][]byte, 0)
	for {
		line, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		// 跳过空行
		if 0 == len(line) || string(line) == "\n" || string(line) == "\r\n" {
			continue
		}
		lines = append(lines, []byte(string(line))) // 此处有诡异，直接将 line 添加进去最终得到的 lines 会混乱，需要 string 转一下才行
	}
	res, err = NewDoc(lines, m3u8).Parse()
	return
}

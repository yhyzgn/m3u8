// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2022-05-28 03:01
// version: 1.0.0
// desc   :

package util

import (
	"regexp"
	"strings"
)

func BuildRealURL(url, fileUrl string) string {
	switch true {
	case strings.HasPrefix(fileUrl, "http://") || strings.HasPrefix(fileUrl, "https://"):
		query := strings.Split(url, "?")
		// 如果原 fileUrl 中有参数，但分片中没有，很可能需要把参数加回分片地址中
		if !strings.Contains(fileUrl, "?") && len(query) > 1 {
			fileUrl += "?" + query[1]
		}
		// 绝对路径
		break
	case strings.HasPrefix(fileUrl, "/") && "" != url:
		// 域名根路径
		query := strings.Split(url, "?")
		urlSplit := strings.Split(query[0], "://")
		domain := strings.Split(urlSplit[1], "/")[0]
		// 如果原 fileUrl 中有参数，但分片中没有，很可能需要把参数加回分片地址中
		if !strings.Contains(fileUrl, "?") && len(query) > 1 {
			fileUrl += "?" + query[1]
		}
		fileUrl = urlSplit[0] + "://" + domain + fileUrl
		break
	case regexp.MustCompile("^[\\w+]").MatchString(fileUrl) && "" != url:
		// 域名根路径
		query := strings.Split(url, "?")
		lastIndex := strings.LastIndex(query[0], "/")
		root := url[0 : lastIndex+1]
		// 如果原 fileUrl 中有参数，但分片中没有，很可能需要把参数加回分片地址中
		if !strings.Contains(fileUrl, "?") && len(query) > 1 {
			fileUrl += "?" + query[1]
		}
		fileUrl = root + fileUrl
		break
	default:
		break
	}
	return fileUrl
}

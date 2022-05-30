// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2022-05-28 00:05
// version: 1.0.0
// desc   :

package model

type (
	// CryptMethod 加密方式
	CryptMethod string
)

// M3U8 m3u8 实体
type M3U8 struct {
	URL      string
	FileName string
	Version  int
	PlayList []PlayItem
	TsList   []TS
}

// PlayItem 播放列表条目
type PlayItem struct {
	ProgramID  string
	BandWidth  string
	Resolution string
	CodeCS     string
	URL        string
}

// TS 每个播放片段里的 ts 信息
type TS struct {
	Duration string
	Title    string
	URL      string
	Key      *Key
}

// Key 秘钥信息 #EXT-X-KEY:METHOD=AES-128,URI="key.key"
type Key struct {
	Method CryptMethod
	URI    string
	IV     string
}

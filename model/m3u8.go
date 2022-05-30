// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2022-05-28 00:05
// version: 1.0.0
// desc   :

package model

type (
	CryptMethod string
)

type M3U8 struct {
	URL      string
	FileName string
	Version  int
	PlayList []PlayItem
	TsList   []TS
}

type PlayItem struct {
	ProgramID  string
	BandWidth  string
	Resolution string
	CodeCS     string
	URL        string
}

type TS struct {
	Duration string
	Title    string
	URL      string
	Key      *Key
}

// Key #EXT-X-KEY:METHOD=AES-128,URI="key.key"
type Key struct {
	Method CryptMethod
	URI    string
	IV     string
}

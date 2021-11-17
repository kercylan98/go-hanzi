package hanzi

import (
	"github.com/kercylan98/go-pinyin/pinyin"
)

// 汉字结构
type hanzi struct {
	char   string           // 文字
	pinyin []*pinyin.Pinyin // 拼音信息
	rhyme  []string         // 韵脚（通常韵脚取韵母一位，如 chi : i , li xiang : i ang）
}

func (slf *hanzi) String() string {
	return slf.char
}

func (slf *hanzi) Rhyme() []string {
	return slf.rhyme
}

func (slf *hanzi) Pinyin() []*pinyin.Pinyin {
	return slf.pinyin
}

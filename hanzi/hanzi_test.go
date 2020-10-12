package hanzi

import (
	"fmt"
	"github.com/KercyLAN/secret-dimension-pinyin/pinyin"
	"strconv"
	"testing"
)

func TestTakeWords(t *testing.T) {
	hz := TakeWordsFull("今天是个好日子")
	for _, v := range hz {
		t.Log(v.String(), "(", v.Rhyme(), ")", pinyin.ToString(v.Pinyin()))
	}
}

func TestToHanzi(t *testing.T) {
	hz := ToHanzi("今天是个好日子")
	t.Log("源文本", hz.String())
	t.Log("拼音", pinyin.ToString(hz.Pinyin()))
	t.Log("拼音", pinyin.ToStringTone(hz.Pinyin()))
	t.Log("韵脚", hz.Rhyme())
}

func TestFull(t *testing.T) {
	hz := ToHanzi("您好，我是Xusixxx，这是一段测试代码")
	t.Log("源文本：", hz.String())
	t.Log("韵脚：", hz.Rhyme())
	// 输出拼音不带音调
	for _, py := range hz.Pinyin() {
		print("拼音：")
		fmt.Print(py.Source() + "(" + py.String() + ")")
	}
	fmt.Println()
	fmt.Println()
	// 输出拼音携带音调
	for _, pinyin := range hz.Pinyin() {
		print("拼音：")
		fmt.Print(pinyin.Source() + "(" + pinyin.StringTone() + ")")
	}
	fmt.Println()
	fmt.Println()
	// 输出拼音信息
	for _, py := range hz.Pinyin() {
		t.Log("声母：", py.Source(), py.Initial())
		t.Log("韵母：", py.Source(), py.Finals())
		switch py.Tone() {
		case 0:
			t.Log("音调：", py.Source(), "轻声汉字")
		default:
			t.Log("音调：", py.Source(), "第"+strconv.Itoa(py.Tone())+"声")
		}
	}
	fmt.Println()
	// 输出错误信息
	for _, py := range hz.Pinyin() {
		t.Log("错误信息：", py.Source(), py.Err())
	}
}

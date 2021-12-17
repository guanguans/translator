package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/guanguans/translator/baidu"
	"github.com/guanguans/translator/youdao"
)

func main() {
	baidu := baidu.New("appid", "appSecret")
	spew.Dump(baidu.Translate("你好", "zh", "en"))

	youdao := youdao.New("appid", "appSecret")
	spew.Dump(youdao.Translate("你好", "zh-CHS", "en"))
}

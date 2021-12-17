package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/guanguans/translator/baidu"
)

func main() {
	baidu := baidu.New("appid", "appSecret")

    // (*resty.Response)(0xc00008e0f0)({"from":"zh","to":"en","trans_result":[{"src":"\u4f60\u597d","dst":"Hello"}]})
    // (interface {}) <nil>
	spew.Dump(baidu.Translate("你好", "zh", "en"))
}

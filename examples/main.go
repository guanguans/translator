package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/guanguans/translator"
	"github.com/guanguans/translator/driver/baidu"
	"github.com/guanguans/translator/driver/youdao"
)

func main() {
	// register one(or some) translator driver
	translator.Register(baidu.Name, baidu.New("appid", "appSecret"))
	translator.Register(youdao.Name, youdao.New("appid", "appSecret"))

	// setting default driver name
	translator.DefaultUse(baidu.Name)

	// quick use.(it is default driver)
	spew.Dump(translator.Translate("你好", "zh", "en"))

	// use alone
	spew.Dump(baidu.New("appid", "appSecret").Translate("你好", "zh", "en"))
	spew.Dump(youdao.New("appid", "appSecret").Translate("你好", "zh-CHS", "en"))
}

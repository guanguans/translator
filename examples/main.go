package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/guanguans/translator"
	"github.com/guanguans/translator/client"
)

func main() {
	// register one(or some) translator driver
	translator.Register(client.BaiduName, client.NewBaidu("appid", "appSecret"))
	translator.Register(client.YoudaoName, client.NewYoudao("appid", "appSecret"))

	// setting default driver name
	translator.DefaultUse(client.BaiduName)

	// quick use.(it is default driver)
	spew.Dump(translator.Translate("你好", "zh", "en"))

	// use alone
	spew.Dump(client.NewBaidu("appid", "appSecret").Translate("你好", "zh", "en"))
	spew.Dump(client.NewBaidu("appid", "appSecret").Translate("你好", "zh-CHS", "en"))
}

# translator

[ENGLISH](README-EN.md) | [简体中文](README.md)

> Translation SDK(baidu、youdao). - 翻译 SDK(百度翻译、有道翻译)。

[![tests](https://github.com/guanguans/translator/workflows/tests/badge.svg)](https://github.com/guanguans/translator/actions)
[![gocover.io](https://gocover.io/_badge/github.com/guanguans/translator)](https://gocover.io/github.com/guanguans/translator)
[![Go Report Card](https://goreportcard.com/badge/github.com/guanguans/translator)](https://goreportcard.com/report/github.com/guanguans/translator)
[![GoDoc](https://godoc.org/github.com/guanguans/translator?status.svg)](https://godoc.org/github.com/guanguans/translator)
[![GitHub license](https://img.shields.io/github/license/guanguans/translator.svg)](https://github.com/guanguans/translator/blob/master/LICENSE)
[![Open in Visual Studio Code](https://open.vscode.dev/badges/open-in-vscode.svg)](https://open.vscode.dev/guanguans/translator)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/guanguans/translator)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guanguans/translator)
![GitHub repo size](https://img.shields.io/github/repo-size/guanguans/translator)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/guanguans/translator)

## 平台支持

* [百度翻译](http://api.fanyi.baidu.com/api/trans/product/apidoc)
* [有道翻译](http://ai.youdao.com/docs/doc-trans-api.s#p02)

## 安装

```shell script
$ go get -u github.com/guanguans/translator
```

## 使用

这只是一个快速介绍, 请查看 [GoDoc](https://godoc.org/github.com/guanguans/translator) 获得详细信息。

让我们从一个简单的例子开始：

```go
package main

import (
    "github.com/davecgh/go-spew/spew"
    "github.com/guanguans/translator/driver/baidu"
    "github.com/guanguans/translator/driver/youdao"
)

func main() {
    spew.Dump(baidu.New("appid", "appSecret").Translate("你好", "zh", "en"))
    spew.Dump(youdao.New("appid", "appSecret").Translate("你好", "zh-CHS", "en"))
}
```

## 测试

```shell script
$ make test
$ make bench
```

## 变更日志

请参阅 [CHANGELOG](CHANGELOG.md) 获取最近有关更改的更多信息。

## 贡献指南

请参阅 [CONTRIBUTING](.github/CONTRIBUTING.md) 有关详细信息。

## 安全漏洞

请查看[我们的安全政策](../../security/policy)了解如何报告安全漏洞。

## 贡献者

* [guanguans](https://github.com/guanguans)
* [所有贡献者](../../contributors)

## 协议

MIT 许可证（MIT）。有关更多信息，请参见[协议文件](LICENSE)。

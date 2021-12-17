# go-translator

[ENGLISH](README.md) | [简体中文](README-zh_CN.md)

> A Go package template repository. - 一个 Go 软件包模板存储库。

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

## 功能

* package template

## 环境要求

* Go >= 1.11

## 安装

``` shell script
$ go get -u github.com/guanguans/translator
```

## 使用

这只是一个快速介绍, 请查看 [GoDoc](https://godoc.org/github.com/guanguans/translator) 获得详细信息。

让我们从一个简单的例子开始：

``` go
package main

import (
	"github.com/guanguans/translator"
	"gopkg.in/ffmt.v1"
)

func main() {
	ffmt.P(translator.ReturnSelf("go-translator"))
}
```

## 测试

``` bash
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

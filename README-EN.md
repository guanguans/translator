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

## Platform

* [baidu](http://api.fanyi.baidu.com/api/trans/product/apidoc)
* [youdao](http://ai.youdao.com/docs/doc-trans-api.s#p02)

## Installation

```shell script
$ go get -u github.com/guanguans/translator
```

## Usage

This is just a quick introduction, view the [GoDoc](https://godoc.org/github.com/guanguans/translator) for details.

Let's start with a trivial example:

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

## Testing

```shell script
$ make test
$ make bench
```

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information on what has changed recently.

## Contributing

Please see [CONTRIBUTING](.github/CONTRIBUTING.md) for details.

## Security Vulnerabilities

Please review [our security policy](../../security/policy) on how to report security vulnerabilities.

## Credits

* [guanguans](https://github.com/guanguans)
* [All Contributors](../../contributors)

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.

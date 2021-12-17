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

## Features

* package template

## Requirement

* Go >= 1.11

## Installation

``` shell script
$ go get -u github.com/guanguans/translator
```

## Usage

This is just a quick introduction, view the [GoDoc](https://godoc.org/github.com/guanguans/translator) for details.

Let's start with a trivial example:

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

## Testing

``` bash
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

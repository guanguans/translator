package main

import (
	"github.com/awesee/php2go/php"
	"github.com/davecgh/go-spew/spew"
	"github.com/guanguans/translator"
	"github.com/spf13/cast"
	"github.com/syyongx/php2go"
	"gopkg.in/ffmt.v1"
)

func main() {
	// ffmt
	ffmt.P(translator.ReturnSelf("go-translator"))

	// davecgh/go-spew
	spew.Dump("go-translator")

	// spf13/cast
	ffmt.P(cast.ToInt("go-translator"))

	// syyongx/php2go
	ffmt.P(php.Md5("go-translator"))

	// awesee/php2go
	ffmt.P(php2go.Md5("go-translator"))
}

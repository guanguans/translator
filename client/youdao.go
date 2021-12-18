package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/guanguans/translator/pkg"
)

const YoudaoName = "youdao"

// Youdao http://ai.youdao.com/docs/doc-trans-api.s#p02
type Youdao struct {
	AppID      string
	AppSecret  string
	HttpClient *resty.Client
}

func NewYoudao(appID string, appSecret string) *Youdao {
	httpClient := resty.New().SetBaseURL(YoudaoBaseURL)
	return &Youdao{
		AppID:      appID,
		AppSecret:  appSecret,
		HttpClient: httpClient,
	}
}

func (y *Youdao) Translate(q string, from string, to string) (*resty.Response, error) {
	return y.HttpClient.R().
		SetFormData(y.buildRequestParams(q, from, to)).
		Post("")
}

func (y *Youdao) buildRequestParams(q string, from string, to string) map[string]string {
	salt := pkg.RandomStr(RandomStrLen)
	return map[string]string{
		"q":      q,
		"from":   from,
		"to":     to,
		"appKey": y.AppID,
		"salt":   salt,
		"ext":    "mp3",
		"voice":  "0",
		"sign":   pkg.Md5(y.AppID + q + salt + y.AppSecret),
	}
}

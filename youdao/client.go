package youdao

import (
	"github.com/go-resty/resty/v2"
	"github.com/guanguans/translator/util"
)

const BaseURL = "https://openapi.youdao.com/api"

// http://ai.youdao.com/docs/doc-trans-api.s#p02
type youdao struct {
	AppID      string
	AppSecret  string
	HttpClient *resty.Client
}

func New(appID string, appSecret string) *youdao {
	httpClient := resty.New().SetBaseURL(BaseURL)

	return &youdao{
		AppID:      appID,
		AppSecret:  appSecret,
		HttpClient: httpClient,
	}
}

func (y *youdao) Translate(q string, from string, to string) (*resty.Response, error) {
	return y.HttpClient.R().
		SetFormData(y.buildRequestParams(q, from, to)).
		Post("")
}

func (y *youdao) buildRequestParams(q string, from string, to string) map[string]string {
	salt := util.RandomStr(1 << 3)

	return map[string]string{
		"q":      q,
		"from":   from,
		"to":     to,
		"appKey": y.AppID,
		"salt":   salt,
		"ext":    "mp3",
		"voice":  "0",
		"sign":   util.Md5(y.AppID + q + salt + y.AppSecret),
	}
}

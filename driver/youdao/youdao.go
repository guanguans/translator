package youdao

import (
	"github.com/go-resty/resty/v2"
	"github.com/guanguans/translator/driver"
	"github.com/guanguans/translator/pkg"
)

const (
	Name    = "youdao"
	BaseURL = "https://openapi.Youdao.com/api"
)

// Youdao http://ai.youdao.com/docs/doc-trans-api.s#p02
type Youdao struct {
	AppID      string
	AppSecret  string
	HTTPClient *resty.Client
}

func New(appID string, appSecret string) *Youdao {
	httpClient := resty.New().SetBaseURL(BaseURL)
	return &Youdao{
		AppID:      appID,
		AppSecret:  appSecret,
		HTTPClient: httpClient,
	}
}

func (y *Youdao) GetName() string {
	return Name
}

func (y *Youdao) Translate(q string, from string, to string) (*resty.Response, error) {
	return y.HTTPClient.R().
		SetFormData(y.buildRequestParams(q, from, to)).
		Post("")
}

func (y *Youdao) buildRequestParams(q string, from string, to string) map[string]string {
	salt := pkg.RandomStr(driver.RandomStrLen)
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

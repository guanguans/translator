package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/guanguans/translator/pkg"
)

const BaiduName = "baidu"

// Baidu http://api.fanyi.baidu.com/api/trans/product/apidoc
type Baidu struct {
	AppID      string
	AppSecret  string
	HttpClient *resty.Client
}

func NewBaidu(appID string, appSecret string) *Baidu {
	httpClient := resty.New().
		SetBaseURL(BaiduBaseURL).
		SetHeader("Content-Type", "application/x-www-form-urlencoded")
	return &Baidu{
		AppID:      appID,
		AppSecret:  appSecret,
		HttpClient: httpClient,
	}
}

func (c *Baidu) Translate(q string, from string, to string) (*resty.Response, error) {
	return c.HttpClient.R().
		SetFormData(c.buildRequestParams(q, from, to)).
		Post("")
}

func (c *Baidu) buildRequestParams(q string, from string, to string) map[string]string {
	salt := pkg.RandomStr(RandomStrLen)
	return map[string]string{
		"q":     q,
		"from":  from,
		"to":    to,
		"appid": c.AppID,
		"salt":  salt,
		"sign":  pkg.Md5(c.AppID + q + salt + c.AppSecret),
	}
}

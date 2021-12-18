package baidu

import (
	"github.com/go-resty/resty/v2"
	"github.com/guanguans/translator/driver"
	"github.com/guanguans/translator/pkg"
)

const (
	Name    = "baidu"
	BaseURL = "https://fanyi-api.baidu.com/api/trans/vip/translate"
)

// Baidu http://api.fanyi.baidu.com/api/trans/product/apidoc
type Baidu struct {
	AppID      string
	AppSecret  string
	HTTPClient *resty.Client
}

func New(appID string, appSecret string) *Baidu {
	httpClient := resty.New().
		SetBaseURL(BaseURL).
		SetHeader("Content-Type", "application/x-www-form-urlencoded")
	return &Baidu{
		AppID:      appID,
		AppSecret:  appSecret,
		HTTPClient: httpClient,
	}
}

func (b *Baidu) GetName() string {
	return Name
}

func (b *Baidu) Translate(q string, from string, to string) (*resty.Response, error) {
	return b.HTTPClient.R().
		SetFormData(b.buildRequestParams(q, from, to)).
		Post("")
}

func (b *Baidu) buildRequestParams(q string, from string, to string) map[string]string {
	salt := pkg.RandomStr(driver.RandomStrLen)
	return map[string]string{
		"q":     q,
		"from":  from,
		"to":    to,
		"appid": b.AppID,
		"salt":  salt,
		"sign":  pkg.Md5(b.AppID + q + salt + b.AppSecret),
	}
}

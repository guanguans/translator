package baidu

import (
	"github.com/go-resty/resty/v2"
	"github.com/guanguans/translator/util"
)

const (
	BaseURL      = "https://fanyi-api.baidu.com/api/trans/vip/translate"
	RandomStrLen = 1 << 3
)

// Client http://api.fanyi.baidu.com/api/trans/product/apidoc
type Client struct {
	AppID      string
	AppSecret  string
	HttpClient *resty.Client
}

func New(appID string, appSecret string) *Client {
	httpClient := resty.New().
		SetBaseURL(BaseURL).
		SetHeader("Content-Type", "application/x-www-form-urlencoded")

	return &Client{
		AppID:      appID,
		AppSecret:  appSecret,
		HttpClient: httpClient,
	}
}

func (c *Client) Translate(q string, from string, to string) (*resty.Response, error) {
	return c.HttpClient.R().
		SetFormData(c.buildRequestParams(q, from, to)).
		Post("")
}

func (c *Client) buildRequestParams(q string, from string, to string) map[string]string {
	salt := util.RandomStr(RandomStrLen)

	return map[string]string{
		"q":     q,
		"from":  from,
		"to":    to,
		"appid": c.AppID,
		"salt":  salt,
		"sign":  util.Md5(c.AppID + q + salt + c.AppSecret),
	}
}

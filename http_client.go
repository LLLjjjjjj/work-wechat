package work

import (
	"bytes"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/url"
)

var HttpClient = NewHttpClient()

type httpClient struct {
	client *ghttp.Client
}

func NewHttpClient() *httpClient {
	return &httpClient{
		client: g.Client(),
	}
}

// DefaultQYAPIHost 默认企业微信 API Host
const DefaultQYAPIHost = "https://qyapi.weixin.qq.com"

func (h httpClient) httpGet(path string, req urlValuer) (string, error) {

	url := h.composeQyapiURL(path, req)

	resp, err := h.
		client.
		SetHeader("Content-Type", "application/json").
		Get(url.String())

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp == nil {
		return "", err
	}

	return resp.ReadAllString(), nil
}

func (h httpClient) httpPost(path string, req bodyer) (string, error) {

	url := h.composeQyapiURL(path, req)
	urlStr := url.String()

	body, err := req.intoBody()
	if err != nil {
		return "", err
	}

	resp, err := h.client.Post(urlStr, "application/json", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp == nil {
		return "", err
	}

	return resp.ReadAllString(), nil
}

// 将请求结构体和请求路径拼接成统一的url
func (h httpClient) composeQyapiURL(path string, req interface{}) *url.URL {
	values := url.Values{}
	if valuer, ok := req.(urlValuer); ok {
		values = valuer.intoURLValues()
	}

	base, err := url.Parse(DefaultQYAPIHost)
	if err != nil {
		panic(fmt.Sprintf("qyapiHost invalid: host=%s err=%+v", DefaultQYAPIHost, err))
	}

	base.Path = path
	base.RawQuery = values.Encode()

	return base
}

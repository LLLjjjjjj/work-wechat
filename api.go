/*
@Time : 2021/7/6 3:46 下午
@Author : 21
@File : api
@Software: GoLand
*/
package work

import (
	"context"
	"errors"
	"github.com/gogf/gf/frame/g"
	"net/http"
	"time"
)

type HTTPOption func(s *httpSettings)

const (
	HttpGet  = "GET"
	HttpPost = "POST"
)

// httpSettings http request options
type httpSettings struct {
	headers map[string]string
	cookies []*http.Cookie
	close   bool
	timeout time.Duration
}

type HTTPMethod string

//  企业微信 api 结构体
type weWorkApi struct {
	reqURL string
	method HTTPMethod
	//query      url.Values
	//wxml       func(appid, mchid, nonce string) (WXML, error)
	body func() ([]byte, error)
	//uploadForm UploadForm
	//decode     func(resp []byte) error
	timeout time.Duration
}
type ActionOption func(w *weWorkApi)

func (w weWorkApi) GetRequestUrl() string {
	return w.reqURL
}

func (w weWorkApi) GetRequestBody() ([]byte, error) {
	return w.body()
}

func (w weWorkApi) GetHttpMethod() HTTPMethod {
	return w.method
}
func (w weWorkApi) GetTimeOut() time.Duration {
	return w.timeout
}

func (w weWorkApi) DoRequest(ctx context.Context) ([]byte, error) {
	requestUrl := w.GetRequestUrl()
	if len(requestUrl) < 1 {
		return nil, errors.New("设置http请求有误，请设置httpUrl后重试")
	}

	method := w.GetHttpMethod()

	if len(method) < 1 {
		return nil, errors.New("设置httpMethod请求有误，设置httpMethod请求有误后重试")
	}
	httpClient := g.Client()
	// 超时包
	httpClient.SetCtx(ctx)
	// 超时时间
	timeOut := w.GetTimeOut()
	httpClient.SetTimeout(timeOut)

	switch method {
	case HttpGet:
		r, err := httpClient.Get(requestUrl)
		if err != nil {
			return nil, err
		}
		defer r.Close()
		return r.ReadAll(), nil
	case HttpPost:
		body, err := w.GetRequestBody()
		if err != nil {
			return nil, err
		}
		r, err := httpClient.Post(requestUrl, body)
		if err != nil {
			return nil, err
		}
		defer r.Close()
		return r.ReadAll(), nil
	}

	return nil, errors.New("获取类型出错")
}

// 设置方法
func WitchMethod(method HTTPMethod) ActionOption {
	return func(w *weWorkApi) {
		w.method = method
	}
}

// 超时时间
func WitchTimeOut(timeOut time.Duration) ActionOption {
	return func(w *weWorkApi) {
		w.timeout = timeOut
	}
}

// 设置请求体
func WitchBody(bodyFunc func() (bytes []byte, e error)) ActionOption {
	return func(w *weWorkApi) {
		w.body = bodyFunc
	}
}

func NewWeWordApi(reqUrl string, opts ...ActionOption) Action {
	defaultWeWorkApi := weWorkApi{
		reqURL: reqUrl,
		method: HttpPost,
		body: func() (bytes []byte, e error) {
			return nil, nil
		},
		// 默认超时时间为 3s
		timeout: time.Duration(time.Second * 3),
	}
	for _, v := range opts {
		v(&defaultWeWorkApi)
	}
	return defaultWeWorkApi
}

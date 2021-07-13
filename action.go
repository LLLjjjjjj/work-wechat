/*
@Time : 2021/7/6 3:10 下午
@Author : 21
@File : action
@Software: GoLand
*/
package work

import (
	"context"
	"time"
)

type Action interface {
	// 获取 设置请求url
	GetRequestUrl() string

	// 获取 设置请求的 body
	GetRequestBody() ([]byte, error)

	// 真正的发出请求
	DoRequest(ctx context.Context) ([]byte, error)

	// 获取httpMethod
	GetHttpMethod() HTTPMethod

	// 获取超时时间
	GetTimeOut() time.Duration

	// 获取头部
	GetHeader() map[string]string
}

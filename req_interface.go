package work

import "net/url"

// urlValuer 可转化为 url.Values 类型的 trait
type urlValuer interface {
	intoURLValues() url.Values
}

// bodyer 可转化为 API 请求体的 trait
type bodyer interface {
	intoBody() ([]byte, error)
}


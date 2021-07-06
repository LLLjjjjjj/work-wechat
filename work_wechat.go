package work

import "context"

type Options func(w *workWechat) *workWechat

type workWechat struct {
	// 服务商的corpid
	ProviderCorpID string

	// 服务商的secret，在服务商管理后台可见
	ProviderSecret string

	// 以ww或wx开头应用id
	SuiteID string

	// 应用secret
	SuiteSecret string

	// suite_ticket
	SuiteTicket string

	// 授权企业corpId
	CorpId string

	// 企业永久授权码
	PermanentCode string
}

var defaultWorkWechat = workWechat{
	ProviderCorpID: "",
	ProviderSecret: "",
	SuiteID:        "",
	SuiteSecret:    "",
	SuiteTicket:    "",
	CorpId:         "",
	PermanentCode:  "",
}

func SetProviderCorpID(ProviderCorpID string) Options {
	return func(w *workWechat) *workWechat {
		w.ProviderCorpID = ProviderCorpID
		return w
	}
}

// todo
func SetProviderSecret() {

}

// todo 选项设计模式
func NewWeWork(opts ...Options) *workWechat {
	defaultWorkInfo := defaultWorkWechat
	for _, v := range opts {
		v(&defaultWorkInfo)
	}
	return &defaultWorkInfo
}

func NewWorkWechat(config Config) *workWechat {
	return &workWechat{
		ProviderCorpID: config.ProviderCorpID,
		ProviderSecret: config.ProviderSecret,
		SuiteID:        config.SuiteID,
		SuiteSecret:    config.SuiteSecret,
		SuiteTicket:    config.SuiteTicket,
		CorpId:         config.CorpId,
		PermanentCode:  config.PermanentCode,
	}
}

// 返回原包数据  直接进行链式炒作
func (w workWechat) Do(ctx context.Context, weWorkAction Action) ([]byte, error) {
	//todo  context包超时 释放原则
	return weWorkAction.DoRequest()
}

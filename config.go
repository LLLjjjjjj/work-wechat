package work

type Config struct {
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

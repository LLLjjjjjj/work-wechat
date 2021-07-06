package work

type AuthInfo struct {
	Agent []Agent `json:"agent"`
}

// 授权的应用信息，注意是一个数组，但仅旧的多应用套件授权时会返回多个agent，对新的单应用授权，永远只返回一个agent
type Agent struct {
	AgentID       int        `json:"agentid"`
	Name          string     `json:"name"`
	SquareLogoUrl string     `json:"square_logo_url"`
	RoundLogoUrl  string     `json:"round_logo_url"`
	AppID         int        `json:"appid"`
	AuthMode      int        `json:"auth_mode"`
	Privilege     Privilege  `json:"privilege"`
	SharedFrom    SharedFrom `json:"shared_from"`
}

// 应用对应的权限
type Privilege struct {
	AllowParty []int    `json:"allow_party"`
	AllowTag   []int    `json:"allow_tag"`
	AllowUser  []string `json:"allow_user"`
	ExtraParty []int    `json:"extra_party"`
	ExtraUser  []string `json:"extra_user"`
	ExtraTag   []int    `json:"extra_tag"`
	Level      int      `json:"level"`
}

// 共享了应用的互联企业信息，仅当由互联的企业共享应用触发的安装时才返回
type SharedFrom struct {
	CorpID string `json:"corpid"`
}

// 授权公司信息
type AuthCorpInfo struct {
	CorpID            string `json:"corpid"`
	CorpName          string `json:"corp_name"`
	CorpType          string `json:"corp_type"`
	CorpSquareLogoUrl string `json:"corp_square_logo_url"`
	CorpUserMax       int    `json:"corp_user_max"`
	CorpAgentMax      int    `json:"corp_agent_max"`
	CorpFullName      string `json:"corp_full_name"`
	VerifiedEndTime   int    `json:"verified_end_time"`
	SubjectType       int    `json:"subject_type"`
	CorpWxQrcode      string `json:"corp_wxqrcode"`
	CorpScale         string `json:"corp_scale"`
	CorpIndustry      string `json:"corp_industry"`
	CorpSubIndustry   string `json:"corp_sub_industry"`
	Location          string `json:"location"`
}

// 代理服务商企业信息
type DealerCorpInfo struct {
	CorpID   string `json:"corpid"`
	CorpName string `json:"corp_name"`
}

// 推广二维码安装相关信息，扫推广二维码安装时返回。（注：无论企业是否新注册，只要通过扫推广二维码安装，都会返回该字段）
type RegisterCodeInfo struct {
	RegisterCode string `json:"register_code"`
	TemplateId   string `json:"template_id"`
	State        string `json:"state"`
}

// 授权管理员的信息，可能不返回（企业互联由上级企业共享第三方应用给下级时，不返回授权的管理员信息）
type AuthUserInfo struct {
	UserID     string `json:"userid"`
	OpenUserID string `json:"open_userid"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
}

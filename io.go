package work

import (
	"encoding/json"
	"net/url"
)

type respCommon struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

/**
获取服务商凭证校验
 */
// reqGetProviderToken 获取服务商凭证校验
type reqGetProviderToken struct {
	CorpId         string `json:"corpid"`
	ProviderSecret string `json:"provider_secret"`
}

var _ bodyer = reqGetProviderToken{}

func (x reqGetProviderToken) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// respGetProviderToken 获取服务商凭证校验
type respGetProviderToken struct {
	respCommon
	ProviderToken
}

// ProviderToken 服务商凭证
type ProviderToken struct {
	ProviderAccessToken string `json:"provider_access_token"`
	ExpiresIn           int    `json:"expires_in"`
}


/**
获取第三方应用凭证校验
*/
// reqGetSuiteToken 获取第三方应用凭证校验
type reqGetSuiteToken struct {
	SuiteID     string `json:"suite_id"`
	SuitSecret  string `json:"suite_secret"`
	SuiteTicket string `json:"suite_ticket"`
}

var _ bodyer = reqGetSuiteToken{}

func (x reqGetSuiteToken) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// respGetSuiteToken 获取第三方应用凭证校验
type respGetSuiteToken struct {
	respCommon
	SuiteToken
}

// SuiteToken 第三方应用凭证
type SuiteToken struct {
	SuiteAccessToken string `json:"suite_access_token"`
	ExpiresIn        int    `json:"expires_in"`
}

/**
获取企业凭证校验
 */
// reqGetCorpToken 获取企业凭证校验
type reqGetCorpToken struct {
	AuthCorpID    string `json:"auth_corpid"`
	PermanentCode string `json:"permanent_code"`
}

var _ bodyer = reqGetCorpToken{}

func (x reqGetCorpToken) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// respGetSuiteToken 获取企业凭证校验
type respGetCorpToken struct {
	respCommon
	CorpToken
}

// CorpToken 企业凭证
type CorpToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}


/**
获取企业授权码
 */
// reqGetPreAuthCode 获取预授权码校验
type reqGetPreAuthCode struct {
	SuiteAccessToken string `json:"suite_access_token"`
}

var _ urlValuer = reqGetPreAuthCode{}

func (x reqGetPreAuthCode) intoURLValues() url.Values {
	return url.Values{
		"suite_access_token": {x.SuiteAccessToken},
	}
}

// respGetPreAuthCode 获取预授权码校验
type respGetPreAuthCode struct {
	respCommon
	PreAuthCode
}

// PreAuthCode 预授权码
type PreAuthCode struct {
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
}

/**
永久授权码
 */
// reqGetPermanentCode 获取永久授权码校验
type reqGetPermanentCode struct {
	AuthCode string `json:"auth_code"`
}

var _ bodyer = reqGetPermanentCode{}

func (x reqGetPermanentCode) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// respGetPermanentCode 永久授权码校验
type respGetPermanentCode struct {
	respCommon
	PermanentCode
}

// PermanentCode 永久授权码
type PermanentCode struct {
	AccessToken      string           `json:"access_token"`
	ExpiresIn        int              `json:"expires_in"`
	PermanentCode    string           `json:"permanent_code"`
	AuthCorpInfo     AuthCorpInfo     `json:"auth_corp_info"`
	AuthInfo         AuthInfo         `json:"auth_info"`
	AuthUserInfo     AuthUserInfo     `json:"auth_user_info"`
	DealerCorpInfo   DealerCorpInfo   `json:"dealer_corp_info"`
	RegisterCodeInfo RegisterCodeInfo `json:"register_code_info"`
}
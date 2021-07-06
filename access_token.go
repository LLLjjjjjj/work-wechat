package work

import (
	"context"
	"errors"
	"github.com/arden/easy"
	"github.com/arden/easy/database/redis"
	"github.com/gogf/gf/util/gconv"
)

// 服务商token 后面拼接服务商id
const ProviderAccessTokenRedisKey = "hdcj:provider_access_token:"

// 第三方应用token 后面拼接应用id
const SuitAccessTokenRedisKey = "hdcj:provider_access_token:"

// 第三方应用ticket 后面拼接应用id
const SuitTicketRedisKey = "hdcj:suit_ticket:"

// 授权的企业token 后面拼接企业id
const CorpAccessTokenRedisKey = "hdcj:provider_access_token:"

type accessToken struct {
	workWechat workWechat
	cache      *redis.Redis
}

func (w workWechat) NewAccessToken() *accessToken {
	return &accessToken{
		w,
		easy.Redis(),
	}
}

// getProviderAccessToken 获取服务商access_token
func (a *accessToken) getProviderAccessToken() string {
	return a.cache.Get(context.Background(), ProviderAccessTokenRedisKey+a.workWechat.ProviderCorpID).Val()
}

// getSuitAccessToken 获取应用access_token
func (a *accessToken) getSuitAccessToken() string {
	return a.cache.Get(context.Background(), SuitAccessTokenRedisKey+a.workWechat.SuiteID).Val()
}

// getSuiteTicket 获取应用ticket
func (a *accessToken) getSuiteTicket() string {
	return a.cache.Get(context.Background(), SuitTicketRedisKey+a.workWechat.SuiteID).Val()
}

// getCorpAccessToken 获取企业access_token
func (a *accessToken) getCorpAccessToken() string {
	return a.cache.Get(context.Background(), CorpAccessTokenRedisKey+a.workWechat.CorpId).Val()
}

func (a *accessToken) providerAccessToken() (respGetProviderToken, error) {
	var req = reqGetProviderToken{
		CorpId:         a.workWechat.ProviderCorpID,
		ProviderSecret: a.workWechat.ProviderSecret,
	}
	var respGetProviderToken = respGetProviderToken{}

	resp, err := HttpClient.httpPost("/cgi-bin/service/get_provider_token", req)
	if err != nil {
		return respGetProviderToken, err
	}

	if resp == "" {
		return respGetProviderToken, errors.New("请求错误")
	}

	err = gconv.Struct(resp, respGetProviderToken)
	if err != nil {
		return respGetProviderToken, err
	}

	return respGetProviderToken, nil
}

func (a *accessToken) suitAccessToken() (reqGetSuiteToken, error) {
	var req = reqGetSuiteToken{
		SuiteID:     a.workWechat.SuiteID,
		SuitSecret:  a.workWechat.SuiteSecret,
		SuiteTicket: a.workWechat.SuiteTicket,
	}
	var reqGetSuiteToken = reqGetSuiteToken{}

	resp, err := HttpClient.httpPost("/cgi-bin/service/get_suite_token", req)
	if err != nil {
		return reqGetSuiteToken, err
	}

	if resp == "" {
		return reqGetSuiteToken, errors.New("请求错误")
	}

	err = gconv.Struct(resp, reqGetSuiteToken)
	if err != nil {
		return reqGetSuiteToken, err
	}

	return reqGetSuiteToken, nil
}

func (a *accessToken) corpAccessToken() (reqGetCorpToken, error) {
	var req = reqGetCorpToken{
		AuthCorpID:    a.workWechat.CorpId,
		PermanentCode: a.workWechat.PermanentCode,
	}
	var reqGetCorpToken = reqGetCorpToken{}

	resp, err := HttpClient.httpPost("/cgi-bin/service/get_suite_token?suite_access_token="+a.getSuitAccessToken(), req)
	if err != nil {
		return reqGetCorpToken, err
	}

	if resp == "" {
		return reqGetCorpToken, errors.New("请求错误")
	}

	err = gconv.Struct(resp, reqGetCorpToken)
	if err != nil {
		return reqGetCorpToken, err
	}

	return reqGetCorpToken, nil
}

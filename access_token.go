package work

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/arden/easy"
	"github.com/arden/easy/database/redis"
	"time"
)

// 服务商token 后面拼接服务商id
const ProviderAccessTokenRedisKey = "hdcj:provider_access_token:"

// 第三方应用token 后面拼接应用id
const SuiteAccessTokenRedisKey = "hdcj:suite_access_token:"

// 第三方应用ticket 后面拼接应用id
const SuiteTicketRedisKey = "hdcj:suite_ticket:"

// 授权的企业token 后面拼接企业id
const CorpAccessTokenRedisKey = "hdcj:corp_access_token:"

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

// GetProviderAccessTokenByCache 获取服务商access_token
func (a *accessToken) GetProviderAccessTokenByCache() string {
	accessToken := a.cache.Get(context.Background(), ProviderAccessTokenRedisKey+a.workWechat.ProviderCorpID).Val()
	if accessToken == "" {
		accessTokenResp, err := a.GetProviderAccessToken()
		if err != nil {
			return ""
		}
		a.cache.Set(context.Background(), ProviderAccessTokenRedisKey+a.workWechat.ProviderCorpID, accessTokenResp.ProviderAccessToken, time.Second * time.Duration(accessTokenResp.ExpiresIn-200))
		return accessTokenResp.ProviderAccessToken
	}
	return accessToken
}

// GetSuiteAccessTokenByCache 获取应用access_token
func (a *accessToken) GetSuiteAccessTokenByCache() string {
	accessToken := a.cache.Get(context.Background(), SuiteAccessTokenRedisKey+a.workWechat.SuiteID).Val()
	if accessToken == "" {
		accessTokenResp, err := a.GetSuiteAccessToken()
		if err != nil {
			return ""
		}
		a.cache.Set(context.Background(), SuiteAccessTokenRedisKey+a.workWechat.SuiteID, accessTokenResp.SuiteAccessToken, time.Second * time.Duration(accessTokenResp.ExpiresIn-200))
		return accessTokenResp.SuiteAccessToken
	}
	return accessToken
}

// GetSuiteTicketByCache 获取应用ticket
func (a *accessToken) GetSuiteTicketByCache() string {
	return a.cache.Get(context.Background(), SuiteTicketRedisKey+a.workWechat.SuiteID).Val()
}

// GetCorpAccessTokenByCache 获取企业access_token
func (a *accessToken) GetCorpAccessTokenByCache() string {
	accessToken := a.cache.Get(context.Background(), CorpAccessTokenRedisKey+a.workWechat.CorpId).Val()
	if accessToken == "" {
		accessTokenResp, err := a.GetCorpAccessToken()
		if err != nil {
			return ""
		}
		a.cache.Set(context.Background(), CorpAccessTokenRedisKey+a.workWechat.CorpId, accessTokenResp.AccessToken, time.Second * time.Duration(accessTokenResp.ExpiresIn-200))
		return accessTokenResp.AccessToken
	}
	return accessToken
}

func GetProviderAccessTokenAction(providerCorpId string, providerSecret string) Action {
	reqUrl := BaseWeWorkUrl + "/cgi-bin/service/get_provider_token"
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			var req = reqGetProviderToken{
				CorpId:         providerCorpId,
				ProviderSecret: providerSecret,
			}
			jsonInfo, err := json.Marshal(req)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

func GetSuitAccessTokenAction(suiteId string, suiteSecret string, suiteTicket string) Action {
	reqUrl := BaseWeWorkUrl + "/cgi-bin/service/get_suite_token"
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			var req = reqGetSuiteToken{
				SuiteID:     suiteId,
				SuitSecret:  suiteSecret,
				SuiteTicket: suiteTicket,
			}
			jsonInfo, err := json.Marshal(req)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

func GetCorpAccessTokenAction(suitAccessToken string, corpId string, permanentCode string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/service/get_suite_token?suite_access_token=%s", suitAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			var req = reqGetCorpToken{
				AuthCorpID:    corpId,
				PermanentCode: permanentCode,
			}
			jsonInfo, err := json.Marshal(req)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

/**
 * @Description:
 * @author:21
 * @receiver w
 * @return *RespGetProviderToken
 * @return error
 */
func (a *accessToken) GetProviderAccessToken() (*RespGetProviderToken, error) {
	if len(a.workWechat.ProviderCorpID) < 1 {
		return nil, errors.New("设置ProviderCorpID出错")
	}

	if len(a.workWechat.ProviderSecret) < 1 {
		return nil, errors.New("设置ProviderSecret出错")
	}

	var resp = &RespGetProviderToken{}
	err := a.workWechat.Scan(context.Background(),
		GetProviderAccessTokenAction(a.workWechat.ProviderCorpID, a.workWechat.ProviderSecret),
		resp,
	)

	if err != nil {
		return nil, err
	}

	if resp.respCommon.ErrCode != 0 {
		return nil, errors.New("获取响应数据失败")
	}
	return resp, nil
}

/**
 * @Description:获取第三方应用access_token
 * @author:ljj
 * @receiver w
 * @return *RespGetSuiteToken
 * @return error
 */
func (a *accessToken) GetSuiteAccessToken() (*RespGetSuiteToken, error) {
	if len(a.workWechat.SuiteID) < 1 {
		return nil, errors.New("设置SuiteID出错")
	}

	if len(a.workWechat.SuiteSecret) < 1 {
		return nil, errors.New("设置SuiteSecret出错")
	}

	var resp = &RespGetSuiteToken{}
	err := a.workWechat.Scan(context.Background(),
		GetSuitAccessTokenAction(a.workWechat.SuiteID, a.workWechat.SuiteSecret, a.GetSuiteTicketByCache()),
		resp,
	)

	if err != nil {
		return nil, err
	}

	if resp.respCommon.ErrCode != 0 {
		return nil, errors.New("获取响应数据失败")
	}
	return resp, nil
}


/**
 * @Description:获取授权企业应用access_token
 * @author:ljj
 * @receiver w
 * @return *RespGetCorpToken
 * @return error
 */
func (a *accessToken) GetCorpAccessToken() (*RespGetCorpToken, error) {
	if len(a.workWechat.CorpId) < 1 {
		return nil, errors.New("设置CorpId出错")
	}

	if len(a.workWechat.PermanentCode) < 1 {
		return nil, errors.New("设置PermanentCode出错")
	}

	var resp = &RespGetCorpToken{}
	err := a.workWechat.Scan(context.Background(),
		GetCorpAccessTokenAction(a.GetSuiteAccessTokenByCache(), a.workWechat.CorpId, a.workWechat.PermanentCode),
		resp,
	)

	if err != nil {
		return nil, err
	}

	if resp.respCommon.ErrCode != 0 {
		return nil, errors.New("获取响应数据失败")
	}
	return resp, nil
}

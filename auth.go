package work

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type auth struct {
	workWechat      workWechat
	SuitAccessToken string
}

func (w workWechat) NewAuth() *auth {
	return &auth{
		workWechat: w,
	}
}

// GetPreAuthCode 获取预授权码 https://work.weixin.qq.com/api/doc/90001/90143/90601
func NewGetPreAuthCode(suitAccessToken string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/service/get_pre_auth_code?suite_access_token=%s", suitAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpGet),
	)
}

// GetPermanentCode 获取企业永久授权码 https://work.weixin.qq.com/api/doc/90001/90143/90603
func NewGetPermanentCode(suitAccessToken string, authCode string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/service/get_permanent_code?suite_access_token=%s", suitAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			reqInfo := reqGetPermanentCode{
				AuthCode: authCode,
			}
			jsonInfo, err := json.Marshal(reqInfo)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

// NewSetSessionInfo 设置授权配置 https://work.weixin.qq.com/api/doc/90001/90143/90602
func NewSetSessionInfo(suitAccessToken string, preAuthCode string, info SessionInfo) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/service/set_session_info?suite_access_token=%s", suitAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			reqInfo := reqSetSessionInfo{
				PreAuthCode: preAuthCode,
				SessionInfo: info,
			}
			jsonInfo, err := json.Marshal(reqInfo)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

/**
 * @Description: 设置授权配置
 * @author:ljj
 * @receiver w
 * @param preAuthCode string 预授权码
 * @param info SessionInfo 设置的授权配置
 * @return *RespSetSessionInfo
 * @return error
 */
func (a *auth) SetSessionInfo(preAuthCode string, info SessionInfo) (*RespSetSessionInfo, error) {

	suiteAccessToken := a.workWechat.NewAccessToken().GetSuiteAccessTokenByCache()

	opt := &RespSetSessionInfo{}
	err := a.workWechat.Scan(context.Background(), NewSetSessionInfo(suiteAccessToken, preAuthCode, info), opt)
	if err != nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("设置授权配置失败")
	}
	return opt, nil
}

/**
 * @Description: 获取企业永久授权码
 * @author:ljj
 * @receiver w
 * @param authCode string 企业授权后请求到回调地址产生的授权码
 * @return *RespGetPreAuthCode
 * @return error
 */
func (a *auth) GetPermanentCode(authCode string) (*RespGetPreAuthCode, error) {

	suiteAccessToken := a.workWechat.NewAccessToken().GetSuiteAccessTokenByCache()

	opt := &RespGetPreAuthCode{}
	err := a.workWechat.Scan(context.Background(), NewGetPermanentCode(suiteAccessToken, authCode), opt)
	if err != nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("获取永久授权码失败")
	}
	return opt, nil
}

/**
 * @Description: 获取预授权码
 * @author:21
 * @receiver w
 * @return *RespGetPreAuthCode
 * @return error
 */
func (a *auth) GetPreAuthCode() (*RespGetPreAuthCode, error) {

	suiteAccessToken := a.workWechat.NewAccessToken().GetSuiteAccessTokenByCache()

	opt := &RespGetPreAuthCode{}
	err := a.workWechat.Scan(context.Background(), NewGetPreAuthCode(suiteAccessToken), opt)
	if err != nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("获取预授权码失败")
	}
	return opt, nil
}

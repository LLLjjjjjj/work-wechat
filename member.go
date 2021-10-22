/*
@Time : 2021/7/7 11:35 上午
@Author : 21
@File : member
@Software: GoLand
*/
package work

import (
	"context"
	"errors"
	"fmt"
)

//成员模块
type member struct {
	workWechat workWechat
}

func (w workWechat) NewMember() *member {
	return &member{
		w,
	}
}

// 获取成员信息
func NewGetUserInfoAction(accessToken string, userId string) Action   {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/user/get?access_token=%s&userid=%s", accessToken, userId)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpGet),
	)
}

func (m *member) GetUserInfoAction(userId string) (*UserInfo, error) {
	cropAccessToken := m.workWechat.NewAccessToken().GetCorpAccessTokenByCache()

	opt := &UserInfo{}
	err := m.workWechat.Scan(context.Background(), NewGetUserInfoAction(cropAccessToken, userId), opt)
	if err != nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("获取成员信息失败" + opt.ErrMsg)
	}
	return opt, nil
}

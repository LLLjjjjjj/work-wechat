/*
@Time : 2021/7/7 11:38 上午
@Author : 21
@File : member_test
@Software: GoLand
*/
package work

import (
	"context"
	"testing"
)

func TestGetUserInfoAction(t *testing.T) {
	testConfig := Config{
		ProviderCorpID: "",
		ProviderSecret: "",
		SuiteID:        "",
		SuiteSecret:    "",
		SuiteTicket:    "",
		CorpId:         "",
		PermanentCode:  "",
	}
	accessToken := "11111"
	userId := "11111"
	classInfo := NewWorkWechat(testConfig)
	var resp = &UserInfo{}
	err := classInfo.Scan(context.Background(),
		GetUserInfoAction(accessToken, userId),
		resp,
	)
	t.Log(err)
	t.Log(resp.ErrCode)
}

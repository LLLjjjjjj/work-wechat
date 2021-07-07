/*
@Time : 2021/7/7 12:01 下午
@Author : 21
@File : external_userid_test
@Software: GoLand
*/
package work

import (
	"context"
	"testing"
)

func TestGetExternalContactList(t *testing.T) {
	//GetProviderAccessTokenAction
	testConfig := Config{
		ProviderCorpID: "",
		ProviderSecret: "",
		SuiteID:        "",
		SuiteSecret:    "",
		SuiteTicket:    "",
		CorpId:         "",
		PermanentCode:  "",
	}
	classInfo := NewWorkWechat(testConfig)
	accessToken := "11111"
	userId := "11111"
	var resp = &ExternalContactList{}
	err := classInfo.Scan(context.Background(),
		GetExternalContactList(accessToken,userId),
		resp,
	)
	t.Log(err)
	t.Log(resp.ErrCode)
}


func TestGetExternalContactUserInfo(t *testing.T) {
	testConfig := Config{
		ProviderCorpID: "",
		ProviderSecret: "",
		SuiteID:        "",
		SuiteSecret:    "",
		SuiteTicket:    "",
		CorpId:         "",
		PermanentCode:  "",
	}
	classInfo := NewWorkWechat(testConfig)
	accessToken := "11111"
	userId := "11111"
	cursors := "11111"
	var resp = &ExternalContactUserInfo{}
	err := classInfo.Scan(context.Background(),
		GetExternalContactUserInfo(accessToken,userId, cursors),
		resp,
	)
	t.Log(err)
	t.Log(resp.ErrCode)



}

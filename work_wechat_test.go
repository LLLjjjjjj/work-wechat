/*
@Time : 2021/7/6 11:24 上午
@Author : 21
@File : we_work_test
@Software: GoLand
*/
package work

import (
	"context"
	"fmt"
	"testing"
)

func TestNewWeWork(t *testing.T) {
	testClass := NewWeWork(SetProviderCorpID("1"))
	fmt.Println(testClass)
}

func TestNewWorkWechat(t *testing.T) {
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
	// TODO 获取 suitAccessToken
	suitAccessToken := "suitAccessToken"
	// 获取 企业预先授权码
	opt := &RespGetPreAuthCode{}
	err := classInfo.Scan(context.Background(), NewGetPreAuthCode(suitAccessToken), opt)
	fmt.Println(opt)
	t.Log(err)
	// 获取 企业的永久授权码
	authCode := "1111"
	req :=  reqGetPermanentCode{
		AuthCode: authCode,
	}
	err = classInfo.Scan(context.Background(), NewGetPermanentCode(suitAccessToken, authCode),req )
	t.Log(err)

}

func TestGetProviderAccessTokenAction(t *testing.T) {
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
	var respGetProviderToken = RespGetProviderToken{}
	err := classInfo.Scan(context.Background(),
		GetProviderAccessTokenAction(testConfig.ProviderCorpID,testConfig.ProviderSecret),
		respGetProviderToken,
	)
	fmt.Println(err)
}

func TestGetSuitAccessTokenAction(t *testing.T) {
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
	var resp = reqGetSuiteToken{}
	err := classInfo.Scan(context.Background(),
		GetSuitAccessTokenAction(testConfig.SuiteID,testConfig.SuiteSecret, testConfig.SuiteTicket),
		resp,
	)
	fmt.Println(err)
	fmt.Println(resp)
}

func TestGetCorpAccessTokenAction(t *testing.T) {
	weworkChatConfig := Config{
		ProviderCorpID: "",
		ProviderSecret: "",
		SuiteID:        "",
		SuiteSecret:    "",
		SuiteTicket:    "",
		CorpId:         "",
		PermanentCode:  "",
	}
	var req = reqGetCorpToken{}
	suitAccessToken := "1111"


	classInfo := NewWorkWechat(weworkChatConfig)
	err := classInfo.Scan(context.Background(),
		GetCorpAccessTokenAction(suitAccessToken,weworkChatConfig.CorpId, weworkChatConfig.PermanentCode),
		req)
	t.Log(err)
	t.Log(req)

}



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
	fmt.Println(classInfo)
	// TODO 获取 suitAccessToken
	//suitAccessToken := "suitAccessToken"
	//res , err := classInfo.GetPreAuthCode(suitAccessToken)
	//if err != nil{
	//	t.Log(err)
	//}
	//if res == nil{
	//	t.Log(res)
	//	return
	//}
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
	res , err := classInfo.NewAccessToken().GetProviderAccessToken()
	if err != nil{
		return
	}
	fmt.Println(res)
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



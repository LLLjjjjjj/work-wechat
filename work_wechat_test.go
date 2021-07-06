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
	res, err := classInfo.Do(context.Background(), NewGetPreAuthCode(suitAccessToken))
	fmt.Println(string(res))
	// 获取 企业的永久授权码
	authCode := "1111"
	res, err = classInfo.Do(context.Background(), NewGetPermanentCode(suitAccessToken, authCode))
	fmt.Println(err)
	fmt.Println(res)

	fmt.Println(classInfo)

}

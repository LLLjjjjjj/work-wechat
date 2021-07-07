/*
@Time : 2021/7/7 11:35 上午
@Author : 21
@File : member
@Software: GoLand
*/
package work

import (
	"fmt"
)

//成员模块

// 获取成员信息
func GetUserInfoAction(accessToken string, userId string) Action   {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/user/get?access_token=%s&userid=%s", accessToken, userId)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpGet),
	)
}

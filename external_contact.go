package work

import "fmt"

// 外部联系人相关


// 获取客户列表
func GetExternalContactList(accessToken string , userId string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/externalcontact/list?access_token=%s&userid=%s", accessToken, userId)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpGet),
	)
}

// 获取客户详细信息
func GetExternalContactUserInfo(accessToken string , externalUserId string, cursors string ) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/externalcontact/get?access_token=%s&external_userid=%s&cursor=%s",accessToken,externalUserId, cursors)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpGet),
	)
}


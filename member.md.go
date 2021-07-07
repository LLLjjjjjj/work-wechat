/*
@Time : 2021/7/7 11:40 上午
@Author : 21
@File : member.md.go
@Software: GoLand
*/
package work

type UserInfo struct {
	respCommon
	userInfoDetails
}

type userInfoDetails struct {
	Userid string `json:"userid"`
	Name string `json:"name"`
	Department []int `json:"department"`
	Order []int `json:"order"`
	Position string `json:"position"`
	Mobile string `json:"mobile"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	IsLeaderInDept []int `json:"is_leader_in_dept"`
	Avatar string `json:"avatar"`
	ThumbAvatar string `json:"thumb_avatar"`
	Telephone string `json:"telephone"`
	Alias string `json:"alias"`
	Address string `json:"address"`
	OpenUserid string `json:"open_userid"`
	MainDepartment int `json:"main_department"`
	Extattr struct {
		Attrs []struct {
			Type int `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				URL string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
		} `json:"attrs"`
	} `json:"extattr"`
	Status int `json:"status"`
	QrCode string `json:"qr_code"`
	ExternalPosition string `json:"external_position"`
	ExternalProfile struct {
		ExternalCorpName string `json:"external_corp_name"`
		ExternalAttr []struct {
			Type int `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				URL string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			Miniprogram struct {
				Appid string `json:"appid"`
				Pagepath string `json:"pagepath"`
				Title string `json:"title"`
			} `json:"miniprogram,omitempty"`
		} `json:"external_attr"`
	} `json:"external_profile"`
}

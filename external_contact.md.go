package work

// 外部联系人列表
type ExternalContactList struct {
	respCommon
	externalContactListDetails
}

type externalContactListDetails struct {
	ExternalUserId []string `json:"external_userid"`
}

// 外部联系人用户信息
type ExternalContactUserInfo struct {
	respCommon
	ExternalContact externalContactStruct `json:"external_contact"`
	FollowUser []*followUser `json:"follow_user"`
	NextCursor string `json:"next_cursor"`
}

type externalContactStruct struct {
	ExternalUserid string `json:"external_userid"`
	Name string `json:"name"`
	Position string `json:"position"`
	Avatar string `json:"avatar"`
	CorpName string `json:"corp_name"`
	CorpFullName string `json:"corp_full_name"`
	Type int `json:"type"`
	Gender int `json:"gender"`
	Unionid string `json:"unionid"`
	ExternalProfile struct {
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

type followUser struct {
	Userid string `json:"userid"`
	Remark string `json:"remark"`
	Description string `json:"description"`
	Createtime int `json:"createtime"`
	Tags []struct {
		GroupName string `json:"group_name"`
		TagName string `json:"tag_name"`
		TagID string `json:"tag_id"`
		Type int `json:"type"`
	} `json:"tags,omitempty"`
	RemarkCorpName string `json:"remark_corp_name,omitempty"`
	RemarkMobiles []string `json:"remark_mobiles,omitempty"`
	OperUserid string `json:"oper_userid"`
	AddWay int `json:"add_way"`
	State string `json:"state,omitempty"`
}

// 发送欢迎语
type reqSendWelcomeMsg struct {
	WelcomeCode string `json:"welcome_code"`
	Text Text `json:"text"`
	Attachments []Attachments `json:"attachments"`
}
type Text struct {
	Content string `json:"content"`
}
type Image struct {
	MediaID string `json:"media_id"`
	PicURL string `json:"pic_url"`
}
type Link struct {
	Title string `json:"title"`
	Picurl string `json:"picurl"`
	Desc string `json:"desc"`
	URL string `json:"url"`
}
type Miniprogram struct {
	Title string `json:"title"`
	PicMediaID string `json:"pic_media_id"`
	Appid string `json:"appid"`
	Page string `json:"page"`
}
type Video struct {
	MediaID string `json:"media_id"`
}
type Attachments struct {
	Msgtype string `json:"msgtype"`
	Image Image `json:"image,omitempty"`
	Link Link `json:"link,omitempty"`
	Miniprogram Miniprogram `json:"miniprogram,omitempty"`
	Video Video `json:"video,omitempty"`
}

type RespSendWelcomeMsg struct {
	respCommon
}
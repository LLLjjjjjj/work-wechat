package work

// 外部联系人列表 17
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
	FollowUser      []*followUser         `json:"follow_user"`
	NextCursor      string                `json:"next_cursor"`
}

type externalContactStruct struct {
	ExternalUserid  string `json:"external_userid"`
	Name            string `json:"name"`
	Position        string `json:"position"`
	Avatar          string `json:"avatar"`
	CorpName        string `json:"corp_name"`
	CorpFullName    string `json:"corp_full_name"`
	Type            int    `json:"type"`
	Gender          int    `json:"gender"`
	Unionid         string `json:"unionid"`
	ExternalProfile struct {
		ExternalAttr []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			Miniprogram struct {
				Appid    string `json:"appid"`
				Pagepath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		} `json:"external_attr"`
	} `json:"external_profile"`
}

type followUser struct {
	Userid      string `json:"userid"`
	Remark      string `json:"remark"`
	Description string `json:"description"`
	Createtime  int    `json:"createtime"`
	Tags        []struct {
		GroupName string `json:"group_name"`
		TagName   string `json:"tag_name"`
		TagID     string `json:"tag_id"`
		Type      int    `json:"type"`
	} `json:"tags,omitempty"`
	RemarkCorpName string   `json:"remark_corp_name,omitempty"`
	RemarkMobiles  []string `json:"remark_mobiles,omitempty"`
	OperUserid     string   `json:"oper_userid"`
	AddWay         int      `json:"add_way"`
	State          string   `json:"state,omitempty"`
}

// 发送欢迎语
type reqSendWelcomeMsg struct {
	WelcomeCode string        `json:"welcome_code"`
	Text        Text          `json:"text"`
	Attachments []Attachments `json:"attachments"`
}
type Text struct {
	Content string `json:"content"`
}
type Image struct {
	MediaID string `json:"media_id"`
	PicURL  string `json:"pic_url"`
}
type Link struct {
	Title  string `json:"title"`
	Picurl string `json:"picurl"`
	Desc   string `json:"desc"`
	URL    string `json:"url"`
}
type Miniprogram struct {
	Title      string `json:"title"`
	PicMediaID string `json:"pic_media_id"`
	Appid      string `json:"appid"`
	Page       string `json:"page"`
}
type Video struct {
	MediaID string `json:"media_id"`
}
type Attachments struct {
	Msgtype     string      `json:"msgtype"`
	Image       Image       `json:"image,omitempty"`
	Link        Link        `json:"link,omitempty"`
	Miniprogram Miniprogram `json:"miniprogram,omitempty"`
	Video       Video       `json:"video,omitempty"`
}

type RespSendWelcomeMsg struct {
	respCommon
}

type RespGetFollowUserList struct {
	respCommon
	FollowUser []string `json:"follow_user"`
}

type AddContactWayResp struct {
	respCommon
	ConfigID string `json:"config_id"`
	QrCode   string `json:"qr_code"`
}

type OpenGidToChatIdReq struct {
	OpenGid string `json:"opengid"` //群的opengid,由前端加密数据解得opengid
}

type OpenGidToChatIdResp struct {
	respCommon
	ChatID string `json:"chat_id"` //群chat_id
}

type AddContactWayReq struct {
	Type          int      `json:"type"`
	Scene         int      `json:"scene"`
	Style         int      `json:"style"`
	Remark        string   `json:"remark"`
	SkipVerify    bool     `json:"skip_verify"`
	State         string   `json:"state"`
	User          []string `json:"user"`
	Party         []int    `json:"party"`
	IsTemp        bool     `json:"is_temp"`
	ExpiresIn     int      `json:"expires_in"`
	ChatExpiresIn int      `json:"chat_expires_in"`
	Unionid       string   `json:"unionid"`
	Conclusions   struct {
		Text struct {
			Content string `json:"content"`
		} `json:"text"`
		Image struct {
			MediaID string `json:"media_id"`
		} `json:"image"`
		Link struct {
			Title  string `json:"title"`
			Picurl string `json:"picurl"`
			Desc   string `json:"desc"`
			URL    string `json:"url"`
		} `json:"link"`
		Miniprogram struct {
			Title      string `json:"title"`
			PicMediaID string `json:"pic_media_id"`
			Appid      string `json:"appid"`
			Page       string `json:"page"`
		} `json:"miniprogram"`
	} `json:"conclusions"`
}

type AddGroupWelcomeTemplateReq struct {
	Text        *GroupWelcomeTemplateText        `json:"text,omitempty"`
	Image       *GroupWelcomeTemplateImage       `json:"image,omitempty"`
	Link        *GroupWelcomeTemplateLink        `json:"link,omitempty"`
	MiniProgram *GroupWelcomeTemplateMiniProgram `json:"miniprogram,omitempty"`
	File        *GroupWelcomeTemplateFile        `json:"file,omitempty"`
	Video       *GroupWelcomeTemplateVideo       `json:"video,omitempty"`
	AgentID     int                              `json:"agentid"`
	Notify      int                              `json:"notify"`
}

// GroupWelcomeTemplateText defines the text part of the welcome template.
type GroupWelcomeTemplateText struct {
	Content string `json:"content,omitempty"`
}

// GroupWelcomeTemplateImage defines the image part of the welcome template.
type GroupWelcomeTemplateImage struct {
	MediaID string `json:"media_id,omitempty"`
	PicURL  string `json:"pic_url,omitempty"`
}

// GroupWelcomeTemplateLink defines the link part of the welcome template.
type GroupWelcomeTemplateLink struct {
	Title  string `json:"title,omitempty"`
	PicURL string `json:"picurl,omitempty"`
	Desc   string `json:"desc,omitempty"`
	URL    string `json:"url,omitempty"`
}

// GroupWelcomeTemplateMiniProgram defines the mini program part of the welcome template.
type GroupWelcomeTemplateMiniProgram struct {
	Title      string `json:"title,omitempty"`
	PicMediaID string `json:"pic_media_id,omitempty"`
	AppID      string `json:"appid,omitempty"`
	Page       string `json:"page,omitempty"`
}

// GroupWelcomeTemplateFile defines the file part of the welcome template.
type GroupWelcomeTemplateFile struct {
	MediaID string `json:"media_id,omitempty"`
}

// GroupWelcomeTemplateVideo defines the video part of the welcome template.
type GroupWelcomeTemplateVideo struct {
	MediaID string `json:"media_id,omitempty"`
}

type AddGroupWelcomeTemplateResp struct {
	respCommon
	TemplateId string `json:"template_id"`
}

type EditGroupWelcomeTemplateReq struct {
	TemplateID  string                           `json:"template_id"`
	Text        *GroupWelcomeTemplateText        `json:"text,omitempty"`
	Image       *GroupWelcomeTemplateImage       `json:"image,omitempty"`
	Link        *GroupWelcomeTemplateLink        `json:"link,omitempty"`
	MiniProgram *GroupWelcomeTemplateMiniProgram `json:"miniprogram,omitempty"`
	File        *GroupWelcomeTemplateFile        `json:"file,omitempty"`
	Video       *GroupWelcomeTemplateVideo       `json:"video,omitempty"`
	AgentID     int                              `json:"agentid"`
}

type GetUnionidToExternalUserid3rdReq struct {
	Unionid string `json:"unionid"`
	Openid  string `json:"openid"`
	Corpid  string `json:"corpid,omitempty"`
}

type ExternalUserInfo struct {
	CorpID         string `json:"corpid"`
	ExternalUserID string `json:"external_userid"`
}

type GetUnionidToExternalUserid3rdResp struct {
	respCommon
	ExternalUserIDInfo []*ExternalUserInfo `json:"external_userid_info"`
}

type EditGroupWelcomeTemplateResp struct {
	respCommon
}

type GetGroupWelcomeTemplateReq struct {
	TemplateId string `json:"template_id"`
}

type GetGroupWelcomeTemplateResp struct {
	respCommon
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Image struct {
		PicURL string `json:"pic_url"`
	} `json:"image"`
	Link struct {
		Title  string `json:"title"`
		PicURL string `json:"picurl"`
		Desc   string `json:"desc"`
		URL    string `json:"url"`
	} `json:"link"`
	MiniProgram struct {
		Title      string `json:"title"`
		PicMediaID string `json:"pic_media_id"`
		AppID      string `json:"appid"`
		Page       string `json:"page"`
	} `json:"miniprogram"`
	File struct {
		MediaID string `json:"media_id"`
	} `json:"file"`
	Video struct {
		MediaID string `json:"media_id"`
	} `json:"video"`
}

type DeleteGroupWelcomeTemplateReq struct {
	TemplateId string `json:"template_id"`
	AgentId    int    `json:"agentid"`
}

type DeleteGroupWelcomeTemplateResp struct {
	respCommon
}

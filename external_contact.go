package work

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

// 外部联系人相关
type externalContact struct {
	workWechat workWechat
}

func (w workWechat) NewExternalContact() *externalContact {
	return &externalContact{
		w,
	}
}

// 获取客户列表
func NewGetExternalContactList(accessToken string, userId string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/externalcontact/list?access_token=%s&userid=%s", accessToken, userId)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpGet),
	)
}

/**
 * @Description: 获取客户列表
 * @author:ljj
 * @receiver w
 * @param userId string 企业成员的userid
 * @return *ExternalContactList
 * @return error
 */
func (e *externalContact) GetExternalContactList(userId string) (*ExternalContactList, error) {

	cropAccessToken := e.workWechat.NewAccessToken().GetCorpAccessTokenByCache()

	opt := &ExternalContactList{}
	err := e.workWechat.Scan(context.Background(), NewGetExternalContactList(cropAccessToken, userId), opt)
	if err != nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("获取客户列表失败" + opt.ErrMsg)
	}
	return opt, nil
}

// 获取客户详细信息
func NewGetExternalContactUserInfo(accessToken string, externalUserId string, cursors string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/externalcontact/get?access_token=%s&external_userid=%s&cursor=%s", accessToken, externalUserId, cursors)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpGet),
	)
}

/**
 * @Description: 获取客户详细信息
 * @author:ljj
 * @receiver w
 * @param externalUserId string 外部联系人的userid，注意不是企业成员的帐号
 * @param cursor string 上次请求返回的next_cursor
 * @return *ExternalContactUserInfo
 * @return error
 */
func (e *externalContact) GetExternalContactUserInfo(externalUserId string, cursor string) (*ExternalContactUserInfo, error) {

	cropAccessToken := e.workWechat.NewAccessToken().GetCorpAccessTokenByCache()

	opt := &ExternalContactUserInfo{}
	err := e.workWechat.Scan(context.Background(), NewGetExternalContactUserInfo(cropAccessToken, externalUserId, cursor), opt)
	if err != nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("获取客户详细信息失败" + opt.ErrMsg)
	}
	return opt, nil
}

// 发送新客户欢迎语
func NewSendWelcomeMsg(corpAccessToken string, welcomeCode string, text string, attachments []Attachments) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/externalcontact/send_welcome_msg?access_token=%s", corpAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			reqInfo := reqSendWelcomeMsg{
				WelcomeCode: welcomeCode,
				Text: Text{
					Content: text,
				},
				Attachments: attachments,
			}
			jsonInfo, err := json.Marshal(reqInfo)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

/**
 * @Description: 发送新客户欢迎语
 * @author:ljj
 * @receiver w
 * @param welcomeCode string 通过添加外部联系人事件推送给企业的发送欢迎语的凭证，有效期为20秒
 * @param text string 消息文本内容,最长为4000字节
 * @param attachments []Attachments 附件，最多可添加9个附件
 * @return *RespSendWelcomeMsg
 * @return error
 */
func (e *externalContact) SendWelcomeMsg(welcomeCode string, text string, attachments []Attachments) (*RespSendWelcomeMsg, error) {

	cropAccessToken := e.workWechat.NewAccessToken().GetCorpAccessTokenByCache()

	opt := &RespSendWelcomeMsg{}
	err := e.workWechat.Scan(context.Background(), NewSendWelcomeMsg(
		cropAccessToken,
		welcomeCode,
		text,
		attachments,
	), opt)
	if err != nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("发送新客户欢迎语失败" + opt.ErrMsg)
	}
	return opt, nil
}

// 获取配置了客户联系功能的成员列表
func NewGetFollowUserList(corpAccessToken string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/externalcontact/get_follow_user_list?access_token=%s", corpAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpGet),
	)
}

/**
 * @Description: 获取配置了客户联系功能的成员列表
 * @author:ljj
 * @receiver w
 * @return *RespGetFollowUserList
 * @return error
 */
func (e *externalContact) GetFollowUserList() (*RespGetFollowUserList, error) {

	cropAccessToken := e.workWechat.NewAccessToken().GetCorpAccessTokenByCache()

	opt := &RespGetFollowUserList{}
	err := e.workWechat.Scan(context.Background(), NewGetFollowUserList(
		cropAccessToken,
	), opt)
	if err != nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("获取配置了客户联系功能的成员列表失败" + opt.ErrMsg)
	}
	return opt, nil
}

// 添加配置客户联系「联系我」方式
func NewAddContactWay(corpAccessToken string, userId []string, remark string, state string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/externalcontact/add_contact_way?access_token=%s", corpAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			reqInfo := AddContactWayReq{
				Type:       1,
				Scene:      2,
				Remark:     remark,
				SkipVerify: true,
				State:      state,
				User:       userId,
			}
			jsonInfo, err := json.Marshal(reqInfo)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

// 添加配置客户联系「联系我」方式
func (e *externalContact) AddContactWay(userId []string, remark string, state string) (*AddContactWayResp, error) {
	cropAccessToken := e.workWechat.NewAccessToken().GetCorpAccessTokenByCache()

	opt := &AddContactWayResp{}
	err := e.workWechat.Scan(context.Background(), NewAddContactWay(
		cropAccessToken,
		userId,
		remark,
		state,
	), opt)
	if err != nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("配置客户联系「联系我」方式失败" + opt.ErrMsg)
	}
	return opt, nil
}

func NewOpenGidToChatId(corpAccessToken, openGid string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/externalcontact/opengid_to_chatid?access_token=%s", corpAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			reqInfo := &OpenGidToChatIdReq{
				OpenGid: openGid,
			}
			jsonInfo, err := json.Marshal(reqInfo)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

// 客户群opengid转换ChatId
func (e *externalContact) GetOpenGidToChatId(openGid string) (*OpenGidToChatIdResp, error) {
	cropAccessToken := e.workWechat.NewAccessToken().GetCorpAccessTokenByCache()
	opt := &OpenGidToChatIdResp{}
	err := e.workWechat.Scan(context.Background(), NewOpenGidToChatId(
		cropAccessToken,
		openGid,
	), opt)
	if err != nil || opt == nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("客户群opengid转换ChatId" + opt.ErrMsg)
	}
	return opt, nil
}

func NewAddGroupWelcomeTemplate(corpAccessToken string, req *AddGroupWelcomeTemplateReq) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/externalcontact/group_welcome_template/add?access_token=%s", corpAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			jsonInfo, err := json.Marshal(req)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

// 添加入群欢迎语素材
func (e *externalContact) AddGroupWelcomeTemplate(req *AddGroupWelcomeTemplateReq) (*AddGroupWelcomeTemplateResp, error) {
	cropAccessToken := e.workWechat.NewAccessToken().GetCorpAccessTokenByCache()
	opt := &AddGroupWelcomeTemplateResp{}
	err := e.workWechat.Scan(context.Background(), NewAddGroupWelcomeTemplate(
		cropAccessToken,
		req,
	), opt)
	if err != nil || opt == nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("添加入群欢迎语素材" + opt.ErrMsg)
	}
	return opt, nil
}

func NewEditGroupWelcomeTemplate(corpAccessToken string, req *EditGroupWelcomeTemplateReq) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/externalcontact/group_welcome_template/edit?access_token=%s", corpAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			jsonInfo, err := json.Marshal(req)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

// 编辑入群欢迎语素材
func (e *externalContact) EditGroupWelcomeTemplate(req *EditGroupWelcomeTemplateReq) (*EditGroupWelcomeTemplateResp, error) {
	cropAccessToken := e.workWechat.NewAccessToken().GetCorpAccessTokenByCache()
	opt := &EditGroupWelcomeTemplateResp{}
	err := e.workWechat.Scan(context.Background(), NewEditGroupWelcomeTemplate(
		cropAccessToken,
		req,
	), opt)
	if err != nil || opt == nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("编辑入群欢迎语素材" + opt.ErrMsg)
	}
	return opt, nil
}

func NewGetGroupWelcomeTemplate(corpAccessToken string, templateId string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/externalcontact/group_welcome_template/get?access_token=%s", corpAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			req := &GetGroupWelcomeTemplateReq{TemplateId: templateId}
			jsonInfo, err := json.Marshal(req)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

// 获取入群欢迎语素材
func (e *externalContact) GetGroupWelcomeTemplate(templateId string) (*GetGroupWelcomeTemplateResp, error) {
	cropAccessToken := e.workWechat.NewAccessToken().GetCorpAccessTokenByCache()
	opt := &GetGroupWelcomeTemplateResp{}
	err := e.workWechat.Scan(context.Background(), NewGetGroupWelcomeTemplate(
		cropAccessToken,
		templateId,
	), opt)
	if err != nil {
		return nil, err
	}
	if opt == nil || opt.ErrCode != 0 {
		return nil, errors.New("获取入群欢迎语素材" + opt.ErrMsg)
	}
	return opt, nil
}

func NewDelGroupWelcomeTemplate(corpAccessToken string, templateId string, agentId int) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/externalcontact/group_welcome_template/del?access_token=%s", corpAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			req := &DeleteGroupWelcomeTemplateReq{TemplateId: templateId, AgentId: agentId}
			jsonInfo, err := json.Marshal(req)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

// 删除入群欢迎语素材
func (e *externalContact) DelGroupWelcomeTemplate(templateId string, agentId int) (*DeleteGroupWelcomeTemplateResp, error) {
	cropAccessToken := e.workWechat.NewAccessToken().GetCorpAccessTokenByCache()
	opt := &DeleteGroupWelcomeTemplateResp{}
	err := e.workWechat.Scan(context.Background(), NewDelGroupWelcomeTemplate(
		cropAccessToken,
		templateId,
		agentId,
	), opt)
	if err != nil || opt == nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("删除入群欢迎语素材" + opt.ErrMsg)
	}
	return opt, nil
}

func NewUnionidToExternalUserid3rd(suiteAccessToken string, req *GetUnionidToExternalUserid3rdReq) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/service/externalcontact/unionid_to_external_userid_3rd?suite_access_token=%s", suiteAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			jsonInfo, err := json.Marshal(req)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

// 第三方主体unionid转换为第三方external_userid
func (e *externalContact) GetUnionidToExternalUserid3rd(req *GetUnionidToExternalUserid3rdReq) (*GetUnionidToExternalUserid3rdResp, error) {
	suiteAccessToken := e.workWechat.NewAccessToken().GetSuiteAccessTokenByCache()
	opt := &GetUnionidToExternalUserid3rdResp{}
	err := e.workWechat.Scan(context.Background(), NewUnionidToExternalUserid3rd(suiteAccessToken, req), opt)
	if err != nil || opt == nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("第三方主体unionid转换为第三方external_userid" + opt.ErrMsg)
	}
	return opt, nil
}

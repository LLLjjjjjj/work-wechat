package work

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

// 素材管理
type media struct {
	workWechat workWechat
}

func (w workWechat) NewMedia() *media {
	return &media{
		w,
	}
}

// 上传临时素材
func NewTemporaryUpload(accessToken string, filePath, fileType string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/media/upload?access_token=%s&type=%s", accessToken, fileType)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpUpload),
		WitchUploadForm(fmt.Sprintf("upload-file=@file:%v", filePath)),
	)
}

/**
 * @Description: 上传临时素材
 * @author:ljj
 * @receiver w
 * @param fileSteam string 二进制流
 * @param fileType string 媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件（file）
 * @return *RespTemporaryUpload
 * @return error
 */
func (m *media) TemporaryUpload(filePath, fileType string) (*RespTemporaryUpload, error) {

	cropAccessToken := m.workWechat.NewAccessToken().GetCorpAccessTokenByCache()

	opt := &RespTemporaryUpload{}
	err := m.workWechat.Scan(context.Background(), NewTemporaryUpload(
		cropAccessToken,
		filePath,
		fileType,
	), opt)
	if err != nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("上传临时素材失败" + opt.ErrMsg)
	}
	return opt, nil
}


// 上传图片素材
func NewImgUpload(accessToken string, filePath string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/media/uploadimg?access_token=%s", accessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpUpload),
		WitchUploadForm(fmt.Sprintf("upload-file=@file:%v", filePath)),
	)
}

/**
 * @Description: 上传图片素材
 * @author:ljj
 * @receiver w
 * @param fileSteam string 二进制流
 * @return *RespImgUpload
 * @return error
 */
func (m *media) ImgUpload(filePath string) (*RespImgUpload, error) {

	cropAccessToken := m.workWechat.NewAccessToken().GetCorpAccessTokenByCache()

	opt := &RespImgUpload{}
	err := m.workWechat.Scan(context.Background(), NewImgUpload(
		cropAccessToken,
		filePath,
	), opt)
	if err != nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("上传图片失败" + opt.ErrMsg)
	}
	return opt, nil
}

// 获取素材
func NewMediaGet(accessToken string, mediaId string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/media/get?access_token=%s&media_id=%s", accessToken, mediaId)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpGet),
	)
}


/**
 * @Description: 获取素材
 * @author:ljj
 * @receiver w
 * @param mediaId string 媒体id
 * @return *RespImgUpload
 * @return error
 */
func (m *media) MediaGet(mediaId string) (*RespMediaGet, error) {
	cropAccessToken := m.workWechat.NewAccessToken().GetCorpAccessTokenByCache()
	res, err := m.workWechat.Do(context.Background(), NewMediaGet(cropAccessToken,mediaId))
	if err != nil {
		return nil, err
	}

	opt := &RespMediaGet{}
	err = json.Unmarshal(res, opt)
	if err == nil { // 无素材
		return opt, nil
	}

	// 有素材二进制文件
	opt.ErrCode = 0
	opt.FileStream = res

	return opt, nil
}


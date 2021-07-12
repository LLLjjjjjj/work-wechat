package work

type reqTemporaryUpload struct {

}

type RespTemporaryUpload struct {
	respCommon
	Type string `json:"type"`
	MediaId string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

// 上传图片
type reqImgUpload struct {

}

type RespImgUpload struct {
	respCommon
	Url string `json:"url"`
}

// 获取图文素材
type RespMediaGet struct {
	respCommon
}
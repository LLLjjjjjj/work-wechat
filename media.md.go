package work

import (
	"io"
	"mime/multipart"
	"os"
)

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
	FileStream []byte
}

type Media struct {
	filename string
	filesize int64
	stream   io.Reader
}

// NewMediaFromFile 从操作系统级文件创建一个欲上传的素材对象
func NewMediaFromFile(f *os.File) (*Media, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	return &Media{
		filename: stat.Name(),
		filesize: stat.Size(),
		stream:   f,
	}, nil
}

func (m *Media) writeTo(w *multipart.Writer) error {
	wr, err := w.CreateFormFile("media", m.filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(wr, m.stream)
	if err != nil {
		return err
	}

	return nil
}

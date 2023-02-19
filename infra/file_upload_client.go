package infra

import (
	"bytes"
	"github.com/empenguin1186/cobra-demo/domain"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"time"
)

type FileUploadClient struct {
	host string
}

func (f *FileUploadClient) Upload(fileModel *domain.FileModel) error {
	// マルチパートリクエストの準備
	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)
	defer mw.Close()

	// csv ファイルをリクエストに含める
	fw, _ := mw.CreateFormFile(fileModel.FileName(), fileModel.FileName())
	_, _ = io.Copy(fw, fileModel.Data())

	// ログインIDとパスワードをリクエストに含める
	mw.WriteField("filename", fileModel.FileName())

	path := "/path/to/api"
	contentType := mw.FormDataContentType()

	// リクエストオブジェクト作成
	request, err := http.NewRequest("POST", "https://"+f.host+path, body)
	if err != nil {
		log.Printf("failed to create request object. err=%v", err)
		return err
	}
	request.Header.Set("Content-Type", contentType)

	// リクエスト情報をダンプ
	dumpedRequest, _ := httputil.DumpRequest(request, true)
	log.Printf(string(dumpedRequest))

	// リクエスト実行
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("failed to request. err=%v", err)
		return err
	}
	defer response.Body.Close()
	log.Printf("Status Code: %d", response.StatusCode)

	// レスポンス構築
	dumpedResponse, _ := httputil.DumpResponse(response, true)
	log.Printf(string(dumpedResponse))

	return nil
}

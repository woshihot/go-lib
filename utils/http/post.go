package http

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/woshihot/go-lib/utils/log"
	"io"
	"mime/multipart"
	"net/http"
	u "net/url"
	"os"
	"strings"
)

func PostUrl(url string) ([]byte, error) {
	return post(&client, url, nil, map[string]string{"Content-Type": "application/json"}, 0)
}

func PostFile(url string, file *os.File) ([]byte, error) {
	bodyBuf := bytes.NewBufferString("")
	bodyWriter := multipart.NewWriter(bodyBuf)

	headers := make(map[string]string)

	fileInfo, err := file.Stat()
	if err != nil {
		log.EF(TAG_ERROR, "PostFile GetFileStat %s\n", err.Error())
		return nil, err
	}
	_, err = bodyWriter.CreateFormFile("file", fileInfo.Name())
	if nil != err {
		log.EF(TAG_ERROR, "PostFile CreateFormFile %s\n", err.Error())
		return nil, err
	}
	boundary := bodyWriter.Boundary()
	closeBuf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))
	requestReader := io.MultiReader(bodyBuf, file, closeBuf)
	headers["Content-Type"] = "multipart/form-data;  boundary=" + boundary
	length := fileInfo.Size() + int64(bodyBuf.Len()) + int64(closeBuf.Len())
	return post(http.DefaultClient, url, requestReader, headers, length)
}

func PostForm(url string, form map[string]string) ([]byte, error) {
	value := make(u.Values)
	for k, v := range form {
		value.Set(k, v)
	}
	return post(&client, url, strings.NewReader(value.Encode()), map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, 0)
}

func PostJson(url string, params map[string]interface{}) ([]byte, error) {
	content, err := jsoniter.Marshal(params)
	if nil != err {
		return nil, err
	}
	return post(&client, url, bytes.NewBuffer(content), map[string]string{"Content-Type": "application/json"}, 0)
}

func post(client *http.Client, url string, reader io.Reader, headers map[string]string, length int64) ([]byte, error) {

	return do(http.MethodPost, client, url, reader, headers, length)
}

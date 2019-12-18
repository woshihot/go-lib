package http

import (
	"fmt"
	"github.com/woshihot/go-lib/utils/log"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	TAG_ERROR = "[Http-error]"
	TAG_DEBUG = "[Http]"
)

var client http.Client

func init() {

	client = http.Client{
		Transport: http.DefaultTransport,
		Timeout:   60 * time.Second,
	}
}

func do(method string, client *http.Client, url string, reader io.Reader, headers map[string]string, length int64) ([]byte, error) {
	return doWithReadBody(method, client, url, reader, headers, length, func(resp *http.Response) (bytes []byte, e error) {
		b, err := ioutil.ReadAll(resp.Body)
		if nil != err {
			log.EF(TAG_ERROR, "%s readBody: %s\n", method, err.Error())
			return nil, err
		}
		return b, nil
	})
}

func doWithReadBody(method string, client *http.Client, url string,
	reader io.Reader, headers map[string]string, length int64,
	readFunc func(resp *http.Response) ([]byte, error)) ([]byte, error) {
	if "" == url {
		err := fmt.Errorf("%s url is empty", method)
		log.EF(TAG_ERROR, "%s\n", err.Error())
		return nil, err
	}
	request, err := http.NewRequest(method, url, reader)
	if nil != err {
		log.EF(TAG_ERROR, "%s NewRequest: %s\n", method, err.Error())
		return nil, err
	}
	request.Close = true
	if nil != headers {
		for key, value := range headers {
			request.Header.Set(key, value)
		}
	}
	if length > 0 {
		request.ContentLength = length
	}
	resp, err := client.Do(request)
	if nil != err {
		log.EF(TAG_ERROR, "%s Do: %s\n", method, err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	return readFunc(resp)
}

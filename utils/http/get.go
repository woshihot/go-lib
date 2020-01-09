package http

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/woshihot/go-lib/utils/log"
	"io"
	"net/http"
	u "net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func Get(url string, params map[string]string) ([]byte, error) {
	URL, err := ParseQuery(url, params)
	if nil != err {
		return nil, err
	}
	return do(http.MethodGet, &client, URL, nil, nil, 0)

}

func GetFile(url, downPath string) error {
	if "" == downPath {
		return fmt.Errorf("downPath is empty")
	}
	fileName := filepath.Base(downPath)
	_, err := doWithReadBody(http.MethodGet, &client, url, nil, nil, 0, func(resp *http.Response) (bytes []byte, e error) {
		log.DF(TAG_DEBUG, "now downloading: [%s]\n", fileName)
		wrap := func(err error) ([]byte, error) {
			return nil, err
		}

		f, err := os.Create(downPath)
		if err != nil {
			return wrap(err)
		}
		defer func() {
			f.Sync()
			f.Close()
		}()

		length, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
		source := resp.Body
		bar := pb.New(length).Set(pb.Bytes, true).SetRefreshRate(time.Millisecond * 10).SetWidth(80)
		bar.Start()
		defer bar.Finish()
		_, err = io.Copy(f, bar.NewProxyReader(source))
		if nil != err {
			log.EF(TAG_ERROR, "download file %s ,%s\n", fileName, err.Error())
			return wrap(err)
		}
		log.DF(TAG_DEBUG, "download file %s ,success\n", fileName)
		return nil, nil
	})
	return err
}

func ParseQuery(url string, params map[string]string) (string, error) {
	URL, err := u.Parse(url)
	if err != nil {
		log.EF(TAG_ERROR, "Get parseQueryUrl %s\n", err.Error())
		return url, err
	}
	q := URL.Query()
	if nil != params {
		for key, value := range params {
			q.Set(key, value)
		}
	}
	URL.RawQuery = q.Encode()
	return URL.String(), nil
}

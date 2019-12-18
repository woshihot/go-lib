package http

import (
	"fmt"
	"github.com/cheggaaa/pb"
	"github.com/woshihot/go-lib/utils/file"
	"github.com/woshihot/go-lib/utils/log"
	"io"
	"net/http"
	u "net/url"
	"path"
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
	fileName := path.Base(downPath)
	_, err := doWithReadBody(http.MethodGet, &client, url, nil, nil, 0, func(resp *http.Response) (bytes []byte, e error) {
		log.DF(TAG_DEBUG, "now downloading: [%s]\n", fileName)
		wrap := func(err error) ([]byte, error) {
			return nil, err
		}
		f, err := file.CreateFile(downPath)
		if err != nil {
			return wrap(err)
		}
		defer func() {
			f.Sync()
			f.Close()
		}()

		length, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
		source := resp.Body

		bar := pb.New(length).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
		bar.Start()
		defer bar.Finish()

		bar.ShowSpeed = true
		bar.ShowFinalTime = true
		bar.SetMaxWidth(80)

		writer := io.MultiWriter(f, bar)
		_, err = io.Copy(writer, source)
		if nil != err {
			log.EF(TAG_ERROR, "download file %s ,%s\n", fileName, err.Error())
		}
		log.DF(TAG_DEBUG, "download file %s ,success\n", fileName)
		return wrap(err)
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

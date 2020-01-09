package http

import (
	"fmt"
	"testing"
)

func TestGetFile(t *testing.T) {
	err := GetFile("http://127.0.0.1:8080/data/agentServer-1.6.3.tar.gz", "c:\\box\\file\\agentServer-1.6.3.tar.gz")
	if nil != err {
		fmt.Printf("getFile error = %s\n", err.Error())
	}
}

func TestGet(t *testing.T) {
	r, e := Get("https://api.i5sesol.com/agentServer/verify/mqtt", map[string]string{"machineNo": "BOX0318060106", "macAddr": "1", "licenseKey": "1"})
	if nil != e {
		fmt.Printf("get error = %s\n", e.Error())
	} else {
		fmt.Printf("get result = %s", string(r))
	}
}

func TestParseQuery(t *testing.T) {
	url, e := ParseQuery("https://api.i5sesol.com/agentServer/verify/mqtt", map[string]string{"macAddr": "1", "licenseKey": "1", "machineNo": "BOX0318060106"})
	if nil != e {
		fmt.Printf("parseQuery error = %s\n", e.Error())
	} else {
		fmt.Printf("parseQuery result = %s\n", url)
	}
}

func TestPostUrl(t *testing.T) {
	/*
	 * 查询手机号码信息
	 *
	 * Params:
	 *      - url : https://tcc.taobao.com/cc/json/mobile_tel_segment.htm?tel=号码
	 *
	 * Return:
	 *     - []byte
	 *     - error
	 */
	result, err := PostUrl("https://tcc.taobao.com/cc/json/mobile_tel_segment.htm?tel=17301851831")
	if nil != err {
		fmt.Printf("post url error = %s\n", err)
	} else {
		fmt.Println(string(result))
	}
}

func TestPostForm(t *testing.T) {
	r, e := PostForm("https://trust.baidu.com/vstar/feedback/getcaptchaajax", map[string]string{"cellphone": "17002777731", "type": ""})
	if nil != e {
		fmt.Printf("post form error = %s\n", e.Error())
	} else {
		fmt.Println(string(r))
	}
}

func TestPostFile(t *testing.T) {

}

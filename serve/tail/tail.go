package tail

import (
	m "github.com/gorilla/mux"
	"github.com/hpcloud/tail"
	"github.com/woshihot/go-lib/serve/mux"
	"github.com/woshihot/go-lib/utils/log"
	"golang.org/x/net/websocket"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var grepKey = "grep"

func CreateTailHandler(tailPath, tailRoutePath, followRoutePath, grepSuffix string) mux.MuxHandlers {
	result := make(mux.MuxHandlers)
	result.Append(tailRoutePath, genTailHandler(tailPath, followRoutePath, grepSuffix))
	result.Append(tailRoutePath+grepSuffix+"/{"+grepKey+"}", genTailHandler(tailPath, followRoutePath, grepSuffix))
	result.Append(followRoutePath, getFollowHandler(tailPath))
	result.Append(followRoutePath+grepSuffix+"/{"+grepKey+"}", getFollowHandler(tailPath))

	return result
}

func genTailHandler(tailFile, followRoutePath, grepSuffix string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := "tail -f " + tailFile
		followPath := "ws://" + r.Host + followRoutePath
		grep := m.Vars(r)[grepKey]
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t := template.Must(template.New("base").Parse(index))
		if "" != grep {
			path += " |grep " + grep
			followPath += grepSuffix + "/" + grep
		}
		v := struct {
			Host       string
			Log        string
			FollowPath string
		}{
			r.Host,
			path,
			followPath,
		}
		if err := t.Execute(w, &v); err != nil {
			log.Ef("Template execute failed, err: %v", err)
			return
		}
	})
}

func getFollowHandler(tailFile string) http.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		grep := m.Vars(ws.Request())[grepKey]
		DefaultLogger := log.New(os.Stderr, log.Verbose, log.LstdFlags)
		tailConfig := tail.Config{
			Poll:   true,
			Follow: true,
			ReOpen: true,
			Logger: DefaultLogger,
			Location: &tail.SeekInfo{
				Offset: 0,
				Whence: os.SEEK_END,
			},
		}
		t, err := tail.TailFile(tailFile, tailConfig)
		if err != nil {
			log.Ef("[tailFile-error] %s\n", err)
			return
		}
		for line := range t.Lines {
			lineStr := line.Text
			if "" == grep || strings.Contains(lineStr, grep) {
				ws.Write([]byte(line.Text))
			}

		}

	})
}

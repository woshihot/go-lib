package mux

import (
	"github.com/gorilla/mux"
	"github.com/hpcloud/tail"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"github.com/woshihot/go-lib/serve"
	"github.com/woshihot/go-lib/utils/log"
	"golang.org/x/net/websocket"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func Init(port string, handles map[string]http.Handler) {
	n := negroni.New()
	r := mux.NewRouter().StrictSlash(true)
	if nil != handles {
		for key, handle := range handles {
			r.Handle(key, handle)
		}
	}
	r.HandleFunc("/tailg/{grep}", HandleTail)
	r.HandleFunc("/tail", HandleTail)
	r.HandleFunc("/follow", websocket.Handler(HandleFollow).ServeHTTP)
	r.HandleFunc("/followg/{grep}", websocket.Handler(HandleFollow).ServeHTTP)

	h := cors.AllowAll().Handler(r)
	recovery := negroni.NewRecovery()
	recovery.Formatter = &panicFormatter{}
	n.Use(recovery)
	n.UseHandler(h)
	n.Run(":" + port)

}

type panicFormatter struct {
}

func (formatter *panicFormatter) FormatPanicError(w http.ResponseWriter, r *http.Request, info *negroni.PanicInformation) {

	serve.Render.JSON(w, http.StatusInternalServerError, map[string]interface{}{})
}

func HandleTail(w http.ResponseWriter, r *http.Request) {
	path := "tail -f " + tailFile
	grep := mux.Vars(r)["grep"]
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t := template.Must(template.New("base").Parse(index))
	if "" != grep {
		path += " |grep " + grep
	}
	v := struct {
		Host string
		Log  string
		Grep string
	}{
		r.Host,
		path,
		grep,
	}
	if err := t.Execute(w, &v); err != nil {
		log.Ef("Template execute failed, err: %v", err)
		return
	}
}

func HandleFollow(ws *websocket.Conn) {
	log.D("ws connected")
	grep := mux.Vars(ws.Request())["grep"]
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
}

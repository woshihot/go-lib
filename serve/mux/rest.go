package mux

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"github.com/woshihot/go-lib/serve"
	"net/http"
)

func Init(port string, handles ...map[string]http.Handler) {
	n := negroni.New()
	r := mux.NewRouter().StrictSlash(true)
	if nil != handles {
		for _, hs := range handles {
			for key, handle := range hs {
				r.Handle(key, handle)
			}
		}
	}
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

type MuxHandlers map[string]http.Handler

func (m1 MuxHandlers) Merge(handlers MuxHandlers) {

	for p, h := range handlers {
		m1.Append(p, h)
	}
}

func (m1 MuxHandlers) Append(routePath string, handler http.Handler) {
	m1[routePath] = handler
}

package mux

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"github.com/woshihot/go-lib/serve"
	"github.com/woshihot/go-lib/utils/log"
	"net/http"
	"strconv"
)

func Init(o Options) {
	port := strconv.Itoa(o.Port)
	if "0" == port {
		log.EF(TAG_ERROR, "service port can not be zero\n")
		return
	}
	n := negroni.New()
	r := mux.NewRouter().StrictSlash(true)
	if nil != o.PathRouterHandlers {
		for key, handle := range o.PathRouterHandlers {
			r.Handle(key, handle)
		}
	}
	if nil != o.PrefixRouterHandlers {
		for key, handle := range o.PrefixRouterHandlers {
			r.PathPrefix(key).Handler(handle)
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

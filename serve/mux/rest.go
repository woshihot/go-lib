package mux

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"github.com/woshihot/go-lib/serve"
	"net/http"
	"strconv"
)

func Init(router *mux.Router, port int) {
	portStr := strconv.Itoa(port)
	// if "0" == port {
	// 	log.EF(TAG_ERROR, "service port can not be zero\n")
	// 	return
	// }
	n := negroni.New()
	// r := mux.NewRouter().StrictSlash(true)
	// if nil != o.PathRouterHandlers {
	// 	for key, handle := range o.PathRouterHandlers {
	// 		r.Handle(key, handle)
	// 	}
	// }
	// if nil != o.PrefixRouterHandlers {
	// 	for key, handle := range o.PrefixRouterHandlers {
	// 		r.PathPrefix(key).Handler(handle).Methods()
	// 	}
	// }
	h := cors.AllowAll().Handler(router)
	recovery := negroni.NewRecovery()
	recovery.Formatter = &panicFormatter{}
	n.Use(recovery)
	n.UseHandler(h)
	n.Run(":" + portStr)

}

func NewRouter() *mux.Router {
	return mux.NewRouter()
}

type panicFormatter struct {
}

func (formatter *panicFormatter) FormatPanicError(w http.ResponseWriter, r *http.Request, info *negroni.PanicInformation) {

	serve.Render.JSON(w, http.StatusInternalServerError, map[string]interface{}{})
}

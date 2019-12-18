package serve

import "github.com/unrolled/render"

func New() *render.Render {
	return render.New()
}

var Render *render.Render

func init() {
	Render = New()
}

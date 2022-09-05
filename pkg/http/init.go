package http

import (
	"discount/database"
	"discount/database/mysql"
	"discount/models"
	"discount/pkg/http/handlers"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

type Route struct {
	Path     string
	Method   string
	Handlers []martini.Handler
	Param    interface{}
}

var httpRoutes []Route = []Route{
	{"/discount/submit", http.MethodPost, []martini.Handler{handlers.SubmitDiscount}, models.SubmitDiscountParam{}},
	{"/discount", http.MethodPut, []martini.Handler{handlers.AddDiscount}, models.AddDiscountParam{}},
	{"/discount/:id/report", http.MethodGet, []martini.Handler{handlers.GetDiscoutUsageReport}, nil},
}

func Init() {
	db := database.IDatabase(&mysql.MysqlDatabase{})
	c := martini.Classic()
	c.Use(render.Renderer())
	c.Map(db)

	for _, r := range httpRoutes {
		ParamHandlerConcat := make([]martini.Handler, 0)
		if r.Param != nil {
			ParamHandlerConcat = append(ParamHandlerConcat, binding.Bind(r.Param))
		}
		ParamHandlerConcat = append(ParamHandlerConcat, r.Handlers...)
		c.AddRoute(r.Method, r.Path, ParamHandlerConcat...)
	}

	db.Init()
	c.Run()
}

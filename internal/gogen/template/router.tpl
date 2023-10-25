package router

import (
	"github.com/gin-gonic/gin"
	"zeuscloud/internal/{{.appName}}/controller"
)

func (r *Router) init{{.Name}}Router(routerGroup *gin.RouterGroup, api *controller.{{.Name}}, pathcomp string) {
	r{{.Name}} := routerGroup.Group(pathcomp)
	{
		r{{.Name}}.GET("", api.Query)
		r{{.Name}}.GET(":id", api.Get)
		r{{.Name}}.POST("", api.Create)
		r{{.Name}}.PUT(":id", api.Update)
		r{{.Name}}.DELETE(":id", api.Delete)
	}
}


package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"fengyin/internal/{{lowCamelCase .appName}}/dto"
	"fengyin/internal/{{lowCamelCase .appName}}/service"
	"fengyin/pkg/api"
	"fengyin/pkg/global"
	"fengyin/pkg/handler"
)

// {{.controllerName}}Set 注入
var {{.controllerName}}Set = wire.NewSet(wire.Struct(new({{.controllerName}}), "*"))

type {{.controllerName}} struct {
	{{.appName}}Srv *service.{{.appName}}Service
}

// List 查询列表
func (s {{.controllerName}}) List(c *gin.Context) {
	ctx := handler.NewContext(c)
	defer func() { handler.JSONResponse(c, ctx) }()
	var param dto.List{{.controllerName}}Req
	valid, errs := api.BindAndValid(c, &param)
	if !valid {
		global.LOG.Infof("app.BindAndValid errs: %v", errs)
		ctx.Err = errs
		return
	}
	ctx.Resp, ctx.Err = s.{{.appName}}Srv.List{{.controllerName}}(c, param)
}

// Get 查询指定数据
func (s {{.controllerName}}) Get(c *gin.Context) {
	ctx := handler.NewContext(c)
	defer func() { handler.JSONResponse(c, ctx) }()
	var param dto.Get{{.controllerName}}Req
	valid, errs := api.BindAndValid(c, &param)
	if !valid {
		global.LOG.Infof("app.BindAndValid errs: %v", errs)
		ctx.Err = errs
		return
	}
	ctx.Resp, ctx.Err = s.{{.appName}}Srv.Get{{.controllerName}}(c, param)
}

// Create 创建数据
func (s {{.controllerName}}) Create(c *gin.Context) {
	ctx := handler.NewContext(c)
	defer func() { handler.JSONResponse(c, ctx) }()
	var param dto.Create{{.controllerName}}Req
	valid, errs := api.BindAndValid(c, &param)
	if !valid {
		global.LOG.Infof("app.BindAndValid errs: %v", errs)
		ctx.Err = errs
		return
	}
	ctx.Resp, ctx.Err = s.{{.appName}}Srv.Create{{.controllerName}}(c, param)
}

// Update 更新数据
func (s {{.controllerName}}) Update(c *gin.Context) {
	ctx := handler.NewContext(c)
	defer func() { handler.JSONResponse(c, ctx) }()

}

// Delete 删除数据
func (s {{.controllerName}}) Delete(c *gin.Context) {
	ctx := handler.NewContext(c)
	defer func() { handler.JSONResponse(c, ctx) }()
}

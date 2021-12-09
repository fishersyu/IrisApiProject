package perm

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin/server/web/web_iris/middleware"
)

// Party 权限模块
// - 定义模型 struct
// - 定义请求 struct
// - 定义响应 struct
// - 定义 module
// - 实现控制器逻辑
func Party() func(index iris.Party) {
	return func(index iris.Party) {
		index.Use(middleware.MultiHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Get("/", GetAll).Name = "权限列表"
		index.Get("/{id:uint}", First).Name = "权限详情"
		index.Post("/", CreatePerm).Name = "创建权限"
		index.Post("/{id:uint}", UpdatePerm).Name = "编辑权限"
		index.Delete("/{id:uint}", DeletePerm).Name = "删除权限"
	}
}

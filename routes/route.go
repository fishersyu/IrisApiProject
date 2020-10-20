package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/IrisAdminApi/config"
	"github.com/snowlyg/IrisAdminApi/controllers"
	"github.com/snowlyg/IrisAdminApi/middleware"
	"github.com/snowlyg/IrisAdminApi/sysinit"
)

func App(api *iris.Application) {
	api.UseRouter(middleware.CrsAuth())
	app := api.Party("/").AllowMethods(iris.MethodOptions)
	{
		// 二进制模式 ， 启用项目入口
		if config.Config.Bindata {
			app.Get("/", func(ctx iris.Context) { // 首页模块
				_ = ctx.View("index.html")
			})
		}

		admin := app.Party("/v1")
		{
			admin.Post("/admin/login", controllers.UserLogin)
			admin.PartyFunc("/admin", func(app iris.Party) {
				casbinMiddleware := middleware.New(sysinit.Enforcer)               //casbin for gorm                                                   // <- IMPORTANT, register the middleware.
				app.Use(middleware.JwtHandler().Serve, casbinMiddleware.ServeHTTP) //登录验证
				app.Post("/logout", controllers.UserLogout).Name = "退出"
				app.Get("/profile", controllers.GetProfile).Name = "个人信息"

				app.PartyFunc("/users", func(users iris.Party) {
					users.Get("/", controllers.GetAllUsers).Name = "用户列表"
					users.Get("/{id:uint}", controllers.GetUser).Name = "用户详情"
					users.Post("/", controllers.CreateUser).Name = "创建用户"
					users.Put("/{id:uint}", controllers.UpdateUser).Name = "编辑用户"
					users.Delete("/{id:uint}", controllers.DeleteUser).Name = "删除用户"

				})
				app.PartyFunc("/roles", func(roles iris.Party) {
					roles.Get("/", controllers.GetAllRoles).Name = "角色列表"
					roles.Get("/{id:uint}", controllers.GetRole).Name = "角色详情"
					roles.Post("/", controllers.CreateRole).Name = "创建角色"
					roles.Put("/{id:uint}", controllers.UpdateRole).Name = "编辑角色"
					roles.Delete("/{id:uint}", controllers.DeleteRole).Name = "删除角色"
				})
				app.PartyFunc("/permissions", func(permissions iris.Party) {
					permissions.Get("/", controllers.GetAllPermissions).Name = "权限列表"
					permissions.Get("/{id:uint}", controllers.GetPermission).Name = "权限详情"
					permissions.Post("/", controllers.CreatePermission).Name = "创建权限"
					permissions.Put("/{id:uint}", controllers.UpdatePermission).Name = "编辑权限"
					permissions.Delete("/{id:uint}", controllers.DeletePermission).Name = "删除权限"
				})
			})
		}
	}

}

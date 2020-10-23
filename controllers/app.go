package controllers

import (
	"github.com/snowlyg/blog/libs"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/blog/models"
	"github.com/snowlyg/blog/validates"
)

/**
* @api {post} /admin/login 用户登陆
* @apiName 用户登陆
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 用户登陆
* @apiSampleRequest /admin/login
* @apiParam {string} username 用户名
* @apiParam {string} password 密码
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func UserLogin(ctx iris.Context) {
	aul := new(validates.LoginRequest)

	if err := ctx.ReadJSON(aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(ApiResource(400, nil, err.Error()))
		return
	}

	err := validates.Validate.Struct(*aul)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.ValidateTrans) {
			if len(e) > 0 {
				ctx.StatusCode(iris.StatusOK)
				_, _ = ctx.JSON(ApiResource(400, nil, e))
				return
			}
		}
	}

	ctx.Application().Logger().Infof("%s 登录系统", aul.Username)
	ctx.StatusCode(iris.StatusOK)

	user, err := models.GetUserByUsername(aul.Username)
	if err != nil {
		_, _ = ctx.JSON(ApiResource(400, nil, err.Error()))
		return
	}

	response, code, msg := user.CheckLogin(aul.Password)

	_, _ = ctx.JSON(ApiResource(code, response, msg))
	return

}

/**
* @api {get} /logout 用户退出登陆
* @apiName 用户退出登陆
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 用户退出登陆
* @apiSampleRequest /logout
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func UserLogout(ctx iris.Context) {
	value := ctx.Values().Get("jwt").(*jwt.Token)
	conn := libs.GetRedisClusterClient()
	defer conn.Close()
	sess, err := models.GetRedisSessionV2(conn, value.Raw)
	if err != nil {
		ctx.StatusCode(http.StatusOK)
		_, _ = ctx.JSON(ApiResource(400, nil, err.Error()))
	}
	if sess != nil {
		if err := sess.DelUserTokenCache(conn, value.Raw); err != nil {
			ctx.StatusCode(http.StatusOK)
			_, _ = ctx.JSON(ApiResource(400, nil, err.Error()))
		}
		ctx.Application().Logger().Infof("%d 退出系统", sess.UserId)
	}

	ctx.StatusCode(http.StatusOK)
	_, _ = ctx.JSON(ApiResource(200, nil, "退出"))
}

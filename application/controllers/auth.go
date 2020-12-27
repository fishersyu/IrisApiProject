package controllers

import (
	"fmt"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/jameskeane/bcrypt"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/blog/application/libs"
	"github.com/snowlyg/blog/application/libs/easygorm"
	"github.com/snowlyg/blog/application/libs/logging"
	"github.com/snowlyg/blog/application/libs/response"
	"github.com/snowlyg/blog/application/models"
	"github.com/snowlyg/blog/service/auth"
	"strings"
)

type LoginRe struct {
	Username string `json:"username" validate:"required,gte=2,lte=50" comment:"用户名"`
	Password string `json:"password" validate:"required,gte=6,lte=30"  comment:"密码"`
}

type User struct {
	Id       uint
	Username string
	Password string
}

type Token struct {
	AccessToken string
}

func Login(ctx iris.Context) {
	loginReq := new(LoginRe)
	if err := ctx.ReadJSON(loginReq); err != nil {
		logging.ErrorLogger.Errorf("login read request json err ", err)
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, response.SystemErr.Msg))
		return
	}

	logging.DebugLogger.Debugf("login user ", loginReq)

	validErr := libs.Validate.Struct(*loginReq)
	errs := libs.ValidRequest(validErr)
	if len(errs) > 0 {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, strings.Join(errs, ";")))
		return
	}

	user := User{Username: loginReq.Username}
	err := easygorm.EasyGorm.DB.Model(models.User{}).Find(&user).Error
	if err != nil {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, err.Error()))
		return
	}

	logging.DebugLogger.Debugf("user", user)

	if user.Id == 0 {
		ctx.JSON(response.NewResponse(response.DataEmptyErr.Code, nil, fmt.Sprintf("用户 %s 不存在", user.Username)))
		return
	}

	if ok := bcrypt.Match(loginReq.Password, user.Password); !ok {
		ctx.JSON(response.NewResponse(response.AuthErr.Code, nil, "用户名或密码错误"))
		return
	}

	authDriver := auth.NewAuthDriver()
	defer authDriver.Close()

	var token string
	token, err = auth.Login(authDriver, uint64(user.Id))
	if err != nil {
		ctx.JSON(response.NewResponse(response.AuthErr.Code, nil, err.Error()))
		return
	}

	logging.DebugLogger.Debugf("user token %s", token)

	ctx.JSON(response.NewResponse(response.NoErr.Code, &Token{AccessToken: token}, response.NoErr.Msg))
	return
}

func Logout(ctx iris.Context) {
	value := ctx.Values().Get("jwt").(*jwt.Token)
	authDriver := auth.NewAuthDriver()
	defer authDriver.Close()

	err := auth.Logout(authDriver, value.Raw)
	if err != nil {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, err.Error()))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, nil, response.NoErr.Msg))
	return
}

func Expire(ctx iris.Context) {
	value := ctx.Values().Get("jwt").(*jwt.Token)
	authDriver := auth.NewAuthDriver()
	defer authDriver.Close()
	if err := auth.Expire(authDriver, value.Raw); err != nil {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, err.Error()))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, nil, response.NoErr.Msg))
	return
}

func Clear(ctx iris.Context) {
	value := ctx.Values().Get("jwt").(*jwt.Token)
	authDriver := auth.NewAuthDriver()
	defer authDriver.Close()
	if err := auth.Clear(authDriver, value.Raw); err != nil {
		ctx.JSON(response.NewResponse(response.SystemErr.Code, nil, err.Error()))
		return
	}
	ctx.JSON(response.NewResponse(response.NoErr.Code, nil, response.NoErr.Msg))
	return
}

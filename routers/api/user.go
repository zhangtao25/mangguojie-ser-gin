package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangtao25/mangostreet-ser-gin/pkg/app"
	"github.com/zhangtao25/mangostreet-ser-gin/pkg/e"
	"github.com/zhangtao25/mangostreet-ser-gin/service/user_service"
	"net/http"
)


type AuthUsersByVerificationCodeForm struct {
	Username       string `form:"username" valid:"Required;MaxSize(100)"`
	Vcode          string `form:"vcode"    valid:"Required;MaxSize(255)"`
}


func AuthUsersByVerificationCode(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = AuthUsersByVerificationCodeForm{}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userService := user_service.User{
		Username:form.Username,
		Vcode:form.Vcode,
	}

	exists,err := userService.ExistByUsername()
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	token, err := userService.Get()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, token)
}
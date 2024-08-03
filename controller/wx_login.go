package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/pkg/utils"
	"sgin/service"
)

type WxLoginController struct {
	VerificationCodeService *service.VerificationCodeService
	WxUserService           *service.WxUserService
}

// 验证码登录
func (w *WxLoginController) VerificationCodeLoginPhone(ctx *app.Context) {
	param := &model.ReqVerificationCodeParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	if param.Email == "" && param.Phone == "" {
		ctx.JSONError(http.StatusBadRequest, "邮箱和手机号码不能同时为空")
		return
	}

	if param.Code == "" {
		ctx.JSONError(http.StatusBadRequest, "验证码不能为空")
		return
	}

	ok, rcode, err := w.VerificationCodeService.CheckVerificationCode(ctx, param.Code, param.Email, param.Phone)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	if ok == false {
		ctx.JSONError(http.StatusBadRequest, "验证码错误")
		return
	}

	// 更新验证码状态
	err = w.VerificationCodeService.UpdateVerificationCode(ctx, rcode.UUID)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	rwxUser, err := w.WxUserService.GetWxUserOrCreateByPhone(ctx, &model.WxUser{Phone: param.Phone})
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	rtoken, err := utils.GenerateWxUserToken(rwxUser.Uuid)

	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(rtoken)

}

func (w *WxLoginController) LoginByPassword(ctx *app.Context) {
	param := &model.ReqWxUserLoginByPasswordParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	if param.Phone == "" {
		ctx.JSONError(http.StatusBadRequest, "手机号码不能为空")
		return
	}

	if param.Password == "" {
		ctx.JSONError(http.StatusBadRequest, "密码不能为空")
		return
	}
	param.Password = utils.HashPasswordWithSalt(param.Password, ctx.Config.PasswdKey)

	rwxUser, err := w.WxUserService.GetWxUserByPhoneAndPassword(ctx, param.Phone, param.Password)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	rtoken, err := utils.GenerateWxUserToken(rwxUser.Uuid)

	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(rtoken)
}

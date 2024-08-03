package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/pkg/utils"
	"sgin/service"
)

type WxUserController struct {
	WxUserService           *service.WxUserService
	VerificationCodeService *service.VerificationCodeService
}

func (c *WxUserController) GetMyWxUserInfo(ctx *app.Context) {

	userId := ctx.GetString("wx_user_id")

	if userId == "" {
		ctx.JSONError(http.StatusBadRequest, "请先登录")
		return
	}
	wxUser, err := c.WxUserService.GetWxUserByUUID(ctx, userId)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(wxUser)
}

func (c *WxUserController) ChangePassword(ctx *app.Context) {

	param := &model.ReqWxUserChangePasswordParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	if param.Phone == "" {
		ctx.JSONError(http.StatusBadRequest, "手机号码不能为空")
		return
	}

	if param.Code == "" {
		ctx.JSONError(http.StatusBadRequest, "验证码不能为空")
		return
	}

	ok, rcode, err := c.VerificationCodeService.CheckVerificationCode(ctx, param.Code, "", param.Phone)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	if ok == false {
		ctx.JSONError(http.StatusBadRequest, "验证码错误")
		return
	}

	userId := ctx.GetString("wx_user_id")

	if userId == "" {
		ctx.JSONError(http.StatusBadRequest, "请先登录")
		return
	}
	wxUser, err := c.WxUserService.GetWxUserByUUID(ctx, userId)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	// 更新验证码状态
	err = c.VerificationCodeService.UpdateVerificationCode(ctx, rcode.UUID)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	wxUser.Password = utils.HashPasswordWithSalt(param.Password, ctx.Config.PasswdKey)

	err = c.WxUserService.UpdateWxUser(ctx, wxUser)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess("修改密码成功")

}

// 获取所有的微信用户
func (c *WxUserController) GetWxUserList(ctx *app.Context) {

	wxUsers, err := c.WxUserService.GetAllWxUsers(ctx)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(wxUsers)
}

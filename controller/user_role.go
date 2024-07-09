package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type UserRoleController struct {
	UserRoleService *service.UserRoleService
}

// @Summary 创建用户角色
// @Description 创建用户角色
// @Tags 用户角色
// @Accept json
// @Produce json
// @Param params body model.ReqUserRole true "用户角色信息"
// @Success 200 {string} string "Successfully fetched user data"
// @Router /api/v1/user_role/create [post]
func (u *UserRoleController) CreateUserRole(ctx *app.Context) {
	userRole := &model.ReqUserRole{}
	if err := ctx.ShouldBindJSON(userRole); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	err := u.UserRoleService.CreateUserRole(ctx, userRole)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("Successfully created user role")
}

// @Summary 删除用户角色
// @Description 删除用户角色
// @Tags 用户角色
// @Accept json
// @Produce json
// @Param params body model.ReqUserRole true "用户角色信息"
// @Success 200 {string} string "Successfully fetched user data"
// @Router /api/v1/user_role/delete [post]
func (u *UserRoleController) DeleteUserRole(ctx *app.Context) {
	userRole := &model.ReqUserRole{}
	if err := ctx.ShouldBindJSON(userRole); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	err := u.UserRoleService.DeleteUserRole(ctx, userRole)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("Successfully deleted user role")
}

// @Summary 获取用户自身角色
// @Description 获取用户自身角色
// @Tags 用户角色
// @Accept json
// @Produce json
// @Param user_uuid query string true "用户UUID"
// @Success 200 {string} string "Successfully fetched user data"
// @Router /api/v1/user_role/getme [get]
func (u *UserRoleController) GetMeUserRole(ctx *app.Context) {
	userUuid := ctx.GetString("user_id")
	userRoles, err := u.UserRoleService.GetUserRoleByUserID(ctx, userUuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(userRoles)
}

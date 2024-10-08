package service

import (
	"errors"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/pkg/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VerificationCodeService struct {
}

// CreateVerificationCode 创建验证码
func (v *VerificationCodeService) CreateVerificationCode(ctx *app.Context, email string, phone string) (string, error) {

	// 先获取最新的一条验证码是否过期
	var vcode model.VerificationCode
	err := ctx.DB.Where("email = ? OR phone = ?", email, phone).Order("created_at desc").First(&vcode).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.Logger.Error("Failed to get verification code", err)
		return "", err
	}

	if err == nil && vcode.Status == 0 && time.Now().Sub(vcode.CreatedAt).Minutes() < 1 {
		err := errors.New("验证码已发送，请稍后再试")
		return "", err
	}

	code := utils.GenerateVerificationCode()

	vcode1 := model.VerificationCode{
		UUID:      uuid.New().String(),
		Code:      code,
		Email:     email,
		Phone:     phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    0,
	}

	err = ctx.DB.Create(&vcode1).Error
	return code, err

}

// CheckVerificationCode 检查验证码
func (v *VerificationCodeService) CheckVerificationCode(ctx *app.Context, code string, email string, phone string) (bool, model.VerificationCode, error) {

	var vcode model.VerificationCode
	err := ctx.DB.Where("code = ? AND (email = ? OR phone = ?)", code, email, phone).Order("id DESC").First(&vcode).Error
	if err != nil {
		return false, vcode, err
	}

	if vcode.Status == 1 {
		return false, vcode, nil
	}

	if time.Now().Sub(vcode.CreatedAt).Minutes() > 5 {
		return false, vcode, nil
	}

	return true, vcode, nil
}

// UpdateVerificationCode 更新验证码
func (v *VerificationCodeService) UpdateVerificationCode(ctx *app.Context, uuidstr string) error {

	err := ctx.DB.Model(&model.VerificationCode{}).Where("uuid = ?", uuidstr).Update("status", 1).Error

	if err != nil {
		ctx.Logger.Error("Failed to update verification code", err)
		return err
	}

	return err
}

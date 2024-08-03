package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WxUserService struct {
}

func NewWxUserService() *WxUserService {
	return &WxUserService{}
}

func (s *WxUserService) GetWxUserOrCreateByPhone(ctx *app.Context, wxUser *model.WxUser) (*model.WxUser, error) {
	var wxUser1 model.WxUser
	err := ctx.DB.Where("phone = ?", wxUser.Phone).First(&wxUser1).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.Logger.Error("Failed to get wxUser", err)
		return nil, errors.New("failed to get wxUser")
	}

	if err == gorm.ErrRecordNotFound {
		wxUser.Uuid = uuid.New().String()
		wxUser.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		wxUser.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		err = ctx.DB.Create(wxUser).Error
		if err != nil {
			ctx.Logger.Error("Failed to create wxUser", err)
			return nil, errors.New("failed to create wxUser")
		}
		return wxUser, nil
	}

	return &wxUser1, nil
}

func (s *WxUserService) GetWxUserByPhoneAndPassword(ctx *app.Context, phone, password string) (*model.WxUser, error) {
	wxUser := &model.WxUser{}
	err := ctx.DB.Where("phone = ? AND password = ?", phone, password).First(wxUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("wxUser not found")
		}
		ctx.Logger.Error("Failed to get wxUser by phone and password", err)
		return nil, errors.New("failed to get wxUser by phone and password")
	}
	return wxUser, nil
}

func (s *WxUserService) CreateWxUser(ctx *app.Context, wxUser *model.WxUser) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	wxUser.CreatedAt = now
	wxUser.UpdatedAt = now
	wxUser.Uuid = uuid.New().String()

	err := ctx.DB.Create(wxUser).Error
	if err != nil {
		ctx.Logger.Error("Failed to create wxUser", err)
		return errors.New("failed to create wxUser")
	}
	return nil
}

func (s *WxUserService) GetWxUserByUUID(ctx *app.Context, uuid string) (*model.WxUser, error) {
	wxUser := &model.WxUser{}
	err := ctx.DB.Where("uuid = ?", uuid).First(wxUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("wxUser not found")
		}
		ctx.Logger.Error("Failed to get wxUser by UUID", err)
		return nil, errors.New("failed to get wxUser by UUID")
	}
	return wxUser, nil
}

func (s *WxUserService) UpdateWxUser(ctx *app.Context, wxUser *model.WxUser) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	wxUser.UpdatedAt = now
	err := ctx.DB.Where("uuid = ?", wxUser.Uuid).Updates(wxUser).Error
	if err != nil {
		ctx.Logger.Error("Failed to update wxUser", err)
		return errors.New("failed to update wxUser")
	}

	return nil
}

func (s *WxUserService) DeleteWxUser(ctx *app.Context, uuid string) error {
	err := ctx.DB.Model(&model.WxUser{}).Where("uuid = ?", uuid).Update("is_deleted", 1).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete wxUser", err)
		return errors.New("failed to delete wxUser")
	}

	return nil
}

func (s *WxUserService) GetWxUserList(ctx *app.Context, params *model.ReqWxUserQueryParam) (*model.PagedResponse, error) {
	var (
		wxUsers []*model.WxUser
		total   int64
	)

	db := ctx.DB.Model(&model.WxUser{})

	if params.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%"+params.NickName+"%")
	}

	db = db.Where("is_deleted = ?", 0)

	err := db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to get wxUser count", err)
		return nil, errors.New("failed to get wxUser count")
	}

	err = db.Offset(params.GetOffset()).Limit(params.PageSize).Find(&wxUsers).Error
	if err != nil {
		ctx.Logger.Error("Failed to get wxUser list", err)
		return nil, errors.New("failed to get wxUser list")
	}

	return &model.PagedResponse{
		Total: total,
		Data:  wxUsers,
	}, nil
}

// 根据uuid列表获取用户列表
func (s *WxUserService) GetWxUserListByUUIDs(ctx *app.Context, uuids []string) ([]*model.WxUser, error) {
	var wxUsers []*model.WxUser
	err := ctx.DB.Where("uuid IN ?", uuids).Find(&wxUsers).Error
	if err != nil {
		ctx.Logger.Error("Failed to get wxUser list by uuids", err)
		return nil, errors.New("failed to get wxUser list by uuids")
	}
	return wxUsers, nil
}

// 获取所有的微信用户
func (s *WxUserService) GetAllWxUsers(ctx *app.Context) ([]*model.WxUser, error) {
	var wxUsers []*model.WxUser
	err := ctx.DB.Find(&wxUsers).Error
	if err != nil {
		ctx.Logger.Error("Failed to get all wxUsers", err)
		return nil, errors.New("failed to get all wxUsers")
	}
	return wxUsers, nil
}

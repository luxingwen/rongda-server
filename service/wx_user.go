package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"
	"sgin/pkg/utils"

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

		if wxUser.Password != "" {
			wxUser.Password = utils.HashPasswordWithSalt(wxUser.Password, ctx.Config.PasswdKey)
		}

		err = ctx.DB.Create(wxUser).Error
		if err != nil {
			ctx.Logger.Error("Failed to create wxUser", err)
			return nil, errors.New("failed to create wxUser")
		}
		return wxUser, nil
	}

	return &wxUser1, nil
}

// 根据openid获取用户信息
func (s *WxUserService) GetWxUserByOpenid(ctx *app.Context, openid string) (*model.WxUser, error) {
	var wxUser model.WxUser
	err := ctx.DB.Where("openid = ?", openid).First(&wxUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("wxUser not found")
		}
		ctx.Logger.Error("Failed to get wxUser by openid", err)
		return nil, errors.New("failed to get wxUser by openid")
	}
	return &wxUser, nil
}

func (s *WxUserService) GetWxUserOrCreateByOpenid(ctx *app.Context, openid string, req *model.ReqWXUserInfo) (*model.WxUser, error) {
	var wxUser model.WxUser
	err := ctx.DB.Where("openid = ?", openid).First(&wxUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.Logger.Error("Failed to get wxUser", err)
		return nil, errors.New("failed to get wxUser")
	}

	if err == gorm.ErrRecordNotFound {
		wxUser.Uuid = uuid.New().String()
		wxUser.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		wxUser.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		wxUser.Openid = openid
		wxUser.Unionid = req.UserInfo.UnionID
		wxUser.NickName = req.UserInfo.NickName
		wxUser.Avatar = req.UserInfo.AvatarUrl
		wxUser.Gender = req.UserInfo.Gender
		wxUser.City = req.UserInfo.City
		wxUser.Status = 1
		wxUser.IsDeleted = 0

		err = ctx.DB.Create(&wxUser).Error
		if err != nil {
			ctx.Logger.Error("Failed to create wxUser", err)
			return nil, errors.New("failed to create wxUser")
		}
		return &wxUser, nil
	}

	// 更新用户信息
	wxUser.NickName = req.UserInfo.NickName
	wxUser.Avatar = req.UserInfo.AvatarUrl
	wxUser.City = req.UserInfo.City

	err = ctx.DB.Where("openid = ?", openid).Updates(&wxUser).Error
	if err != nil {
		ctx.Logger.Error("Failed to update wxUser", err)
		return nil, errors.New("failed to update wxUser")
	}

	return &wxUser, nil
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

	wxUser.Password = ""
	return wxUser, nil
}

func (s *WxUserService) UpdateWxUser(ctx *app.Context, wxUser *model.WxUser) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	wxUser.UpdatedAt = now

	if wxUser.Password != "" {
		wxUser.Password = utils.HashPasswordWithSalt(wxUser.Password, ctx.Config.PasswdKey)
	}

	err := ctx.DB.Where("uuid = ?", wxUser.Uuid).Updates(wxUser).Error
	if err != nil {
		ctx.Logger.Error("Failed to update wxUser", err)
		return errors.New("failed to update wxUser")
	}

	return nil
}

func (s *WxUserService) DeleteWxUser(ctx *app.Context, uuid string) error {
	err := ctx.DB.Model(&model.WxUser{}).Where("uuid = ?", uuid).Delete(&model.WxUser{}).Error
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

	for _, wxUser := range wxUsers {
		wxUser.Password = ""
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

// UpdateWxUserIsRealNameAuth
func (s *WxUserService) UpdateWxUserIsRealNameAuth(ctx *app.Context, params *model.ReqWxUserUpdateIsRealNameAuthParam) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Model(&model.WxUser{}).Where("uuid = ?", params.Uuid).Updates(map[string]interface{}{
		"is_real_name": params.IsRealName,
		"updated_at":   now,
	}).Error

	if err != nil {
		ctx.Logger.Error("Failed to update wxUser is real name auth", err)
		return errors.New("failed to update wxUser is real name auth")
	}

	return nil
}

// JoinTeamByInviteCode
func (s *WxUserService) JoinTeamByInviteCode(ctx *app.Context, param *model.ReqInviteCodeParam) (*model.WxUser, error) {
	// 获取用户
	// 根据电话号码获取用户
	wxUser := &model.WxUser{}
	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		// 获取teamRef
		var teamRef model.TeamRef
		err := tx.Where("team_uuid = ? AND category = ?", param.TeamUuid, model.TeamCategoryCustomer).First(&teamRef).Error
		if err != nil {
			tx.Rollback()
			ctx.Logger.Error("Failed to get teamRef", err)
			return errors.New("failed to get teamRef")
		}

		err = tx.Where("phone = ?", param.Phone).First(wxUser).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// 创建用户
				wxUser.Uuid = uuid.New().String()
				wxUser.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
				wxUser.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
				wxUser.Phone = param.Phone
				wxUser.IsRealName = 0
				wxUser.IsDeleted = 0
				wxUser.Status = 1
				wxUser.Name = param.Name
				wxUser.Email = param.Email

				err = tx.Create(wxUser).Error
				if err != nil {
					tx.Rollback()
					ctx.Logger.Error("Failed to create wxUser", err)
					return errors.New("failed to create wxUser")
				}
			} else {
				tx.Rollback()
				ctx.Logger.Error("Failed to get wxUser by phone", err)
				return errors.New("failed to get wxUser by phone")
			}
		}

		// 先查询是否已经加入团队
		var teamMember model.TeamMember
		err = tx.Where("team_uuid = ? AND user_uuid = ?", teamRef.TeamUuid, wxUser.Uuid).First(&teamMember).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			tx.Rollback()
			ctx.Logger.Error("Failed to get team member", err)
			return errors.New("failed to get team member")
		}

		if teamMember.Id > 0 {
			tx.Rollback()
			return errors.New("已经加入过该团队")
		}

		// 加入团队
		teamMember = model.TeamMember{
			UUID:       uuid.New().String(),
			TeamUUID:   teamRef.TeamUuid,
			UserUUID:   wxUser.Uuid,
			Role:       "成员",
			Status:     "未审核",
			InviteCode: param.Code,
			CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		}
		err = tx.Create(&teamMember).Error
		if err != nil {
			tx.Rollback()
			ctx.Logger.Error("Failed to create team member", err)
			return errors.New("failed to create team member")
		}
		return nil
	})
	if err != nil {
		if err.Error() == "已经加入过该团队" {
			return wxUser, nil
		}
		ctx.Logger.Error("Failed to join team by invite code", err)
		return nil, errors.New("failed to join team by invite code")
	}

	return wxUser, nil
}

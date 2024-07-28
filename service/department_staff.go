package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepartmentStaffService struct {
}

func NewDepartmentStaffService() *DepartmentStaffService {
	return &DepartmentStaffService{}
}

func (s *DepartmentStaffService) CreateDepartmentStaff(ctx *app.Context, staff *model.DepartmentStaff) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	staff.CreatedAt = now
	staff.UpdatedAt = now
	staff.Uuid = uuid.New().String()

	// 先查询是否已经存在关系
	staffExist := &model.DepartmentStaff{}

	ctx.Logger.Info("staff.DepartmentUuid:", staff.DepartmentUuid)
	ctx.Logger.Info("staff.StaffUuid:", staff.StaffUuid)

	err := ctx.DB.Where("department_uuid = ? AND   staff_uuid = ?", staff.DepartmentUuid, staff.StaffUuid).First(staffExist).Error

	if err == nil && staffExist.Id > 0 {
		return errors.New("department staff already exists")
	}

	err = ctx.DB.Create(staff).Error
	if err != nil {
		ctx.Logger.Error("Failed to create department staff", err)
		return errors.New("failed to create department staff")
	}
	return nil
}

func (s *DepartmentStaffService) GetDepartmentStaffByUUID(ctx *app.Context, uuid string) (*model.DepartmentStaff, error) {
	staff := &model.DepartmentStaff{}
	err := ctx.DB.Where("uuid = ?", uuid).First(staff).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("department staff not found")
		}
		ctx.Logger.Error("Failed to get department staff by UUID", err)
		return nil, errors.New("failed to get department staff by UUID")
	}
	return staff, nil
}

func (s *DepartmentStaffService) UpdateDepartmentStaff(ctx *app.Context, staff *model.DepartmentStaff) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	staff.UpdatedAt = now
	err := ctx.DB.Where("uuid = ?", staff.Uuid).Updates(staff).Error
	if err != nil {
		ctx.Logger.Error("Failed to update department staff", err)
		return errors.New("failed to update department staff")
	}
	return nil
}

func (s *DepartmentStaffService) DeleteDepartmentStaff(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.DepartmentStaff{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete department staff", err)
		return errors.New("failed to delete department staff")
	}
	return nil
}

// 查询部门下的员工列表
func (s *DepartmentStaffService) DepartmentStaffList(ctx *app.Context, params *model.ReqDepartmentStaffQueryParam) (r *model.PagedResponse, err error) {
	var (
		staffList []*model.DepartmentStaff
		total     int64
	)

	db := ctx.DB.Model(&model.DepartmentStaff{})

	if params.DepartmentUuid != "" {
		db = db.Where("department_uuid = ?", params.DepartmentUuid)
	}

	err = db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to count department staff", err)
		return nil, errors.New("failed to count department staff")
	}

	err = db.Offset(params.GetOffset()).Limit(params.PageSize).Find(&staffList).Error
	if err != nil {
		ctx.Logger.Error("Failed to get department staff list", err)
		return nil, errors.New("failed to get department staff list")
	}

	return &model.PagedResponse{
		Total: total,
		Data:  staffList,
	}, nil
}

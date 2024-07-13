package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepartmentService struct {
}

func NewDepartmentService() *DepartmentService {
	return &DepartmentService{}
}

func (s *DepartmentService) CreateDepartment(ctx *app.Context, department *model.Department) error {
	department.UUID = uuid.New().String()

	department.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	department.UpdatedAt = department.CreatedAt

	err := ctx.DB.Create(department).Error
	if err != nil {
		ctx.Logger.Error("Failed to create department", err)
		return errors.New("failed to create department")
	}
	return nil
}

func (s *DepartmentService) GetDepartmentByUUID(ctx *app.Context, uuid string) (*model.Department, error) {
	department := &model.Department{}
	err := ctx.DB.Where("uuid = ?", uuid).First(department).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("department not found")
		}
		ctx.Logger.Error("Failed to get department by UUID", err)
		return nil, errors.New("failed to get department by UUID")
	}
	return department, nil
}

func (s *DepartmentService) UpdateDepartment(ctx *app.Context, department *model.Department) error {
	department.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Where("uuid = ?", department.UUID).Updates(department).Error
	if err != nil {
		ctx.Logger.Error("Failed to update department", err)
		return errors.New("failed to update department")
	}

	return nil
}

func (s *DepartmentService) DeleteDepartment(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.Department{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete department", err)
		return errors.New("failed to delete department")
	}

	return nil
}

// 获取部门列表
func (s *DepartmentService) GetDepartmentList(ctx *app.Context, param *model.ReqDepartmentQueryParam) (r *model.PagedResponse, err error) {
	var (
		departmentList []*model.Department
		total          int64
	)

	db := ctx.DB.Model(&model.Department{})

	if param.Name != "" {
		db = db.Where("name like ?", "%"+param.Name+"%")
	}

	if err = db.Find(&departmentList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     departmentList,
	}

	return
}

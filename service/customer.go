package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerService struct {
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) CreateCustomer(ctx *app.Context, customer *model.Customer) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	customer.CreatedAt = now
	customer.UpdatedAt = now
	customer.Uuid = uuid.New().String()

	err := ctx.DB.Create(customer).Error
	if err != nil {
		ctx.Logger.Error("Failed to create customer", err)
		return errors.New("failed to create customer")
	}
	return nil
}

func (s *CustomerService) GetCustomerByUUID(ctx *app.Context, uuid string) (*model.Customer, error) {
	customer := &model.Customer{}
	err := ctx.DB.Where("uuid = ?", uuid).First(customer).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("customer not found")
		}
		ctx.Logger.Error("Failed to get customer by UUID", err)
		return nil, errors.New("failed to get customer by UUID")
	}
	return customer, nil
}

func (s *CustomerService) UpdateCustomer(ctx *app.Context, customer *model.Customer) error {
	now := time.Now()
	customer.UpdatedAt = now.Format("2006-01-02 15:04:05")
	err := ctx.DB.Where("uuid = ?", customer.Uuid).Updates(customer).Error
	if err != nil {
		ctx.Logger.Error("Failed to update customer", err)
		return errors.New("failed to update customer")
	}

	return nil
}

func (s *CustomerService) DeleteCustomer(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.Customer{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete customer", err)
		return errors.New("failed to delete customer")
	}

	return nil
}

// GetCustomerList retrieves a list of customers based on query parameters
func (s *CustomerService) GetCustomerList(ctx *app.Context, params *model.ReqCustomerQueryParam) (*model.PagedResponse, error) {
	var (
		customers []*model.Customer
		total     int64
	)

	db := ctx.DB.Model(&model.Customer{})

	if params.Name != "" {
		db = db.Where("name LIKE ?", "%"+params.Name+"%")
	}

	err := db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to get customer count", err)
		return nil, errors.New("failed to get customer count")
	}

	err = db.Find(&customers).Error
	if err != nil {
		ctx.Logger.Error("Failed to get customer list", err)
		return nil, errors.New("failed to get customer list")
	}

	return &model.PagedResponse{
		Total: total,
		Data:  customers,
	}, nil
}
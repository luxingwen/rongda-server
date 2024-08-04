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

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(customer).Error
		if err != nil {
			ctx.Logger.Error("Failed to create customer", err)
			return errors.New("failed to create customer")
		}
		teamRef := model.TeamRef{
			TeamUuid:  customer.Uuid,
			Category:  model.TeamCategoryCustomer,
			CreatedAt: now,
			UpdatedAt: now,
		}
		err = tx.Create(&teamRef).Error
		if err != nil {
			ctx.Logger.Error("Failed to create teamRef", err)
			return errors.New("failed to create teamRef")
		}
		return nil
	})

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
	err := ctx.DB.Where("uuid = ?", uuid).Update("is_deleted", 1).Error
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

	db = db.Where("is_deleted = ?", 0)

	err := db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to get customer count", err)
		return nil, errors.New("failed to get customer count")
	}

	err = db.Offset(params.GetOffset()).Limit(params.PageSize).Find(&customers).Error
	if err != nil {
		ctx.Logger.Error("Failed to get customer list", err)
		return nil, errors.New("failed to get customer list")
	}

	return &model.PagedResponse{
		Total: total,
		Data:  customers,
	}, nil
}

// 获取所用可用的客户
func (s *CustomerService) GetAllCustomers(ctx *app.Context) ([]*model.Customer, error) {
	var customers []*model.Customer
	err := ctx.DB.Find(&customers).Error
	if err != nil {
		ctx.Logger.Error("Failed to get all customers", err)
		return nil, errors.New("failed to get all customers")
	}
	return customers, nil
}

// 根据uuid列表获取客户列表
func (s *CustomerService) GetCustomerListByUUIDs(ctx *app.Context, uuids []string) (map[string]*model.Customer, error) {
	var customers []*model.Customer
	err := ctx.DB.Where("uuid IN (?)", uuids).Find(&customers).Error
	if err != nil {
		ctx.Logger.Error("Failed to get customer list by UUIDs", err)
		return nil, errors.New("failed to get customer list by UUIDs")
	}

	customerMap := make(map[string]*model.Customer)
	for _, customer := range customers {
		customerMap[customer.Uuid] = customer
	}

	return customerMap, nil
}

// 获取客户订单
func (s *CustomerService) GetCustomerOrders(ctx *app.Context, params *model.ReqSalesOrderQueryParam) (*model.PagedResponse, error) {

	if params.CustomerUuid == "" {
		return nil, errors.New("customer uuid is required")
	}

	var (
		orders []*model.SalesOrder
		total  int64
	)

	db := ctx.DB.Model(&model.SalesOrder{})
	db = db.Where("customer_uuid = ?", params.CustomerUuid)

	if params.Status != "" {
		db = db.Where("status = ?", params.Status)
	}

	if params.StartDate != "" {
		db = db.Where("created_at >= ?", params.StartDate)
	}

	if params.EndDate != "" {
		db = db.Where("created_at <= ?", params.EndDate)
	}

	err := db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to get customer order count", err)
		return nil, errors.New("failed to get customer order count")
	}

	err = db.Offset(params.GetOffset()).Limit(params.PageSize).Find(&orders).Error
	if err != nil {
		ctx.Logger.Error("Failed to get customer order list", err)
		return nil, errors.New("failed to get customer order list")
	}

	purchaseOrderUuids := make([]string, 0)
	for _, order := range orders {
		purchaseOrderUuids = append(purchaseOrderUuids, order.PurchaseOrderNo)
	}

	purchaseOrderMap, err := NewPurchaseOrderService().GetPurchaseOrderListByOrderNos(ctx, purchaseOrderUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get purchase order list by order nos", err)
		return nil, errors.New("failed to get purchase order list by order nos")
	}

	res := make([]*model.CustomerSalesOrderRes, 0)
	for _, order := range orders {
		customerOrderResItem := model.CustomerSalesOrderRes{
			SalesOrder: *order,
		}

		if purchaseOrder, ok := purchaseOrderMap[order.PurchaseOrderNo]; ok {
			customerOrderResItem.PurchaseOrderInfo = purchaseOrder
		}

		res = append(res, &customerOrderResItem)
	}

	return &model.PagedResponse{
		Total: total,
		Data:  res,
	}, nil

}

// UpdateOrderStatus
func (s *CustomerService) UpdateOrderStatus(ctx *app.Context, params *model.ReqSalesOrderUpdateStatusParam) error {
	nowstr := time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Model(&model.SalesOrder{}).Where("order_no = ?", params.OrderNo).Updates(map[string]interface{}{
		"status":     params.Status,
		"updated_at": nowstr,
	}).Error
	if err != nil {
		ctx.Logger.Error("Failed to update order status", err)
		return errors.New("failed to update order status")
	}
	return nil
}

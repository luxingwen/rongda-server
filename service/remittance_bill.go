package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RemittanceBillService struct {
}

func NewRemittanceBillService() *RemittanceBillService {
	return &RemittanceBillService{}
}

func (s *RemittanceBillService) CreateRemittanceBill(ctx *app.Context, remittanceBill *model.RemittanceBill) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	remittanceBill.Uuid = uuid.New().String()
	remittanceBill.CreatedAt = now
	remittanceBill.UpdatedAt = now

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(remittanceBill).Error
		if err != nil {
			ctx.Logger.Error("Failed to create remittance bill", err)
			tx.Rollback()
			return errors.New("failed to create remittance bill")
		}
		return nil
	})

	if err != nil {
		ctx.Logger.Error("Failed to create remittance bill", err)
		return errors.New("failed to create remittance bill")
	}
	return nil
}

func (s *RemittanceBillService) GetRemittanceBillByUUID(ctx *app.Context, uuid string) (*model.RemittanceBillRes, error) {
	remittanceBill := &model.RemittanceBill{}
	err := ctx.DB.Where("uuid = ?", uuid).First(remittanceBill).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("remittance bill not found")
		}
		ctx.Logger.Error("Failed to get remittance bill by UUID", err)
		return nil, errors.New("failed to get remittance bill by UUID")
	}

	salesOrder, err := NewSalesOrderService().GetSalesOrder(ctx, remittanceBill.OrderNo)
	if err != nil {
		ctx.Logger.Error("Failed to get sales order", err)
		return nil, errors.New("failed to get sales order")
	}

	supplier, err := NewSupplierService().GetSupplierByUUID(ctx, remittanceBill.Supplier)
	if err != nil {
		ctx.Logger.Error("Failed to get supplier", err)
		return nil, errors.New("failed to get supplier")
	}
	res := &model.RemittanceBillRes{
		RemittanceBill: *remittanceBill,
		SalesOrderInfo: salesOrder,
		SupplierInfo:   supplier,
	}

	return res, nil
}

func (s *RemittanceBillService) UpdateRemittanceBill(ctx *app.Context, remittanceBill *model.RemittanceBill) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	remittanceBill.UpdatedAt = now
	err := ctx.DB.Where("uuid = ?", remittanceBill.Uuid).Updates(remittanceBill).Error
	if err != nil {
		ctx.Logger.Error("Failed to update remittance bill", err)
		return errors.New("failed to update remittance bill")
	}

	return nil
}

func (s *RemittanceBillService) DeleteRemittanceBill(ctx *app.Context, uuid string) error {
	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.RemittanceBill{}).Where("uuid = ?", uuid).Update("is_deleted", 1).Error
		if err != nil {
			ctx.Logger.Error("Failed to delete remittance bill", err)
			tx.Rollback()
			return errors.New("failed to delete remittance bill")
		}
		return nil
	})

	return err
}

// GetRemittanceBillList retrieves a list of remittance bills based on query parameters
func (s *RemittanceBillService) GetRemittanceBillList(ctx *app.Context, params *model.ReqRemittanceBillQueryParam) (*model.PagedResponse, error) {
	var (
		remittanceBills []*model.RemittanceBill
		total           int64
	)

	db := ctx.DB.Model(&model.RemittanceBill{})

	if params.OrderNo != "" {
		db = db.Where("order_no = ?", params.OrderNo)
	}
	if params.AgreementNo != "" {
		db = db.Where("agreement_no = ?", params.AgreementNo)
	}

	if params.TeamUuid != "" {
		db = db.Where("team_uuid = ?", params.TeamUuid)
	}

	if params.Status != "" {
		db = db.Where("status = ?", params.Status)
	}

	if params.Type != "" {
		db = db.Where("type = ?", params.Type)
	}

	db = db.Where("is_deleted = ?", 0)

	err := db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to get remittance bill count", err)
		return nil, errors.New("failed to get remittance bill count")
	}

	err = db.Offset(params.GetOffset()).Limit(params.PageSize).Find(&remittanceBills).Error
	if err != nil {
		ctx.Logger.Error("Failed to get remittance bill list", err)
		return nil, errors.New("failed to get remittance bill list")
	}

	// Additional processing if required

	salesOrderNos := make([]string, 0)
	suppliereUuids := make([]string, 0)

	for _, remittanceBill := range remittanceBills {
		salesOrderNos = append(salesOrderNos, remittanceBill.OrderNo)
		suppliereUuids = append(suppliereUuids, remittanceBill.Supplier)
	}

	salesOrderMap, err := NewSalesOrderService().GetSalesOrdersByUUIDs(ctx, salesOrderNos)
	if err != nil {
		ctx.Logger.Error("Failed to get sales orders by UUIDs", err)
		return nil, errors.New("failed to get sales orders by UUIDs")
	}

	supplierMap, err := NewSupplierService().GetSupplierListByUUIDs(ctx, suppliereUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get suppliers by UUIDs", err)
		return nil, errors.New("failed to get suppliers by UUIDs")
	}

	remittanceBillRes := make([]*model.RemittanceBillRes, 0)

	for _, remittanceBill := range remittanceBills {
		item := &model.RemittanceBillRes{
			RemittanceBill: *remittanceBill,
		}

		if salesOrder, ok := salesOrderMap[remittanceBill.OrderNo]; ok {
			item.SalesOrderInfo = salesOrder
		}

		if supplier, ok := supplierMap[remittanceBill.Supplier]; ok {
			item.SupplierInfo = supplier
		}

		remittanceBillRes = append(remittanceBillRes, item)
	}

	return &model.PagedResponse{
		Total: total,
		Data:  remittanceBillRes,
	}, nil
}

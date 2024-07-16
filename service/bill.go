package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BillService struct {
}

func NewBillService() *BillService {
	return &BillService{}
}

func (s *BillService) CreateBill(ctx *app.Context, bill *model.Bill) error {
	bill.Uuid = uuid.New().String()
	bill.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	bill.UpdatedAt = bill.CreatedAt

	if err := ctx.DB.Create(bill).Error; err != nil {
		ctx.Logger.Error("Failed to create bill", err)
		return errors.New("failed to create bill")
	}
	return nil
}

func (s *BillService) GetBill(ctx *app.Context, uuid string) (*model.Bill, error) {
	bill := &model.Bill{}
	err := ctx.DB.Where("uuid = ?", uuid).First(bill).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("bill not found")
		}
		ctx.Logger.Error("Failed to get bill by UUID", err)
		return nil, errors.New("failed to get bill by UUID")
	}
	return bill, nil
}

func (s *BillService) UpdateBill(ctx *app.Context, bill *model.Bill) error {
	bill.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	if err := ctx.DB.Save(bill).Error; err != nil {
		ctx.Logger.Error("Failed to update bill", err)
		return errors.New("failed to update bill")
	}
	return nil
}

func (s *BillService) DeleteBill(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.Bill{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete bill", err)
		return errors.New("failed to delete bill")
	}
	return nil
}

func (s *BillService) ListBills(ctx *app.Context, param *model.ReqBillQueryParam) (r *model.PagedResponse, err error) {
	var (
		billList []*model.Bill
		total    int64
	)

	db := ctx.DB.Model(&model.Bill{})

	if param.InvoiceCompany != "" {
		db = db.Where("invoice_company = ?", param.InvoiceCompany)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&billList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     billList,
	}
	return
}

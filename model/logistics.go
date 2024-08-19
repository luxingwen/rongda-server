package model

type Logistics struct {
	ID   uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`            // 主键ID
	Uuid string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"` // UUID
	// 出库单号
	OutboundOrderNo string `json:"outbound_order_no" gorm:"comment:'出库单号'"` // 出库单号
	OrderNo         string `json:"order_no" gorm:"comment:'订单号'"`           // 订单号
	Shipper         string `json:"shipper" gorm:"comment:'承运商'"`            // 承运商
	ShipperPhone    string `json:"shipper_phone" gorm:"comment:'承运商电话'"`    // 承运商电话
	ShipperNo       string `json:"shipper_no" gorm:"comment:'承运单号'"`        // 承运单号
	// 运输方式
	TransportMode string `json:"transport_mode" gorm:"comment:'运输方式'"` // 运输方式
	// 运输状态
	TransportStatus string `json:"transport_status" gorm:"comment:'运输状态'"` // 运输状态 1:未发货 2:已发货 3:已到货

	// 运输车辆信息
	VehicleNo string `json:"vehicle_no" gorm:"comment:'车牌号'"` // 车牌号
	// 司机
	Driver string `json:"driver" gorm:"comment:'司机'"` // 司机
	// 司机电话
	DriverPhone string `json:"driver_phone" gorm:"comment:'司机电话'"` // 司机电话

	Receiver        string `json:"receiver" gorm:"comment:'收货人'"`          // 收货人
	ReceiverPhone   string `json:"receiver_phone" gorm:"comment:'收货人电话'"`  // 收货人电话
	ReceiverAddress string `json:"receiver_address" gorm:"comment:'收货地址'"` // 收货地址
	// 发货日期
	ShipDate string `json:"ship_date" gorm:"comment:'发货日期'"` // 发货日期
	// 到货日期
	ArrivalDate string `json:"arrival_date" gorm:"comment:'到货日期'"` // 到货日期
	// 运费
	Freight float64 `json:"freight" gorm:"comment:'运费'"` // 运费
	// 运费支付方
	FreightPayer string `json:"freight_payer" gorm:"comment:'运费支付方'"` // 运费支付方 1:发货方 2:收货方
	// 运费支付状态
	FreightStatus string `json:"freight_status" gorm:"comment:'运费支付状态'"` // 运费支付状态 1:未支付 2:已支付
	// 备注
	Remark string `json:"remark" gorm:"comment:'备注'"` // 备注
	// 附件
	Attachment string `json:"attachment" gorm:"comment:'附件'"`                  // 附件
	CreatedAt  string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt  string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

type ReqLogisticsQueryParam struct {
	OrderNo string `json:"order_no" form:"order_no"` // 订单号
	Pagination
}

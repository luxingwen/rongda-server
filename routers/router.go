package routers

import (
	"sgin/controller"
	"sgin/middleware"
	"sgin/pkg/app"

	"io/ioutil"
	"net/http"
	"sgin/service"

	"github.com/gin-gonic/gin"
)

func InitRouter(ctx *app.App) {
	InitSwaggerRouter(ctx)
	InitUserRouter(ctx)
	InitMenuRouter(ctx)
	InitAppRouter(ctx)
	InitVerificationCodeRouter(ctx)
	InitRegisterRouter(ctx)
	InitLoginRouter(ctx)
	InitServerRouter(ctx)
	InitTeamRouter(ctx)
	InitCustomerRouter(ctx)
	InitAgentRouter(ctx)
	InitSupplierRouter(ctx)
	InitSettlementCurrencyRouter(ctx)
	InitSkuRouter(ctx)
	InitDepartmentRouter(ctx)
	InitProductRouter(ctx)
	InitStorehouseRouter(ctx)
	InitStorehouseInboundRouter(ctx)
	InitStorehouseProductRouter(ctx)
	InitStorehouseOutRouter(ctx)
	InitStorehouseInventoryCheckRouter(ctx)
	InitPurchaseOrderRouter(ctx)
	InitAgreementRouter(ctx)
	InitPurchaseArrivalRouter(ctx)
	InitPurchaseBillRouter(ctx)
	InitSalesOrderRouter(ctx)
	InitSalesOutOfStockRouter(ctx)
	InitSalesSettlementRouter(ctx)
	InitBillRouter(ctx)

}

func InitUserRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	v1.Use(middleware.LoginCheck())
	{
		userController := &controller.UserController{
			Service: &service.UserService{},
		}

		v1.POST("/user/create", userController.CreateUser)
		v1.POST("/user/info", userController.GetUserByUUID)
		v1.POST("/user/list", userController.GetUserList)
		v1.POST("/user/update", userController.UpdateUser)
		v1.POST("/user/delete", userController.DeleteUser)
		v1.GET("/user/myinfo", userController.GetMyInfo)
		v1.POST("/user/avatar", userController.UpdateAvatar)

	}

	{
		roleController := &controller.RoleController{
			RoleService: &service.RoleService{},
		}

		v1.POST("/role/create", roleController.CreateRole)
		v1.POST("/role/list", roleController.GetRoleList)
		v1.POST("/role/update", roleController.UpdateRole)
		v1.POST("/role/delete", roleController.DeleteRole)

	}
}

func InitMenuRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		menuController := &controller.MenuController{
			MenuService: &service.MenuService{},
		}
		v1.POST("/menu/create", menuController.CreateMenu)
		v1.POST("/menu/list", menuController.GetMenuList)
		v1.POST("/menu/update", menuController.UpdateMenu)
		v1.POST("/menu/delete", menuController.DeleteMenu)
	}
}

func InitAppRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		appController := &controller.AppController{
			AppService: &service.AppService{},
		}
		v1.POST("/app/list", appController.GetAppList)
		v1.POST("/app/create", appController.CreateApp)
		v1.POST("/app/update", appController.UpdateApp)
		v1.POST("/app/delete", appController.DeleteApp)

	}
}

func InitPurchaseOrderRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		purchaseOrderController := &controller.PurchaseOrderController{
			PurchaseOrderService: &service.PurchaseOrderService{},
		}
		v1.POST("/purchase_order/create", purchaseOrderController.CreatePurchaseOrder)
		v1.POST("/purchase_order/update", purchaseOrderController.UpdatePurchaseOrder)
		v1.POST("/purchase_order/delete", purchaseOrderController.DeletePurchaseOrder)
		v1.POST("/purchase_order/info", purchaseOrderController.GetPurchaseOrder)
		v1.POST("/purchase_order/list", purchaseOrderController.GetPurchaseOrderList)
		v1.POST("/purchase_order/item/list", purchaseOrderController.GetPurchaseOrderProducts)
	}
}

func InitVerificationCodeRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		verificationCodeController := &controller.VerificationCodeController{
			VerificationCodeService: &service.VerificationCodeService{},
		}
		v1.POST("/verification_code/create", verificationCodeController.CreateVerificationCode)
	}
}

// 注册的路由
func InitRegisterRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		registerController := &controller.RegisterController{
			UserService:             &service.UserService{},
			VerificationCodeService: &service.VerificationCodeService{},
		}
		v1.POST("/register", registerController.Register)
	}
}

func InitLoginRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		loginController := &controller.LoginController{
			UserService: &service.UserService{},
		}
		v1.POST("/login", loginController.Login)
	}
}

// 服务的路由
func InitServerRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		serverController := &controller.ServerController{
			ServerService: &service.ServerService{},
		}
		v1.POST("/server/create", serverController.CreateServer)
		v1.POST("/server/update", serverController.UpdateServer)
		v1.POST("/server/delete", serverController.DeleteServer)
		v1.POST("/server/info", serverController.GetServerInfo)
		v1.POST("/server/list", serverController.GetServerList)
	}
}

// 团队的路由
func InitTeamRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		teamController := &controller.TeamController{
			TeamService: &service.TeamService{},
		}
		v1.POST("/team/create", teamController.CreateTeam)
		v1.POST("/team/update", teamController.UpdateTeam)
		v1.POST("/team/delete", teamController.DeleteTeam)
		v1.POST("/team/info", teamController.GetTeamInfo)
		v1.POST("/team/list", teamController.GetTeamList)
	}
}

func InitCustomerRouter(ctx *app.App) {

	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		customerController := &controller.CustomerController{
			CustomerService: &service.CustomerService{},
		}
		v1.POST("/customer/create", customerController.CreateCustomer)
		v1.POST("/customer/update", customerController.UpdateCustomer)
		v1.POST("/customer/delete", customerController.DeleteCustomer)
		v1.POST("/customer/info", customerController.GetCustomerInfo)
		v1.POST("/customer/list", customerController.GetCustomerList)
		v1.POST("/customer/all", customerController.GetAllCustomerList)
	}
}
func InitAgentRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		agentController := &controller.AgentController{
			AgentService: &service.AgentService{},
		}
		v1.POST("/agent/create", agentController.CreateAgent)
		v1.POST("/agent/update", agentController.UpdateAgent)
		v1.POST("/agent/delete", agentController.DeleteAgent)
		v1.POST("/agent/info", agentController.GetAgentInfo)
		v1.POST("/agent/list", agentController.GetAgentList)
	}
}

func InitSupplierRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		suppliersController := &controller.SupplierController{
			SupplierService: &service.SupplierService{},
		}
		v1.POST("/supplier/create", suppliersController.CreateSupplier)
		v1.POST("/supplier/update", suppliersController.UpdateSupplier)
		v1.POST("/supplier/delete", suppliersController.DeleteSupplier)
		v1.POST("/supplier/info", suppliersController.GetSupplierInfo)
		v1.POST("/supplier/list", suppliersController.GetSupplierList)
		v1.POST("/supplier/all", suppliersController.GetAllSupplier)
	}
}

func InitSettlementCurrencyRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		settlementController := &controller.SettlementCurrencyController{
			SettlementCurrencyService: &service.SettlementCurrencyService{},
		}
		v1.POST("/settlement_currency/create", settlementController.CreateSettlementCurrency)
		v1.POST("/settlement_currency/update", settlementController.UpdateSettlementCurrency)
		v1.POST("/settlement_currency/delete", settlementController.DeleteSettlementCurrency)
		v1.POST("/settlement_currency/info", settlementController.GetSettlementCurrencyInfo)
		v1.POST("/settlement_currency/list", settlementController.GetSettlementCurrencyList)
		v1.GET("/settlement_currency/all", settlementController.GetAllSettlementCurrency)
	}
}

func InitSkuRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		skuController := &controller.SkuController{
			SkuService: &service.SkuService{},
		}
		v1.POST("/sku/create", skuController.CreateSku)
		v1.POST("/sku/update", skuController.UpdateSku)
		v1.POST("/sku/delete", skuController.DeleteSku)
		v1.POST("/sku/info", skuController.GetSkuInfo)
		v1.POST("/sku/list", skuController.GetSkuList)
	}
}

func InitDepartmentRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		departmentController := &controller.DepartmentController{
			DepartmentService: &service.DepartmentService{},
		}
		v1.POST("/department/create", departmentController.CreateDepartment)
		v1.POST("/department/update", departmentController.UpdateDepartment)
		v1.POST("/department/delete", departmentController.DeleteDepartment)
		v1.POST("/department/info", departmentController.GetDepartmentInfo)
		v1.POST("/department/list", departmentController.GetDepartmentList)
	}
}

func InitProductRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		productController := &controller.ProductController{
			ProductService: &service.ProductService{},
			SkuService:     &service.SkuService{},
		}
		v1.POST("/product/create", productController.CreateProduct)
		v1.POST("/product/update", productController.UpdateProduct)
		v1.POST("/product/delete", productController.DeleteProduct)
		v1.POST("/product/info", productController.GetProductInfo)
		v1.POST("/product/list", productController.GetProductList)
		v1.GET("/product/all", productController.GetAllProduct)
		v1.POST("/product/sku/list", productController.GetProductSkuList)
	}
}

func InitStorehouseRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	v1.Use(middleware.LoginCheck())
	{
		storehouseController := &controller.StorehouseController{
			StorehouseService: &service.StorehouseService{},
		}
		v1.POST("/storehouse/create", storehouseController.CreateStorehouse)
		v1.POST("/storehouse/update", storehouseController.UpdateStorehouse)
		v1.POST("/storehouse/delete", storehouseController.DeleteStorehouse)
		v1.POST("/storehouse/info", storehouseController.GetStorehouseInfo)
		v1.POST("/storehouse/list", storehouseController.GetStorehouseList)
		v1.POST("/storehouse/all", storehouseController.GetAllStorehouse)
	}
}

func InitStorehouseInboundRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	v1.Use(middleware.LoginCheck())
	{
		storehouseInboundController := &controller.StorehouseInboundController{
			InboundService: &service.StorehouseInboundService{},
		}
		v1.POST("/storehouse_inbound/create", storehouseInboundController.CreateInbound)
		v1.POST("/storehouse_inbound/update", storehouseInboundController.UpdateInbound)
		v1.POST("/storehouse_inbound/delete", storehouseInboundController.DeleteInbound)
		v1.POST("/storehouse_inbound/info", storehouseInboundController.GetInboundInfo)
		v1.POST("/storehouse_inbound/list", storehouseInboundController.GetInboundList)
		v1.POST("/storehouse_inbound/detail", storehouseInboundController.GetInboundDetail)
	}
}

func InitStorehouseProductRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	v1.Use(middleware.LoginCheck())
	{
		storehouseProductController := &controller.StorehouseProductController{
			ProductService: &service.StorehouseProductService{},
		}
		v1.POST("/storehouse_product/create", storehouseProductController.CreateProduct)
		v1.POST("/storehouse_product/update", storehouseProductController.UpdateProduct)
		v1.POST("/storehouse_product/delete", storehouseProductController.DeleteProduct)
		v1.POST("/storehouse_product/info", storehouseProductController.GetProductInfo)
		v1.POST("/storehouse_product/list", storehouseProductController.GetProductList)
		v1.POST("/storehouse_product/op_log", storehouseProductController.GetProductOpLog)
	}
}

func InitStorehouseOutRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	v1.Use(middleware.LoginCheck())
	{
		storehouseOutController := &controller.StorehouseOutboundController{
			OutboundService: &service.StorehouseOutboundService{},
		}
		v1.POST("/storehouse_outbound/create", storehouseOutController.CreateOutbound)
		v1.POST("/storehouse_outbound/update", storehouseOutController.UpdateOutbound)
		v1.POST("/storehouse_outbound/delete", storehouseOutController.DeleteOutbound)
		v1.POST("/storehouse_outbound/info", storehouseOutController.GetOutboundInfo)
		v1.POST("/storehouse_outbound/list", storehouseOutController.GetOutboundList)
		v1.POST("/storehouse_outbound/detail", storehouseOutController.GetOutboundDetail)
	}
}

func InitStorehouseInventoryCheckRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	v1.Use(middleware.LoginCheck())
	{
		storehouseInventoryCheckController := &controller.StorehouseInventoryCheckController{
			InventoryCheckService: &service.StorehouseInventoryCheckService{},
		}
		v1.POST("/storehouse_inventory_check/create", storehouseInventoryCheckController.CreateInventoryCheck)
		v1.POST("/storehouse_inventory_check/update", storehouseInventoryCheckController.UpdateInventoryCheck)
		v1.POST("/storehouse_inventory_check/delete", storehouseInventoryCheckController.DeleteInventoryCheck)
		v1.POST("/storehouse_inventory_check/info", storehouseInventoryCheckController.GetInventoryCheck)
		v1.POST("/storehouse_inventory_check/list", storehouseInventoryCheckController.GetInventoryCheckList)
		v1.POST("/storehouse_inventory_check/detail", storehouseInventoryCheckController.GetInventoryCheckDetail)
	}

}

func InitAgreementRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		agreementController := &controller.AgreementController{
			AgreementService: &service.AgreementService{},
		}
		v1.POST("/agreement/create", agreementController.CreateAgreement)
		v1.POST("/agreement/update", agreementController.UpdateAgreement)
		v1.POST("/agreement/delete", agreementController.DeleteAgreement)
		v1.POST("/agreement/info", agreementController.GetAgreement)
		v1.POST("/agreement/list", agreementController.ListAgreements)
	}
}

func InitPurchaseArrivalRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		purchaseArrivalController := &controller.PurchaseArrivalController{
			PurchaseArrivalService: &service.PurchaseArrivalService{},
		}
		v1.POST("/purchase_arrival/create", purchaseArrivalController.CreatePurchaseArrival)
		v1.POST("/purchase_arrival/update", purchaseArrivalController.UpdatePurchaseArrival)
		v1.POST("/purchase_arrival/delete", purchaseArrivalController.DeletePurchaseArrival)
		v1.POST("/purchase_arrival/info", purchaseArrivalController.GetPurchaseArrival)
		v1.POST("/purchase_arrival/list", purchaseArrivalController.GetPurchaseArrivalList)
	}
}

func InitPurchaseBillRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		purchaseBillController := &controller.PurchaseBillController{
			PurchaseBillService: &service.PurchaseBillService{},
		}
		v1.POST("/purchase_bill/create", purchaseBillController.CreatePurchaseBill)
		v1.POST("/purchase_bill/update", purchaseBillController.UpdatePurchaseBill)
		v1.POST("/purchase_bill/delete", purchaseBillController.DeletePurchaseBill)
		v1.POST("/purchase_bill/info", purchaseBillController.GetPurchaseBill)
		v1.POST("/purchase_bill/list", purchaseBillController.GetPurchaseBillList)

	}
}

func InitSalesOrderRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		salesOrderController := &controller.SalesOrderController{
			SalesOrderService: &service.SalesOrderService{},
		}
		v1.POST("/sales_order/create", salesOrderController.CreateSalesOrder)
		v1.POST("/sales_order/update", salesOrderController.UpdateSalesOrder)
		v1.POST("/sales_order/delete", salesOrderController.DeleteSalesOrder)
		v1.POST("/sales_order/info", salesOrderController.GetSalesOrder)
		v1.POST("/sales_order/list", salesOrderController.GetSalesOrderList)
		v1.POST("/sales_order/all", salesOrderController.GetAllSalesOrder)
		v1.POST("/sales_order/product_item/list", salesOrderController.GetSalesOrderProducts)
	}
}

func InitSalesOutOfStockRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		salesOutOfStockController := &controller.SalesOutOfStockController{
			SalesOutOfStockService: &service.SalesOutOfStockService{},
		}
		v1.POST("/sales_out_of_stock/create", salesOutOfStockController.CreateSalesOutOfStock)
		v1.POST("/sales_out_of_stock/update", salesOutOfStockController.UpdateSalesOutOfStock)
		v1.POST("/sales_out_of_stock/delete", salesOutOfStockController.DeleteSalesOutOfStock)
		v1.POST("/sales_out_of_stock/info", salesOutOfStockController.GetSalesOutOfStock)
		v1.POST("/sales_out_of_stock/list", salesOutOfStockController.GetSalesOutOfStockList)
	}
}

func InitSalesSettlementRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		salesSettlementController := &controller.SalesSettlementController{
			SalesSettlementService: &service.SalesSettlementService{},
		}
		v1.POST("/sales_settlement/create", salesSettlementController.CreateSalesSettlement)
		v1.POST("/sales_settlement/update", salesSettlementController.UpdateSalesSettlement)
		v1.POST("/sales_settlement/delete", salesSettlementController.DeleteSalesSettlement)
		v1.POST("/sales_settlement/info", salesSettlementController.GetSalesSettlement)
		v1.POST("/sales_settlement/list", salesSettlementController.GetSalesSettlementList)
	}
}

func InitBillRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		billController := &controller.BillController{
			BillService: &service.BillService{},
		}
		v1.POST("/bill/create", billController.CreateBill)
		v1.POST("/bill/update", billController.UpdateBill)
		v1.POST("/bill/delete", billController.DeleteBill)
		v1.POST("/bill/info", billController.GetBill)
		v1.POST("/bill/list", billController.GetBillList)
	}
}

func InitSwaggerRouter(ctx *app.App) {
	ctx.GET("/swagger/doc.json", func(c *app.Context) {
		jsonFile, err := ioutil.ReadFile("./docs/swagger.json") // Replace with your actual json file path
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Data(http.StatusOK, "application/json", jsonFile)
	})

	ctx.GET("/swagger/redoc.standalone.js", func(c *app.Context) {
		b, err := ioutil.ReadFile("./swagger/redoc.standalone.js") // Replace with your actual json file path
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", b)
	})

	ctx.GET("/swagger/index.html", func(c *app.Context) {
		b, err := ioutil.ReadFile("./swagger/swagger.html") // Replace with your actual json file path
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", b)
	})
}

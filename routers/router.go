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
		}
		v1.POST("/product/create", productController.CreateProduct)
		v1.POST("/product/update", productController.UpdateProduct)
		v1.POST("/product/delete", productController.DeleteProduct)
		v1.POST("/product/info", productController.GetProductInfo)
		v1.POST("/product/list", productController.GetProductList)
	}
}

func InitStorehouseRouter(ctx *app.App) {
	v1 := ctx.Group(ctx.Config.ApiPrefix + "/v1")
	{
		storehouseController := &controller.StorehouseController{
			StorehouseService: &service.StorehouseService{},
		}
		v1.POST("/storehouse/create", storehouseController.CreateStorehouse)
		v1.POST("/storehouse/update", storehouseController.UpdateStorehouse)
		v1.POST("/storehouse/delete", storehouseController.DeleteStorehouse)
		v1.POST("/storehouse/info", storehouseController.GetStorehouseInfo)
		v1.POST("/storehouse/list", storehouseController.GetStorehouseList)
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

INSERT INTO `apis` (`uuid`, `module`, `name`, `permission_level`, `status`, `path`, `method`, `created_at`, `updated_at`) VALUES 
(UUID(), '用户信息', '创建用户', 2, 1, '/api/v1/user/create', 'POST', NOW(), NOW()),
(UUID(), '用户信息', '获取用户信息', 2, 1, '/api/v1/user/info', 'POST', NOW(), NOW()),
(UUID(), '用户信息', '获取用户列表', 2, 1, '/api/v1/user/list', 'POST', NOW(), NOW()),
(UUID(), '用户信息', '更新用户信息', 2, 1, '/api/v1/user/update', 'POST', NOW(), NOW()),
(UUID(), '用户信息', '删除用户', 2, 1, '/api/v1/user/delete', 'POST', NOW(), NOW()),
(UUID(), '用户信息', '获取我的信息', 2, 1, '/api/v1/user/myinfo', 'GET', NOW(), NOW()),
(UUID(), '用户信息', '上传头像', 2, 1, '/api/v1/user/avatar', 'POST', NOW(), NOW());



INSERT INTO `apis` (`uuid`, `module`, `name`, `permission_level`, `status`, `path`, `method`, `created_at`, `updated_at`) VALUES 
(UUID(), '菜单', '创建菜单', 2, 1, '/api/v1/menu/create', 'POST', NOW(), NOW()),
(UUID(), '菜单', '获取菜单列表', 2, 1, '/api/v1/menu/list', 'POST', NOW(), NOW()),
(UUID(), '菜单', '更新菜单', 2, 1, '/api/v1/menu/update', 'POST', NOW(), NOW()),
(UUID(), '菜单', '删除菜单', 2, 1, '/api/v1/menu/delete', 'POST', NOW(), NOW());


INSERT INTO `apis` (`uuid`, `module`, `name`, `permission_level`, `status`, `path`, `method`, `created_at`, `updated_at`) VALUES 
(UUID(), '通用', '用户登录', 1, 1, '/api/v1/login', 'POST', NOW(), NOW());


INSERT INTO `apis` (`uuid`, `module`, `name`, `permission_level`, `status`, `path`, `method`, `created_at`, `updated_at`) VALUES 
(UUID(), '客户管理', '创建客户', 2, 1, '/api/v1/customer/create', 'POST', NOW(), NOW()),
(UUID(), '客户管理', '更新客户', 2, 1, '/api/v1/customer/update', 'POST', NOW(), NOW()),
(UUID(), '客户管理', '删除客户', 2, 1, '/api/v1/customer/delete', 'POST', NOW(), NOW()),
(UUID(), '客户管理', '获取客户信息', 2, 1, '/api/v1/customer/info', 'POST', NOW(), NOW()),
(UUID(), '客户管理', '获取客户列表', 2, 1, '/api/v1/customer/list', 'POST', NOW(), NOW()),
(UUID(), '客户管理', '获取所有可用客户-SELECT选择使用', 2, 1, '/api/v1/customer/all', 'POST', NOW(), NOW());



INSERT INTO `apis` (`uuid`, `module`, `name`, `permission_level`, `status`, `path`, `method`, `created_at`, `updated_at`) VALUES 
(UUID(), '代理', '创建代理', 2, 1, '/api/v1/agent/create', 'POST', NOW(), NOW()),
(UUID(), '代理', '更新代理', 2, 1, '/api/v1/agent/update', 'POST', NOW(), NOW()),
(UUID(), '代理', '删除代理', 2, 1, '/api/v1/agent/delete', 'POST', NOW(), NOW()),
(UUID(), '代理', '获取代理信息', 2, 1, '/api/v1/agent/info', 'POST', NOW(), NOW()),
(UUID(), '代理', '获取代理列表', 2, 1, '/api/v1/agent/list', 'POST', NOW(), NOW());

INSERT INTO `apis` (`uuid`, `module`, `name`, `permission_level`, `status`, `path`, `method`, `created_at`, `updated_at`) VALUES 
(UUID(), '供应商管理', '创建供应商', 2, 1, '/api/v1/supplier/create', 'POST', NOW(), NOW()),
(UUID(), '供应商管理', '更新供应商', 2, 1, '/api/v1/supplier/update', 'POST', NOW(), NOW()),
(UUID(), '供应商管理', '删除供应商', 2, 1, '/api/v1/supplier/delete', 'POST', NOW(), NOW()),
(UUID(), '供应商管理', '获取供应商信息', 2, 1, '/api/v1/supplier/info', 'POST', NOW(), NOW()),
(UUID(), '供应商管理', '获取供应商列表', 2, 1, '/api/v1/supplier/list', 'POST', NOW(), NOW()),
(UUID(), '供应商管理', '获取所有可用供应商-SELECT选择使用', 2, 1, '/api/v1/supplier/all', 'POST', NOW(), NOW());



INSERT INTO `apis` (`uuid`, `module`, `name`, `permission_level`, `status`, `path`, `method`, `created_at`, `updated_at`) VALUES 
(UUID(), 'SKU管理', '创建SKU', 2, 1, '/api/v1/sku/create', 'POST', NOW(), NOW()),
(UUID(), 'SKU管理', '更新SKU', 2, 1, '/api/v1/sku/update', 'POST', NOW(), NOW()),
(UUID(), 'SKU管理', '删除SKU', 2, 1, '/api/v1/sku/delete', 'POST', NOW(), NOW()),
(UUID(), 'SKU管理', '获取SKU信息', 2, 1, '/api/v1/sku/info', 'POST', NOW(), NOW()),
(UUID(), 'SKU管理', '获取SKU列表', 2, 1, '/api/v1/sku/list', 'POST', NOW(), NOW());


INSERT INTO `apis` (`uuid`, `module`, `name`, `permission_level`, `status`, `path`, `method`, `created_at`, `updated_at`) VALUES 
(UUID(), '产品管理', '创建产品', 2, 1, '/api/v1/product/create', 'POST', NOW(), NOW()),
(UUID(), '产品管理', '更新产品', 2, 1, '/api/v1/product/update', 'POST', NOW(), NOW()),
(UUID(), '产品管理', '删除产品', 2, 1, '/api/v1/product/delete', 'POST', NOW(), NOW()),
(UUID(), '产品管理', '获取产品信息', 2, 1, '/api/v1/product/info', 'POST', NOW(), NOW()),
(UUID(), '产品管理', '获取产品列表', 2, 1, '/api/v1/product/list', 'POST', NOW(), NOW()),
(UUID(), '产品管理', '获取所有产品', 2, 1, '/api/v1/product/all', 'GET', NOW(), NOW()),
(UUID(), '产品管理', '获取产品SKU列表', 2, 1, '/api/v1/product/sku/list', 'POST', NOW(), NOW());


INSERT INTO `apis` (`uuid`, `module`, `name`, `permission_level`, `status`, `path`, `method`, `created_at`, `updated_at`) VALUES 
(UUID(), '仓库信息管理', '创建仓库', 2, 1, '/api/v1/storehouse/create', 'POST', NOW(), NOW()),
(UUID(), '仓库信息管理', '更新仓库', 2, 1, '/api/v1/storehouse/update', 'POST', NOW(), NOW()),
(UUID(), '仓库信息管理', '删除仓库', 2, 1, '/api/v1/storehouse/delete', 'POST', NOW(), NOW()),
(UUID(), '仓库信息管理', '获取仓库信息', 2, 1, '/api/v1/storehouse/info', 'POST', NOW(), NOW()),
(UUID(), '仓库信息管理', '获取仓库列表', 2, 1, '/api/v1/storehouse/list', 'POST', NOW(), NOW()),
(UUID(), '仓库信息管理', '获取所有仓库', 2, 1, '/api/v1/storehouse/all', 'POST', NOW(), NOW());


INSERT INTO `apis` (`uuid`, `module`, `name`, `permission_level`, `status`, `path`, `method`, `created_at`, `updated_at`) VALUES 
(UUID(), '仓库入库管理', '创建入库', 2, 1, '/api/v1/storehouse_inbound/create', 'POST', NOW(), NOW()),
(UUID(), '仓库入库管理', '更新入库', 2, 1, '/api/v1/storehouse_inbound/update', 'POST', NOW(), NOW()),
(UUID(), '仓库入库管理', '删除入库', 2, 1, '/api/v1/storehouse_inbound/delete', 'POST', NOW(), NOW()),
(UUID(), '仓库入库管理', '获取入库信息', 2, 1, '/api/v1/storehouse_inbound/info', 'POST', NOW(), NOW()),
(UUID(), '仓库入库管理', '获取入库列表', 2, 1, '/api/v1/storehouse_inbound/list', 'POST', NOW(), NOW()),
(UUID(), '仓库入库管理', '获取入库详情', 2, 1, '/api/v1/storehouse_inbound/detail', 'POST', NOW(), NOW());


INSERT INTO `apis` (`uuid`, `module`, `name`, `permission_level`, `status`, `path`, `method`, `created_at`, `updated_at`) VALUES 
(UUID(), '仓库库存管理', '创建库存产品', 2, 1, '/api/v1/storehouse_product/create', 'POST', NOW(), NOW()),
(UUID(), '仓库库存管理', '更新库存产品', 2, 1, '/api/v1/storehouse_product/update', 'POST', NOW(), NOW()),
(UUID(), '仓库库存管理', '删除库存产品', 2, 1, '/api/v1/storehouse_product/delete', 'POST', NOW(), NOW()),
(UUID(), '仓库库存管理', '获取库存产品信息', 2, 1, '/api/v1/storehouse_product/info', 'POST', NOW(), NOW()),
(UUID(), '仓库库存管理', '获取库存产品列表', 2, 1, '/api/v1/storehouse_product/list', 'POST', NOW(), NOW()),
(UUID(), '仓库库存管理', '获取库存操作日志', 2, 1, '/api/v1/storehouse_product/op_log', 'POST', NOW(), NOW());



INSERT INTO `apis` (`uuid`, `module`, `name`, `permission_level`, `status`, `path`, `method`, `created_at`, `updated_at`) VALUES 
(UUID(), '仓库出库管理', '创建出库', 2, 1, '/api/v1/storehouse_outbound/create', 'POST', NOW(), NOW()),
(UUID(), '仓库出库管理', '更新出库', 2, 1, '/api/v1/storehouse_outbound/update', 'POST', NOW(), NOW()),
(UUID(), '仓库出库管理', '删除出库', 2, 1, '/api/v1/storehouse_outbound/delete', 'POST', NOW(), NOW()),
(UUID(), '仓库出库管理', '获取出库信息', 2, 1, '/api/v1/storehouse_outbound/info', 'POST', NOW(), NOW()),
(UUID(), '仓库出库管理', '获取出库列表', 2, 1, '/api/v1/storehouse_outbound/list', 'POST', NOW(), NOW()),
(UUID(), '仓库出库管理', '获取出库详情', 2, 1, '/api/v1/storehouse_outbound/detail', 'POST', NOW(), NOW());



INSERT INTO `apis` (`uuid`, `module`, `name`, `permission_level`, `status`, `path`, `method`, `created_at`, `updated_at`) VALUES 
(UUID(), '仓库库存盘点管理', '创建库存盘点', 2, 1, '/api/v1/storehouse_inventory_check/create', 'POST', NOW(), NOW()),
(UUID(), '仓库库存盘点管理', '更新库存盘点', 2, 1, '/api/v1/storehouse_inventory_check/update', 'POST', NOW(), NOW()),
(UUID(), '仓库库存盘点管理', '删除库存盘点', 2, 1, '/api/v1/storehouse_inventory_check/delete', 'POST', NOW(), NOW()),
(UUID(), '仓库库存盘点管理', '获取库存盘点信息', 2, 1, '/api/v1/storehouse_inventory_check/info', 'POST', NOW(), NOW()),
(UUID(), '仓库库存盘点管理', '获取库存盘点列表', 2, 1, '/api/v1/storehouse_inventory_check/list', 'POST', NOW(), NOW()),
(UUID(), '仓库库存盘点管理', '获取库存盘点详情', 2, 1, '/api/v1/storehouse_inventory_check/detail', 'POST', NOW(), NOW());

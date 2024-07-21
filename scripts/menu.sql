-- 插入一级菜单
INSERT INTO `menus` (`uuid`, `name`, `link`, `parent_uuid`, `created_at`, `updated_at`, `icon`, `order`, `is_show`, `type`)
VALUES 
(UUID(), '首页', '/home', NULL, NOW(), NOW(), 'home', 1, 1, 1),
(UUID(), '登录', '/login', NULL, NOW(), NOW(), 'login', 2, 0, 2),
(UUID(), '个人中心', '/user/profile', NULL, NOW(), NOW(), 'user', 3, 0, 2),
(UUID(), '资料管理', '/resource', NULL, NOW(), NOW(), 'solution', 4, 1, 1),
(UUID(), '员工管理', '/staff', NULL, NOW(), NOW(), 'user', 5, 1, 1),
(UUID(), '仓库管理', '/storehouse', NULL, NOW(), NOW(), 'container', 6, 1, 1),
(UUID(), '采购管理', '/purchase', NULL, NOW(), NOW(), 'shopping', 7, 1, 1),
(UUID(), '销售管理', '/sales', NULL, NOW(), NOW(), 'shop', 8, 1, 1),
(UUID(), '系统管理', '/system', NULL, NOW(), NOW(), 'setting', 9, 1, 1);

-- 获取各个一级菜单的 UUID
SET @resource_uuid = (SELECT `uuid` FROM `menus` WHERE `name` = '资料管理');
SET @staff_uuid = (SELECT `uuid` FROM `menus` WHERE `name` = '员工管理');
SET @storehouse_uuid = (SELECT `uuid` FROM `menus` WHERE `name` = '仓库管理');
SET @purchase_uuid = (SELECT `uuid` FROM `menus` WHERE `name` = '采购管理');
SET @sales_uuid = (SELECT `uuid` FROM `menus` WHERE `name` = '销售管理');
SET @system_uuid = (SELECT `uuid` FROM `menus` WHERE `name` = '系统管理');

-- 插入资料管理子菜单
INSERT INTO `menus` (`uuid`, `name`, `link`, `parent_uuid`, `created_at`, `updated_at`, `icon`, `order`, `is_show`, `type`)
VALUES 
(UUID(), '客户管理', '/resource/customer', @resource_uuid, NOW(), NOW(), NULL, 1, 1, 2),
(UUID(), '客户详细', '/resource/customer/detail/:uuid', @resource_uuid, NOW(), NOW(), NULL, 2, 0, 3),
(UUID(), '代理机构', '/resource/agent', @resource_uuid, NOW(), NOW(), NULL, 3, 1, 2),
(UUID(), '代理机构详情', '/resource/agent/detail/:uuid', @resource_uuid, NOW(), NOW(), NULL, 4, 0, 3),
(UUID(), '供应商管理', '/resource/supplier', @resource_uuid, NOW(), NOW(), NULL, 5, 1, 2),
(UUID(), '供应商详情', '/resource/supplier/detail/:uuid', @resource_uuid, NOW(), NOW(), NULL, 6, 0, 3),
(UUID(), 'SKU管理', '/resource/sku', @resource_uuid, NOW(), NOW(), NULL, 7, 1, 2),
(UUID(), '商品管理', '/resource/product-manage', @resource_uuid, NOW(), NOW(), NULL, 8, 1, 2),
(UUID(), '商品类别', '/resource/product-category', @resource_uuid, NOW(), NOW(), NULL, 9, 1, 2);

-- 插入员工管理子菜单
INSERT INTO `menus` (`uuid`, `name`, `link`, `parent_uuid`, `created_at`, `updated_at`, `icon`, `order`, `is_show`, `type`)
VALUES 
(UUID(), '部门管理', '/staff/department', @staff_uuid, NOW(), NOW(), NULL, 1, 1, 2),
(UUID(), '用户管理', '/staff/user', @staff_uuid, NOW(), NOW(), NULL, 2, 1, 2),
(UUID(), '权限管理', '/staff/permission', @staff_uuid, NOW(), NOW(), NULL, 3, 1, 2);

-- 插入仓库管理子菜单
INSERT INTO `menus` (`uuid`, `name`, `link`, `parent_uuid`, `created_at`, `updated_at`, `icon`, `order`, `is_show`, `type`)
VALUES 
(UUID(), '仓库信息管理', '/storehouse/manage', @storehouse_uuid, NOW(), NOW(), NULL, 1, 1, 2),
(UUID(), '仓库信息详情', '/storehouse/detail/:uuid', @storehouse_uuid, NOW(), NOW(), NULL, 2, 0, 3),
(UUID(), '库存管理', '/storehouse/inventory', @storehouse_uuid, NOW(), NOW(), NULL, 3, 1, 2);

-- 获取刚刚插入的 `库存管理` 菜单的 UUID
SET @inventory_uuid = (SELECT `uuid` FROM `menus` WHERE `name` = '库存管理' AND `parent_uuid` = @storehouse_uuid);

INSERT INTO `menus` (`uuid`, `name`, `link`, `parent_uuid`, `created_at`, `updated_at`, `icon`, `order`, `is_show`, `type`)
VALUES 
(UUID(), '库存查询', '/storehouse/inventory/query', @inventory_uuid, NOW(), NOW(), NULL, 1, 1, 2),
(UUID(), '库存详情', '/storehouse/inventory/storehouse-product-detail/:uuid', @inventory_uuid, NOW(), NOW(), NULL, 2, 0, 3),
(UUID(), '入库', '/storehouse/inventory/in', @inventory_uuid, NOW(), NOW(), NULL, 3, 1, 2),
(UUID(), '入库详情', '/storehouse/inventory/inbound-detail/:uuid', @inventory_uuid, NOW(), NOW(), NULL, 4, 0, 3),
(UUID(), '添加入库', '/storehouse/inventory/inbound-add', @inventory_uuid, NOW(), NOW(), NULL, 5, 0, 3),
(UUID(), '出库', '/storehouse/inventory/out', @inventory_uuid, NOW(), NOW(), NULL, 6, 1, 2),
(UUID(), '添加出库', '/storehouse/inventory/outbound-add', @inventory_uuid, NOW(), NOW(), NULL, 7, 0, 3),
(UUID(), '出库详情', '/storehouse/inventory/outbound-detail/:uuid', @inventory_uuid, NOW(), NOW(), NULL, 8, 0, 3),
(UUID(), '库存盘点', '/storehouse/inventory/check', @inventory_uuid, NOW(), NOW(), NULL, 9, 1, 2),
(UUID(), '添加盘点', '/storehouse/inventory/check-add', @inventory_uuid, NOW(), NOW(), NULL, 10, 0, 3),
(UUID(), '盘点明细', '/storehouse/inventory/check-detail/:uuid', @inventory_uuid, NOW(), NOW(), NULL, 11, 0, 3);

-- 插入采购管理子菜单
INSERT INTO `menus` (`uuid`, `name`, `link`, `parent_uuid`, `created_at`, `updated_at`, `icon`, `order`, `is_show`, `type`)
VALUES 
(UUID(), '合同管理', '/purchase/agreement', @purchase_uuid, NOW(), NOW(), NULL, 1, 1, 2),
(UUID(), '添加合同', '/purchase/agreement/add', @purchase_uuid, NOW(), NOW(), NULL, 2, 0, 3),
(UUID(), '合同详情', '/purchase/agreement/detail/:uuid', @purchase_uuid, NOW(), NOW(), NULL, 3, 0, 3),
(UUID(), '编辑合同', '/purchase/agreement/edit/:uuid', @purchase_uuid, NOW(), NOW(), NULL, 4, 0, 3),
(UUID(), '采购订单', '/purchase/order', @purchase_uuid, NOW(), NOW(), NULL, 5, 1, 2),
(UUID(), '创建采购订单（期货）', '/purchase/order/add-futures', @purchase_uuid, NOW(), NOW(), NULL, 6, 0, 3),
(UUID(), '创建采购订单（现货）', '/purchase/order/add-spot', @purchase_uuid, NOW(), NOW(), NULL, 7, 0, 3),
(UUID(), '采购订单详情', '/purchase/order/detail/:uuid', @purchase_uuid, NOW(), NOW(), NULL, 8, 0, 3),
(UUID(), '采购订单(现货)详情', '/purchase/order/detail-spot/:uuid', @purchase_uuid, NOW(), NOW(), NULL, 9, 0, 3),
(UUID(), '到库登记', '/purchase/arrival', @purchase_uuid, NOW(), NOW(), NULL, 10, 1, 2),
(UUID(), '添加到库登记', '/purchase/arrival/add', @purchase_uuid, NOW(), NOW(), NULL, 11, 0, 3),
(UUID(), '编辑到库登记', '/purchase/arrival/edit/:uuid', @purchase_uuid, NOW(), NOW(), NULL, 12, 0, 3),
(UUID(), '到库详情', '/purchase/arrival/detail/:uuid', @purchase_uuid, NOW(), NOW(), NULL, 13, 0, 3),
(UUID(), '结算', '/purchase/settlement', @purchase_uuid, NOW(), NOW(), NULL, 14, 1, 2),
(UUID(), '添加结算', '/purchase/settlement/add', @purchase_uuid, NOW(), NOW(), NULL, 15, 0, 3),
(UUID(), '编辑结算', '/purchase/settlement/edit/:uuid', @purchase_uuid, NOW(), NOW(), NULL, 16, 0, 3),
(UUID(), '结算详情', '/purchase/settlement/detail/:uuid', @purchase_uuid, NOW(), NOW(), NULL, 17, 0, 3);

-- 插入销售管理子菜单
INSERT INTO `menus` (`uuid`, `name`, `link`, `parent_uuid`, `created_at`, `updated_at`, `icon`, `order`, `is_show`, `type`)
VALUES 
(UUID(), '合同管理', '/sales/agreement', @sales_uuid, NOW(), NOW(), NULL, 1, 1, 2),
(UUID(), '销售订单', '/sales/order', @sales_uuid, NOW(), NOW(), NULL, 2, 1, 2),
(UUID(), '订单详情', '/sales/order/detail/:uuid', @sales_uuid, NOW(), NOW(), NULL, 3, 0, 3),
(UUID(), '添加订单', '/sales/order/add', @sales_uuid, NOW(), NOW(), NULL, 4, 0, 3),
(UUID(), '编辑订单', '/sales/order/edit/:uuid', @sales_uuid, NOW(), NOW(), NULL, 5, 0, 3),
(UUID(), '出库登记', '/sales/outbound', @sales_uuid, NOW(), NOW(), NULL, 6, 1, 2),
(UUID(), '添加出库登记', '/sales/outbound/add', @sales_uuid, NOW(), NOW(), NULL, 7, 0, 3),
(UUID(), '编辑出库登记', '/sales/outbound/edit/:uuid', @sales_uuid, NOW(), NOW(), NULL, 8, 0, 3),
(UUID(), '出库登记详情', '/sales/outbound/detail/:uuid', @sales_uuid, NOW(), NOW(), NULL, 9, 0, 3),
(UUID(), '结算', '/sales/settlement', @sales_uuid, NOW(), NOW(), NULL, 10, 1, 2),
(UUID(), '发票', '/sales/bill', @sales_uuid, NOW(), NOW(), NULL, 11, 1, 2),
(UUID(), '添加发票', '/sales/bill/add', @sales_uuid, NOW(), NOW(), NULL, 12, 0, 3),
(UUID(), '编辑发票', '/sales/bill/edit/:uuid', @sales_uuid, NOW(), NOW(), NULL, 13, 0, 3),
(UUID(), '详情', '/sales/bill/detail/:uuid', @sales_uuid, NOW(), NOW(), NULL, 14, 0, 3);

-- 插入系统管理子菜单
INSERT INTO `menus` (`uuid`, `name`, `link`, `parent_uuid`, `created_at`, `updated_at`, `icon`, `order`, `is_show`, `type`)
VALUES 
(UUID(), '菜单管理', '/system/menu', @system_uuid, NOW(), NOW(), NULL, 1, 1, 2),
(UUID(), '结算币种', '/system/currency', @system_uuid, NOW(), NOW(), NULL, 2, 1, 2),
(UUID(), '银行信息', '/system/bankinfo', @system_uuid, NOW(), NOW(), NULL, 3, 1, 2),
(UUID(), '登录日志', '/system/loginlog', @system_uuid, NOW(), NOW(), NULL, 4, 1, 2),
(UUID(), 'API管理', '/system/api', @system_uuid, NOW(), NOW(), NULL, 5, 1, 2);

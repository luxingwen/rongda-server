
INSERT INTO `permissions` (`uuid`, `name`, `bit`, `parent_uuid`, `created_at`, `updated_at`)
VALUES 
(UUID(), '资料管理', 4, "", NOW(), NOW()),
(UUID(), '员工管理', 4, "", NOW(), NOW()),
(UUID(), '仓库管理', 4, "", NOW(), NOW()),
(UUID(), '采购管理', 4, "", NOW(), NOW()),
(UUID(), '销售管理', 4, "", NOW(), NOW()),
(UUID(), '系统管理', 4, "", NOW(), NOW());



-- 获取各个一级菜单的 UUID
SET @resource_uuid = (SELECT `uuid` FROM `permissions` WHERE `name` = '资料管理');
SET @staff_uuid = (SELECT `uuid` FROM `permissions` WHERE `name` = '员工管理');
SET @storehouse_uuid = (SELECT `uuid` FROM `permissions` WHERE `name` = '仓库管理');
SET @purchase_uuid = (SELECT `uuid` FROM `permissions` WHERE `name` = '采购管理');
SET @sales_uuid = (SELECT `uuid` FROM `permissions` WHERE `name` = '销售管理');
SET @system_uuid = (SELECT `uuid` FROM `permissions` WHERE `name` = '系统管理');



-- 插入资料管理子菜单
INSERT INTO `permissions` (`uuid`, `name`, `bit`, `parent_uuid`, `created_at`, `updated_at`)
VALUES 
(UUID(), '客户管理', 4, @resource_uuid, NOW(), NOW()),
(UUID(), '代理机构', 4, @resource_uuid, NOW(), NOW()),
(UUID(), '供应商管理', 4, @resource_uuid, NOW(), NOW()),
(UUID(), 'SKU管理', 4, @resource_uuid, NOW(), NOW()),
(UUID(), '商品管理', 4, @resource_uuid, NOW(), NOW()),
(UUID(), '商品类别', 4, @resource_uuid, NOW(), NOW());



SET @customer_uuid = (SELECT `uuid` FROM `permissions` WHERE `name` = '客户管理' AND `parent_uuid` = @resource_uuid);

INSERT INTO `permissions` (`uuid`, `name`, `bit`, `parent_uuid`, `created_at`, `updated_at`)
VALUES
(UUID(), '创建客户', 8, @customer_uuid, NOW(), NOW()),
(UUID(), '编辑客户资料', 2, @customer_uuid, NOW(), NOW()),
(UUID(), '删除客户', 1, @customer_uuid, NOW(), NOW());


SET @agent_uuid = (SELECT `uuid` FROM `permissions` WHERE `name` = '代理机构' AND `parent_uuid` = @resource_uuid);

INSERT INTO `permissions` (`uuid`, `name`, `bit`, `parent_uuid`, `created_at`, `updated_at`)
VALUES
(UUID(), '创建代理机构', 8, @agent_uuid, NOW(), NOW()),
(UUID(), '编辑代理机构资料', 2, @agent_uuid, NOW(), NOW()),
(UUID(), '删除代理机构', 1, @agent_uuid, NOW(), NOW());


SET @supplier_uuid = (SELECT `uuid` FROM `permissions` WHERE `name` = '供应商管理' AND `parent_uuid` = @resource_uuid);
INSERT INTO `permissions` (`uuid`, `name`, `bit`, `parent_uuid`, `created_at`, `updated_at`)
VALUES
(UUID(), '创建供应商', 8, @supplier_uuid, NOW(), NOW()),
(UUID(), '编辑供应商资料', 2, @supplier_uuid, NOW(), NOW()),
(UUID(), '删除供应商', 1, @supplier_uuid, NOW(), NOW());


SET @sku_uuid = (SELECT `uuid` FROM `permissions` WHERE `name` = 'SKU管理' AND `parent_uuid` = @resource_uuid);
INSERT INTO `permissions` (`uuid`, `name`, `bit`, `parent_uuid`, `created_at`, `updated_at`)
VALUES
(UUID(), '创建SKU', 8, @sku_uuid, NOW(), NOW()),
(UUID(), '编辑SKU资料', 2, @sku_uuid, NOW(), NOW()),
(UUID(), '删除SKU', 1, @sku_uuid, NOW(), NOW());


SET @product_manage_uuid = (SELECT `uuid` FROM `permissions` WHERE `name` = '商品管理' AND `parent_uuid` = @resource_uuid);
INSERT INTO `permissions` (`uuid`, `name`, `bit`, `parent_uuid`, `created_at`, `updated_at`)
VALUES
(UUID(), '创建商品', 8, @product_manage_uuid, NOW(), NOW()),
(UUID(), '编辑商品资料', 2, @product_manage_uuid, NOW(), NOW()),
(UUID(), '删除商品', 1, @product_manage_uuid, NOW(), NOW());


SET @product_category_uuid = (SELECT `uuid` FROM `permissions` WHERE `name` = '商品类别' AND `parent_uuid` = @resource_uuid);
INSERT INTO `permissions` (`uuid`, `name`, `bit`, `parent_uuid`, `created_at`, `updated_at`)
VALUES
(UUID(), '创建商品类别', 8, @product_category_uuid, NOW(), NOW()),
(UUID(), '编辑商品类别', 2, @product_category_uuid, NOW(), NOW()),
(UUID(), '删除商品类别', 1, @product_category_uuid, NOW(), NOW());



-- 插入员工管理子菜单
INSERT INTO `permissions` (`uuid`, `name`, `bit`, `parent_uuid`, `created_at`, `updated_at`)
VALUES
(UUID(), '部门管理', 4, @staff_uuid, NOW(), NOW()),
(UUID(), '用户管理', 4, @staff_uuid, NOW(), NOW()),
(UUID(), '权限管理', 4, @staff_uuid, NOW(), NOW());

SET @department_uuid = (SELECT `uuid` FROM `permissions` WHERE `name` = '部门管理' AND `parent_uuid` = @staff_uuid);
INSERT INTO `permissions` (`uuid`, `name`, `bit`, `parent_uuid`, `created_at`, `updated_at`)
VALUES
(UUID(), '创建部门', 8, @department_uuid, NOW(), NOW()),
(UUID(), '编辑部门', 2, @department_uuid, NOW(), NOW()),
(UUID(), '删除部门', 1, @department_uuid, NOW(), NOW());


SET @user_uuid = (SELECT `uuid` FROM `permissions` WHERE `name` = '用户管理' AND `parent_uuid` = @staff_uuid);
INSERT INTO `permissions` (`uuid`, `name`, `bit`, `parent_uuid`, `created_at`, `updated_at`)
VALUES
(UUID(), '创建用户', 8, @user_uuid, NOW(), NOW()),
(UUID(), '编辑用户', 2, @user_uuid, NOW(), NOW()),
(UUID(), '删除用户', 1, @user_uuid, NOW(), NOW());


SET @permission_uuid = (SELECT `uuid` FROM `permissions` WHERE `name` = '权限管理' AND `parent_uuid` = @staff_uuid);
INSERT INTO `permissions` (`uuid`, `name`, `bit`, `parent_uuid`, `created_at`, `updated_at`)
VALUES
(UUID(), '授权用户', 2, @permission_uuid, NOW(), NOW());



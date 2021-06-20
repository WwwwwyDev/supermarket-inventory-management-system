# 程序设计报告

姓名：吴文益 班级：软件193 学号：2019082323

## 题目

超市进销存管理系统

## 代码及版本管理地址

[gitee码云版本地址(点击跳转)](https://gitee.com/wu_wen_yi/supermarket-inventory-management-system)

## 问题描述

本项目学生需要通过此项目完成项目数据库设计，信息系统开发等。 

### 系统介绍

系统分为两个权限部分：管理员用户登录模块，普通用户登陆模块。管理员用户登录模块包括：供应商管理，员工管理，商品管理，会员管理，采购管理。普通用户登陆模块包括：个人信息，商品信息，采购信息。本项目实现以下功能：

### 管理员用户登录模块：

#### 供应商管理

批量录入供应商数据，手动输入供应商编号，供应商名称，供应商简称，地址，公司电话，邮件，联系人，联系人电话，备注。

查询并列出表中所有的供应商信息。

修改指定供应商的信息。

删除指定供应商数据。

#### 商品管理

批量录入商品数据，手动输入商品编号，商品名称，商品单价，供应商编号，商品简介，备注等。

查询并列出所有商品信息。

修改制定商品的商品信息。

删除指定商品数据。

#### 员工管理

批量录入员工数据，手动输入员工编号，员工姓名，员工密码，员工级别，员工电话，员工工资，备注等。

查询并列出所有员工信息。

修改指定员工信息。

删除指定员工数据。

#### 采购管理

录入采购主表数据，手动输入采购清单号，员工编号（外码），采购数量，采购总价，采购时间，备注。

录入采购明细数据，手动输入采购明细号，采购清单号（外码），商品编号（外码），采购数量，商品单价，商品总价，备注。

查询并列出所有采购主表和采购明细信息。

修改指定采购主表信息和采购明细信息。

删除指定采购主表信息和采购明细信息。

### 普通用户登陆模块：

#### 员工查询

查询自己的员工信息。

#### 商品查询

查询所有商品信息。

#### 采购查询

查询所有的采购信息。

## 软件设计

### 软件架构图

![WEB开发技术架构图](/Users/wuwenyi/Desktop/WEB开发技术架构图.png)

### 技术栈

- 视图层:vue3.0
- 控制层:gin+gorm
- 持久化:postgresql

### 持久化

#### 数据库概念模型设计(ER图)

![ER](/Users/wuwenyi/Desktop/ER.png)

#### PostgreSQL数据库物理模型设计

![模型](/Users/wuwenyi/Desktop/模型.png)

#### PostgreSQL

##### 触发器

###### 删除goods实体后触发

```sql
CREATE OR REPLACE FUNCTION "public"."del_goods_func"()
  RETURNS "pg_catalog"."trigger" AS $BODY$  
    BEGIN
				IF NEW.is_del = TRUE
				THEN
					UPDATE erp_purchase_details SET is_del = TRUE where purchase_details_goods = NEW.goods_id;
				END IF;
				RETURN NEW;
    END;  
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100
```

```sql
CREATE TRIGGER "del_goods_trigger" AFTER UPDATE OF "is_del" ON "public"."erp_goods"
FOR EACH ROW
EXECUTE PROCEDURE "public"."del_goods_func"();
```

###### 删除staff实体后触发

```sql
CREATE OR REPLACE FUNCTION "public"."del_staff_func"()
  RETURNS "pg_catalog"."trigger" AS $BODY$  
    BEGIN
				IF NEW.is_del = TRUE
				THEN
					UPDATE erp_purchase_list SET is_del = TRUE where purchase_list_staff = NEW.staff_id;
				END IF;
				RETURN NEW;
    END;  
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100
```

```sql
CREATE TRIGGER "del_staff_trigger" AFTER UPDATE OF "is_del" ON "public"."erp_staff"
FOR EACH ROW
EXECUTE PROCEDURE "public"."del_staff_func"();
```

###### 删除supplier实体后触发

```sql
CREATE OR REPLACE FUNCTION "public"."del_supplier_func"()
  RETURNS "pg_catalog"."trigger" AS $BODY$  
    BEGIN
				IF NEW.is_del = TRUE
				THEN
					UPDATE erp_goods SET is_del = TRUE where goods_supplier = NEW.supplier_id;
				END IF;
				RETURN NEW;
    END;  
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100
```

```sql
CREATE TRIGGER "del_supplier_trigger" AFTER UPDATE OF "is_del" ON "public"."erp_supplier"
FOR EACH ROW
EXECUTE PROCEDURE "public"."del_supplier_func"();
```

###### 删除purchase_list实体后触发

```sql
CREATE OR REPLACE FUNCTION "public"."del_purchase_list_func"()
  RETURNS "pg_catalog"."trigger" AS $BODY$  
    BEGIN
				IF NEW.is_del = TRUE
				THEN
					UPDATE erp_purchase_details SET is_del = TRUE where purchase_details_list = NEW.purchase_list_id;
				END IF;
				RETURN NEW;
    END;  
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100
```

```sql
CREATE TRIGGER "del_purchase_list_trigger" AFTER UPDATE OF "is_del" ON "public"."erp_purchase_list"
FOR EACH ROW
EXECUTE PROCEDURE "public"."del_purchase_list_func"();
```

###### 插入purchase_details实体后触发

```sql
CREATE OR REPLACE FUNCTION "public"."insert_details_func"()
  RETURNS "pg_catalog"."trigger" AS $BODY$  
    BEGIN  
        UPDATE erp_purchase_list SET purchase_list_number = purchase_list_number+NEW.purchase_details_number where purchase_list_id = NEW.purchase_details_list;
				UPDATE erp_purchase_list SET purchase_list_price = purchase_list_price+NEW.purchase_details_price where purchase_list_id = NEW.purchase_details_list;
				RETURN NEW;
    END;  
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100
```

```sql
CREATE TRIGGER "insert_detail_trigger" AFTER INSERT ON "public"."erp_purchase_details"
FOR EACH ROW
EXECUTE PROCEDURE "public"."insert_details_func"();
```

###### 更新purchase_details实体后触发

```sql
CREATE OR REPLACE FUNCTION "public"."update_details_func"()
  RETURNS "pg_catalog"."trigger" AS $BODY$  
    BEGIN  
				IF NEW.is_del = TRUE
				THEN
					UPDATE erp_purchase_list SET purchase_list_number = purchase_list_number-OLD.purchase_details_number where purchase_list_id = NEW.purchase_details_list;
					UPDATE erp_purchase_list SET purchase_list_price = purchase_list_price-OLD.purchase_details_price where purchase_list_id = NEW.purchase_details_list;
				ELSE
					UPDATE erp_purchase_list SET purchase_list_number = purchase_list_number+(NEW.purchase_details_number-OLD.purchase_details_number) where purchase_list_id = NEW.purchase_details_list;
					UPDATE erp_purchase_list SET purchase_list_price = purchase_list_price+(NEW.purchase_details_price-OLD.purchase_details_price) where purchase_list_id = NEW.purchase_details_list;
				END IF;
				RETURN NEW;
    END;  
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100
```

```sql
CREATE TRIGGER "update_detail_trigger" AFTER UPDATE ON "public"."erp_purchase_details"
FOR EACH ROW
EXECUTE PROCEDURE "public"."update_details_func"();
```

##### 索引

###### 对goods实体的goods_name字段建立唯一索引

```sql
CREATE UNIQUE INDEX "unip_name" ON "public"."erp_goods" USING btree (
  "goods_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
```

###### 对staff实体的staff_ name字段建立唯一索引

```sql
CREATE UNIQUE INDEX "uniq_staff_name" ON "public"."erp_staff" USING btree (
  "staff_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
```

###### 对supplier实体的supplier_name字段建立唯一索引

```sql
CREATE UNIQUE INDEX "uniq_supplier_name" ON "public"."erp_supplier" USING btree (
  "supplier_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
```

###### 对purchase_details实体的details_list和details_goods两字段建立联合索引

```sql
CREATE INDEX "list_goods_btree" ON "public"."erp_purchase_details" USING btree (
  "purchase_details_list" "pg_catalog"."int8_ops" ASC NULLS LAST,
  "purchase_details_goods" "pg_catalog"."int8_ops" ASC NULLS LAST
);
```

#### Redis

用于缓存postgresql中的sys_menu实体

### 控制层

#### models层

建立关系实体与数据库映射

##### Goods

| 属性              | 映射字段            |
| ----------------- | ------------------- |
| GoodsId           | goods_id            |
| GoodsName         | goods_name          |
| GoodsPrice        | goods_price         |
| GoodsSupplier     | goods_supplier      |
| GoodsSupplierName | goods_supplier_name |
| GoodsSynopsis     | goods_synopsis      |
| GoodsRemarks      | goods_remarks       |
| Isdel             | is_del              |
| UpdateTime        | update_time         |

##### Menu

| 属性        | 映射字段     |
| ----------- | ------------ |
| MenuId      | menu_id      |
| MenuName    | menu_name    |
| MenuAddress | menu_address |
| MenuLevel   | menu_level   |

##### PurchaseDetails

| 属性                     | 映射字段                    |
| ------------------------ | --------------------------- |
| PurchaseDetailsId        | purchase_details_id         |
| PurchaseDetailsList      | purchase_details_list       |
| PurchaseDetailsGoods     | purchase_details_goods      |
| PurchaseDetailsGoodsName | purchase_details_goods_name |
| PurchaseDetailsNumber    | purchase_details_number     |
| PurchaseDetailsPrice     | purchase_details_price      |
| PurchaseDetailsRemarks   | purchase_details_remarks    |
| UpdateTime               | update_time                 |
| IsDel                    | is_del                      |

##### PurchaseList

| 属性                  | 映射字段                 |
| --------------------- | ------------------------ |
| PurchaseListId        | purchase_list_id         |
| PurchaseListStaff     | purchase_list_staff      |
| PurchaseListStaffName | purchase_list_staff_name |
| PurchaseListNumber    | purchase_list_number     |
| PurchaseListPrice     | purchase_list_price      |
| PurchaseListTime      | purchase_list_time       |
| PurchaseListRemarks   | purchase_list_remarks    |
| UpdateTime            | update_time              |
| IsDel                 | is_del                   |

##### Staff

| 属性           | 映射字段        |
| -------------- | --------------- |
| StaffId        | staff_id        |
| StaffName      | staff_name      |
| StaffPassword  | staff_password  |
| StaffLevel     | staff_level     |
| StaffTelephone | staff_telephone |
| StaffSalary    | staff_salary    |
| StaffRemarks   | staff_remarks   |
| IsDel          | is_del          |
| UpdateTime     | update_time     |

##### Supplier

| 属性               | 映射字段            |
| ------------------ | ------------------- |
| SupplierId         | supplier_id         |
| SupplierName       | supplier_name       |
| SupplierShortname  | supplier_shortname  |
| SupplierAddress    | supplier_address    |
| SupplierCtelephone | supplier_ctelephone |
| SupplierEmail      | supplier_email      |
| SupplierLiaisonman | supplier_liaisonman |
| SupplierLtelephone | supplier_ltelephone |
| SupplierRemarks    | supplier_remarks    |
| IsDel              | is_del              |
| UpdateTime         | update_time         |

#### dao层

##### GoodsDao

| 函数名           | 函数功能                     |
| ---------------- | ---------------------------- |
| GetAllGoods      | 与supplier实体进行连接查询   |
| GetGoodPriceById | 通过id获取goods实体的price值 |
| DelGoods         | 删除goods实体                |
| UpdateGoods      | 更新goods实体                |
| AddGoods         | 添加goods实体                |

###### 部分多表查询代码

内连supplier表，对goods_name字段进行模糊分页查询，并计算查询总条数

```go
db.Table("erp_goods g").Select("g.goods_id, g.goods_name, g.goods_price, g.goods_supplier,s.supplier_name goods_supplier_name, g.goods_remarks,g.goods_synopsis,g.is_del, g.update_time").Joins("INNER JOIN erp_supplier s ON g.goods_supplier = s.supplier_id").Where("g.goods_name like ? and g.is_del = false", searchName+"%").Order("goods_id ASC").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&goodsData).Error
```

##### MenuDao

| 函数名         | 函数功能                 |
| -------------- | ------------------------ |
| GetMenuByLevel | 通过等级获取menu路由菜单 |

##### PurchaseDetailsDao

| 函数名                | 函数功能                                 |
| --------------------- | ---------------------------------------- |
| GetAllPurchaseDetails | 对purchase_details实体的单表分页模糊查询 |
| DelPurchaseDetails    | 删除purchase_details实体                 |
| UpdatePurchaseDetails | 更新purchase_details实体                 |
| AddPurchaseDetails    | 添加purchase_details实体                 |

##### PurchaseListDao

| 函数名             | 函数功能                              |
| ------------------ | ------------------------------------- |
| GetAllPurchaseList | 对purchase_list实体的单表分页模糊查询 |
| DelPurchaseList    | 删除purchase_list实体                 |
| UpdatePurchaseList | 更新purchase_list实体                 |
| AddPurchaseList    | 添加purchase_list实体                 |


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

![软件架构图](https://images.gitee.com/uploads/images/2021/0620/104705_3051979e_6512114.png "WEB开发技术架构图.png")

### 技术栈

- 视图层:vue3.0
- 控制层:gin+gorm
- 持久化:postgresql

### 持久化

#### 数据库概念模型设计(ER图)

![ER](https://images.gitee.com/uploads/images/2021/0620/104730_4125dced_6512114.png "ER.png")

#### PostgreSQL数据库物理模型设计

![模型](https://images.gitee.com/uploads/images/2021/0620/104838_1734bd16_6512114.png "模型.png")

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

按照最左优先原则，将details_list字段放在左边

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

##### goodsDao

| 函数名           | 函数功能                     |
| ---------------- | ---------------------------- |
| GetAllGoods      | 与supplier实体的多表查询     |
| GetGoodPriceById | 通过id获取goods实体的price值 |
| DelGoods         | 删除goods实体                |
| UpdateGoods      | 更新goods实体                |
| AddGoods         | 添加goods实体                |

###### 部分多表查询代码

内连supplier表，对goods_name字段进行模糊分页查询，并计算查询总条数

```go
db.Table("erp_goods g").Select("g.goods_id, g.goods_name, g.goods_price, g.goods_supplier,s.supplier_name goods_supplier_name, g.goods_remarks,g.goods_synopsis,g.is_del, g.update_time").Joins("INNER JOIN erp_supplier s ON g.goods_supplier = s.supplier_id").Where("g.goods_name like ? and g.is_del = false", searchName+"%").Order("goods_id ASC").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&goodsData).Error
```

##### menuDao

| 函数名         | 函数功能                 |
| -------------- | ------------------------ |
| GetMenuByLevel | 通过等级获取menu路由菜单 |

##### purchaseDetailsDao

| 函数名                | 函数功能                 |
| --------------------- | ------------------------ |
| GetAllPurchaseDetails | 与goods实体的多表查询    |
| DelPurchaseDetails    | 删除purchase_details实体 |
| UpdatePurchaseDetails | 更新purchase_details实体 |
| AddPurchaseDetails    | 添加purchase_details实体 |

###### 部分多表查询代码

内连goods表，对purchase_details_list字段和goods_name字段进行分页查询，并计算查询总条数

```go
err := db.Table("erp_purchase_details d").Select("d.purchase_details_id, d.purchase_details_list, d.purchase_details_goods, g.goods_name purchase_details_goods_name,d.purchase_details_number, d.purchase_details_price, d.purchase_details_remarks, d.update_time, d.is_del").Joins("INNER JOIN erp_goods g ON d.purchase_details_goods = g.goods_id").Where("d.purchase_details_list = ?", searchList).Where("d.is_del = false").Where("g.goods_name like ?",searchGoodsName+"%").Order("purchase_details_id ASC").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&purchaseDetailsData).Error
```

##### purchaseListDao

| 函数名             | 函数功能              |
| ------------------ | --------------------- |
| GetAllPurchaseList | 与staff实体的多表查询 |
| DelPurchaseList    | 删除purchase_list实体 |
| UpdatePurchaseList | 更新purchase_list实体 |
| AddPurchaseList    | 添加purchase_list实体 |

###### 部分多表查询代码

内连staff表，对staff_name字段进行模糊分页查询，并计算查询总条数

```go
err := db.Table("erp_purchase_list p").Select("p.purchase_list_id, p.purchase_list_number, p.purchase_list_price, p.purchase_list_remarks, p.purchase_list_staff, s.staff_name purchase_list_staff_name, p.purchase_list_time, p.update_time, p.is_del").Joins("INNER JOIN erp_staff s ON p.purchase_list_staff = s.staff_id").Where("s.staff_name like ? and p.is_del = false", searchName+"%").Order("purchase_list_id ASC").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&purchaseListData).Error
```

##### staffDao

| 函数名       | 函数功能                      |
| ------------ | ----------------------------- |
| GetAllStaff  | 对staff实体的单表分页模糊查询 |
| GetStaffById | 通过id查询staff实体           |
| AddStaff     | 添加staff实体                 |
| UpdateStaff  | 更新staff实体                 |
| DelStaff     | 删除staff实体                 |

##### supplierDao

| 函数名         | 函数功能                         |
| -------------- | -------------------------------- |
| GetAllSupplier | 对supplier实体的单表分页模糊查询 |
| AddSupplier    | 添加supplier实体                 |
| UpdateSupplier | 更新supplier实体                 |
| DelSupplier    | 删除supplier实体                 |

#### service层

##### goodsService

| 函数名           | 函数功能            |
| ---------------- | ------------------- |
| GetAllGoods      | 获取Goods对象数组   |
| DelGoods         | 删除Goods对象       |
| UpdateGoods      | 更新Goods对象       |
| AddGoods         | 添加Goods对象       |
| GetGoodPriceById | 通过id获取Goods对象 |

##### menuService

| 函数名         | 函数功能             |
| -------------- | -------------------- |
| GetMenuByLevel | 通过等级获取Menu对象 |

##### purchaseDetailsService

| 函数名                | 函数功能                    |
| --------------------- | --------------------------- |
| GetAllPurchaseDetails | 获取PurchaseDetails对象数组 |
| DelPurchaseDetails    | 删除PurchaseDetails对象     |
| UpdatePurchaseDetails | 更新PurchaseDetails对象     |
| AddPurchaseDetails    | 添加PurchaseDetails对象     |

##### purchaseListService

| 函数名                | 函数功能                    |
| --------------------- | --------------------------- |
| GetAllPurchaseDetails | 获取PurchaseDetails对象数组 |
| DelPurchaseDetails    | 删除PurchaseDetails对象     |
| UpdatePurchaseDetails | 更新PurchaseDetails对象     |
| AddPurchaseDetails    | 添加PurchaseDetails对象     |

##### staffService

| 函数名         | 函数功能                            |
| -------------- | ----------------------------------- |
| GetStaffByName | 通过name获取Staff对象(用于权限校验) |
| AddStaff       | 添加Staff对象                       |
| UpdateStaff    | 更新Staff对象                       |
| DelStaff       | 删除Staff对象                       |
| GetStaffById   | 通过id获取Staff对象                 |
| GetAllStaff    | 获取Staff对象数组                   |

##### supplierService

| 函数名         | 函数功能             |
| -------------- | -------------------- |
| GetAllSupplier | 获取Supplier对象数组 |
| DelSupplier    | 删除Supplier对象     |
| UpdateSupplier | 更新Supplier对象     |
| AddSupplier    | 添加Supplier对象     |

#### controller层

##### goodsController

| 函数名      | 函数功能                                                     |
| ----------- | ------------------------------------------------------------ |
| GetGoods    | 接受get请求头的"page","limit","searchName"参数，进行查询操作 |
| DelGoods    | 接受delete请求头的“id”参数，进行删除操作                     |
| UpdateGoods | 接受put请求体的json数据，进行更新操作                        |
| AddGoods    | 接受post请求体的json数据，进行添加操作                       |

##### loginController

| 函数名 | 函数功能                           |
| ------ | ---------------------------------- |
| Login  | 接受post请求体的json数据，进行登录 |

##### menuController

| 函数名  | 函数功能                             |
| ------- | ------------------------------------ |
| GetMenu | 接受get请求头的“level”参数，进行登录 |

##### purchaseDetailsController

| 函数名                | 函数功能                                                     |
| --------------------- | ------------------------------------------------------------ |
| GetPurchaseDetails    | 接受get请求头的"page","limit",""searchList",searchGoodsName"参数，进行查询操作 |
| DelPurchaseDetails    | 接受delete请求头的“id”参数，进行删除操作                     |
| UpdatePurchaseDetails | 接受put请求体的json数据，进行更新操作                        |
| AddPurchaseDetails    | 接受post请求体的json数据，进行添加操作                       |

##### purchaseListController

| 函数名             | 函数功能                                                     |
| ------------------ | ------------------------------------------------------------ |
| GetPurchaseList    | 接受get请求头的"page","limit","searchName"参数，进行查询操作 |
| DelPurchaseList    | 接受delete请求头的“id”参数，进行删除操作                     |
| UpdatePurchaseList | 接受put请求体的json数据，进行更新操作                        |
| AddPurchaseList    | 接受post请求体的json数据，进行添加操作                       |

##### staffController

| 函数名      | 函数功能                                                     |
| ----------- | ------------------------------------------------------------ |
| GetStaff    | 接受get请求头的"page","limit","searchName"参数，进行查询操作 |
| DelStaff    | 接受delete请求头的“id”参数，进行删除操作                     |
| UpdateStaff | 接受put请求体的json数据，进行更新操作                        |
| AddStaff    | 接受post请求体的json数据，进行添加操作                       |

##### supplierController

| 函数名         | 函数功能                                                     |
| -------------- | ------------------------------------------------------------ |
| GetSupplier    | 接受get请求头的"page","limit","searchName"参数，进行查询操作 |
| DelSupplier    | 接受delete请求头的“id”参数，进行删除操作                     |
| UpdateSupplier | 接受put请求体的json数据，进行更新操作                        |
| AddSupplier    | 接受post请求体的json数据，进行添加操作                       |

#### router层

请求接口根路径“/system/apis/”，入口为每个实体类名小写，如“/system/apis/supplier”

对每一种资源CRUD操作，依次对应HTTP协议中的四类请求：GET,PUT,POST,DELETE

### 视图层

使用element-plus进行视图美化

#### 登录界面
![输入图片说明](https://images.gitee.com/uploads/images/2021/0620/105047_a2fab8b5_6512114.png "截屏2021-06-20 上午10.38.21.png")

#### 动态路由
![输入图片说明](https://images.gitee.com/uploads/images/2021/0620/105103_5143c807_6512114.png "截屏2021-06-20 上午10.40.13.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0620/105116_b37e0bef_6512114.png "截屏2021-06-20 上午10.40.28.png")

#### 管理员视图
![输入图片说明](https://images.gitee.com/uploads/images/2021/0620/105140_7da87ff1_6512114.png "截屏2021-06-20 上午10.38.49.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0620/105151_69167852_6512114.png "截屏2021-06-20 上午10.38.54.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0620/105203_45a9aa03_6512114.png "截屏2021-06-20 上午10.38.58.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0620/105213_a7ae1561_6512114.png "截屏2021-06-20 上午10.39.03.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0620/105223_fdf649fa_6512114.png "截屏2021-06-20 上午10.39.08.png")

#### 普通用户视图
![输入图片说明](https://images.gitee.com/uploads/images/2021/0620/105241_a55b6a91_6512114.png "截屏2021-06-20 上午10.39.48.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0620/105251_3b1e2951_6512114.png "截屏2021-06-20 上午10.39.53.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0620/105259_9aa50050_6512114.png "截屏2021-06-20 上午10.39.57.png")

#### 查看个人信息视图
![输入图片说明](https://images.gitee.com/uploads/images/2021/0620/105312_d8ed057e_6512114.png "截屏2021-06-20 上午10.39.22.png")

#### 添加视图
![输入图片说明](https://images.gitee.com/uploads/images/2021/0620/105331_34c9f92f_6512114.png "截屏2021-06-20 上午10.40.46.png")

#### 编辑视图
![输入图片说明](https://images.gitee.com/uploads/images/2021/0620/105349_55cdbfad_6512114.png "截屏2021-06-20 上午10.41.01.png")
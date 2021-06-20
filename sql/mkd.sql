/*
 Navicat Premium Data Transfer

 Source Server         : 本机p2
 Source Server Type    : PostgreSQL
 Source Server Version : 90426
 Source Host           : localhost:5432
 Source Catalog        : test3
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 90426
 File Encoding         : 65001

 Date: 20/06/2021 11:07:29
*/


-- ----------------------------
-- Sequence structure for erp_goods_goods_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."erp_goods_goods_id_seq";
CREATE SEQUENCE "public"."erp_goods_goods_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."erp_goods_goods_id_seq" OWNER TO "wuwenyi";

-- ----------------------------
-- Sequence structure for erp_purchase_details_purchase_details_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."erp_purchase_details_purchase_details_id_seq";
CREATE SEQUENCE "public"."erp_purchase_details_purchase_details_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."erp_purchase_details_purchase_details_id_seq" OWNER TO "wuwenyi";

-- ----------------------------
-- Sequence structure for erp_purchase_list_purchase_list_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."erp_purchase_list_purchase_list_id_seq";
CREATE SEQUENCE "public"."erp_purchase_list_purchase_list_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."erp_purchase_list_purchase_list_id_seq" OWNER TO "wuwenyi";

-- ----------------------------
-- Sequence structure for erp_staff_staff_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."erp_staff_staff_id_seq";
CREATE SEQUENCE "public"."erp_staff_staff_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."erp_staff_staff_id_seq" OWNER TO "wuwenyi";

-- ----------------------------
-- Sequence structure for erp_supplier_supplier_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."erp_supplier_supplier_id_seq";
CREATE SEQUENCE "public"."erp_supplier_supplier_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."erp_supplier_supplier_id_seq" OWNER TO "wuwenyi";

-- ----------------------------
-- Table structure for erp_goods
-- ----------------------------
DROP TABLE IF EXISTS "public"."erp_goods";
CREATE TABLE "public"."erp_goods" (
  "goods_id" int8 NOT NULL DEFAULT nextval('erp_goods_goods_id_seq'::regclass),
  "goods_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "goods_price" int8 NOT NULL,
  "goods_supplier" int8 NOT NULL,
  "goods_synopsis" varchar(255) COLLATE "pg_catalog"."default",
  "goods_remarks" varchar(255) COLLATE "pg_catalog"."default",
  "is_del" bool NOT NULL,
  "update_time" timestamptz(6) NOT NULL
)
;
ALTER TABLE "public"."erp_goods" OWNER TO "wuwenyi";
COMMENT ON COLUMN "public"."erp_goods"."goods_id" IS '编号';
COMMENT ON COLUMN "public"."erp_goods"."goods_name" IS '名称';
COMMENT ON COLUMN "public"."erp_goods"."goods_price" IS '单价(分)';
COMMENT ON COLUMN "public"."erp_goods"."goods_supplier" IS '供应商编号(外键)';
COMMENT ON COLUMN "public"."erp_goods"."goods_synopsis" IS '简介';
COMMENT ON COLUMN "public"."erp_goods"."goods_remarks" IS '备注';
COMMENT ON TABLE "public"."erp_goods" IS '商品表';

-- ----------------------------
-- Records of erp_goods
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for erp_purchase_details
-- ----------------------------
DROP TABLE IF EXISTS "public"."erp_purchase_details";
CREATE TABLE "public"."erp_purchase_details" (
  "purchase_details_id" int8 NOT NULL DEFAULT nextval('erp_purchase_details_purchase_details_id_seq'::regclass),
  "purchase_details_list" int8 NOT NULL,
  "purchase_details_goods" int8 NOT NULL,
  "purchase_details_number" int4 NOT NULL,
  "purchase_details_price" int8 NOT NULL,
  "purchase_details_remarks" varchar(255) COLLATE "pg_catalog"."default",
  "update_time" timestamptz(6) NOT NULL,
  "is_del" bool NOT NULL
)
;
ALTER TABLE "public"."erp_purchase_details" OWNER TO "wuwenyi";
COMMENT ON COLUMN "public"."erp_purchase_details"."purchase_details_id" IS '编号';
COMMENT ON COLUMN "public"."erp_purchase_details"."purchase_details_list" IS '采购清单';
COMMENT ON COLUMN "public"."erp_purchase_details"."purchase_details_goods" IS '商品';
COMMENT ON COLUMN "public"."erp_purchase_details"."purchase_details_number" IS '数量';
COMMENT ON COLUMN "public"."erp_purchase_details"."purchase_details_price" IS '总价(分)';
COMMENT ON COLUMN "public"."erp_purchase_details"."purchase_details_remarks" IS '备注';
COMMENT ON COLUMN "public"."erp_purchase_details"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."erp_purchase_details"."is_del" IS '是否删除';
COMMENT ON TABLE "public"."erp_purchase_details" IS '采购明细表';

-- ----------------------------
-- Records of erp_purchase_details
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for erp_purchase_list
-- ----------------------------
DROP TABLE IF EXISTS "public"."erp_purchase_list";
CREATE TABLE "public"."erp_purchase_list" (
  "purchase_list_id" int8 NOT NULL DEFAULT nextval('erp_purchase_list_purchase_list_id_seq'::regclass),
  "purchase_list_staff" int8 NOT NULL,
  "purchase_list_number" int4 NOT NULL,
  "purchase_list_price" int8 NOT NULL,
  "purchase_list_time" timestamptz(6) NOT NULL,
  "purchase_list_remarks" varchar(255) COLLATE "pg_catalog"."default",
  "update_time" timestamptz(6) NOT NULL,
  "is_del" bool NOT NULL
)
;
ALTER TABLE "public"."erp_purchase_list" OWNER TO "wuwenyi";
COMMENT ON COLUMN "public"."erp_purchase_list"."purchase_list_id" IS '编号';
COMMENT ON COLUMN "public"."erp_purchase_list"."purchase_list_staff" IS '员工编号(外键)';
COMMENT ON COLUMN "public"."erp_purchase_list"."purchase_list_number" IS '数量';
COMMENT ON COLUMN "public"."erp_purchase_list"."purchase_list_price" IS '总价(分)';
COMMENT ON COLUMN "public"."erp_purchase_list"."purchase_list_time" IS '时间';
COMMENT ON COLUMN "public"."erp_purchase_list"."purchase_list_remarks" IS '备注';
COMMENT ON COLUMN "public"."erp_purchase_list"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."erp_purchase_list"."is_del" IS '是否删除';
COMMENT ON TABLE "public"."erp_purchase_list" IS '采购清单表';

-- ----------------------------
-- Records of erp_purchase_list
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for erp_staff
-- ----------------------------
DROP TABLE IF EXISTS "public"."erp_staff";
CREATE TABLE "public"."erp_staff" (
  "staff_id" int8 NOT NULL DEFAULT nextval('erp_staff_staff_id_seq'::regclass),
  "staff_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "staff_password" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "staff_level" int2 NOT NULL,
  "staff_telephone" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "staff_salary" int8 NOT NULL,
  "staff_remarks" varchar(255) COLLATE "pg_catalog"."default",
  "is_del" bool NOT NULL DEFAULT false,
  "update_time" timestamptz(0) NOT NULL
)
;
ALTER TABLE "public"."erp_staff" OWNER TO "wuwenyi";
COMMENT ON COLUMN "public"."erp_staff"."staff_id" IS '编号';
COMMENT ON COLUMN "public"."erp_staff"."staff_name" IS '姓名';
COMMENT ON COLUMN "public"."erp_staff"."staff_password" IS '密码';
COMMENT ON COLUMN "public"."erp_staff"."staff_level" IS '级别';
COMMENT ON COLUMN "public"."erp_staff"."staff_telephone" IS '电话';
COMMENT ON COLUMN "public"."erp_staff"."staff_salary" IS '工资(分)';
COMMENT ON COLUMN "public"."erp_staff"."staff_remarks" IS '备注';
COMMENT ON COLUMN "public"."erp_staff"."is_del" IS '是否删除';
COMMENT ON COLUMN "public"."erp_staff"."update_time" IS '更新日期';
COMMENT ON TABLE "public"."erp_staff" IS '员工表';

-- ----------------------------
-- Records of erp_staff
-- ----------------------------
BEGIN;
INSERT INTO "public"."erp_staff" VALUES (1, 'admin', 'admin', 2, '123456', 1223456, '管理员', 'f', '2021-06-20 11:04:59+08');
INSERT INTO "public"."erp_staff" VALUES (267868, 'testtest', 'testtest', 1, '123456', 31324, '普通用户', 'f', '2021-06-20 11:06:24+08');
COMMIT;

-- ----------------------------
-- Table structure for erp_supplier
-- ----------------------------
DROP TABLE IF EXISTS "public"."erp_supplier";
CREATE TABLE "public"."erp_supplier" (
  "supplier_id" int8 NOT NULL DEFAULT nextval('erp_supplier_supplier_id_seq'::regclass),
  "supplier_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "supplier_shortname" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "supplier_address" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "supplier_ctelephone" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "supplier_email" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "supplier_liaisonman" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "supplier_ltelephone" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "supplier_remarks" varchar(255) COLLATE "pg_catalog"."default",
  "is_del" bool NOT NULL,
  "update_time" timestamptz(6) NOT NULL
)
;
ALTER TABLE "public"."erp_supplier" OWNER TO "wuwenyi";
COMMENT ON COLUMN "public"."erp_supplier"."supplier_id" IS '序号';
COMMENT ON COLUMN "public"."erp_supplier"."supplier_name" IS '名称';
COMMENT ON COLUMN "public"."erp_supplier"."supplier_shortname" IS '简称';
COMMENT ON COLUMN "public"."erp_supplier"."supplier_address" IS '地址';
COMMENT ON COLUMN "public"."erp_supplier"."supplier_ctelephone" IS '公司电话';
COMMENT ON COLUMN "public"."erp_supplier"."supplier_email" IS '邮箱';
COMMENT ON COLUMN "public"."erp_supplier"."supplier_liaisonman" IS '联系人';
COMMENT ON COLUMN "public"."erp_supplier"."supplier_ltelephone" IS '联系人电话';
COMMENT ON COLUMN "public"."erp_supplier"."supplier_remarks" IS '备注';
COMMENT ON COLUMN "public"."erp_supplier"."is_del" IS '是否删除';
COMMENT ON COLUMN "public"."erp_supplier"."update_time" IS '更新时间';
COMMENT ON TABLE "public"."erp_supplier" IS '供应商表';

-- ----------------------------
-- Records of erp_supplier
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_menu";
CREATE TABLE "public"."sys_menu" (
  "menu_id" int4 NOT NULL,
  "menu_address" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "menu_level" int4 NOT NULL,
  "menu_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "public"."sys_menu" OWNER TO "wuwenyi";
COMMENT ON COLUMN "public"."sys_menu"."menu_id" IS '编号';
COMMENT ON COLUMN "public"."sys_menu"."menu_address" IS '路由地址';
COMMENT ON COLUMN "public"."sys_menu"."menu_level" IS '权限';
COMMENT ON COLUMN "public"."sys_menu"."menu_name" IS '名称';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_menu" VALUES (200, '/goodsa', 2, '商品管理');
INSERT INTO "public"."sys_menu" VALUES (201, '/pda', 2, '采购明细管理');
INSERT INTO "public"."sys_menu" VALUES (202, '/pla', 2, '采购主表管理');
INSERT INTO "public"."sys_menu" VALUES (203, '/suppliera', 2, '供应商管理');
INSERT INTO "public"."sys_menu" VALUES (204, '/staffa', 2, '员工管理');
INSERT INTO "public"."sys_menu" VALUES (100, '/goods', 1, '商品信息');
INSERT INTO "public"."sys_menu" VALUES (101, '/pl', 1, '采购主表');
INSERT INTO "public"."sys_menu" VALUES (102, '/pd', 1, '采购明细');
COMMIT;

-- ----------------------------
-- Function structure for del_goods_func
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."del_goods_func"();
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
  COST 100;
ALTER FUNCTION "public"."del_goods_func"() OWNER TO "wuwenyi";

-- ----------------------------
-- Function structure for del_purchase_list_func
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."del_purchase_list_func"();
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
  COST 100;
ALTER FUNCTION "public"."del_purchase_list_func"() OWNER TO "wuwenyi";

-- ----------------------------
-- Function structure for del_staff_func
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."del_staff_func"();
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
  COST 100;
ALTER FUNCTION "public"."del_staff_func"() OWNER TO "wuwenyi";

-- ----------------------------
-- Function structure for del_supplier_func
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."del_supplier_func"();
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
  COST 100;
ALTER FUNCTION "public"."del_supplier_func"() OWNER TO "wuwenyi";

-- ----------------------------
-- Function structure for insert_details_func
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."insert_details_func"();
CREATE OR REPLACE FUNCTION "public"."insert_details_func"()
  RETURNS "pg_catalog"."trigger" AS $BODY$  
    BEGIN  
        UPDATE erp_purchase_list SET purchase_list_number = purchase_list_number+NEW.purchase_details_number where purchase_list_id = NEW.purchase_details_list;
				UPDATE erp_purchase_list SET purchase_list_price = purchase_list_price+NEW.purchase_details_price where purchase_list_id = NEW.purchase_details_list;
				RETURN NEW;
    END;  
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION "public"."insert_details_func"() OWNER TO "wuwenyi";

-- ----------------------------
-- Function structure for update_details_func
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."update_details_func"();
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
  COST 100;
ALTER FUNCTION "public"."update_details_func"() OWNER TO "wuwenyi";

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."erp_goods_goods_id_seq"
OWNED BY "public"."erp_goods"."goods_id";
SELECT setval('"public"."erp_goods_goods_id_seq"', 12, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."erp_purchase_details_purchase_details_id_seq"
OWNED BY "public"."erp_purchase_details"."purchase_details_id";
SELECT setval('"public"."erp_purchase_details_purchase_details_id_seq"', 17, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."erp_purchase_list_purchase_list_id_seq"
OWNED BY "public"."erp_purchase_list"."purchase_list_id";
SELECT setval('"public"."erp_purchase_list_purchase_list_id_seq"', 25, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."erp_staff_staff_id_seq"
OWNED BY "public"."erp_staff"."staff_id";
SELECT setval('"public"."erp_staff_staff_id_seq"', 267869, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."erp_supplier_supplier_id_seq"
OWNED BY "public"."erp_supplier"."supplier_id";
SELECT setval('"public"."erp_supplier_supplier_id_seq"', 10, true);

-- ----------------------------
-- Indexes structure for table erp_goods
-- ----------------------------
CREATE UNIQUE INDEX "unip_name" ON "public"."erp_goods" USING btree (
  "goods_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Triggers structure for table erp_goods
-- ----------------------------
CREATE TRIGGER "del_goods_trigger" AFTER UPDATE OF "is_del" ON "public"."erp_goods"
FOR EACH ROW
EXECUTE PROCEDURE "public"."del_goods_func"();

-- ----------------------------
-- Primary Key structure for table erp_goods
-- ----------------------------
ALTER TABLE "public"."erp_goods" ADD CONSTRAINT "_copy_3" PRIMARY KEY ("goods_id");

-- ----------------------------
-- Indexes structure for table erp_purchase_details
-- ----------------------------
CREATE INDEX "list_goods_btree" ON "public"."erp_purchase_details" USING btree (
  "purchase_details_list" "pg_catalog"."int8_ops" ASC NULLS LAST,
  "purchase_details_goods" "pg_catalog"."int8_ops" ASC NULLS LAST
);

-- ----------------------------
-- Triggers structure for table erp_purchase_details
-- ----------------------------
CREATE TRIGGER "insert_detail_trigger" AFTER INSERT ON "public"."erp_purchase_details"
FOR EACH ROW
EXECUTE PROCEDURE "public"."insert_details_func"();
CREATE TRIGGER "update_detail_trigger" AFTER UPDATE ON "public"."erp_purchase_details"
FOR EACH ROW
EXECUTE PROCEDURE "public"."update_details_func"();

-- ----------------------------
-- Primary Key structure for table erp_purchase_details
-- ----------------------------
ALTER TABLE "public"."erp_purchase_details" ADD CONSTRAINT "_copy_1" PRIMARY KEY ("purchase_details_id");

-- ----------------------------
-- Triggers structure for table erp_purchase_list
-- ----------------------------
CREATE TRIGGER "del_purchase_list_trigger" AFTER UPDATE OF "is_del" ON "public"."erp_purchase_list"
FOR EACH ROW
EXECUTE PROCEDURE "public"."del_purchase_list_func"();

-- ----------------------------
-- Primary Key structure for table erp_purchase_list
-- ----------------------------
ALTER TABLE "public"."erp_purchase_list" ADD CONSTRAINT "_copy_2" PRIMARY KEY ("purchase_list_id");

-- ----------------------------
-- Indexes structure for table erp_staff
-- ----------------------------
CREATE UNIQUE INDEX "uniq_staff_name" ON "public"."erp_staff" USING btree (
  "staff_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Triggers structure for table erp_staff
-- ----------------------------
CREATE TRIGGER "del_staff_trigger" AFTER UPDATE OF "is_del" ON "public"."erp_staff"
FOR EACH ROW
EXECUTE PROCEDURE "public"."del_staff_func"();

-- ----------------------------
-- Primary Key structure for table erp_staff
-- ----------------------------
ALTER TABLE "public"."erp_staff" ADD CONSTRAINT "_copy_4" PRIMARY KEY ("staff_id");

-- ----------------------------
-- Indexes structure for table erp_supplier
-- ----------------------------
CREATE UNIQUE INDEX "uniq_supplier_name" ON "public"."erp_supplier" USING btree (
  "supplier_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Triggers structure for table erp_supplier
-- ----------------------------
CREATE TRIGGER "del_supplier_trigger" AFTER UPDATE OF "is_del" ON "public"."erp_supplier"
FOR EACH ROW
EXECUTE PROCEDURE "public"."del_supplier_func"();

-- ----------------------------
-- Primary Key structure for table erp_supplier
-- ----------------------------
ALTER TABLE "public"."erp_supplier" ADD CONSTRAINT "erp_supplier_pkey" PRIMARY KEY ("supplier_id");

-- ----------------------------
-- Primary Key structure for table sys_menu
-- ----------------------------
ALTER TABLE "public"."sys_menu" ADD CONSTRAINT "sys_router_pkey" PRIMARY KEY ("menu_id");

-- ----------------------------
-- Foreign Keys structure for table erp_goods
-- ----------------------------
ALTER TABLE "public"."erp_goods" ADD CONSTRAINT "fk_erp_goods_erp_supplier_1" FOREIGN KEY ("goods_supplier") REFERENCES "public"."erp_supplier" ("supplier_id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table erp_purchase_details
-- ----------------------------
ALTER TABLE "public"."erp_purchase_details" ADD CONSTRAINT "fk_erp_purchase_detail_erp_purchase_list_1" FOREIGN KEY ("purchase_details_list") REFERENCES "public"."erp_purchase_list" ("purchase_list_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."erp_purchase_details" ADD CONSTRAINT "fk_erp_purchase_detail_erp_purchase_list_2" FOREIGN KEY ("purchase_details_goods") REFERENCES "public"."erp_goods" ("goods_id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table erp_purchase_list
-- ----------------------------
ALTER TABLE "public"."erp_purchase_list" ADD CONSTRAINT "fk_erp_purchase_list_erp_staff_1" FOREIGN KEY ("purchase_list_staff") REFERENCES "public"."erp_staff" ("staff_id") ON DELETE NO ACTION ON UPDATE NO ACTION;

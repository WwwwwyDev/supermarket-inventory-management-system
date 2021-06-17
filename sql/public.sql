/*
 Navicat Premium Data Transfer

 Source Server         : 本机p2
 Source Server Type    : PostgreSQL
 Source Server Version : 90426
 Source Host           : localhost:5432
 Source Catalog        : test2
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 90426
 File Encoding         : 65001

 Date: 17/06/2021 23:27:12
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
-- Sequence structure for sys_router_router_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_router_router_id_seq";
CREATE SEQUENCE "public"."sys_router_router_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."sys_router_router_id_seq" OWNER TO "wuwenyi";

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
  "goods_remarks" varchar(255) COLLATE "pg_catalog"."default"
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
  "purchase_details_remarks" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."erp_purchase_details" OWNER TO "wuwenyi";
COMMENT ON COLUMN "public"."erp_purchase_details"."purchase_details_id" IS '编号';
COMMENT ON COLUMN "public"."erp_purchase_details"."purchase_details_list" IS '采购清单';
COMMENT ON COLUMN "public"."erp_purchase_details"."purchase_details_goods" IS '商品';
COMMENT ON COLUMN "public"."erp_purchase_details"."purchase_details_number" IS '数量';
COMMENT ON COLUMN "public"."erp_purchase_details"."purchase_details_price" IS '总价(分)';
COMMENT ON COLUMN "public"."erp_purchase_details"."purchase_details_remarks" IS '备注';
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
  "purchase_list_time" timestamp(6) NOT NULL,
  "purchase_list_remarks" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."erp_purchase_list" OWNER TO "wuwenyi";
COMMENT ON COLUMN "public"."erp_purchase_list"."purchase_list_id" IS '编号';
COMMENT ON COLUMN "public"."erp_purchase_list"."purchase_list_staff" IS '员工编号(外键)';
COMMENT ON COLUMN "public"."erp_purchase_list"."purchase_list_number" IS '数量';
COMMENT ON COLUMN "public"."erp_purchase_list"."purchase_list_price" IS '总价(分)';
COMMENT ON COLUMN "public"."erp_purchase_list"."purchase_list_time" IS '时间';
COMMENT ON COLUMN "public"."erp_purchase_list"."purchase_list_remarks" IS '备注';
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
INSERT INTO "public"."erp_staff" VALUES (247834, 'test123', 'test123', 1, '2124781', 234236, '普通用户', 'f', '2021-06-17 21:10:47+08');
INSERT INTO "public"."erp_staff" VALUES (247836, 'union', 'union', 1, '14443', 12314, '普通用户', 'f', '2021-06-17 22:43:34+08');
INSERT INTO "public"."erp_staff" VALUES (247833, 'admin', 'admin', 2, '2189324', 213912, '管理员用户', 'f', '2021-06-17 21:10:03+08');
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
  "supplier_remarks" varchar(255) COLLATE "pg_catalog"."default"
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
INSERT INTO "public"."sys_menu" VALUES (100, '/goods', 1, '商品信息');
INSERT INTO "public"."sys_menu" VALUES (201, '/suppliera', 2, '供应商管理');
INSERT INTO "public"."sys_menu" VALUES (202, '/staffa', 2, '员工管理');
INSERT INTO "public"."sys_menu" VALUES (203, '/pda', 2, '采购主表管理');
INSERT INTO "public"."sys_menu" VALUES (204, '/pla', 2, '采购明细管理');
INSERT INTO "public"."sys_menu" VALUES (102, '/pl', 1, '采购明细');
INSERT INTO "public"."sys_menu" VALUES (101, '/pd', 1, '采购主表');
COMMIT;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."erp_goods_goods_id_seq"
OWNED BY "public"."erp_goods"."goods_id";
SELECT setval('"public"."erp_goods_goods_id_seq"', 2, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."erp_purchase_details_purchase_details_id_seq"
OWNED BY "public"."erp_purchase_details"."purchase_details_id";
SELECT setval('"public"."erp_purchase_details_purchase_details_id_seq"', 2, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."erp_purchase_list_purchase_list_id_seq"
OWNED BY "public"."erp_purchase_list"."purchase_list_id";
SELECT setval('"public"."erp_purchase_list_purchase_list_id_seq"', 2, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."erp_staff_staff_id_seq"
OWNED BY "public"."erp_staff"."staff_id";
SELECT setval('"public"."erp_staff_staff_id_seq"', 247837, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."erp_supplier_supplier_id_seq"
OWNED BY "public"."erp_supplier"."supplier_id";
SELECT setval('"public"."erp_supplier_supplier_id_seq"', 2, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_router_router_id_seq"
OWNED BY "public"."sys_menu"."menu_id";
SELECT setval('"public"."sys_router_router_id_seq"', 10, true);

-- ----------------------------
-- Primary Key structure for table erp_goods
-- ----------------------------
ALTER TABLE "public"."erp_goods" ADD CONSTRAINT "_copy_3" PRIMARY KEY ("goods_id");

-- ----------------------------
-- Primary Key structure for table erp_purchase_details
-- ----------------------------
ALTER TABLE "public"."erp_purchase_details" ADD CONSTRAINT "_copy_1" PRIMARY KEY ("purchase_details_id");

-- ----------------------------
-- Primary Key structure for table erp_purchase_list
-- ----------------------------
ALTER TABLE "public"."erp_purchase_list" ADD CONSTRAINT "_copy_2" PRIMARY KEY ("purchase_list_id");

-- ----------------------------
-- Uniques structure for table erp_staff
-- ----------------------------
ALTER TABLE "public"."erp_staff" ADD CONSTRAINT "uniq_name" UNIQUE ("staff_name");

-- ----------------------------
-- Primary Key structure for table erp_staff
-- ----------------------------
ALTER TABLE "public"."erp_staff" ADD CONSTRAINT "_copy_4" PRIMARY KEY ("staff_id");

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
-- Foreign Keys structure for table erp_purchase_list
-- ----------------------------
ALTER TABLE "public"."erp_purchase_list" ADD CONSTRAINT "fk_erp_purchase_list_erp_staff_1" FOREIGN KEY ("purchase_list_staff") REFERENCES "public"."erp_staff" ("staff_id") ON DELETE NO ACTION ON UPDATE NO ACTION;

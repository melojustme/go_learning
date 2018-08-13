/*
 Navicat Premium Data Transfer

 Source Server         : pgsql
 Source Server Type    : PostgreSQL
 Source Server Version : 90305
 Source Host           : 192.168.99.100:32769
 Source Catalog        : test
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 90305
 File Encoding         : 65001

 Date: 11/07/2018 13:50:26
*/


-- ----------------------------
-- Table structure for birds
-- ----------------------------
DROP TABLE IF EXISTS "public"."birds";
CREATE TABLE "public"."birds" (
  "id" int4 NOT NULL DEFAULT nextval('birds_id_seq'::regclass),
  "bird" varchar(256) COLLATE "pg_catalog"."default",
  "description" varchar(1024) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."birds" OWNER TO "postgres";

-- ----------------------------
-- Primary Key structure for table birds
-- ----------------------------
ALTER TABLE "public"."birds" ADD CONSTRAINT "birds_pkey" PRIMARY KEY ("id");

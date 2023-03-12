/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 12/03/2023 19:53:59
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
                                `id` int unsigned NOT NULL AUTO_INCREMENT,
                                `tag_id` int unsigned DEFAULT '0' COMMENT '标签ID',
                                `title` varchar(100) DEFAULT '' COMMENT '文章标题',
                                `desc` varchar(255) DEFAULT '' COMMENT '简述',
                                `content` text COMMENT '内容',
                                `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
                                `created_on` int unsigned DEFAULT '0' COMMENT '新建时间',
                                `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
                                `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
                                `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
                                `deleted_on` int unsigned DEFAULT '0',
                                `state` tinyint unsigned DEFAULT '1' COMMENT '删除时间',
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='文章管理';

-- ----------------------------
-- Records of blog_article
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth` (
                             `id` int unsigned NOT NULL AUTO_INCREMENT,
                             `username` varchar(50) DEFAULT '' COMMENT '账号',
                             `password` varchar(50) DEFAULT '' COMMENT '密码',
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of blog_auth
-- ----------------------------
BEGIN;
INSERT INTO `blog_auth` (`id`, `username`, `password`) VALUES (1, 'test', 'test123');
INSERT INTO `blog_auth` (`id`, `username`, `password`) VALUES (2, 'admin', 'admin');
COMMIT;

-- ----------------------------
-- Table structure for blog_comment
-- ----------------------------
DROP TABLE IF EXISTS `blog_comment`;
CREATE TABLE `blog_comment` (
                                `id` int unsigned NOT NULL AUTO_INCREMENT,
                                `username` varchar(100) NOT NULL,
                                `content` text NOT NULL,
                                `created_at` datetime NOT NULL,
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of blog_comment
-- ----------------------------
BEGIN;
INSERT INTO `blog_comment` (`id`, `username`, `content`, `created_at`) VALUES (1, 'q we', 'www', '2023-03-11 00:07:11');
INSERT INTO `blog_comment` (`id`, `username`, `content`, `created_at`) VALUES (2, 'q we', 'e e', '2023-03-11 00:07:37');
INSERT INTO `blog_comment` (`id`, `username`, `content`, `created_at`) VALUES (3, '123', '123', '2023-03-11 00:26:41');
INSERT INTO `blog_comment` (`id`, `username`, `content`, `created_at`) VALUES (11, 'xxx', '<a>xxxx</a>', '2023-03-11 10:02:02');
INSERT INTO `blog_comment` (`id`, `username`, `content`, `created_at`) VALUES (12, 'ss', '</p><a>s<a/>', '2023-03-11 10:02:44');
INSERT INTO `blog_comment` (`id`, `username`, `content`, `created_at`) VALUES (13, '222', '</p><a>s<a/><p>', '2023-03-11 10:03:21');
INSERT INTO `blog_comment` (`id`, `username`, `content`, `created_at`) VALUES (14, '<p></p><a>s<a/><p></p>', '<p></p><a>s</a><p></p>', '2023-03-11 10:03:48');
INSERT INTO `blog_comment` (`id`, `username`, `content`, `created_at`) VALUES (15, '<a href=\"https://www.example.com\">Link to example.com</a>', '<a href=\"https://www.example.com\">Link to example.com</a>', '2023-03-11 10:05:04');
COMMIT;

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
                            `id` int unsigned NOT NULL AUTO_INCREMENT,
                            `name` varchar(100) DEFAULT '' COMMENT '标签名称',
                            `created_on` int unsigned DEFAULT '0' COMMENT '创建时间',
                            `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
                            `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
                            `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
                            `deleted_on` int unsigned DEFAULT '0' COMMENT '删除时间',
                            `state` tinyint unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='文章标签管理';

-- ----------------------------
-- Records of blog_tag
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;

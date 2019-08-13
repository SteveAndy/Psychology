/*
Navicat MySQL Data Transfer

Source Server         : 心理咨询
Source Server Version : 50557
Source Host           : 122.114.151.4:3306
Source Database       : nxxlzx

Target Server Type    : MYSQL
Target Server Version : 50557
File Encoding         : 65001

Date: 2019-08-13 16:41:47
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) NOT NULL,
  `portrait` varchar(255) NOT NULL,
  `telephone` varchar(11) NOT NULL COMMENT '电话',
  `openid` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for banner
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'banner的ID',
  `img_urls` varchar(255) NOT NULL COMMENT 'banner的URL',
  `url` varchar(255) NOT NULL COMMENT 'banner对应的超链接',
  `type` int(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for community
-- ----------------------------
DROP TABLE IF EXISTS `community`;
CREATE TABLE `community` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userId` int(11) NOT NULL,
  `classifyId` int(11) NOT NULL,
  `content` varchar(255) NOT NULL,
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `see` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `id_2` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for community_class
-- ----------------------------
DROP TABLE IF EXISTS `community_class`;
CREATE TABLE `community_class` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `class_name` varchar(50) NOT NULL,
  `color` varchar(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for community_class_sub
-- ----------------------------
DROP TABLE IF EXISTS `community_class_sub`;
CREATE TABLE `community_class_sub` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `father` int(11) NOT NULL,
  `title` varchar(50) NOT NULL,
  `icon` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for community_reply
-- ----------------------------
DROP TABLE IF EXISTS `community_reply`;
CREATE TABLE `community_reply` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cid` int(11) NOT NULL COMMENT '所属的帖子的id',
  `user_id` int(11) NOT NULL,
  `content` varchar(255) NOT NULL,
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for expert
-- ----------------------------
DROP TABLE IF EXISTS `expert`;
CREATE TABLE `expert` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '专家的ID',
  `openid` varchar(255) NOT NULL DEFAULT '',
  `photo` varchar(255) NOT NULL,
  `icon` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `phone_num` varchar(11) NOT NULL,
  `classify_id` int(11) NOT NULL,
  `address` varchar(255) NOT NULL,
  `info` varchar(255) NOT NULL,
  `gender` int(1) NOT NULL DEFAULT '2' COMMENT '1为男生、2为女生',
  `workAge` int(11) NOT NULL COMMENT '从业年龄',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `id_2` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=150 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for expert_auth
-- ----------------------------
DROP TABLE IF EXISTS `expert_auth`;
CREATE TABLE `expert_auth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `address` varchar(255) NOT NULL COMMENT '地址',
  `classify_id` int(11) NOT NULL COMMENT '专家所属分类ID',
  `photo` varchar(255) NOT NULL COMMENT '专家本人头像url',
  `name` varchar(255) NOT NULL COMMENT '专家姓名',
  `info` varchar(255) NOT NULL COMMENT '专家简介',
  `age` varchar(255) NOT NULL COMMENT '从业年龄',
  `phone_num` varchar(255) NOT NULL COMMENT '手机号码',
  `idcard_f` varchar(255) NOT NULL COMMENT '身份证正面url',
  `idcard_b` varchar(255) NOT NULL COMMENT '身份证反面url',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '提交时间',
  `status` int(2) NOT NULL COMMENT '是否通过认证',
  `reject` varchar(255) NOT NULL COMMENT '驳回原因',
  `openid` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for expert_class
-- ----------------------------
DROP TABLE IF EXISTS `expert_class`;
CREATE TABLE `expert_class` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `class_name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for info
-- ----------------------------
DROP TABLE IF EXISTS `info`;
CREATE TABLE `info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `classifyId` int(11) NOT NULL,
  `title` varchar(255) NOT NULL COMMENT '标题',
  `content` mediumtext NOT NULL COMMENT '内容',
  `icon` varchar(255) NOT NULL COMMENT '缩略图',
  `uid` int(11) NOT NULL,
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `author_type` int(2) NOT NULL COMMENT '1专家，2管理员',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for info_class
-- ----------------------------
DROP TABLE IF EXISTS `info_class`;
CREATE TABLE `info_class` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `class_name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `img_urls` varchar(255) NOT NULL,
  `title` varchar(255) NOT NULL,
  `type` int(2) NOT NULL,
  `url` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for opinion
-- ----------------------------
DROP TABLE IF EXISTS `opinion`;
CREATE TABLE `opinion` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT '提交反馈的用户',
  `content` varchar(255) NOT NULL COMMENT '反馈内容',
  `telephone` varchar(255) NOT NULL COMMENT '用户联系方式',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `openid` varchar(255) NOT NULL COMMENT '用户的openid',
  `user_name` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '用户名',
  `year` varchar(255) DEFAULT NULL COMMENT '年龄',
  `telephone` varchar(255) DEFAULT NULL COMMENT '联系电话',
  `portrait` varchar(255) DEFAULT NULL COMMENT '头像的URL',
  `address` varchar(255) DEFAULT NULL COMMENT '用户的地址',
  `gender` varchar(255) DEFAULT NULL COMMENT '性别：1代表男，2代表女，0未知',
  `is_super` varchar(1) NOT NULL DEFAULT '0' COMMENT '是否为特殊用户',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=695 DEFAULT CHARSET=utf8;

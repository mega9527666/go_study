create database if not exists db_account Character Set utf8mb4;
use db_account;


SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `t_accounts`;
CREATE TABLE `t_accounts` (
   `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `account` varchar(30) NOT NULL,
  `pass` varchar(32) NOT NULL,
  `token` varchar(255) DEFAULT NULL,
  `account_type` int(11) DEFAULT '1' COMMENT '账号类型',
  `status` int(11) DEFAULT '1' COMMENT '账号状态',
  `ip` varchar(255) DEFAULT NULL,
  `nick_name` varchar(64) DEFAULT NULL COMMENT '呢称',
  `channel` int(11) DEFAULT NULL COMMENT '渠道',
  `os` varchar(30) CHARACTER SET utf8 DEFAULT NULL COMMENT '系统',
  `phone_type` varchar(30) CHARACTER SET utf8 DEFAULT NULL COMMENT '手机型号',
  `bundle_name` varchar(100) CHARACTER SET utf8 DEFAULT NULL COMMENT '包名',
  `system_version` varchar(50) CHARACTER SET utf8 DEFAULT NULL COMMENT '系统版本号',
  `create_time` BIGINT DEFAULT NULL COMMENT '创建时间',
  `last_login_time` BIGINT DEFAULT NULL COMMENT '最后登录时间',
  `phone` varchar(30) DEFAULT NULL COMMENT '手机号码',
  `sex` int(11) DEFAULT NULL,
  `headimgurl` varchar(255) DEFAULT NULL,
   PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_account` (`account`) USING BTREE,
  KEY `idx_last_login_time` (`last_login_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


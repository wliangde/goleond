
CREATE DATABASE IF NOT EXISTS star DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
use star;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for gm_uc_auth
-- ----------------------------
DROP TABLE IF EXISTS `gm_uc_auth`;
CREATE TABLE `gm_uc_auth`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `pid` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级ID，0为顶级',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '权限名称',
  `auth_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT 'URL地址',
  `sort` int(1) UNSIGNED NOT NULL DEFAULT 999 COMMENT '排序，越小越前',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态，1-正常 0禁用',
  `site_bar` tinyint(4) NOT NULL DEFAULT 1 COMMENT '是否在侧边栏显示',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6015 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限因子' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gm_uc_auth
-- ----------------------------
INSERT INTO `gm_uc_auth` VALUES (1, 0, '所有权限', '/', 0, '', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (2, 1, 'GM操作', '/', 1, 'fa-cubes', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (3, 1, '操作记录', '/', 2, 'fa-cubes', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (4, 1, '权限管理', '/', 998, 'fa-id-card', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (5, 1, '个人中心', '/', 999, 'fa-user-circle-o', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (6, 1, '服务器管理', '/', 3, 'fa-cogs', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (2001, 2, '玩家查询', '/gm/searchuser', 1, 'fa-edit', 0, 1);
INSERT INTO `gm_uc_auth` VALUES (2002, 2, 'GM-加资源', '/gm/addresource', 0, '', 1, 0);
INSERT INTO `gm_uc_auth` VALUES (2003, 2, 'GM-任意GM', '/gm/dogm', 0, '', 1, 0);
INSERT INTO `gm_uc_auth` VALUES (3001, 3, '日志记录', '/log/list', 999, 'fa-cubes', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (4001, 4, '用户管理', '/backenduser/list', 1, '', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (4002, 4, '用户组管理', '/group/list', 2, '', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (4003, 4, '权限管理', '/auth/list', 3, '', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (5001, 5, '修改信息', '/backenduser/pwd', 1, 'fa-edit', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (6001, 6, 'Game服', '/servergame/list', 2, 'fa-server', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (6002, 6, 'Cross服', '/servercross/list', 3, 'fa-cog', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (6003, 6, 'Battle服', '/serverbattle/list', 4, 'fa-cog', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (6004, 6, 'Gate服', '/servergate/list', 5, 'fa-cog', 0, 1);
INSERT INTO `gm_uc_auth` VALUES (6005, 6, 'launcher配置', '/launcher/list', 1, 'fa-cog', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (6007, 2, '邮件操作', '/gm/mail', 2, 'fa-edit', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (6008, 2, '客户端包版本', '/gm/versionlist', 7, 'fa-edit', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (6009, 2, '服务器版本', '/gm/serverversion', 8, 'fa-edit', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (6010, 2, '更新配置', '/gm/updatecfg', 5, 'fa-edit', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (6011, 2, '更新配置_后台', '/gm/ajaxupdatecfg', 0, 'fa-edit', 1, 0);
INSERT INTO `gm_uc_auth` VALUES (6012, 2, '一键养成', '/gm/onekeydev', 3, 'fa-edit', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (6013, 2, '一键导号', '/gm/onekeydump', 4, 'fa-edit', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (6014, 2, '账号封禁', '/gm/forbiduser', 6, 'fa-edit', 1, 1);
INSERT INTO `gm_uc_auth` VALUES (6015, 6, '版本服', '/versionsvr/list', 5, 'fa-cog', 1, 1);

-- ----------------------------
-- Table structure for gm_uc_backend_user
-- ----------------------------
DROP TABLE IF EXISTS `gm_uc_backend_user`;
CREATE TABLE `gm_uc_backend_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录名',
  `real_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `pwd` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `group_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限组',
  `is_super` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是超级管理员',
  `salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码盐',
  `last_login` int(11) NOT NULL DEFAULT 0 COMMENT '最后登录时间',
  `last_ip` char(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态，1-正常 0禁用',
  `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '后台用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gm_uc_backend_user
-- ----------------------------
INSERT INTO `gm_uc_backend_user` VALUES (1, 'star', '超级管理员', '0a477a7d3e186c1a081931d4703f49d1', 1, 1, '2VBR', 1550644531, '30.97.196.11', 1, 1523934413, 1524021420);

-- ----------------------------
-- Table structure for gm_uc_group
-- ----------------------------
DROP TABLE IF EXISTS `gm_uc_group`;
CREATE TABLE `gm_uc_group`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组名',
  `desc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态，1-正常 0禁用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理组表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gm_uc_group
-- ----------------------------
INSERT INTO `gm_uc_group` VALUES (1, '超级管理员', '所有权限', 1);

-- ----------------------------
-- Table structure for gm_uc_group_auth
-- ----------------------------
DROP TABLE IF EXISTS `gm_uc_group_auth`;
CREATE TABLE `gm_uc_group_auth`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `group_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '组ID',
  `auth_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `i_group_id_auth_id`(`group_id`, `auth_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '组权限' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

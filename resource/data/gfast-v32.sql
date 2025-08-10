/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : gfast-v32-github

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 19/01/2023 11:13:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `ptype` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', '1', '27', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '28', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '29', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '30', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '1', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '2', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '3', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '4', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '11', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '10', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '12', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '13', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '14', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '15', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '19', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '20', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '21', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '22', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '23', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '24', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '25', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '26', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_3', '1', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_3', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '31', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '32', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '34', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_31', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '35', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '33', 'All', '', '', '');

-- ----------------------------
-- Table structure for sys_auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_auth_rule`;
CREATE TABLE `sys_auth_rule`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'ParentID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Rule name',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Rule title',
  `icon` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Icon',
  `condition` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Condition',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Remark',
  `menu_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Type 0Directory 1Menu 2Button',
  `weigh` int(10) NOT NULL DEFAULT 0 COMMENT 'Weight',
  `is_hide` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Display status',
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Route path',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Component path',
  `is_link` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Is external link 1Yes 0No',
  `module_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Module type',
  `model_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'ModelID',
  `is_iframe` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Is embeddediframe',
  `is_cached` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Is cached',
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Route redirect address',
  `is_affix` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Is affixed',
  `link_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Link URL',
  `created_at` datetime NULL DEFAULT NULL COMMENT 'Creation date',
  `updated_at` datetime NULL DEFAULT NULL COMMENT 'Modification date',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `weigh`(`weigh`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 59 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Menu node table' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_auth_rule
-- ----------------------------
INSERT INTO `sys_auth_rule` VALUES (1, 0, 'api/v1/system/auth', 'Permission Management', 'ele-Stamp', '', '', 0, 30, 0, '/system/auth', 'layout/routerView/parent', 0, '', 0, 0, 1, '0', 0, '', '2022-03-24 15:03:37', '2022-04-14 16:29:19');
INSERT INTO `sys_auth_rule` VALUES (2, 1, 'api/v1/system/auth/menuList', 'Menu Management', 'ele-Calendar', '', '', 1, 0, 0, '/system/auth/menuList', 'system/menu/index', 0, '', 0, 0, 1, '', 0, '', '2022-03-24 17:24:13', '2022-03-29 10:54:49');
INSERT INTO `sys_auth_rule` VALUES (3, 2, 'api/v1/system/menu/add', 'Add Menu', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-03-29 16:48:43', '2022-03-29 17:05:19');
INSERT INTO `sys_auth_rule` VALUES (4, 2, 'api/v1/system/menu/update', 'Edit Menu', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-03-29 17:04:25', '2022-03-29 18:11:36');
INSERT INTO `sys_auth_rule` VALUES (10, 1, 'api/v1/system/role/list', 'Role Management', 'iconfont icon-juxingkaobei', '', '', 1, 0, 0, '/system/auth/roleList', 'system/role/index', 0, '', 0, 0, 1, '', 0, '', '2022-03-29 18:15:03', '2022-03-30 10:25:34');
INSERT INTO `sys_auth_rule` VALUES (11, 2, 'api/v1/system/menu/delete', 'Delete Menu', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:49:10', '2022-04-06 14:49:17');
INSERT INTO `sys_auth_rule` VALUES (12, 10, 'api/v1/system/role/add', 'Add Role', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:49:46', '2022-04-06 14:49:46');
INSERT INTO `sys_auth_rule` VALUES (13, 10, '/api/v1/system/role/edit', 'Edit Role', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:50:08', '2022-04-06 14:50:08');
INSERT INTO `sys_auth_rule` VALUES (14, 10, '/api/v1/system/role/delete', 'Delete Role', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:50:22', '2022-04-06 14:50:22');
INSERT INTO `sys_auth_rule` VALUES (15, 1, 'api/v1/system/dept/list', 'Department Management', 'iconfont icon-siweidaotu', '', '', 1, 0, 0, '/system/auth/deptList', 'system/dept/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:52:23', '2022-04-07 22:59:20');
INSERT INTO `sys_auth_rule` VALUES (16, 17, 'aliyun', 'Alibaba Cloud-iframe', 'iconfont icon-diannao1', '', '', 1, 0, 0, '/demo/outLink/aliyun', 'layout/routerView/iframes', 1, '', 0, 1, 1, '', 0, 'https://www.aliyun.com/daily-act/ecs/activity_selection?spm=5176.8789780.J_3965641470.5.568845b58KHj51', '2022-04-06 17:26:29', '2022-04-07 15:27:17');
INSERT INTO `sys_auth_rule` VALUES (17, 0, 'outLink', 'External Link Test', 'iconfont icon-zhongduancanshu', '', '', 0, 20, 0, '/demo/outLink', 'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 15:20:51', '2022-04-14 16:29:07');
INSERT INTO `sys_auth_rule` VALUES (18, 17, 'tenyun', 'Tencent Cloud-External Link', 'iconfont icon-shouye_dongtaihui', '', '', 1, 0, 0, '/demo/outLink/tenyun', 'layout/routerView/link', 1, '', 0, 0, 1, '', 0, 'https://cloud.tencent.com/act/new?cps_key=20b1c3842f74986b2894e2c5fcde7ea2&fromSource=gwzcw.3775555.3775555.3775555&utm_id=gwzcw.3775555.3775555.3775555&utm_medium=cpc', '2022-04-07 15:23:52', '2022-04-07 15:27:25');
INSERT INTO `sys_auth_rule` VALUES (19, 15, 'api/v1/system/dept/add', 'Add Department', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:56:39', '2022-04-07 22:56:39');
INSERT INTO `sys_auth_rule` VALUES (20, 15, 'api/v1/system/dept/edit', 'Edit Department', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:57:00', '2022-04-07 22:57:00');
INSERT INTO `sys_auth_rule` VALUES (21, 15, 'api/v1/system/dept/delete', 'Delete Department', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:57:30', '2022-04-07 22:57:30');
INSERT INTO `sys_auth_rule` VALUES (22, 1, 'api/v1/system/post/list', 'Position Management', 'iconfont icon-neiqianshujuchucun', '', '', 1, 0, 0, '/system/auth/postList', 'system/post/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:58:46', '2022-04-09 14:26:15');
INSERT INTO `sys_auth_rule` VALUES (23, 22, 'api/v1/system/post/add', 'Add Position', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:14:49', '2022-04-09 14:14:49');
INSERT INTO `sys_auth_rule` VALUES (24, 22, 'api/v1/system/post/edit', 'Edit Position', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:15:25', '2022-04-09 14:15:25');
INSERT INTO `sys_auth_rule` VALUES (25, 22, 'api/v1/system/post/delete', 'Delete Position', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:15:47', '2022-04-09 14:15:47');
INSERT INTO `sys_auth_rule` VALUES (26, 1, 'api/v1/system/user/list', 'User Management', 'ele-User', '', '', 1, 0, 0, '/system/auth/user/list', 'system/user/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:19:10', '2022-04-09 14:19:58');
INSERT INTO `sys_auth_rule` VALUES (27, 0, 'api/v1/system/dict', 'System Configuration', 'iconfont icon-shuxingtu', '', '', 0, 40, 0, '/system/dict', 'layout/routerView/parent', 0, '', 0, 0, 1, '654', 0, '', '2022-04-14 16:28:51', '2022-04-18 14:40:56');
INSERT INTO `sys_auth_rule` VALUES (28, 27, 'api/v1/system/dict/type/list', 'Dictionary Management', 'iconfont icon-crew_feature', '', '', 1, 0, 0, '/system/dict/type/list', 'system/dict/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-14 16:32:10', '2022-04-16 17:02:50');
INSERT INTO `sys_auth_rule` VALUES (29, 27, 'api/v1/system/dict/dataList', 'Dictionary Data Management', 'iconfont icon-putong', '', '', 1, 0, 1, '/system/dict/data/list/:dictType', 'system/dict/dataList', 0, '', 0, 0, 1, '', 0, '', '2022-04-18 12:04:17', '2022-04-18 14:58:43');
INSERT INTO `sys_auth_rule` VALUES (30, 27, 'api/v1/system/config/list', 'Parameter Management', 'ele-Cherry', '', '', 1, 0, 0, '/system/config/list', 'system/config/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-18 21:05:20', '2022-04-18 21:13:19');
INSERT INTO `sys_auth_rule` VALUES (31, 0, 'api/v1/system/monitor', 'System Monitoring', 'iconfont icon-xuanzeqi', '', '', 0, 30, 0, '/system/monitor', 'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-04-19 10:40:19', '2022-04-19 10:44:38');
INSERT INTO `sys_auth_rule` VALUES (32, 31, 'api/v1/system/monitor/server', 'Server Monitoring', 'iconfont icon-shuju', '', '', 1, 0, 0, '/system/monitor/server', 'system/monitor/server/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-19 10:43:32', '2022-04-19 10:44:47');
INSERT INTO `sys_auth_rule` VALUES (33, 35, 'api/swagger', 'apiDocumentation', 'iconfont icon--chaifenlie', '', '', 1, 0, 0, '/system/swagger', 'layout/routerView/iframes', 1, '', 0, 1, 1, '', 0, 'http://localhost:8808/swagger', '2022-04-21 09:23:43', '2022-11-29 17:10:35');
INSERT INTO `sys_auth_rule` VALUES (34, 31, 'api/v1/system/loginLog/list', 'Login Logs', 'ele-Finished', '', '', 1, 0, 0, '/system/monitor/loginLog', 'system/monitor/loginLog/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-28 09:59:47', '2022-04-28 09:59:47');
INSERT INTO `sys_auth_rule` VALUES (35, 0, 'api/v1/system/tools', 'System Tools', 'iconfont icon-zujian', '', '', 0, 25, 0, '/system/tools', 'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-10-26 09:29:08', '2022-10-26 10:11:25');
INSERT INTO `sys_auth_rule` VALUES (38, 31, 'api/v1/system/operLog/list', 'Operation Logs', 'iconfont icon-bolangnengshiyanchang', '', '', 1, 0, 0, '/system/monitor/operLog', 'system/monitor/operLog/index', 0, '', 0, 0, 1, '', 0, '', '2022-12-23 16:19:05', '2022-12-23 16:21:50');
INSERT INTO `sys_auth_rule` VALUES (39, 31, 'api/v1/system/online/list', 'Online Users', 'iconfont icon-skin', '', '', 1, 0, 0, '/system/monitor/userOnlie', 'system/monitor/userOnline/index', 0, '', 0, 0, 1, '', 0, '', '2023-01-11 15:48:06', '2023-01-11 17:02:39');

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `config_id` int(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Parameter ID',
  `config_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Parameter name',
  `config_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Parameter key',
  `config_value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Parameter value',
  `config_type` tinyint(1) NULL DEFAULT 0 COMMENT 'System built-in（1=Yes 0=No）',
  `create_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT 'Creator',
  `update_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT 'Updater',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Remark',
  `created_at` datetime NULL DEFAULT NULL COMMENT 'Creation time',
  `updated_at` datetime NULL DEFAULT NULL COMMENT 'Update time',
  PRIMARY KEY (`config_id`) USING BTREE,
  UNIQUE INDEX `uni_config_key`(`config_key`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, 'File Upload-File Size', 'sys.uploadFile.fileSize', '50M', 1, 31, 31, 'File upload size limit', NULL, '2021-07-06 14:57:35');
INSERT INTO `sys_config` VALUES (2, 'File Upload-File Types', 'sys.uploadFile.fileType', 'doc,docx,zip,xls,xlsx,rar,jpg,jpeg,gif,npm,png,mp4', 1, 31, 31, 'Allowed file extensions for upload', NULL, '2022-12-16 09:52:45');
INSERT INTO `sys_config` VALUES (3, 'Image Upload-Image Types', 'sys.uploadFile.imageType', 'jpg,jpeg,gif,npm,png', 1, 31, 0, 'Allowed image extensions for upload', NULL, NULL);
INSERT INTO `sys_config` VALUES (4, 'Image Upload-Image Size', 'sys.uploadFile.imageSize', '50M', 1, 31, 31, 'Image upload size limit', NULL, NULL);
INSERT INTO `sys_config` VALUES (11, 'Static Resources', 'static.resource', '/', 1, 2, 0, '', NULL, NULL);

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'Departmentid',
  `parent_id` bigint(20) NULL DEFAULT 0 COMMENT 'Parent departmentid',
  `ancestors` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Ancestor hierarchy',
  `dept_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Department name',
  `order_num` int(4) NULL DEFAULT 0 COMMENT 'Display order',
  `leader` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Manager',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Phone number',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Email',
  `status` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT 'Department status（0Normal 1Disabled）',
  `created_by` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT 'Created by',
  `updated_by` bigint(20) NULL DEFAULT NULL COMMENT 'Updated by',
  `created_at` datetime NULL DEFAULT NULL COMMENT 'Creation time',
  `updated_at` datetime NULL DEFAULT NULL COMMENT 'Update time',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT 'Deletion time',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 204 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Departments table' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (100, 0, '0', 'Qixun Technology', 0, NULL, '15888888888', 'ry@qq.com', 1, 0, 31, '2021-07-13 15:56:52', '2021-07-13 15:57:05', NULL);
INSERT INTO `sys_dept` VALUES (101, 100, '0,100', 'Shenzhen Headquarters', 1, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (102, 100, '0,100', 'Changsha Branch', 2, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (103, 101, '0,100,101', 'R&D Department', 1, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (104, 101, '0,100,101', 'Marketing Department', 2, NULL, '15888888888', 'ry@qq.com', 1, 0, 31, '2021-07-13 15:56:52', '2021-11-04 09:16:38', NULL);
INSERT INTO `sys_dept` VALUES (105, 101, '0,100,101', 'Testing Department', 3, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (106, 101, '0,100,101', 'Finance Department', 4, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (107, 101, '0,100,101', 'Operations Department', 5, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (108, 102, '0,100,102', 'Marketing Department', 1, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (109, 102, '0,100,102', 'Finance Department', 2, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (200, 100, '', 'Big Data', 1, '', '18888888888', 'liou@qq.com', 0, 0, 31, '2021-07-13 15:56:52', '2022-09-16 16:46:57', NULL);
INSERT INTO `sys_dept` VALUES (201, 100, '', 'Development', 1, NULL, '18888888888', 'li@qq.com', 0, 31, NULL, '2021-07-13 15:56:52', '2022-04-07 22:35:21', NULL);
INSERT INTO `sys_dept` VALUES (202, 108, '', 'Field Operations', 1, NULL, '18888888888', 'aa@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (203, 108, '', 'Administration', 0, '', '18888888888', 'aa@qq.com', 1, 0, 31, '2021-07-13 15:56:52', '2022-09-16 16:46:47', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'Dictionary code',
  `dict_sort` int(4) NULL DEFAULT 0 COMMENT 'Sort order',
  `dict_label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Dictionary label',
  `dict_value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Dictionary key value',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Dictionary type',
  `css_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Style attributes（additional style extensions）',
  `list_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Table display style',
  `is_default` tinyint(1) NULL DEFAULT 0 COMMENT 'Is default（1Yes 0No）',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT 'Status（0Normal 1Disabled）',
  `create_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT 'Creator',
  `update_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT 'Updater',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Remarks',
  `created_at` datetime NULL DEFAULT NULL COMMENT 'Creation time',
  `updated_at` datetime NULL DEFAULT NULL COMMENT 'Update time',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 106 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Dictionary data table' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 0, 'Male', '1', 'sys_user_sex', '', '', 0, 1, 31, 2, 'Remarks', '2022-04-18 16:46:22', NULL);
INSERT INTO `sys_dict_data` VALUES (2, 0, 'Female', '2', 'sys_user_sex', '', '', 0, 1, 31, 31, 'Remarks', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (3, 0, 'Private', '0', 'sys_user_sex', '', '', 1, 1, 31, 31, 'Remarks', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (24, 0, 'Channel page', '1', 'cms_category_type', '', '', 0, 1, 31, 31, 'As a channel page，cannot publish articles，but can add subcategories', NULL, '2021-07-21 10:54:22');
INSERT INTO `sys_dict_data` VALUES (25, 0, 'Publish column', '2', 'cms_category_type', '', '', 0, 1, 31, 31, 'As a publishing column，can add articles', NULL, '2021-07-21 10:54:22');
INSERT INTO `sys_dict_data` VALUES (26, 0, 'Redirect column', '3', 'cms_category_type', '', '', 0, 1, 31, 31, 'Does not directly publish content，used for page redirection', NULL, '2021-07-21 10:54:22');
INSERT INTO `sys_dict_data` VALUES (27, 0, 'Single page column', '4', 'cms_category_type', '', '', 0, 1, 31, 31, 'Single page mode，category directly displays as article', NULL, '2021-07-21 10:54:22');
INSERT INTO `sys_dict_data` VALUES (28, 0, 'Normal', '0', 'sys_job_status', '', 'default', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (29, 0, 'Paused', '1', 'sys_job_status', '', 'default', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (30, 0, 'Default', 'DEFAULT', 'sys_job_group', '', 'default', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (31, 0, 'System', 'SYSTEM', 'sys_job_group', '', 'default', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (32, 0, 'Success', '1', 'admin_login_status', '', 'default', 0, 1, 31, 31, '', NULL, '2022-09-16 15:26:01');
INSERT INTO `sys_dict_data` VALUES (33, 0, 'Failed', '0', 'admin_login_status', '', 'default', 0, 1, 31, 0, '', NULL, '2022-09-16 15:26:01');
INSERT INTO `sys_dict_data` VALUES (34, 0, 'Success', '1', 'sys_oper_log_status', '', 'default', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (35, 0, 'Failed', '0', 'sys_oper_log_status', '', 'default', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (36, 0, 'Repeat execution', '1', 'sys_job_policy', '', 'default', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (37, 0, 'Execute once', '2', 'sys_job_policy', '', 'default', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (38, 0, 'Show', '0', 'sys_show_hide', NULL, 'default', 1, 1, 31, 0, NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (39, 0, 'Hide', '1', 'sys_show_hide', NULL, 'default', 0, 1, 31, 0, NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (40, 0, 'Normal', '1', 'sys_normal_disable', '', 'default', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (41, 0, 'Disabled', '0', 'sys_normal_disable', '', 'default', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (49, 0, 'Yes', '1', 'sys_yes_no', '', '', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (50, 0, 'No', '0', 'sys_yes_no', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (51, 0, 'Published', '1', 'cms_article_pub_type', '', '', 1, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (54, 0, 'Unpublished', '0', 'cms_article_pub_type', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (55, 0, 'Pinned', '1', 'cms_article_attr', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (56, 0, 'Recommended', '2', 'cms_article_attr', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (57, 0, 'Regular article', '0', 'cms_article_type', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (58, 0, 'Redirect link', '1', 'cms_article_type', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (59, 0, 'cmsmodel', '6', 'cms_cate_models', '', '', 0, 1, 1, 1, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (61, 0, 'Gov work objectives', '1', 'gov_cate_models', '', '', 0, 1, 2, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (62, 0, 'System backend', 'sys_admin', 'menu_module_type', '', '', 1, 1, 2, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (63, 0, 'Government work', 'gov_work', 'menu_module_type', '', '', 0, 1, 2, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (64, 0, 'Slideshow', '3', 'cms_article_attr', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (65, 0, '[work]Test business table', 'wf_news', 'flow_type', '', '', 0, 1, 2, 2, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (66, 0, 'Rollback modification', '-1', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (67, 0, 'Saving', '0', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (68, 0, 'In progress', '1', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (69, 0, 'Approved', '2', 'flow_status', '', '', 0, 1, 31, 2, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (70, 2, 'Publish column', '2', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (71, 3, 'Redirect column', '3', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (72, 4, 'Single page column', '4', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (73, 2, 'Pinned', '1', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (74, 3, 'Slideshow', '2', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (75, 4, 'Recommended', '3', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (76, 1, 'General', '0', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (77, 1, 'Channel page', '1', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (78, 0, 'Normal', '0', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2021-07-20 08:55:20');
INSERT INTO `sys_dict_data` VALUES (79, 0, 'Urgent', '1', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2021-07-20 08:55:20');
INSERT INTO `sys_dict_data` VALUES (80, 0, 'Emergency', '2', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2021-07-20 08:55:20');
INSERT INTO `sys_dict_data` VALUES (81, 0, 'Critical', '3', 'flow_level', '', '', 0, 1, 31, 31, '', NULL, '2021-07-20 08:55:25');
INSERT INTO `sys_dict_data` VALUES (82, 0, 'Channel page', '1', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (83, 0, 'Publish column', '2', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (84, 0, 'Redirect column', '3', 'sys_blog_type', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (85, 0, 'Single page column', '4', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (87, 0, '[cms]Article table', 'cms_news', 'flow_type', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (91, 0, 'Test it out', '666', 'cms_article_type', '', '', 0, 1, 31, 0, '', '2021-08-03 17:04:12', '2021-08-03 17:04:12');
INSERT INTO `sys_dict_data` VALUES (92, 0, 'Cache test222', '33333', 'cms_article_type', '', '', 0, 1, 31, 31, '', '2021-08-03 17:16:45', '2021-08-03 17:19:41');
INSERT INTO `sys_dict_data` VALUES (93, 0, 'Cache test222', '11111', 'cms_article_type', '', '', 0, 1, 31, 31, '', '2021-08-03 17:26:14', '2021-08-03 17:26:26');
INSERT INTO `sys_dict_data` VALUES (94, 0, '10% off', '10', 'plugin_store_discount', '', '', 0, 1, 31, 0, '', '2021-08-14 11:59:38', '2021-08-14 11:59:38');
INSERT INTO `sys_dict_data` VALUES (95, 0, '50% off', '50', 'plugin_store_discount', '', '', 0, 1, 31, 0, '', '2021-08-14 11:59:49', '2021-08-14 11:59:49');
INSERT INTO `sys_dict_data` VALUES (96, 0, '80% off', '80', 'plugin_store_discount', '', '', 0, 1, 31, 0, '', '2021-08-14 12:00:00', '2021-08-14 12:00:00');
INSERT INTO `sys_dict_data` VALUES (97, 0, '90% off', '90', 'plugin_store_discount', '', '', 0, 1, 31, 0, '', '2021-08-14 12:00:07', '2021-08-14 12:00:07');
INSERT INTO `sys_dict_data` VALUES (98, 0, 'No discount', '100', 'plugin_store_discount', '', '', 0, 1, 31, 0, '', '2021-08-14 12:00:16', '2021-08-14 12:00:16');
INSERT INTO `sys_dict_data` VALUES (99, 0, 'Do not display', 'none', 'cms_nav_position', '', '', 1, 1, 22, 0, '', '2021-08-31 15:37:35', '2021-08-31 15:37:35');
INSERT INTO `sys_dict_data` VALUES (100, 0, 'Top navigation', 'top', 'cms_nav_position', '', '', 0, 1, 22, 0, '', '2021-08-31 15:37:57', '2021-08-31 15:37:57');
INSERT INTO `sys_dict_data` VALUES (101, 0, 'Bottom navigation', 'bottom', 'cms_nav_position', '', '', 0, 1, 22, 0, '', '2021-08-31 15:38:08', '2021-08-31 15:38:08');
INSERT INTO `sys_dict_data` VALUES (102, 0, 'Read', 'GET', 'sys_oper_log_type', '', '', 0, 1, 31, 31, '', '2022-12-21 11:59:10', '2022-12-23 19:03:02');
INSERT INTO `sys_dict_data` VALUES (103, 0, 'Add', 'POST', 'sys_oper_log_type', '', '', 0, 1, 31, 31, '', '2022-12-21 11:59:22', '2022-12-23 19:03:10');
INSERT INTO `sys_dict_data` VALUES (104, 0, 'Modify', 'PUT', 'sys_oper_log_type', '', '', 0, 1, 31, 31, '', '2022-12-21 11:59:32', '2022-12-23 19:03:19');
INSERT INTO `sys_dict_data` VALUES (105, 0, 'Delete', 'DELETE', 'sys_oper_log_type', '', '', 0, 1, 31, 31, '', '2022-12-21 11:59:44', '2022-12-23 19:03:27');

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
  `dict_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Primary key of dictionary',
  `dict_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Dictionary name',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Dictionary type',
  `status` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT 'Status（0Normal 1Disabled）',
  `create_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT 'Creator',
  `update_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT 'Updater',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Remark',
  `created_at` datetime NULL DEFAULT NULL COMMENT 'Creation date',
  `updated_at` datetime NULL DEFAULT NULL COMMENT 'Modification date',
  PRIMARY KEY (`dict_id`) USING BTREE,
  UNIQUE INDEX `dict_type`(`dict_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 51 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Dictionary type table' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (1, 'User gender', 'sys_user_sex', 1, 31, 1, 'For selecting user gender', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (2, 'Category type', 'cms_category_type', 1, 31, 3, 'Article category type', NULL, '2021-07-21 10:54:22');
INSERT INTO `sys_dict_type` VALUES (3, 'Task status', 'sys_job_status', 1, 31, 31, 'Task status list', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (13, 'Task group', 'sys_job_group', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (14, 'Admin login status', 'admin_login_status', 1, 31, 31, '', NULL, '2022-09-16 15:26:01');
INSERT INTO `sys_dict_type` VALUES (15, 'Operation log status', 'sys_oper_log_status', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (16, 'Task policy', 'sys_job_policy', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (17, 'Menu status', 'sys_show_hide', 1, 31, 0, 'Menu status', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (18, 'System switch', 'sys_normal_disable', 1, 31, 31, 'System switch', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (24, 'System built-in', 'sys_yes_no', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (25, 'Article publish status', 'cms_article_pub_type', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (26, 'Article additional status', 'cms_article_attr', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (27, 'Article type', 'cms_article_type', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (28, 'Article column model category', 'cms_cate_models', 1, 1, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (29, 'Government work model category', 'gov_cate_models', 1, 2, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (30, 'Menu module type', 'menu_module_type', 1, 2, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (31, 'Workflow type', 'flow_type', 1, 2, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (32, 'Workflow approval status', 'flow_status', 1, 31, 0, 'Workflow approval status', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (33, 'Blog category type', 'sys_blog_type', 1, 31, 31, 'Flag in blog categories', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (34, 'Blog log flag', 'sys_log_sign', 1, 31, 0, 'Flag data dictionary in blog log management', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (35, 'Workflow urgency level', 'flow_level', 1, 31, 31, '', NULL, '2021-07-20 08:55:20');
INSERT INTO `sys_dict_type` VALUES (48, 'Plugin store discount', 'plugin_store_discount', 1, 31, 0, '', '2021-08-14 11:59:26', '2021-08-14 11:59:26');
INSERT INTO `sys_dict_type` VALUES (49, 'CMScolumn navigation position', 'cms_nav_position', 1, 22, 0, '', '2021-08-31 15:37:04', '2021-08-31 15:37:04');
INSERT INTO `sys_dict_type` VALUES (50, 'Operation log type', 'sys_oper_log_type', 1, 31, 0, '', '2022-12-21 11:55:02', '2022-12-21 11:55:02');

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log`  (
  `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'AccessID',
  `login_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Login account',
  `ipaddr` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'LoginIPaddress',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Login location',
  `browser` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Browser type',
  `os` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Operating system',
  `status` tinyint(4) NULL DEFAULT 0 COMMENT 'Login status（0success 1failure）',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Notification message',
  `login_time` datetime NULL DEFAULT NULL COMMENT 'Login time',
  `module` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Login module',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'System access log' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------
INSERT INTO `sys_login_log` VALUES (1, 'demo', '::1', 'InternalIP', 'Chrome', 'Windows 10', 1, 'Login successful', '2023-01-19 10:17:18', 'System backend');

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
  `oper_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Log primary key',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Module title',
  `business_type` int(2) NULL DEFAULT 0 COMMENT 'Business type（0other 1add 2modify 3delete）',
  `method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Method name',
  `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Request method',
  `operator_type` int(1) NULL DEFAULT 0 COMMENT 'Operator type（0other 1backend user 2mobile user）',
  `oper_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Operator name',
  `dept_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Department name',
  `oper_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'RequestURL',
  `oper_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Host address',
  `oper_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Operation location',
  `oper_param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'Request parameters',
  `error_msg` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Error message',
  `oper_time` datetime NULL DEFAULT NULL COMMENT 'Operation time',
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 57 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Operation log record' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------
INSERT INTO `sys_oper_log` VALUES (1, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=sys_oper_log_type&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"sys_oper_log_type\"}', '', '2023-01-19 10:10:49');
INSERT INTO `sys_oper_log` VALUES (2, 'Operation Log', 0, '/api/v1/system/operLog/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/operLog/list?pageNum=1&pageSize=10', '::1', 'InternalIP', '{\"pageNum\":\"1\",\"pageSize\":\"10\"}', '', '2023-01-19 10:10:49');
INSERT INTO `sys_oper_log` VALUES (3, 'Operation Log', 0, '/api/v1/system/operLog/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/operLog/list?pageNum=1&pageSize=10', '::1', 'InternalIP', '{\"pageNum\":\"1\",\"pageSize\":\"10\"}', '', '2023-01-19 10:11:04');
INSERT INTO `sys_oper_log` VALUES (4, 'Online Users', 0, '/api/v1/system/online/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/online/list?ipaddr=&userName=&pageNum=1&pageSize=10', '::1', 'InternalIP', '{\"ipaddr\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\",\"userName\":\"\"}', '', '2023-01-19 10:16:55');
INSERT INTO `sys_oper_log` VALUES (5, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=sys_oper_log_type&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"sys_oper_log_type\"}', '', '2023-01-19 10:16:57');
INSERT INTO `sys_oper_log` VALUES (6, 'Operation Log', 0, '/api/v1/system/operLog/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/operLog/list?pageNum=1&pageSize=10', '::1', 'InternalIP', '{\"pageNum\":\"1\",\"pageSize\":\"10\"}', '', '2023-01-19 10:16:57');
INSERT INTO `sys_oper_log` VALUES (7, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=admin_login_status&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"admin_login_status\"}', '', '2023-01-19 10:16:59');
INSERT INTO `sys_oper_log` VALUES (8, 'Login Log', 0, '/api/v1/system/loginLog/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/loginLog/list?pageNum=1&pageSize=10&status=&ipaddr=&loginLocation=&userName=', '::1', 'InternalIP', '{\"ipaddr\":\"\",\"loginLocation\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\",\"status\":\"\",\"userName\":\"\"}', '', '2023-01-19 10:16:59');
INSERT INTO `sys_oper_log` VALUES (9, 'Server Monitoring', 0, '/api/v1/system/monitor/server', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/monitor/server', '::1', 'InternalIP', '{}', '', '2023-01-19 10:17:01');
INSERT INTO `sys_oper_log` VALUES (10, 'Login Log', 0, '/api/v1/system/loginLog/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/loginLog/list?pageNum=1&pageSize=10&status=&ipaddr=&loginLocation=&userName=', '::1', 'InternalIP', '{\"ipaddr\":\"\",\"loginLocation\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\",\"status\":\"\",\"userName\":\"\"}', '', '2023-01-19 10:17:05');
INSERT INTO `sys_oper_log` VALUES (11, 'Login Log', 0, '/api/v1/system/loginLog/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/loginLog/list?pageNum=1&pageSize=10&status=&ipaddr=&loginLocation=&userName=', '::1', 'InternalIP', '{\"ipaddr\":\"\",\"loginLocation\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\",\"status\":\"\",\"userName\":\"\"}', '', '2023-01-19 10:17:07');
INSERT INTO `sys_oper_log` VALUES (12, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=admin_login_status&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"admin_login_status\"}', '', '2023-01-19 10:17:20');
INSERT INTO `sys_oper_log` VALUES (13, 'Login Log', 0, '/api/v1/system/loginLog/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/loginLog/list?pageNum=1&pageSize=10&status=&ipaddr=&loginLocation=&userName=', '::1', 'InternalIP', '{\"ipaddr\":\"\",\"loginLocation\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\",\"status\":\"\",\"userName\":\"\"}', '', '2023-01-19 10:17:20');
INSERT INTO `sys_oper_log` VALUES (14, 'Login Log', 0, '/api/v1/system/loginLog/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/loginLog/list?pageNum=1&pageSize=10&status=&ipaddr=&loginLocation=&userName=', '::1', 'InternalIP', '{\"ipaddr\":\"\",\"loginLocation\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\",\"status\":\"\",\"userName\":\"\"}', '', '2023-01-19 10:17:22');
INSERT INTO `sys_oper_log` VALUES (15, 'Server Monitoring', 0, '/api/v1/system/monitor/server', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/monitor/server', '::1', 'InternalIP', '{}', '', '2023-01-19 10:17:25');
INSERT INTO `sys_oper_log` VALUES (16, 'Dictionary Management', 0, '/api/v1/system/dict/type/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/type/list?pageNum=1&pageSize=10&dictName=&dictType=&status=', '::1', 'InternalIP', '{\"dictName\":\"\",\"dictType\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\",\"status\":\"\"}', '', '2023-01-19 10:17:29');
INSERT INTO `sys_oper_log` VALUES (17, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=sys_yes_no&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"sys_yes_no\"}', '', '2023-01-19 10:17:31');
INSERT INTO `sys_oper_log` VALUES (18, 'Parameter Management', 0, '/api/v1/system/config/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/config/list?pageNum=1&pageSize=10&configName=&configKey=&configType=', '::1', 'InternalIP', '{\"configKey\":\"\",\"configName\":\"\",\"configType\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\"}', '', '2023-01-19 10:17:31');
INSERT INTO `sys_oper_log` VALUES (19, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=sys_job_status&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"sys_job_status\"}', '', '2023-01-19 10:18:32');
INSERT INTO `sys_oper_log` VALUES (20, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=sys_job_policy&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"sys_job_policy\"}', '', '2023-01-19 10:18:32');
INSERT INTO `sys_oper_log` VALUES (21, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=sys_job_group&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"sys_job_group\"}', '', '2023-01-19 10:18:32');
INSERT INTO `sys_oper_log` VALUES (22, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=sys_show_hide&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"sys_show_hide\"}', '', '2023-01-19 10:19:34');
INSERT INTO `sys_oper_log` VALUES (23, '', 0, '/api/v1/system/menu/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/menu/list?title=&component=', '::1', 'InternalIP', '{\"component\":\"\",\"title\":\"\"}', '', '2023-01-19 10:19:34');
INSERT INTO `sys_oper_log` VALUES (24, 'Delete Menu', 0, '/api/v1/system/menu/delete', 'DELETE', 1, 'demo', 'Finance Department', '/api/v1/system/menu/delete', '::1', 'InternalIP', '{\"ids\":[36]}', '', '2023-01-19 10:19:43');
INSERT INTO `sys_oper_log` VALUES (25, '', 0, '/api/v1/system/user/getUserMenus', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/user/getUserMenus', '::1', 'InternalIP', '{}', '', '2023-01-19 10:19:44');
INSERT INTO `sys_oper_log` VALUES (26, '', 0, '/api/v1/system/menu/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/menu/list?title=&component=', '::1', 'InternalIP', '{\"component\":\"\",\"title\":\"\"}', '', '2023-01-19 10:19:44');
INSERT INTO `sys_oper_log` VALUES (27, 'Delete Menu', 0, '/api/v1/system/menu/delete', 'DELETE', 1, 'demo', 'Finance Department', '/api/v1/system/menu/delete', '::1', 'InternalIP', '{\"ids\":[53]}', '', '2023-01-19 10:19:48');
INSERT INTO `sys_oper_log` VALUES (28, '', 0, '/api/v1/system/user/getUserMenus', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/user/getUserMenus', '::1', 'InternalIP', '{}', '', '2023-01-19 10:19:49');
INSERT INTO `sys_oper_log` VALUES (29, '', 0, '/api/v1/system/menu/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/menu/list?title=&component=', '::1', 'InternalIP', '{\"component\":\"\",\"title\":\"\"}', '', '2023-01-19 10:19:49');
INSERT INTO `sys_oper_log` VALUES (30, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=sys_show_hide&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"sys_show_hide\"}', '', '2023-01-19 10:19:56');
INSERT INTO `sys_oper_log` VALUES (31, '', 0, '/api/v1/system/menu/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/menu/list?title=&component=', '::1', 'InternalIP', '{\"component\":\"\",\"title\":\"\"}', '', '2023-01-19 10:19:56');
INSERT INTO `sys_oper_log` VALUES (32, 'Server Monitoring', 0, '/api/v1/system/monitor/server', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/monitor/server', '::1', 'InternalIP', '{}', '', '2023-01-19 10:35:29');
INSERT INTO `sys_oper_log` VALUES (33, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=admin_login_status&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"admin_login_status\"}', '', '2023-01-19 10:35:31');
INSERT INTO `sys_oper_log` VALUES (34, 'Login Log', 0, '/api/v1/system/loginLog/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/loginLog/list?pageNum=1&pageSize=10&status=&ipaddr=&loginLocation=&userName=', '::1', 'InternalIP', '{\"ipaddr\":\"\",\"loginLocation\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\",\"status\":\"\",\"userName\":\"\"}', '', '2023-01-19 10:35:31');
INSERT INTO `sys_oper_log` VALUES (35, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=sys_oper_log_type&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"sys_oper_log_type\"}', '', '2023-01-19 10:35:33');
INSERT INTO `sys_oper_log` VALUES (36, 'Operation Log', 0, '/api/v1/system/operLog/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/operLog/list?pageNum=1&pageSize=10', '::1', 'InternalIP', '{\"pageNum\":\"1\",\"pageSize\":\"10\"}', '', '2023-01-19 10:35:33');
INSERT INTO `sys_oper_log` VALUES (37, 'Online Users', 0, '/api/v1/system/online/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/online/list?ipaddr=&userName=&pageNum=1&pageSize=10', '::1', 'InternalIP', '{\"ipaddr\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\",\"userName\":\"\"}', '', '2023-01-19 10:35:35');
INSERT INTO `sys_oper_log` VALUES (38, '', 0, '/api/v1/system/personal/getPersonalInfo', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/personal/getPersonalInfo', '::1', 'InternalIP', '{}', '', '2023-01-19 10:49:47');
INSERT INTO `sys_oper_log` VALUES (39, 'Dictionary Management', 0, '/api/v1/system/dict/type/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/type/list?pageNum=1&pageSize=10&dictName=&dictType=&status=', '::1', 'InternalIP', '{\"dictName\":\"\",\"dictType\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\",\"status\":\"\"}', '', '2023-01-19 11:01:02');
INSERT INTO `sys_oper_log` VALUES (40, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=sys_yes_no&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"sys_yes_no\"}', '', '2023-01-19 11:01:05');
INSERT INTO `sys_oper_log` VALUES (41, 'Parameter Management', 0, '/api/v1/system/config/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/config/list?pageNum=1&pageSize=10&configName=&configKey=&configType=', '::1', 'InternalIP', '{\"configKey\":\"\",\"configName\":\"\",\"configType\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\"}', '', '2023-01-19 11:01:05');
INSERT INTO `sys_oper_log` VALUES (42, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=sys_show_hide&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"sys_show_hide\"}', '', '2023-01-19 11:01:07');
INSERT INTO `sys_oper_log` VALUES (43, '', 0, '/api/v1/system/menu/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/menu/list?title=&component=', '::1', 'InternalIP', '{\"component\":\"\",\"title\":\"\"}', '', '2023-01-19 11:01:08');
INSERT INTO `sys_oper_log` VALUES (44, 'Role Management', 0, '/api/v1/system/role/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/role/list?roleName=&roleStatus=&pageNum=1&pageSize=10', '::1', 'InternalIP', '{\"pageNum\":\"1\",\"pageSize\":\"10\",\"roleName\":\"\",\"roleStatus\":\"\"}', '', '2023-01-19 11:01:11');
INSERT INTO `sys_oper_log` VALUES (45, 'Department Management', 0, '/api/v1/system/dept/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dept/list?pageNum=1&pageSize=10&deptName=&status=', '::1', 'InternalIP', '{\"deptName\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\",\"status\":\"\"}', '', '2023-01-19 11:01:14');
INSERT INTO `sys_oper_log` VALUES (46, 'Position Management', 0, '/api/v1/system/post/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/post/list?postName=&status=&postCode=&pageNum=1&pageSize=10', '::1', 'InternalIP', '{\"pageNum\":\"1\",\"pageSize\":\"10\",\"postCode\":\"\",\"postName\":\"\",\"status\":\"\"}', '', '2023-01-19 11:01:19');
INSERT INTO `sys_oper_log` VALUES (47, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=sys_user_sex&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"sys_user_sex\"}', '', '2023-01-19 11:01:23');
INSERT INTO `sys_oper_log` VALUES (48, '', 0, '/api/v1/system/dept/treeSelect', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dept/treeSelect', '::1', 'InternalIP', '{}', '', '2023-01-19 11:01:23');
INSERT INTO `sys_oper_log` VALUES (49, '', 0, '/api/v1/system/user/params', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/user/params', '::1', 'InternalIP', '{}', '', '2023-01-19 11:01:23');
INSERT INTO `sys_oper_log` VALUES (50, 'User Management', 0, '/api/v1/system/user/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/user/list?pageNum=1&pageSize=10&deptId=&mobile=&status=&keyWords=', '::1', 'InternalIP', '{\"deptId\":\"\",\"keyWords\":\"\",\"mobile\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\",\"status\":\"\"}', '', '2023-01-19 11:01:23');
INSERT INTO `sys_oper_log` VALUES (51, 'Server Monitoring', 0, '/api/v1/system/monitor/server', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/monitor/server', '::1', 'InternalIP', '{}', '', '2023-01-19 11:01:30');
INSERT INTO `sys_oper_log` VALUES (52, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=admin_login_status&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"admin_login_status\"}', '', '2023-01-19 11:01:32');
INSERT INTO `sys_oper_log` VALUES (53, 'Login Log', 0, '/api/v1/system/loginLog/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/loginLog/list?pageNum=1&pageSize=10&status=&ipaddr=&loginLocation=&userName=', '::1', 'InternalIP', '{\"ipaddr\":\"\",\"loginLocation\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\",\"status\":\"\",\"userName\":\"\"}', '', '2023-01-19 11:01:32');
INSERT INTO `sys_oper_log` VALUES (54, '', 0, '/api/v1/system/dict/data/getDictData', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/dict/data/getDictData?dictType=sys_oper_log_type&defaultValue=', '::1', 'InternalIP', '{\"defaultValue\":\"\",\"dictType\":\"sys_oper_log_type\"}', '', '2023-01-19 11:01:34');
INSERT INTO `sys_oper_log` VALUES (55, 'Operation Log', 0, '/api/v1/system/operLog/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/operLog/list?pageNum=1&pageSize=10', '::1', 'InternalIP', '{\"pageNum\":\"1\",\"pageSize\":\"10\"}', '', '2023-01-19 11:01:35');
INSERT INTO `sys_oper_log` VALUES (56, 'Online Users', 0, '/api/v1/system/online/list', 'GET', 1, 'demo', 'Finance Department', '/api/v1/system/online/list?ipaddr=&userName=&pageNum=1&pageSize=10', '::1', 'InternalIP', '{\"ipaddr\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\",\"userName\":\"\"}', '', '2023-01-19 11:01:36');

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `post_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PostID',
  `post_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Post Code',
  `post_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Post Name',
  `post_sort` int(4) NOT NULL COMMENT 'Display Order',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Status（0normal 1disabled）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Remark',
  `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Created By',
  `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Updated By',
  `created_at` datetime NULL DEFAULT NULL COMMENT 'Creation Time',
  `updated_at` datetime NULL DEFAULT NULL COMMENT 'Update Time',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT 'Deletion Time',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Position Information Table' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES (1, 'ceo', 'Chairman', 1, 1, '', 0, 0, '2021-07-11 11:32:58', NULL, NULL);
INSERT INTO `sys_post` VALUES (2, 'se', 'Project Manager', 2, 1, '', 0, 0, '2021-07-12 11:01:26', NULL, NULL);
INSERT INTO `sys_post` VALUES (3, 'hr', 'Human Resources', 3, 1, '', 0, 31, '2021-07-12 11:01:30', '2022-09-16 16:48:18', NULL);
INSERT INTO `sys_post` VALUES (4, 'user', 'Regular Employee', 4, 0, 'Regular employee', 0, 31, '2021-07-12 11:01:33', '2022-04-08 15:32:23', NULL);
INSERT INTO `sys_post` VALUES (5, 'it', 'ITDepartment', 5, 1, 'Information Department', 31, 31, '2021-07-12 11:09:42', '2022-04-09 12:59:12', NULL);
INSERT INTO `sys_post` VALUES (6, '1111', '1111', 0, 1, '11111', 31, 0, '2022-04-08 15:32:44', '2022-04-08 15:32:44', '2022-04-08 15:51:24');
INSERT INTO `sys_post` VALUES (7, '222', '2222', 0, 1, '22222', 31, 0, '2022-04-08 15:32:55', '2022-04-08 15:32:55', '2022-04-08 15:51:24');
INSERT INTO `sys_post` VALUES (8, '33333', '3333', 0, 0, '33333', 31, 0, '2022-04-08 15:33:01', '2022-04-08 15:33:01', '2022-04-08 15:51:40');
INSERT INTO `sys_post` VALUES (9, '222', '111', 0, 1, '2313213', 31, 0, '2022-04-08 15:52:53', '2022-04-08 15:52:53', '2022-04-08 15:52:56');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Status;0:Disabled;1:Active',
  `list_order` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Sort Order',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Role Name',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Remark',
  `data_scope` tinyint(3) UNSIGNED NOT NULL DEFAULT 3 COMMENT 'Data Scope（1：All data permissions 2：Custom data permissions 3：Current department data permissions 4：Current department and subordinate data permissions）',
  `created_at` datetime NULL DEFAULT NULL COMMENT 'Creation Time',
  `updated_at` datetime NULL DEFAULT NULL COMMENT 'Update Time',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Role Table' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, 1, 0, 'Super Admin', 'Remark', 3, '2022-04-01 11:38:39', '2022-04-28 10:00:15');
INSERT INTO `sys_role` VALUES (2, 1, 0, 'Regular Admin', 'Remark', 3, '2022-04-01 11:38:39', '2022-04-28 10:01:34');
INSERT INTO `sys_role` VALUES (3, 1, 0, 'Site Admin', 'Site management personnel', 3, '2022-04-01 11:38:39', '2022-04-01 11:38:39');
INSERT INTO `sys_role` VALUES (4, 1, 0, 'Junior Admin', 'Junior Administrator', 3, '2022-04-01 11:38:39', '2022-04-01 11:38:39');
INSERT INTO `sys_role` VALUES (5, 1, 0, 'Senior Admin', 'Senior Administrator', 2, '2022-04-01 11:38:39', '2022-04-01 11:38:39');
INSERT INTO `sys_role` VALUES (8, 1, 0, 'District Admin', '', 2, '2022-04-01 11:38:39', '2022-04-06 09:53:40');

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept`  (
  `role_id` bigint(20) NOT NULL COMMENT 'RoleID',
  `dept_id` bigint(20) NOT NULL COMMENT 'DepartmentID',
  PRIMARY KEY (`role_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Role and department association table' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
INSERT INTO `sys_role_dept` VALUES (5, 103);
INSERT INTO `sys_role_dept` VALUES (5, 104);
INSERT INTO `sys_role_dept` VALUES (5, 105);
INSERT INTO `sys_role_dept` VALUES (8, 105);
INSERT INTO `sys_role_dept` VALUES (8, 106);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Username',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Chinese mobile number without country code，international format：country code-mobile number',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'User nickname',
  `birthday` int(11) NOT NULL DEFAULT 0 COMMENT 'Birthday',
  `user_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Login password;cmf_passwordencrypted with',
  `user_salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Encryption salt',
  `user_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT 'User status;0:Disabled,1:Active,2:Unverified',
  `user_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'User login email',
  `sex` tinyint(2) NOT NULL DEFAULT 0 COMMENT 'Gender;0:Not specified,1:Male,2:Female',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'User avatar',
  `dept_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Departmentid',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Remark',
  `is_admin` tinyint(4) NOT NULL DEFAULT 1 COMMENT 'Is backend administrator 1 Yes  0   No',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Contact address',
  `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT ' Description',
  `last_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Last loginip',
  `last_login_time` datetime NULL DEFAULT NULL COMMENT 'Last login time',
  `created_at` datetime NULL DEFAULT NULL COMMENT 'Creation time',
  `updated_at` datetime NULL DEFAULT NULL COMMENT 'Update time',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT 'Deletion time',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_login`(`user_name`, `deleted_at`) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile`, `deleted_at`) USING BTREE,
  INDEX `user_nickname`(`user_nickname`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 43 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'User table' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '13578342363', 'Super Administrator', 0, 'c567ae329f9929b518759d3bea13f492', 'f9aZTAa8yz', 1, 'yxh669@qq.com', 1, 'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-07-19/ccwpeuqz1i2s769hua.jpeg', 101, '', 1, 'asdasfdsafasdasfdsaf Test Address', 'Description', '::1', '2022-10-26 03:01:52', '2021-06-22 17:58:00', '2022-11-03 15:44:38', NULL);
INSERT INTO `sys_user` VALUES (2, 'yixiaohu', '13699885599', 'Nice', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'yxh@qq.com', 1, 'upload_file/2022-11-04/co3e5ljknns8jhlp8s.jpg', 102, 'Remark', 1, '', '', '::1', '2022-11-04 09:54:56', '2021-06-22 17:58:00', '2022-11-04 17:54:56', NULL);
INSERT INTO `sys_user` VALUES (3, 'zs', '16399669855', 'Zhang San', 0, '41e3778c20338f4d7d6cc886fd3b2a52', 'redoHIj524', 1, 'zs@qq.com', 0, 'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-08-02/cd8nif79egjg9kbkgk.jpeg', 101, '', 1, '', '', '::1', '2022-04-28 10:01:47', '2021-06-22 17:58:00', '2022-04-28 10:01:47', NULL);
INSERT INTO `sys_user` VALUES (4, 'qlgl', '13758596696', 'Testc', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'qlgl@qq.com', 0, '', 102, '', 1, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2022-11-03 15:44:20', NULL);
INSERT INTO `sys_user` VALUES (5, 'test', '13845696696', 'Test2', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '123@qq.com', 0, '', 101, '', 0, '', '', '::1', '2022-03-30 10:50:39', '2021-06-22 17:58:00', '2022-11-03 15:44:10', NULL);
INSERT INTO `sys_user` VALUES (6, '18999998889', '13755866654', 'Liu Da', 0, '5df78d20315a5af61f45d20f72c184fc', 'lC6OoXDCbM', 1, '1223@qq.com', 0, '', 103, '', 1, '', '', '[::1]', '2022-02-25 14:29:22', '2021-06-22 17:58:00', '2022-11-03 17:05:07', NULL);
INSERT INTO `sys_user` VALUES (7, 'zmm', '13788566696', 'Zhang Mingming', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '11123@qq.com', 0, '', 104, '', 1, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2022-04-12 17:55:42', NULL);
INSERT INTO `sys_user` VALUES (8, 'lxx', '13756566696', 'Li Xiaoxiao', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '123333@qq.com', 0, '', 101, '', 1, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2022-04-12 17:55:45', NULL);
INSERT INTO `sys_user` VALUES (10, 'xmm', '13588999969', 'Little Secret', 0, '2de2a8df703bfc634cfda2cb2f6a59be', 'Frz7LJY7SE', 1, '696@qq.com', 0, '', 101, '', 1, '', '', '[::1]', '2021-07-22 17:08:53', '2021-06-22 17:58:00', '2022-04-12 17:55:50', NULL);
INSERT INTO `sys_user` VALUES (14, 'cd_19', '13699888899', 'Kan Jinli Technology', 0, '1169d5fe4119fd4277a95f02d7036171', '7paigEoedh', 1, '', 0, '', 102, '', 1, '', '', '', NULL, '2021-06-22 17:58:00', '2022-04-12 18:13:22', NULL);
INSERT INTO `sys_user` VALUES (15, 'lmm', '13587754545', 'Liu Minmin', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'a@coc.com', 0, '', 201, '', 1, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2022-04-12 17:56:23', NULL);
INSERT INTO `sys_user` VALUES (16, 'ldn', '13899658874', 'Li Daniu', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'a@ll.con', 0, '', 102, '', 1, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2022-04-12 17:56:27', NULL);
INSERT INTO `sys_user` VALUES (20, 'dbc', '13877555566', 'Big Dictionary', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '', 0, '', 0, '', 1, '', '', '', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL);
INSERT INTO `sys_user` VALUES (22, 'yxfmlbb', '15969423326', 'Big Data Department Test', 0, '66f89b40ee4a10aabaf70c15756429ea', 'mvd2OtUe8f', 1, 'yxh6691@qq.com', 0, 'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-09-29/cem20k3fdciosy7nwo.jpeg', 200, '', 1, '2222233', '1222', '[::1]', '2021-10-28 11:36:07', '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL);
INSERT INTO `sys_user` VALUES (23, 'wangming', '13699888855', 'Wang Ming', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '', 0, '', 0, '', 1, '', '', '', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL);
INSERT INTO `sys_user` VALUES (24, 'zhk', '13699885591', 'General Department', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '', 0, '', 0, '', 1, '', '', '192.168.0.146', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL);
INSERT INTO `sys_user` VALUES (28, 'demo3', '18699888855', 'Test Account1', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '123132@qq.com', 0, '', 109, '', 1, '', '', '192.168.0.229', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL);
INSERT INTO `sys_user` VALUES (31, 'demo', '15334455789', 'Li Si', 0, '6dd68eea81e0fca319add0bd58c3fdf6', '46PvWe1Sl7', 1, '123@qq.com', 2, 'upload_file/2022-11-11/co9copop81co0gysbz.jpg', 109, '3', 1, 'Qujing City, Yunnan Province22223', 'No matter how bad life gets，it doesn\'t stop me from getting better', '::1', '2023-01-19 10:17:18', '2021-06-22 17:58:00', '2022-11-11 17:25:27', NULL);
INSERT INTO `sys_user` VALUES (32, 'demo100', '18699888859', 'Test Account1', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '', 0, '', 0, '', 1, '', '', '[::1]', '2021-11-24 18:01:21', '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL);
INSERT INTO `sys_user` VALUES (33, 'demo110', '18699888853', 'Test Account1', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '', 0, '', 0, '', 1, '', '', '', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL);
INSERT INTO `sys_user` VALUES (34, 'yxfmlbb2', '15969423327', 'R&D Department Test', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '1111@qqq.com', 1, '', 103, '', 0, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL);
INSERT INTO `sys_user` VALUES (35, 'wk666', '18888888888', 'wk', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '396861557@qq.com', 1, '', 100, '', 1, '', '', '[::1]', '2021-12-09 14:52:37', '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL);
INSERT INTO `sys_user` VALUES (36, 'zxd', '13699885565', 'Zhang Xiaodong', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'zxk@qq.com', 1, '', 201, '666', 1, '', '', '', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL);
INSERT INTO `sys_user` VALUES (37, 'yxfmlbb3', '13513513511', 'Zhang San', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '111@qq.com', 0, '', 204, '', 1, '', '', '[::1]', '2021-07-26 14:49:25', '2021-06-22 17:58:00', '2021-07-26 14:49:18', NULL);
INSERT INTO `sys_user` VALUES (38, 'test_user', '18888888880', 'test', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '11@qq.com', 1, '', 200, '111', 0, '', '', '', NULL, '2021-06-22 17:58:00', '2021-07-12 22:05:29', NULL);
INSERT INTO `sys_user` VALUES (39, 'asan', '18687460555', 'A San', 0, '2354837137115700e2adf870ac113dcf', 'drdDvbtYZW', 1, '456654@qq.com', 1, '', 201, '666666', 1, '', '', '', NULL, '2021-07-12 17:21:43', '2021-07-12 21:13:31', '2021-07-12 22:00:44');
INSERT INTO `sys_user` VALUES (40, 'asi', '13655888888', 'A Si', 0, 'fbb755b35d48759dad47bb1540249fd1', '9dfUstcxrz', 1, '5464@qq.com', 1, '', 201, 'adsaasd', 1, '', '', '', NULL, '2021-07-12 17:46:27', '2021-07-12 21:29:41', '2021-07-12 22:00:44');
INSERT INTO `sys_user` VALUES (41, 'awu', '13578556546', 'A Wu', 0, '3b36a96afa0dfd66aa915e0816e0e9f6', '9gHRa9ho4U', 0, '132321@qq.com', 1, '', 201, 'asdasdasd', 1, '', '', '', NULL, '2021-07-12 17:54:31', '2021-07-12 21:46:34', '2021-07-12 21:59:56');
INSERT INTO `sys_user` VALUES (42, 'demo01', '13699888556', 'Test01222', 0, '048dc94116558fb40920f3553ecd5fe8', 'KiVrfzKJQx', 1, '456@qq.com', 2, '', 109, 'Test User', 1, '', '', '', NULL, '2022-04-12 16:15:23', '2022-04-12 17:54:49', NULL);

-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uuid` char(32) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT 'User identifier',
  `token` varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT 'Usertoken',
  `create_time` datetime NULL DEFAULT NULL COMMENT 'Login time',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Username',
  `ip` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Loginip',
  `explorer` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Browser',
  `os` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Operating system',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_token`(`token`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'User online status table' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------
INSERT INTO `sys_user_online` VALUES (15, 'f2d1992a5bff07f46a70490451a225e8', '7ZUSfVIf2HyYjcv86SKPPs29v003ECPEScsdYsYYqO1xEIcOpHEu9FS4ZmjQsf1GCmQAky2EuUzyGJF53YyQWvdOP3vC5KeHSmJ1BX0mSAnnw7CD4fNQF4wbtkE4I78lTUjvovXRSC5oDkWPMe79iQ==', '2023-01-13 14:09:51', 'demo', '::1', 'Chrome', 'Windows 10');
INSERT INTO `sys_user_online` VALUES (16, 'c0ce4001700ef589195c41ef073daa62', '7ZUSfVIf2HyYjcv86SKPPs29v003ECPEScsdYsYYqO0y3Gdni2HPIbjTYvAE1/8jYVxUh0VVfhtbUzIENCClH8vlzKtsEfway1I2p8fkF9NRP0ycB7htjT0UJLDmhMUpMaTXSYnL2PPorrqaf4roHg==', '2023-01-19 10:17:18', 'demo', '::1', 'Chrome', 'Windows 10');

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post`  (
  `user_id` bigint(20) NOT NULL COMMENT 'UserID',
  `post_id` bigint(20) NOT NULL COMMENT 'PostID',
  PRIMARY KEY (`user_id`, `post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'User and post association table' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
INSERT INTO `sys_user_post` VALUES (1, 2);
INSERT INTO `sys_user_post` VALUES (1, 3);
INSERT INTO `sys_user_post` VALUES (2, 1);
INSERT INTO `sys_user_post` VALUES (2, 2);
INSERT INTO `sys_user_post` VALUES (3, 2);
INSERT INTO `sys_user_post` VALUES (4, 1);
INSERT INTO `sys_user_post` VALUES (5, 2);
INSERT INTO `sys_user_post` VALUES (10, 1);
INSERT INTO `sys_user_post` VALUES (10, 2);
INSERT INTO `sys_user_post` VALUES (10, 3);
INSERT INTO `sys_user_post` VALUES (10, 4);
INSERT INTO `sys_user_post` VALUES (10, 5);
INSERT INTO `sys_user_post` VALUES (14, 1);
INSERT INTO `sys_user_post` VALUES (15, 4);
INSERT INTO `sys_user_post` VALUES (16, 2);
INSERT INTO `sys_user_post` VALUES (22, 1);
INSERT INTO `sys_user_post` VALUES (22, 2);
INSERT INTO `sys_user_post` VALUES (31, 2);
INSERT INTO `sys_user_post` VALUES (34, 1);
INSERT INTO `sys_user_post` VALUES (35, 2);
INSERT INTO `sys_user_post` VALUES (35, 3);
INSERT INTO `sys_user_post` VALUES (36, 1);
INSERT INTO `sys_user_post` VALUES (37, 3);
INSERT INTO `sys_user_post` VALUES (38, 2);
INSERT INTO `sys_user_post` VALUES (38, 3);
INSERT INTO `sys_user_post` VALUES (42, 2);
INSERT INTO `sys_user_post` VALUES (42, 3);

SET FOREIGN_KEY_CHECKS = 1;

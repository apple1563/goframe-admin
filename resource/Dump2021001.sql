-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.33 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             12.5.0.6677
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Dumping structure for table starter.button
DROP TABLE IF EXISTS `button`;
CREATE TABLE IF NOT EXISTS `button` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `menu_id` bigint NOT NULL DEFAULT '0' COMMENT '按钮所在菜单id',
  `menu_title` varchar(50) NOT NULL DEFAULT '' COMMENT '按钮所在菜单名称',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '按钮标识符',
  `title` varchar(50) NOT NULL DEFAULT '' COMMENT '按钮名称',
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `FK_button_menu` (`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='  Id             int64       `json:"id"             description:"菜单ID"`\r\n    Pid            int64       `json:"pid"            description:"父菜单ID"`\r\n    Title          string      `json:"title"          description:"菜单名称"`\r\n    Name           string      `json:"name"           description:"名称编码"`';

-- Dumping data for table starter.button: ~0 rows (approximately)
DELETE FROM `button`;
INSERT INTO `button` (`id`, `menu_id`, `menu_title`, `name`, `title`, `remark`) VALUES
	(4, 3, '按钮管理', 'add', '添加', '添加按钮'),
	(5, 3, '按钮管理', 'query', '查询', '查询表单'),
	(6, 3, '按钮管理', 'reset', '重置', '重置表单');

-- Dumping structure for table starter.casbin_policy
DROP TABLE IF EXISTS `casbin_policy`;
CREATE TABLE IF NOT EXISTS `casbin_policy` (
  `ptype` varchar(10) NOT NULL DEFAULT '',
  `v0` varchar(256) NOT NULL DEFAULT '',
  `v1` varchar(256) NOT NULL DEFAULT '',
  `v2` varchar(256) NOT NULL DEFAULT '',
  `v3` varchar(256) NOT NULL DEFAULT '',
  `v4` varchar(256) NOT NULL DEFAULT '',
  `v5` varchar(256) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='policy table';

-- Dumping data for table starter.casbin_policy: ~6 rows (approximately)
DELETE FROM `casbin_policy`;
INSERT INTO `casbin_policy` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
	('p', 'role-menu 1', '1', 'ALL', '', '', ''),
	('p', 'role-menu 1', '2', 'ALL', '', '', ''),
	('p', 'role-menu 2', '1', 'ALL', '', '', ''),
	('p', 'role-menu 2', '2', 'ALL', '', '', ''),
	('p', 'role-menu 3', '1', 'ALL', '', '', ''),
	('p', 'role-menu 3', '2', 'ALL', '', '', '');

-- Dumping structure for table starter.dict
DROP TABLE IF EXISTS `dict`;
CREATE TABLE IF NOT EXISTS `dict` (
  `config_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '参数键值',
  `config_type` tinyint(1) DEFAULT '0' COMMENT '系统内置（Y是 N否）',
  `create_by` int unsigned DEFAULT '0' COMMENT '创建者',
  `update_by` int unsigned DEFAULT '0' COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`config_id`) USING BTREE,
  UNIQUE KEY `uni_config_key` (`config_key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT;

-- Dumping data for table starter.dict: ~0 rows (approximately)
DELETE FROM `dict`;

-- Dumping structure for table starter.login_log
DROP TABLE IF EXISTS `login_log`;
CREATE TABLE IF NOT EXISTS `login_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL,
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `client_agent` varchar(500) DEFAULT NULL COMMENT '注册clientAgen头',
  `role` int DEFAULT '1' COMMENT '1用户2代理3管理',
  `p_role` varchar(45) DEFAULT NULL,
  `pid` int DEFAULT NULL,
  `p_username` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table starter.login_log: ~26 rows (approximately)
DELETE FROM `login_log`;
INSERT INTO `login_log` (`id`, `uid`, `username`, `ip`, `created_at`, `updated_at`, `client_agent`, `role`, `p_role`, `pid`, `p_username`) VALUES
	(1, 1, 'test', '127.0.0.1', '2023-09-27 19:15:29', '2023-09-27 19:15:29', 'Apifox/1.0.0 (https://apifox.com)', 2, '', 0, ''),
	(2, 1, 'test', '127.0.0.1', '2023-09-27 19:17:08', '2023-09-27 19:17:08', 'Apifox/1.0.0 (https://apifox.com)', 2, '', 0, ''),
	(3, 1, 'test', '127.0.0.1', '2023-09-27 19:18:30', '2023-09-27 19:18:30', 'Apifox/1.0.0 (https://apifox.com)', 2, '', 0, ''),
	(4, 1, 'test', '127.0.0.1', '2023-09-27 19:41:16', '2023-09-27 19:41:16', 'Apifox/1.0.0 (https://apifox.com)', 2, '', 0, ''),
	(5, 1, 'test', '127.0.0.1', '2023-09-27 19:42:01', '2023-09-27 19:42:01', 'Apifox/1.0.0 (https://apifox.com)', 2, '', 0, ''),
	(6, 1, 'test', '127.0.0.1', '2023-09-27 19:48:04', '2023-09-27 19:48:04', 'Apifox/1.0.0 (https://apifox.com)', 2, '', 0, ''),
	(7, 2, 'admin', '127.0.0.1', '2023-10-02 19:28:04', '2023-10-02 19:28:04', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(8, 2, 'admin', '127.0.0.1', '2023-10-02 19:43:01', '2023-10-02 19:43:01', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(9, 2, 'admin', '127.0.0.1', '2023-10-02 19:43:28', '2023-10-02 19:43:28', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(10, 2, 'admin', '127.0.0.1', '2023-10-02 19:45:44', '2023-10-02 19:45:44', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(11, 2, 'admin', '127.0.0.1', '2023-10-02 19:49:55', '2023-10-02 19:49:55', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(12, 2, 'admin', '127.0.0.1', '2023-10-02 19:50:24', '2023-10-02 19:50:24', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(13, 2, 'admin', '127.0.0.1', '2023-10-02 19:51:58', '2023-10-02 19:51:58', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(14, 2, 'admin', '127.0.0.1', '2023-10-02 19:54:06', '2023-10-02 19:54:06', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(15, 2, 'admin', '127.0.0.1', '2023-10-04 16:31:41', '2023-10-04 16:31:41', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(16, 2, 'admin', '127.0.0.1', '2023-10-04 16:32:58', '2023-10-04 16:32:58', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(17, 2, 'admin', '127.0.0.1', '2023-10-04 16:36:26', '2023-10-04 16:36:26', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(18, 2, 'admin', '127.0.0.1', '2023-10-04 16:37:13', '2023-10-04 16:37:13', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(19, 2, 'admin', '127.0.0.1', '2023-10-04 16:38:38', '2023-10-04 16:38:38', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(20, 2, 'admin', '127.0.0.1', '2023-10-04 16:39:29', '2023-10-04 16:39:29', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(21, 2, 'admin', '127.0.0.1', '2023-10-04 16:40:59', '2023-10-04 16:40:59', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(22, 2, 'admin', '127.0.0.1', '2023-10-04 16:44:49', '2023-10-04 16:44:49', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(23, 2, 'admin', '127.0.0.1', '2023-10-04 16:46:37', '2023-10-04 16:46:37', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(24, 2, 'admin', '127.0.0.1', '2023-10-04 16:47:57', '2023-10-04 16:47:57', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(25, 2, 'admin', '127.0.0.1', '2023-10-04 16:50:08', '2023-10-04 16:50:08', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, ''),
	(26, 2, 'admin', '127.0.0.1', '2023-10-04 19:36:48', '2023-10-04 19:36:48', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36', 0, '0', 0, '');

-- Dumping structure for table starter.menu
DROP TABLE IF EXISTS `menu`;
CREATE TABLE IF NOT EXISTS `menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `pid` bigint DEFAULT '0' COMMENT '父菜单ID，0表示根级',
  `title` varchar(64) NOT NULL COMMENT '菜单名称',
  `name` varchar(128) NOT NULL COMMENT '名称编码',
  `path` varchar(200) DEFAULT NULL COMMENT '路由地址',
  `icon` varchar(128) DEFAULT NULL COMMENT '菜单图标',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '菜单类型（1目录 2菜单 3按钮）',
  `redirect` varchar(255) DEFAULT NULL COMMENT '重定向地址',
  `permissions` varchar(255) DEFAULT NULL COMMENT '菜单包含权限集合',
  `permission_name` varchar(64) DEFAULT NULL COMMENT '权限名称',
  `component` varchar(255) NOT NULL COMMENT '组件路径',
  `always_show` tinyint(1) DEFAULT '0' COMMENT '取消自动计算根路由模式',
  `active_menu` varchar(255) DEFAULT NULL COMMENT '高亮菜单编码',
  `is_root` tinyint(1) DEFAULT '0' COMMENT '是否跟路由',
  `is_frame` tinyint(1) DEFAULT '1' COMMENT '是否内嵌',
  `frame_src` varchar(512) DEFAULT NULL COMMENT '内联外部地址',
  `keep_alive` tinyint(1) DEFAULT '0' COMMENT '缓存该路由',
  `hidden` tinyint(1) DEFAULT '0' COMMENT '是否隐藏',
  `affix` tinyint(1) DEFAULT '0' COMMENT '是否固定',
  `level` int NOT NULL DEFAULT '1' COMMENT '关系树等级 1根2子3孙',
  `tree` varchar(255) NOT NULL COMMENT '关系树',
  `sort` int DEFAULT '0' COMMENT '排序',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` tinyint(1) DEFAULT '1' COMMENT '菜单状态',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3 COMMENT='管理员_菜单权限';

-- Dumping data for table starter.menu: ~2 rows (approximately)
DELETE FROM `menu`;
INSERT INTO `menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES
	(1, 0, '菜单管理', 'menu', '/menu', 'menu-fold', 1, '/menu/index', '', '', 'Layout', 2, '', 1, 2, '', 2, 2, 2, 0, '', 0, '', 1, '2023-10-03 00:00:27', '2023-10-05 11:11:31'),
	(2, 1, '菜单管理', 'MenuIndex', 'index', '', 2, '', '', '', '/menu/menu/index.vue', 2, '', 2, 2, '', 2, 2, 2, 0, '', 0, '', 1, '2023-10-03 09:06:39', '2023-10-03 09:06:39'),
	(3, 1, '按钮管理', 'ButtonIndex', 'button', '', 2, '', '', '', '/menu/button/index.vue', 2, '', 2, 2, '', 2, 2, 2, 0, '', 0, '', 1, '2023-10-05 16:27:50', '2023-10-05 16:27:50');

-- Dumping structure for table starter.role
DROP TABLE IF EXISTS `role`;
CREATE TABLE IF NOT EXISTS `role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态;1:正常2:禁用',
  `list_order` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `data_scope` tinyint unsigned NOT NULL DEFAULT '3' COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='角色表';

-- Dumping data for table starter.role: ~3 rows (approximately)
DELETE FROM `role`;
INSERT INTO `role` (`id`, `status`, `list_order`, `name`, `remark`, `data_scope`, `created_at`, `updated_at`) VALUES
	(1, 1, 0, '用户', '', 3, NULL, NULL),
	(2, 1, 0, '管理员', '', 3, NULL, NULL),
	(3, 1, 0, '代理', '', 3, NULL, NULL);

-- Dumping structure for table starter.user
DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(45) NOT NULL,
  `password` varchar(255) NOT NULL,
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` int DEFAULT '1' COMMENT '1正常2禁用',
  `client_agent` varchar(500) DEFAULT NULL COMMENT '注册clientAgen头',
  `ip` varchar(50) DEFAULT NULL COMMENT 'IP',
  `role_id` int NOT NULL DEFAULT '1' COMMENT '1用户2代理3管理',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `p_role_id` int DEFAULT NULL,
  `pid` int DEFAULT NULL,
  `p_username` varchar(45) DEFAULT NULL,
  `role_name` varchar(45) NOT NULL DEFAULT '用户',
  `p_role_name` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table starter.user: ~1 rows (approximately)
DELETE FROM `user`;
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `phone`, `status`, `client_agent`, `ip`, `role_id`, `created_at`, `updated_at`, `p_role_id`, `pid`, `p_username`, `role_name`, `p_role_name`) VALUES
	(1, 'test', '$2a$10$EjkNiegGWC5jaNX3gyruAeDta1q6N.GD3XRs8sAsNUwviKlqmikum', NULL, NULL, NULL, 1, NULL, NULL, 2, '2023-09-27 11:02:53', '2023-09-27 11:15:24', NULL, NULL, NULL, '用户', NULL),
	(2, 'admin', '$2a$10$BLb8nmjt5kTQi5AtHDROKOG7MeLEfnONGxorjtEZTwfK3x1HV6O0a', NULL, NULL, NULL, 1, NULL, NULL, 0, '2023-10-02 11:26:00', '2023-10-02 11:26:00', NULL, NULL, NULL, '用户', NULL);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;

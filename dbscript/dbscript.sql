-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        8.0.21 - MySQL Community Server - GPL
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  11.0.0.5919
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- 导出 db_expired_reminder 的数据库结构
CREATE DATABASE IF NOT EXISTS `db_expired_reminder` /*!40100 DEFAULT CHARACTER SET utf8 */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `db_expired_reminder`;

-- 导出  表 db_expired_reminder.te_food 结构
CREATE TABLE IF NOT EXISTS `te_food` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `count` int DEFAULT NULL,
  `production_date` timestamp NULL DEFAULT NULL,
  `expired_date` timestamp NULL DEFAULT NULL,
  `created` timestamp NULL DEFAULT NULL,
  `updated` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `remaining_time` float DEFAULT NULL,
  `comment` varchar(255) DEFAULT NULL,
  `status` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8;

-- 正在导出表  db_expired_reminder.te_food 的数据：~2 rows (大约)
DELETE FROM `te_food`;
/*!40000 ALTER TABLE `te_food` DISABLE KEYS */;
/*!40000 ALTER TABLE `te_food` ENABLE KEYS */;

-- 导出  表 db_expired_reminder.te_user 结构
CREATE TABLE IF NOT EXISTS `te_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `email` varchar(200) DEFAULT NULL,
  `login_count` int DEFAULT NULL,
  `last_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `last_ip` varchar(200) DEFAULT 'current_timestamp()',
  `state` tinyint DEFAULT NULL,
  `created` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 正在导出表  db_expired_reminder.te_user 的数据：~1 rows (大约)
DELETE FROM `te_user`;
/*!40000 ALTER TABLE `te_user` DISABLE KEYS */;
INSERT INTO `te_user` (`id`, `username`, `password`, `email`, `login_count`, `last_time`, `last_ip`, `state`, `created`, `updated`) VALUES
	(1, 'admin', ' e10adc3949ba59abbe56e057f20f883e', '', 0, '2020-10-18 00:31:02', '', 0, NULL, '2020-10-18 00:31:02');
/*!40000 ALTER TABLE `te_user` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;

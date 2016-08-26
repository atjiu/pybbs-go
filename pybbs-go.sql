# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.14)
# Database: pybbs-go
# Generation Time: 2016-08-26 08:44:02 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table permission
# ------------------------------------------------------------

LOCK TABLES `permission` WRITE;
/*!40000 ALTER TABLE `permission` DISABLE KEYS */;

INSERT INTO `permission` (`id`, `pid`, `name`, `description`, `url`)
VALUES
	(1,0,'','话题节点',''),
	(2,1,'topic:add','创建话题','/topic/create'),
	(3,1,'topic:edit','编辑话题','/topic/edit/[0-9]+'),
	(4,1,'topic:delete','删除话题','/topic/delete/[0-9]+'),
	(5,0,'','回复节点',''),
	(6,5,'reply:delete','删除回复','/reply/delete/[0-9]+'),
	(7,5,'reply:save','创建回复','/reply/save'),
	(8,5,'reply:up','点赞回复','/reply/up'),
	(12,0,'','权限节点',''),
	(13,12,'user:list','用户列表','/user/list'),
	(14,12,'user:edit','角色配置','/user/edit/[0-9]+'),
	(15,12,'user:delete','用户删除','/user/delete/[0-9]+'),
	(16,12,'role:list','角色列表','/role/list'),
	(17,12,'role:add','添加角色','/role/add'),
	(18,12,'role:delete','删除角色','/role/delete/[0-9]+'),
	(20,12,'role:edit','编辑角色','/role/edit/[0-9]+'),
	(21,12,'permission:list','权限列表','/permission/list'),
	(22,12,'permission:add','添加权限','/permission/add'),
	(23,12,'permission:edit','编辑权限','/permission/edit/[0-9]+'),
	(24,12,'permission:delete','删除权限','/permission/delete/[0-9]+');

/*!40000 ALTER TABLE `permission` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table reply
# ------------------------------------------------------------

LOCK TABLES `reply` WRITE;
/*!40000 ALTER TABLE `reply` DISABLE KEYS */;

INSERT INTO `reply` (`id`, `topic_id`, `content`, `user_id`, `up`, `in_time`)
VALUES
	(1,1,'分享世界',1,0,'2016-08-26 09:22:52');

/*!40000 ALTER TABLE `reply` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table reply_up_log
# ------------------------------------------------------------



# Dump of table role
# ------------------------------------------------------------

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;

INSERT INTO `role` (`id`, `name`)
VALUES
	(3,'超级管理员'),
	(4,'版主'),
	(5,'普通用户');

/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table role_permissions
# ------------------------------------------------------------

LOCK TABLES `role_permissions` WRITE;
/*!40000 ALTER TABLE `role_permissions` DISABLE KEYS */;

INSERT INTO `role_permissions` (`id`, `role_id`, `permission_id`)
VALUES
	(47,4,3),
	(48,4,4),
	(49,4,6),
	(50,5,2),
	(51,5,7),
	(52,5,8),
	(64,3,2),
	(65,3,3),
	(66,3,4),
	(67,3,6),
	(68,3,7),
	(69,3,8),
	(70,3,13),
	(71,3,14),
	(72,3,15),
	(73,3,16),
	(74,3,17),
	(75,3,18),
	(76,3,20),
	(77,3,21),
	(78,3,22),
	(79,3,23),
	(80,3,24);

/*!40000 ALTER TABLE `role_permissions` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table section
# ------------------------------------------------------------

LOCK TABLES `section` WRITE;
/*!40000 ALTER TABLE `section` DISABLE KEYS */;

INSERT INTO `section` (`id`, `name`)
VALUES
	(1,'分享'),
	(3,'博客'),
	(4,'招聘'),
	(2,'问答');

/*!40000 ALTER TABLE `section` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table topic
# ------------------------------------------------------------

LOCK TABLES `topic` WRITE;
/*!40000 ALTER TABLE `topic` DISABLE KEYS */;

INSERT INTO `topic` (`id`, `title`, `content`, `in_time`, `user_id`, `section_id`, `view`, `reply_count`, `last_reply_user_id`, `last_reply_time`)
VALUES
	(1,'测试话题 ，hello world','你好，世界','2016-08-26 09:22:42',1,1,15,1,NULL,'2016-08-26 09:22:42');

/*!40000 ALTER TABLE `topic` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user
# ------------------------------------------------------------

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;

INSERT INTO `user` (`id`, `username`, `password`, `token`, `avatar`, `email`, `url`, `signature`, `in_time`)
VALUES
	(1,'朋也','123123','fcd1cb8e-b71f-46c3-9974-7225997b40c7','/static/imgs/avatar.png','','https://tomoya.cn','这家伙很懒，什么都没留下~','2016-08-26 09:22:16');

/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user_roles
# ------------------------------------------------------------

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;

INSERT INTO `user_roles` (`id`, `user_id`, `role_id`)
VALUES
	(5,1,3);

/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

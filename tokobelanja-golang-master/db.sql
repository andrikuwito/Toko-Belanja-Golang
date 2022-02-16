# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.5.5-10.4.21-MariaDB)
# Database: tokoonline
# Generation Time: 2022-02-15 11:37:51 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table category
# ------------------------------------------------------------

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(128) DEFAULT NULL,
  `sold_production_amount` int(11) DEFAULT 0,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;

INSERT INTO `category` (`id`, `type`, `sold_production_amount`, `created_at`, `updated_at`)
VALUES
	(2,'Makanan',0,'2022-02-15 02:03:05','2022-02-15 02:11:50'),
	(3,'Minuman',0,'2022-02-15 02:05:33','2022-02-15 02:05:33');

/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table products
# ------------------------------------------------------------

DROP TABLE IF EXISTS `products`;

CREATE TABLE `products` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(256) DEFAULT NULL,
  `price` int(11) DEFAULT NULL,
  `stock` int(11) DEFAULT NULL,
  `category_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;

INSERT INTO `products` (`id`, `title`, `price`, `stock`, `category_id`, `created_at`, `updated_at`)
VALUES
	(3,'Ayam Goreng',30000,10,1,'2022-01-23 12:35:39','2022-02-15 17:59:51'),
	(5,'Mobil-mobilan Truck',25000,90,2,'2022-02-15 02:18:47','2022-02-15 02:18:47'),
	(6,'Robot Mainan',10000,50,2,'2022-02-15 02:22:11','2022-02-15 17:51:51');

/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table transaction_histories
# ------------------------------------------------------------

DROP TABLE IF EXISTS `transaction_histories`;

CREATE TABLE `transaction_histories` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `product_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `quantity` int(11) DEFAULT NULL,
  `total_price` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `transaction_histories` WRITE;
/*!40000 ALTER TABLE `transaction_histories` DISABLE KEYS */;

INSERT INTO `transaction_histories` (`id`, `product_id`, `user_id`, `quantity`, `total_price`, `created_at`, `updated_at`)
VALUES
	(1,3,1,20,600000,'2022-01-23 12:39:37','2022-01-23 12:39:37'),
	(2,3,4,10,300000,'2022-02-11 13:00:39','2022-02-11 13:00:39'),
	(7,3,4,10,300000,'2022-02-11 13:06:33','2022-02-11 13:06:33'),
	(8,3,4,10,300000,'2022-02-11 13:18:49','2022-02-11 13:18:49'),
	(9,3,4,10,300000,'2022-02-11 13:30:10','2022-02-11 13:30:10'),
	(10,3,4,10,300000,'2022-02-11 13:31:16','2022-02-11 13:31:16'),
	(11,3,4,10,300000,'2022-02-11 13:33:33','2022-02-11 13:33:33'),
	(12,3,4,10,300000,'2022-02-11 13:35:57','2022-02-11 13:35:57'),
	(13,3,4,10,300000,'2022-02-11 13:36:17','2022-02-11 13:36:17'),
	(14,3,4,10,300000,'2022-02-11 13:38:23','2022-02-11 13:38:23'),
	(15,3,4,10,300000,'2022-02-11 13:39:59','2022-02-11 13:39:59'),
	(16,3,4,10,300000,'2022-02-15 17:59:51','2022-02-15 17:59:51');

/*!40000 ALTER TABLE `transaction_histories` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `full_name` varchar(256) DEFAULT NULL,
  `email` varchar(256) DEFAULT NULL,
  `password` varchar(256) DEFAULT NULL,
  `role` varchar(256) DEFAULT NULL,
  `balance` varchar(256) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `full_name`, `email`, `password`, `role`, `balance`, `created_at`, `updated_at`)
VALUES
	(1,'Vita','vita@gmail.com','$2a$04$ycPWwBfAR1ZC7m/Flk2TQOmsehuhdUyMwp.ZBdqk8zEKTPVDXWLIK','customer','60000','2022-01-23 11:12:12','2022-01-23 11:27:57'),
	(2,'Admin','admin@gmail.com','$2a$04$ycPWwBfAR1ZC7m/Flk2TQOmsehuhdUyMwp.ZBdqk8zEKTPVDXWLIK','admin','0','2022-01-23 11:12:12','2022-01-23 11:27:57'),
	(4,'Andri Nur Hidayatulloh','andribis13@gmail.com','$2a$04$n.CY5tK2EUDgvuYTNoCuPOu6tZSwB0SokrLXYoa9mHS01EHaBd6hS','customer','80040000','2022-01-30 07:21:19','2022-02-11 12:58:13');

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

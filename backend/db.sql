DROP TABLE IF EXISTS `peer`;
CREATE TABLE `peer` (
  `id` int NOT NULL AUTO_INCREMENT,
  `hash_id` varchar(255) NOT NULL DEFAULT '',
  `online_duration` int NOT NULL DEFAULT '0',
  `claimed` int NOT NULL DEFAULT '0',
  `network` varchar(255) NOT NULL DEFAULT '',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `address` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `hash_id` (`hash_id`,`claimed`,`network`)
) ENGINE=InnoDB AUTO_INCREMENT=1710 DEFAULT CHARSET=utf8mb4;
CREATE TABLE `tiktok_auth`(
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
    `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='JWT认证管理';
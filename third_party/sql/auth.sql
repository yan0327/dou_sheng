CREATE TABLE `tiktok_auth`(
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
    `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='JWT认证管理';

INSERT INTO tiktok_relation (user_id, follower_id,action_type ) VALUES ( 4, 5,1 );
INSERT INTO tiktok_relation (user_id, follower_id,action_type ) VALUES ( 5, 4,1 );
CREATE DATABASE IF NOT EXISTS `tiktok` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `tiktok`;
DROP TABLE IF EXISTS `tiktok_user`;
CREATE TABLE `tiktok_user`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `username`    varchar(20)        NOT NULL DEFAULT '' COMMENT '账号',
    `password`    varchar(64)        NOT NULL DEFAULT '' COMMENT '密码',
    PRIMARY KEY (`id`),
    UNIQUE INDEX username (`username`) USING BTREE 
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';


DROP TABLE IF EXISTS `tiktok_relation`;
CREATE TABLE `tiktok_relation`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id`     bigint(20) unsigned NOT NULL  COMMENT '关注ID',
    `follower_id` bigint(20) unsigned NOT NULL  COMMENT '被关注ID',
    INDEX user_id (`user_id`) USING BTREE,
    INDEX follower_id (`follower_id`) USING BTREE,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='关系表';

DROP TABLE IF EXISTS `tiktok_video`;
CREATE TABLE `tiktok_video`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '视频id',
    `title`       varchar(191) NOT NULL COMMENT '视频标题',
    `author_id`   bigint(20) unsigned NOT NULL COMMENT '作者id',
    `play_url`    varchar(191) DEFAULT NULL COMMENT '播放视频路径',
    `cover_url`   varchar(191) DEFAULT NULL COMMENT '封面图片路径',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    INDEX author_id (`author_id`)
) ENGINE = InnoDB
DEFAULT CHARSET = utf8mb4 COMMENT ='视频表';

  DROP TABLE IF EXISTS `tiktok_video_like`;
CREATE TABLE `tiktok_video_like`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
     `user_id`   bigint(20) unsigned NOT NULL  COMMENT '用户id',
    `video_id`   bigint(20) unsigned NOT NULL  COMMENT '视频id',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `user_video_rel` (`user_id`, `video_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='点赞表';

  DROP TABLE IF EXISTS `tiktok_video_comment`;
CREATE TABLE `tiktok_video_comment`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `video_id`   bigint(20) unsigned NOT NULL  COMMENT '视频id',
     `user_id`   bigint(20) unsigned NOT NULL  COMMENT '用户id',
    `content`  text                NOT NULL COMMENT '评论内容',
    `create_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    INDEX video_id (`video_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='评论表';



             /*
    `follow_count`       bigint(20) unsigned        NOT NULL DEFAULT 0 COMMENT '关注',
    `FollowerCount`       bigint(20) unsigned        NOT NULL DEFAULT 0 COMMENT '粉丝',
    `IsFollow`      tinyint(1)              NOT NULL DEFAULT 0 COMMENT '是否关注',
    */
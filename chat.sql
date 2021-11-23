CREATE DATABASE chat;

DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uuid` varchar(150) NOT NULL COMMENT 'uuid',
    `username` varchar(191) NOT NULL COMMENT '''用户名''',
    `nickname` varchar(255) DEFAULT NULL COMMENT '昵称',
    `email` varchar(80) DEFAULT NULL COMMENT '邮箱',
    `password` varchar(150) NOT NULL COMMENT '密码',
    `avatar` varchar(250) NOT NULL COMMENT '头像',
    `create_at` datetime(3) DEFAULT NULL,
    `update_at` datetime(3) DEFAULT NULL,
    `delete_at` bigint DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`),
    UNIQUE KEY `idx_uuid` (`uuid`),
    UNIQUE KEY `username_2` (`username`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '用户表';


DROP TABLE IF EXISTS `user_friends`;
CREATE TABLE IF NOT EXISTS `user_friends` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
    `deleted_at` bigint unsigned DEFAULT NULL COMMENT '删除时间戳',
    `user_id` int DEFAULT NULL COMMENT '用户ID',
    `friend_id` int DEFAULT NULL COMMENT '好友ID',
    PRIMARY KEY (`id`),
    KEY `idx_user_friends_user_id` (`user_id`),
    KEY `idx_user_friends_friend_id` (`friend_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '好友信息表';


DROP TABLE IF EXISTS `messages`;
CREATE TABLE IF NOT EXISTS `messages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` bigint unsigned DEFAULT NULL COMMENT '删除时间戳',
  `from_user_id` int DEFAULT NULL COMMENT '发送人ID',
  `to_user_id` int DEFAULT NULL COMMENT '发送对象ID',
  `content` varchar(2500) DEFAULT NULL COMMENT '消息内容',
  `url` varchar(350) DEFAULT NULL COMMENT '''文件或者图片地址''',
  `pic` text COMMENT '缩略图',
  `message_type` smallint DEFAULT NULL COMMENT '''消息类型：1单聊，2群聊''',
  `content_type` smallint DEFAULT NULL COMMENT '''消息内容类型：1文字，2语音，3视频''',
  PRIMARY KEY (`id`),
  KEY `idx_messages_deleted_at` (`deleted_at`),
  KEY `idx_messages_from_user_id` (`from_user_id`),
  KEY `idx_messages_to_user_id` (`to_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT '消息表';


DROP TABLE IF EXISTS `groups`;
CREATE TABLE IF NOT EXISTS  `groups` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` bigint unsigned DEFAULT NULL,
  `user_id` int DEFAULT NULL COMMENT '''群主ID''',
  `name` varchar(150) DEFAULT NULL COMMENT '''群名称',
  `notice` varchar(350) DEFAULT NULL COMMENT '''群公告',
  `uuid` varchar(150) NOT NULL COMMENT '''uuid''',
  PRIMARY KEY (`id`),
  KEY `idx_groups_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT '群组表';


DROP TABLE IF EXISTS `group_members`;
CREATE TABLE  IF NOT EXISTS `group_members` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` bigint unsigned DEFAULT NULL,
  `user_id` int DEFAULT NULL COMMENT '''用户ID''',
  `group_id` int DEFAULT NULL COMMENT '''群组ID''',
  `nickname` varchar(350) DEFAULT NULL COMMENT '''昵称',
  `mute` smallint DEFAULT NULL COMMENT '''是否禁言''',
  PRIMARY KEY (`id`),
  KEY `idx_group_members_user_id` (`user_id`),
  KEY `idx_group_members_group_id` (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT '群组成员表';

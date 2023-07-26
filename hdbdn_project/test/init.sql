-- 用户信息表
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `username`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Username',
  `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
  `name`   varchar(128) NOT NULL DEFAULT 'xiaodouyin' COMMENT 'Name',
  
  `follow_count` bigint NOT NULL DEFAULT 0,
  `follower_count` bigint NOT NULL DEFAULT 0,
  `is_follow`      boolean  NOT NULL DEFAULT 1 COMMENT 'IsFollow',
  `avatar`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Avatar',
  `background_image`   varchar(128) NOT NULL DEFAULT '' COMMENT 'BackgroundImage',
  `signature`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Signature',
  `total_favorited`   varchar(128) NOT NULL DEFAULT '' COMMENT 'TotalFavorited',
  `work_count` bigint NOT NULL DEFAULT 0,
  `favorite_count` bigint NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
  PRIMARY KEY (`id`),
  KEY          `idx_username` (`username`) COMMENT 'Username index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';
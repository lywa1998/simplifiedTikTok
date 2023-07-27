-- 用户信息表
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'PrimaryKey', -- 用户id，设为主键
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT 'Username', -- 用户名称
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT 'Password', -- 用户密码
  `follow_count` bigint NOT NULL DEFAULT 0 COMMENT 'FollowCount', -- 关注总数
  `follower_count` bigint NOT NULL DEFAULT 0 COMMENT 'FollowerCount', -- 粉丝总数
  `is_follow` boolean  NOT NULL DEFAULT 1 COMMENT 'IsFollow', -- true-已关注，false-未关注
  `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT 'Avatar', -- 用户头像
  `background_image` varchar(128) NOT NULL DEFAULT '' COMMENT 'BackgroundImage',  -- 用户个人页顶部大图
  `signature` varchar(256) NOT NULL DEFAULT '' COMMENT 'Signature',  -- 个人简介
  `total_favorited` bigint NOT NULL DEFAULT '' COMMENT 'TotalFavorited',  -- 获赞总数
  `work_count` bigint NOT NULL DEFAULT 0 COMMENT 'WorkCount',  -- 作品数量
  `favorite_count` bigint NOT NULL DEFAULT 0 COMMENT 'FavoriteCount',  -- 点赞数量,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
  PRIMARY KEY (`id`),
  KEY          `user_username` (`username`) COMMENT 'Username index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

CREATE TABLE User (
    id BIGINT NOT NULL PRIMARY KEY, -- 用户id，设为主键
    name VARCHAR(255) NOT NULL, -- 用户名称
    follow_count BIGINT, -- 关注总数
    follower_count BIGINT, -- 粉丝总数
    is_follow TINYINT(1) NOT NULL, -- true-已关注，false-未关注
    avatar VARCHAR(255), -- 用户头像
    background_image VARCHAR(255), -- 用户个人页顶部大图
    signature TEXT, -- 个人简介，使用TEXT类型，因为可能包含较长的文本内容
    total_favorited BIGINT, -- 获赞数量
    work_count BIGINT, -- 作品数量
    favorite_count BIGINT -- 点赞数量
);
CREATE TABLE `message` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'PrimaryKey',
    `to_user_id` bigint NOT NULL COMMENT 'ReceiverID',
    `from_user_id` bigint NOT NULL COMMENT 'SenderID',
    `content` text NOT NULL COMMENT 'MessageContent',
    `created_at` datetime(6) NOT NULL COMMENT 'create time',
    PRIMARY KEY(`id`)
    KEY `created_time` (`created_at`) USING BTREE,
    KEY `pair_user_id` (`to_user_id`, `from_user_id`) USING BTREE,
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video table';

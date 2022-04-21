CREATE TABLE `user_info`(
    `user_id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT 'user id',
    `user_name` VARCHAR(128) NOT NULL COMMENT '用户名字',
    `password` VARCHAR(256) NOT NULL COMMENT '用户密码',
    `email` VARCHAR(256) NOT NULL COMMENT '用户邮箱',
    `root_node_id`BIGINT UNSIGNED NOT NULL COMMENT '用户网盘根节点id', 
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`user_id`),
    UNIQUE KEY `unique_name`(`user_name`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';

CREATE TABLE `node`(
    `id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '主键id',
    `node_id` BIGINT UNSIGNED NOT NULL COMMENT '节点id,分享文件使用,64bit',
    `node_type` TINYINT UNSIGNED NOT NULL COMMENT '节点类型',
    `name` VARCHAR(1024) NOT NULL COMMENT '文件名称',
    `content` BLOB NOT NULL COMMENT '文件内容,最大64k',
    `cos_key` VARCHAR(256) NOT NULL COMMENT '文件在cos中的key',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    INDEX `idx_node_id`(`node_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='节点表';

CREATE TABLE `search_index`(
    `index_id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '索引id',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT 'user id',
    `keyword` VARCHAR(128) NOT NULL COMMENT '关键词',
    `file_id` BIGINT UNSIGNED NOT NULL COMMENT '文件id',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`node_id`),
    INDEX `idx_uid_kw`(`uid`,`keyword`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='搜索索引表';

CREATE TABLE `node_rel`(
    `id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '主键id',
    `parent_id` BIGINT UNSIGNED NOT NULL COMMENT '父节点id',
    `child_id` BIGINT UNSIGNED NOT NULL COMMENT '子节点id',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    INDEX `idx_pid`(`parent_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文件节点引用表';
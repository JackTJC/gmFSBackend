-- 用户信息
CREATE TABLE `user_info`(
    `user_id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT 'user id',
    `user_name` VARCHAR(128) NOT NULL COMMENT '用户名字',
    `password` VARCHAR(256) NOT NULL COMMENT '用户密码',
    `email` VARCHAR(256) NOT NULL COMMENT '用户邮箱',
    `root_node_id` BIGINT UNSIGNED NOT NULL COMMENT '用户网盘根节点id',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`user_id`),
    UNIQUE KEY `unique_name`(`user_name`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户信息表';

-- 文件夹节点
CREATE TABLE `node`(
    `id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '主键id',
    `node_id` BIGINT UNSIGNED NOT NULL COMMENT '节点id,分享文件使用,64bit',
    `node_type` TINYINT UNSIGNED NOT NULL COMMENT '节点类型',
    `name` VARCHAR(1024) NOT NULL COMMENT '文件名称',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    INDEX `idx_node_id`(`node_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '节点表';

-- 文件引用关系
CREATE TABLE `node_rel`(
    `id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '主键id',
    `parent_id` BIGINT UNSIGNED NOT NULL COMMENT '父节点id',
    `child_id` BIGINT UNSIGNED NOT NULL COMMENT '子节点id',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    UNIQUE KEY `uniq_parent_child`(`parent_id`, `child_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文件节点引用表';

-- 密钥
CREATE TABLE `secret_key`(
    `id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '主键id',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户id',
    `file_id` BIGINT UNSIGNED NOT NULL COMMENT '文件id',
    `key` BLOB NOT NULL COMMENT '解密密钥,已加密',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    UNIQUE KEY `uniq_uid_fid` (`user_id`,`file_id`)   
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '解密密钥表';

-- 加密索引
CREATE TABLE `search_index`(
    `id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '索引id',
    `keyword` VARCHAR(512) NOT NULL COMMENT '关键词,加密后的',
    `file_id` BIGINT UNSIGNED NOT NULL COMMENT '文件id',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    INDEX `idx_fid`(`file_id`),
    INDEX `idx_kw`(`keyword`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '搜索索引表';

-- 加密索引
CREATE TABLE `share_file`(
    `id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '主键id',
    `from` BIGINT UNSIGNED NOT NULL COMMENT '',
    `to` BIGINT UNSIGNED NOT NULL COMMENT '',
    `file_id` BIGINT UNSIGNED NOT NULL COMMENT '文件id',
    `key` BLOB NOT NULL COMMENT '解密秘钥',
    `status` TINYINT NOT NULL COMMENT '分享是否被处理,1:未处理,2:已处理',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    INDEX `idx_to`(`to`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '分享文件表';


-- for linzefu
CREATE TABLE `file_index`(
    `id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '主键id',
    `file_id` BIGINT UNSIGNED NOT NULL COMMENT '文件id',
    `keyword` VARCHAR(256) NOT NULL COMMENT '关键词W',
    `L` VARCHAR(256) NOT NULL COMMENT '',
    `I_w` VARCHAR(256) NOT NULL COMMENT '',
    `R_w` VARCHAR(256) NOT NULL COMMENT '',
    `C_w` VARCHAR(256) NOT NULL COMMENT '',
    `I_id` VARCHAR(256) NOT NULL COMMENT '',
    `R_id` VARCHAR(256) NOT NULL COMMENT '',
    `C_id` VARCHAR(256) NOT NULL COMMENT '',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    INDEX `idx_L`(`L`),
    INDEX `idx_kw`(`keyword`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文件节点引用表';
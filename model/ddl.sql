-- 用户信息
CREATE TABLE `user_info`(
    `user_id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT 'user id',
    `user_name` VARCHAR(128) NOT NULL COMMENT '用户名字',
    `password` VARCHAR(256) NOT NULL COMMENT '用户密码',
    `email` VARCHAR(256) NOT NULL COMMENT '用户邮箱',
    PRIMARY KEY(`user_id`),
    INDEX `idx_user_name`(`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';

-- 文件节点表
CREATE TABLE `node`(
    `node_id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '节点id',
    `node_type` TINYINT NOT NULL COMMENT '节点类型',
    `node_name` VARCHAR(2048) NOT NULL COMMENT '节点名称',
    `obj_key` VARCHAR(1024) NOT NULL COMMENT '对象在cos中的key',
    `content` BLOB NOT NULL COMMENT '文件内容,最大64k',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`node_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文件节点表';

-- 文件引用关系表
CREATE TABLE `node_rel`(
    `id` BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '主键id',
    `parent_node_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '父节点id',
    `child_node_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '子节点id',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    INDEX `idx_parent_node_id`(`parent_node_id`),
    INDEX `idx_sub_node_id`(`sub_node_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文件节点引用表';
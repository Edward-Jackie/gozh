DROP TABLE IF EXISTS `setting_dict`;
CREATE TABLE `setting_dict`
(
    `id`   bigint(20) NOT NULL,
    `code` longtext,
    `name` longtext,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for setting_dict_values
-- ----------------------------
DROP TABLE IF EXISTS `setting_dict_values`;
CREATE TABLE `setting_dict_values`
(
    `id`      bigint(11) unsigned NOT NULL COMMENT 'ID',
    `dict_id` bigint(11) NOT NULL COMMENT '字典ID',
    `key`     varchar(50)  NOT NULL COMMENT '字典值Key',
    `value`   varchar(100) NOT NULL COMMENT '字典值Value',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='字典值';

-- ----------------------------
-- Table structure for setting_menu
-- ----------------------------
DROP TABLE IF EXISTS `setting_menu`;
CREATE TABLE `setting_menu`
(
    `id`               bigint(20) NOT NULL,
    `name`             longtext,
    `path`             longtext,
    `type`             longtext,
    `title`            longtext,
    `icon`             longtext,
    `component`        longtext,
    `parent_id`        bigint(20) DEFAULT NULL,
    `hidden`           bigint(20) DEFAULT NULL,
    `hide_bread_crumb` bigint(20) DEFAULT NULL,
    `sort`             bigint(20) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for system_role
-- ----------------------------
DROP TABLE IF EXISTS `system_role`;
CREATE TABLE `system_role`
(
    `id`          bigint(20) NOT NULL COMMENT 'ID',
    `code`        varchar(50)  DEFAULT NULL COMMENT '角色编码',
    `name`        varchar(100) NOT NULL COMMENT '角色名称',
    `comment`     varchar(255) DEFAULT NULL COMMENT '备注',
    `created_by`  bigint(20) DEFAULT NULL COMMENT '创建者',
    `updated_by`  bigint(20) DEFAULT NULL COMMENT '编辑者',
    `created_at`  datetime     DEFAULT NULL COMMENT '创建时间',
    `updated_at`  datetime     DEFAULT NULL COMMENT '更新时间',
    `deleted_at`  datetime     DEFAULT NULL COMMENT '删除时间',
    `permissions` text COMMENT '菜单权限',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色信息';

-- ----------------------------
-- Table structure for system_user
-- ----------------------------
DROP TABLE IF EXISTS `system_user`;
CREATE TABLE `system_user`
(
    `id`              bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'Id',
    `user_name`       varchar(100) DEFAULT NULL COMMENT '用户名',
    `nick_name`       varchar(100) DEFAULT NULL COMMENT '姓名',
    `password`        longtext COMMENT '密码',
    `phone`           varchar(20)  DEFAULT NULL COMMENT '手机号',
    `register_time`   datetime     DEFAULT NULL COMMENT '注册时间',
    `last_login_time` datetime     DEFAULT NULL COMMENT '上次登录时间',
    `last_login_ip`   longtext COMMENT '登录IP',
    `created_by`      bigint(20) DEFAULT NULL COMMENT '创建者',
    `updated_by`      bigint(20) DEFAULT NULL COMMENT '编辑者',
    `created_at`      datetime     DEFAULT NULL COMMENT '创建时间',
    `updated_at`      datetime     DEFAULT NULL COMMENT '编辑时间',
    `deleted_at`      datetime     DEFAULT NULL COMMENT '删除时间',
    `role_ids`        text COMMENT '角色',
    PRIMARY KEY (`id`),
    UNIQUE KEY `phone` (`phone`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for video_info
-- ----------------------------
DROP TABLE IF EXISTS `video_info`;
CREATE TABLE `video_info`
(
    `id`         bigint(20) NOT NULL COMMENT 'ID',
    `name`       varchar(255) NOT NULL COMMENT '视频名称',
    `comment`    text COMMENT '备注',
    `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
    `updated_by` bigint(20) DEFAULT NULL COMMENT '编辑者',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '编辑时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频';

-- ----------------------------
-- Table structure for video_info_item
-- ----------------------------
DROP TABLE IF EXISTS `video_info_item`;
CREATE TABLE `video_info_item`
(
    `id`         bigint(20) DEFAULT NULL COMMENT 'ID',
    `video_id`   bigint(20) DEFAULT NULL COMMENT '视频信息ID',
    `etag`       varchar(100) DEFAULT NULL COMMENT 'etag',
    `url`        varchar(255) DEFAULT NULL COMMENT '视频地址',
    `poster`     varchar(255) DEFAULT NULL COMMENT '缩略图',
    `created_at` datetime     DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime     DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime     DEFAULT NULL COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频详情';


INSERT INTO `system_user` (`id`, `user_name`, `nick_name`, `password`, `phone`, `register_time`, `last_login_time`,
                           `last_login_ip`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`,
                           `role_ids`)
VALUES (1595654112953765888, 'admin', '管理员',
        'h/Il9aWiQPzCg9WpYjSubIgfbx2BUC+W+4+EOzBfBeSftwMjPTUQ1PmG7iF3NEwIbJFwxwbA+A/TRmbfY7/44mE=.t/xXhocYqeFdD2kSELezlLp1WuEkXILJLv69hf5CnDU=',
        '', '2022-11-24 13:42:41', '2022-11-24 16:38:59', '127.0.0.1', 1, 1, '2022-11-24 13:42:41',
        '2022-11-24 16:38:59', NULL, '[]');


CREATE TABLE `system_event`
(
    `id`          bigint(20) NOT NULL COMMENT 'ID',
    `event`       varchar(100) NOT NULL COMMENT '埋点事件',
    `ip`          varchar(80)  DEFAULT NULL COMMENT '访问IP',
    `region`      varchar(100) DEFAULT NULL COMMENT 'ip对应的区域',
    `city`        varchar(100) DEFAULT NULL,
    `relation_id` bigint(20)  DEFAULT 0 COMMENT '关联ID',
    `event_time`  datetime     NOT NULL COMMENT '时间',
    `created_at`  datetime     DEFAULT NULL COMMENT '创建时间',
    `updated_at`  datetime     DEFAULT NULL COMMENT '更新时间',
    `deleted_at`  datetime     DEFAULT NULL COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `system_logger` (
                                 `id` bigint(20) NOT NULL COMMENT 'ID',
                                 `user_id` bigint(20) NOT NULL COMMENT '用户ID',
                                 `nick_name` varchar(100) DEFAULT NULL COMMENT '用户昵称',
                                 `ip` varchar(80) DEFAULT NULL COMMENT '访问IP',
                                 `comment` text COMMENT '日志说明',
                                 `logger_time` datetime DEFAULT NULL COMMENT '日志时间',
                                 `logger_type` int(11) DEFAULT NULL COMMENT '类型 1-登录 2-登出 3-用户 4-角色 5-视频',
                                 `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                                 `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
                                 `deleted_at` datetime DEFAULT NULL COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
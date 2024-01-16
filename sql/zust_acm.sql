use
zust_acm;

DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '序号',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
    `username`   varchar(20)  NOT NULL UNIQUE COMMENT '账号',
    `password`   varchar(256) NOT NULL DEFAULT '' COMMENT '密码，加密后',
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  DEFAULT CHARSET = UTF8 COMMENT '用户鉴权表';

DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`
(
    `id`              bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '序号',
    `created_at`      datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`      datetime NULL DEFAULT NULL COMMENT '删除时间',
    `username`        varchar(20)  NOT NULL UNIQUE COMMENT '账号',
    `email`           varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱',
    `avatar`          varchar(512) NOT NULL DEFAULT '' COMMENT '头像',
    `name`            varchar(32)  NOT NULL DEFAULT '' COMMENT 'english name',
    `cname`           varchar(32)  NOT NULL DEFAULT '' COMMENT 'chinese name',
    `cf_id`           varchar(255) NOT NULL DEFAULT '' COMMENT 'codeforces用户名',
    `cf_rating`       int          NOT NULL DEFAULT 0 COMMENT 'codeforces rating',
    `cf_rank`         varchar(32)  NOT NULL DEFAULT '' COMMENT 'codeforces rank',
    `atc_id`          varchar(255) NOT NULL DEFAULT '' COMMENT 'atcoder用户名',
    `nowcoder_id`     varchar(255) NOT NULL DEFAULT '' COMMENT '牛客网id',
    `enrollment_year` int          NOT NULL DEFAULT 0 COMMENT '入学年份',
    `role_id`         bigint unsigned NOT NULL DEFAULT 0 COMMENT '所属角色',
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  DEFAULT CHARSET = UTF8 COMMENT '用户信息表';

DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '序号',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
    `name`       varchar(255) NOT NULL DEFAULT '' COMMENT '名称',
    `code`       varchar(255) NOT NULL DEFAULT '' COMMENT '编码',
    `type`       varchar(255) NOT NULL DEFAULT '' COMMENT '类型',
    `parent_id`  bigint unsigned NOT NULL DEFAULT 0 COMMENT '父节点id',
    `order`      int(10) NOT NULL DEFAULT 0 COMMENT '排序',
    `icon`       varchar(512) NOT NULL DEFAULT '' COMMENT '菜单图标',
    `component`  varchar(512) NOT NULL DEFAULT '' COMMENT '组件路径',
    `path`       varchar(512) NOT NULL DEFAULT '' COMMENT '路由地址',
    `is_show`    BOOLEAN      NOT NULL DEFAULT FALSE COMMENT '是否显示',
    `is_enable`  BOOLEAN      NOT NULL DEFAULT FALSE COMMENT '是否启用',
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  DEFAULT CHARSET = UTF8 COMMENT '资源表';

DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '序号',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
    `code`       varchar(255) NOT NULL DEFAULT '' UNIQUE COMMENT '角色编码',
    `name`       varchar(255) NOT NULL DEFAULT '' UNIQUE COMMENT '角色名',
    `is_enable`  bool         NOT NULL DEFAULT FALSE COMMENT '启用状态:0-禁用;1-启用',
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  DEFAULT CHARSET = UTF8 COMMENT '角色表';

DROP TABLE IF EXISTS `role_resource_rel`;
CREATE TABLE `role_resource_rel`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '序号',
    `created_at`  datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`  datetime NULL DEFAULT NULL COMMENT '删除时间',
    `role_id`     bigint unsigned NOT NULL DEFAULT 0 COMMENT '角色id',
    `resource_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '资源id',
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  DEFAULT CHARSET = UTF8 COMMENT '角色-资源表';

DROP TABLE IF EXISTS `competition`;
CREATE TABLE `competition`
(
    `id`               bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '序号',
    `created_at`       datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`       datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`       datetime NULL DEFAULT NULL COMMENT '删除时间',
    `name`             varchar(512) NOT NULL DEFAULT '' COMMENT '比赛名称',
    `type`             varchar(512) NOT NULL DEFAULT '' COMMENT '比赛类型',
    `season_year`      int          NOT NULL DEFAULT 0 COMMENT '赛季年',
    `start_time`       datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '开始时间',
    `end_time`         datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '结束时间',
    `official_website` varchar(512) NOT NULL DEFAULT '' COMMENT '比赛官网',
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  DEFAULT CHARSET = UTF8 COMMENT '比赛信息表';

DROP TABLE IF EXISTS `record`;
CREATE TABLE `record`
(
    `id`             bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '序号',
    `created_at`     datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`     datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`     datetime NULL DEFAULT NULL COMMENT '删除时间',
    `competition_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '',
    `type`           varchar(512) NOT NULL DEFAULT '' COMMENT '奖项类型',
    `remark`         varchar(512) NOT NULL DEFAULT '' COMMENT '',
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  DEFAULT CHARSET = UTF8 COMMENT '奖项信息表';


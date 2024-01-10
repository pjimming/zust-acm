DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth`
(
    `id`          bigint unsigned     NOT NULL AUTO_INCREMENT COMMENT '序号',
    `gmt_created` datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `gmt_updated` datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_delete`   tinyint(1) unsigned NOT NULL DEFAULT 0 COMMENT '是否删除，0-未删除，1-删除，默认为0',
    `username`    varchar(20)         NOT NULL UNIQUE COMMENT '账号',
    `password`    varchar(256)        NOT NULL DEFAULT '' COMMENT '密码，加密后',
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  DEFAULT CHARSET = UTF8 COMMENT '用户鉴权表';

DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`
(
    `id`          bigint unsigned     NOT NULL AUTO_INCREMENT COMMENT '序号',
    `gmt_created` datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `gmt_updated` datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_delete`   tinyint(1) unsigned NOT NULL DEFAULT 0 COMMENT '是否删除，0-未删除，1-删除，默认为0',
    `username`    varchar(20)         NOT NULL UNIQUE COMMENT '账号',
    `email`       varchar(128)        NOT NULL DEFAULT '' COMMENT '邮箱',
    `title_photo` varchar(512)        NOT NULL DEFAULT '' COMMENT '头像',
    `name`        varchar(32)         NOT NULL DEFAULT '' COMMENT 'english name',
    `cname`       varchar(32)         NOT NULL DEFAULT '' COMMENT 'chinese name',
    `cf_id`       varchar(255)        NOT NULL DEFAULT '' COMMENT 'codeforces用户名',
    `cf_rating`   int                 NOT NULL DEFAULT 0 COMMENT 'codeforces rating',
    `cf_rank`     varchar(32)         NOT NULL DEFAULT '' COMMENT 'codeforces rank',
    `atc_id`      varchar(255)        NOT NULL DEFAULT '' COMMENT 'atcoder用户名',
    `nowcoder_id` varchar(255)        NOT NULL DEFAULT '' COMMENT '牛客网id',
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  DEFAULT CHARSET = UTF8 COMMENT '用户信息表';
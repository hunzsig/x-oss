DROP TABLE IF EXISTS `log`;
CREATE TABLE `log`
(
    `id`          BIGINT    NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_token`  CHAR(255) NOT NULL COMMENT '用户token',
    `msg`         VARCHAR(1024) COMMENT '记录信息',
    `create_time` DATETIME  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = INNODB COMMENT '系统配置';


DROP TABLE IF EXISTS `files`;
CREATE TABLE `files`
(
    `hash`           CHAR(255)                 NOT NULL COMMENT 'hash（这个哈希由二进制数据生成，代表文件唯一）',
    `key`            CHAR(255)                 NOT NULL COMMENT '文件key（这个key是生成的，用于访问）',
    `user_token`     CHAR(255)                 NOT NULL COMMENT '用户token',
    `name`           CHAR(255)                 NOT NULL COMMENT '文件名',
    `md5_name`       CHAR(255)                 NOT NULL COMMENT 'md5名字（这是sha1生成的md5名字）',
    `suffix`         CHAR(255)                 NOT NULL COMMENT '文件后缀',
    `size`           BIGINT UNSIGNED DEFAULT 0 NOT NULL COMMENT '文件大小',
    `content_type`   CHAR(255)                 NOT NULL COMMENT '内容类型',
    `path`           CHAR(255)                 NOT NULL COMMENT '存在路径',
    `uri`            CHAR(255)                 NOT NULL COMMENT 'URI',
    `from_url`       CHAR(255)                 NOT NULL COMMENT '来源地址',
    `call_qty`       INT UNSIGNED    DEFAULT 0 NOT NULL COMMENT '调用次数',
    `call_last_time` DATETIME                  NOT NULL COMMENT '最后一次调用时间',
    `create_time`    DATETIME                  NOT NULL COMMENT '创建日期',
    `update_time`    DATETIME                  NOT NULL COMMENT '更新日期',
    PRIMARY KEY (`hash`),
    INDEX (`key`),
    INDEX (`user_token`),
    INDEX (`name`),
    INDEX (`content_type`)
) ENGINE = INNODB COMMENT '文件表';


DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `token`  CHAR(255) NOT NULL COMMENT '用户唯一token key',
    `status` TINYINT   NOT NULL DEFAULT 1 COMMENT '状态 -1无效 1有效',
    PRIMARY KEY (`token`),
    INDEX (`status`)
) ENGINE = INNODB COMMENT 'users tokens';


INSERT INTO `users` (`token`, `status`)
VALUES ('hunzsig', 1);
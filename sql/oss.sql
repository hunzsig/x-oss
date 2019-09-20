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
    `token`            CHAR(255)     NOT NULL COMMENT '用户唯一token key',
    `status`           TINYINT       NOT NULL DEFAULT 1 COMMENT '状态 -1无效 1有效',
    `level`            INT UNSIGNED  NOT NULL DEFAULT 1 COMMENT '等级',
    `exp`              INT UNSIGNED  NOT NULL DEFAULT 1 COMMENT '经验',
    `allow_size_unit`  CHAR(255)     NOT NULL COMMENT '文件大小单位',
    `allow_size_total` BIGINT        NOT NULL COMMENT '允许储存的总体文件大小',
    `allow_size_one`   BIGINT        NOT NULL COMMENT '允许储存的单个文件大小',
    `allow_qty`        BIGINT        NOT NULL COMMENT '允许储存的文件数',
    `allow_suffix`     VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '允许的文件后缀',
    PRIMARY KEY (`token`),
    INDEX (`status`),
    INDEX (`level`)
) ENGINE = INNODB COMMENT '资源等级参数';


DROP TABLE IF EXISTS `setting`;
CREATE TABLE `setting`
(
    `key`  CHAR(255)     NOT NULL COMMENT 'key',
    `data` varchar(1024) NOT NULL DEFAULT '' COMMENT '数据',
    PRIMARY KEY (`key`)
) ENGINE = INNODB COMMENT '系统配置';



INSERT INTO `users` (`token`, `status`, `level`, `exp`, `allow_size_unit`, `allow_size_total`, `allow_size_one`,
                     `allow_qty`, `allow_suffix`)
VALUES ('hunzsig', '1', '999', '0', 'TB', '99999', '99999', '99999',
        ',,,,,mp4,rmvb,mkv,avi,wmv,mov,mpg,3gp,mp3,wma,ape,flac,wav,ogg,m4a,gif,jpeg,jpg,bmp,png,ico,tga,txt,pdf,doc,docx,ppt,pptx,xls,xlsx,csv,psd,rar,7z,zip,iso,cso');

INSERT INTO `setting` (`key`, `data`)
VALUES ('size_unit', ',,,,,KB,MB,GB,TB');

INSERT INTO `setting` (`key`, `data`)
VALUES ('file_suffix',
        ',,,,,mp4,rmvb,mkv,avi,wmv,mov,mpg,3gp,mp3,wma,ape,flac,wav,ogg,m4a,gif,jpeg,jpg,bmp,png,ico,tga,txt,pdf,doc,docx,ppt,pptx,xls,xlsx,csv,psd,rar,7z,zip,iso,cso');

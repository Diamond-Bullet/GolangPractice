create table book (
    `id`         bigint(20) unsigned not null auto_increment comment '书本ID',
    `book_name`  varchar(64) not null default '' comment '书本名',
    `book_pages` int(11) not null default 0 comment '书本页数',
    `type`       tinyint(4) unsigned NOT NULL default 0 COMMENT '0普通 1搞笑',

    primary key (`id`),
    unique key `idx_book_name` (`book_name`),
    key          `created_on` (`created_on`)
) engine=InnoDB default charset=utf8mb4 comment='书本表';

CREATE TABLE `friends` (
    `id`         bigint(20) unsigned NOT NULL auto_increment,
    `friend1`    bigint(20) unsigned NOT NULL default 0,
    `friend2`    bigint(20) unsigned NOT NULL default 0,

    PRIMARY KEY (`id`),
    unique key idx_friend1_friend2 (friend1,friend2) --联合索引
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


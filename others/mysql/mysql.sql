create table book (
    `id`         bigint(20) unsigned not null auto_increment comment 'book ID',
    `book_name`  varchar(64) not null default '' comment 'book name',
    `book_pages` int(11) not null default 0 comment 'page amount',
    `type`       tinyint(4) unsigned not null default 0 comment '0:normal 1:funny',
    `created_time` timestamp not null default CURRENT_TIMESTAMP comment 'created time',
    `friend1`    bigint(20) unsigned NOT NULL default 0,
    `friend2`    bigint(20) unsigned NOT NULL default 0,
    `data`       json,

    primary key (`id`),
    unique key `idx_book_name` (`book_name`),
    unique key idx_friend1_friend2 (friend1, friend2) #联合索引
) engine=InnoDB auto_increment=1 default charset=utf8mb4 comment='book table';

CREATE DATABASE database_name CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

SHOW TABLES FROM database_name;
SHOW CREATE DATABASE database_name;

SHOW CREATE TABLE book;
SHOW TABLE STATUS LIKE '%table_name%';
SHOW INDEXES FROM book;

# calculate the in-disk size of table
SELECT table_name                                             `Table`,
       round(((data_length + index_length) / 1024 / 1024), 2) `Size in MB`
FROM information_schema.TABLES
WHERE table_schema = 'database'
  AND table_name = 'table';


####################### Procedure #################
# insert records in batches.
# for now, it doesn't work on mariadb.
CREATE FUNCTION `rand_string`(n INT) RETURNS varchar(255) CHARSET latin1
BEGIN
    DECLARE chars_str varchar(100) DEFAULT 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789';
    DECLARE return_str varchar(255) DEFAULT '';
    DECLARE i INT DEFAULT 0;
    WHILE i < n DO
        SET return_str = concat(return_str, substring(chars_str, FLOOR(1 + RAND() * 62), 1));
        SET i = i + 1;
    END WHILE;
    RETURN return_str;
END;

CREATE PROCEDURE `add_data`(IN n int)
BEGIN
    DECLARE i INT DEFAULT 1;
    WHILE i <= n DO
        INSERT into book (book_name, book_pages, created_time) VALUES (rand_string(20),FLOOR(RAND() * 1000), now());
        set i=i+1;
    END WHILE;
END;

CALL add_data(100);


####################### Trifle #################
# `;`, `\G`
select * from book limit 1;
select * from book limit 1\G
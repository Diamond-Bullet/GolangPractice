####################### CRUD #################
create table book (
    `id`            bigint(20) unsigned not null auto_increment comment 'book ID',
    `book_name`     varchar(64) not null default '' comment 'book name',
    `book_pages`    int(11) not null default 0 comment 'page amount',
    `type`          tinyint(4) unsigned not null default 0 comment '0:normal 1:funny',
    `created_time`  timestamp not null default CURRENT_TIMESTAMP comment 'created time',
    `fan1`          bigint(20) unsigned NOT NULL default 0,
    `fan2`          bigint(20) unsigned NOT NULL default 0,
    `data`          json,

    primary key (`id`),
    unique key `idx_book_name` (`book_name`),
    unique key idx_fan1_fan2 (fan1, fan2) #联合索引
) engine=InnoDB auto_increment=1 default charset=utf8mb4 comment='book table';

SHOW CREATE TABLE book;

DROP TABLE book;

CREATE DATABASE playground CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
create database if not exists playground default charset utf8mb4 collate utf8mb4_bin;

SHOW TABLES FROM playground;
SHOW CREATE DATABASE playground;

DROP DATABASE playground;

####################### System #################
SHOW TABLE STATUS LIKE '%table_name%';
SHOW INDEXES FROM book;

# query all indexes of a database, except for primary keys.
SELECT t.TABLE_NAME, t.INDEX_NAME, GROUP_CONCAT(t.COLUMN_NAME) as INDEX_COLUMN
FROM information_schema.statistics t
WHERE table_schema='database' and t.INDEX_NAME != 'PRIMARY'
GROUP BY t.TABLE_NAME, t.INDEX_NAME;

# calculate the in-disk size of table
SELECT table_name `Table`, round(((data_length + index_length) / 1024 / 1024), 2) `Size in MB`
FROM information_schema.TABLES
WHERE table_schema = 'database' AND table_name = 'table';

show variables like '%transaction_isolation%'; 	# show isolation level of current session.
set session transaction isolation level read uncommitted; # set isolation level.
set session transaction isolation level read committed;
set session transaction isolation level repeatable read;
set session transaction isolation level serializable;

# change password.
ALTER user 'root'@'localhost' IDENTIFIED BY '123456';

####################### Procedure #################
show procedure status where db='database';    # show procedures of the database
select routine_name from information_schema.routines where routine_schema='database';

DROP PROCEDURE IF EXISTS add_data;

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

####################### Grammar #################
# `;`, `\G`
select * from book limit 1;
select * from book limit 1\G

# when the result set of select statement is empty, null gets returned.
select ifnull((select * from book), 0);

select round(1.23333, 2); # 四舍五入，保留n位小数

select left('123456', 2); # 截取左边2位

select 1 > 2 as result; # 0 for false, 1 for true

select CAST(123 AS CHAR); # convert integer to string
select CONCAT('123', 'abc'); # concatenate strings

# https://learn.microsoft.com/en-us/sql/t-sql/functions/row-number-transact-sql?view=sql-server-ver16
# ROW_NUMBER(): number the results sequentially, like 1,2,3,4.
# RANK(): 1,2,2,4
# DENSE_RANK(): 1,2,2,3
SELECT
    ROW_NUMBER() OVER(PARTITION BY book_name ORDER BY book_pages) AS 'Row#',
    book_name, created_time
FROM book WHERE id > 5;

####################### Examples #################



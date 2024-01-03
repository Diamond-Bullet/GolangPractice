# mysqldump https://simplebackups.com/blog/the-complete-mysqldump-guide-with-examples/
mysqldump --version
# dump particular table to `file.sql`
mysqldump -h 1.2.3.4 -P 3306 -uroot -p1234 database table > file.sql

# another way to dump table
mysql -h 1.2.3.4 -P 3306 -uroot -p1234 -D database -e "select * from table_name" > file.txt
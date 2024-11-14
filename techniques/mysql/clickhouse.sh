# clickhouse is compatible for SQL.
# it means that you can use it with no difference from MySQL.
# https://clickhouse.com/docs/en/intro

# log in to clickhouse server
./clickhouse client -h 1.2.3.4 --port 9000 --user root --password root1234

# execute commands in a file
./clickhouse client < file.sql
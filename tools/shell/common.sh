## ssh
# generate key pair and copy to remote server.
ssh-keygen -t rsa
ssh-copy-id root@1.2.3.4

# login
ssh user@1.2.3.4

# scp src dst
scp test.txt  root@1.2.3.4:/root
scp root@1.2.3.4:/root/test.txt .


## ln
ln -s target_file link_file

# change target of link_file
ln -snf target_file link_file


## netstat
# install
apt install net-tools

netstat -lnap

# print the amount of different states for tcp.
# https://computingforgeeks.com/how-to-check-tcp-connections-states-in-linux-with-netstat/
netstat -nat | awk '{print $6}' | sort | uniq -c | sort -r


## lsof. list open file
# https://www.cnblogs.com/muchengnanfeng/p/9554993.html
# search by port
lsof -i:1234

# search by process id
lsof -p 1234


####################### Text Processing #################
## sort. 对文件中所有行，默认字典序排序. -n 按数值排序；-k 取某一列排序；-t 指定列的分隔符；-r 倒序；
sort -n -k 5 -t " " text.txt


## wc
# count by line
wc -l text.txt

# count by word
wc -w text.txt


####################### System resources #################
##df
# 查看磁盘使用况
df -hT


##du
# 查看当前文件夹下各个文件占用空间大小
du -h --max-depth=1 ./*

# 查看该文件夹大小
du -hs [path]


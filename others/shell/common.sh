####################### Remote Interaction #################
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
# list all network adapters
netstat -i

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

##uniq
# group by line and count the number of same line.
uniq -c text.txt

## awk
# 参考 菜鸟教程
# BEGIN{开始时执行，仅一次};
# {对每行进行处理};
# END{结束时执行，仅一次}
awk '{for(i=1;i<=NF;i++){if(NR==1){res[i]=$i;}else{res[i]=res[i]" "$i;}}};
    END{for(i=1;i<=NF;i++){print res[i]}}' file.txt # 转置

awk '{if(NR==10){print $0}}' file.txt # 输出第10行

awk '{for(i=1;i<=NF;i++){res[$i]++}};
    END{for(w in res){print w, res[w]};}' file.txt | sort -nr -k 2 # 统计词频

awk '$0 ~ /^([0-9]{3}\-){2}[0-9]{4}$|^\([0-9]{3}\)\s[0-9]{3}\-[0-9]{4}$/' file.txt # 匹配手机号

awk '{if($0 ~ /github\.com/){print $1"@"$2}}' go.mod # 下载go.mod下的github依赖

ls -l | awk '$1 !~ /^d.*/  {print $9}' | xargs wc -l | sort -nr -k 1 # 排列目录下文件行数

## chmod
chmod -R 600 [path]
chmod -R +x [path]

####################### Profiling #################
##df
# 查看磁盘使用况
df -hT

##du
# 查看当前文件夹下各个文件占用空间大小
du -h --max-depth=1 ./*
# 查看该文件夹大小
du -hs [path]

# network traffic analysis
# https://www.scaler.com/topics/linux-monitor-network-traffic/
cat /proc/net/dev

ifconfig
# show realtime traffic for each program
nethogs
# show realtime traffic for each tcp connection
iftop

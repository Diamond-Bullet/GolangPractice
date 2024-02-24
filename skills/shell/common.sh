####################### Text & File #################
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

# -F 指定分隔符
# refer to https://juejin.cn/post/7055242139222949924
grep "parent" log.txt  | awk -F "child\":\"" '{print $2}' | awk -F "\":\"" '{print $1" "$5}' | awk -F "\",\"name" '{print $1}' | sort -k 2 -r | head -n 25

## sed
# -i, replace string and update the file. -e, print content after substituting and not change the file.
sed -i "s/old_string/new_string/" file.txt
# print line 5-10.
sed -n '5,10p' file.txt

## chmod
chmod -R 600 [path]
chmod -R +x [path]

## ln
ln -s target_file link_file
# change target of link_file
ln -snf target_file link_file

## clear content of the file.
true >test.txt
cat /dev/null >test.txt
echo >test.txt

## grep
# -v exclude lines with certain word
grep -v "exclude_word" test.txt
# print lines around pattern line. -A4 4 lines after it. -B before it. -C before and after it.
grep -A4 "pattern" test.txt
# -E regular expression
grep -E "[0-9][a-z]" test.txt

# print first X lines
head -n 10 test.txt
# print last X lines
tail -n 10 test.txt

# print while enabling
echo -e "\t\n"

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

## netstat
# install
apt install net-tools

netstat -lnap
# list all network adapters
netstat -i
# print the amount of different states for tcp.
# https://computingforgeeks.com/how-to-check-tcp-connections-states-in-linux-with-netstat/
netstat -nat | awk '{print $6}' | sort | uniq -c | sort -r
# print the amount of each tcp connection status
netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'

## lsof. list open file
# https://www.cnblogs.com/muchengnanfeng/p/9554993.html
# search by port
lsof -i:1234
# search by process id
lsof -p 1234

# show processes https://www.geeksforgeeks.org/ps-command-in-linux-with-examples/
# -aux, format: USER PID %CPU %MEM VSZ RSS TTY STAT START TIME COMMAND
# TTY: short for teletype. the terminal file this process connects to.
ps -aux

# list directory tree
tree /tmp
# -a, show hidden files and folders. -L [num] display depth of directory tree.
tree -L 2 -a /tmp

# print absolute path. output: /root/folder/test.txt
realpath test.txt

####################### Run #################
# no hang up. continue running after existing the shell.
# > redirect the output.
# 0 standard input; 1 standard output; 2 standard error.
# & run program in the background.
nohup tmp.exe >log.txt 2>&1 &

## jobs. show background tasks.
jobs

## kill
# kill the No.1 background task. see the number in `jobs`.
kill %1

####################### Remote Interaction #################
# generate key pair and copy to remote server.
ssh-keygen -t rsa
ssh-copy-id root@1.2.3.4
# login
ssh user@1.2.3.4
# scp src dst
scp test.txt root@1.2.3.4:/root
scp root@1.2.3.4:/root/test.txt .

## To generate more types of certificate, refer to internet docs.
# generate key
openssl genrsa -des3 -out server.key 2048
# show key's content
openssl rsa -text -in server.key
# generate doc for applying for certificate
openssl req -new -key server.key -out server.csr
# show doc's content
openssl req -text -in server.csr -noout
# generate certificate
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
# generate random base64 string
openssl rand -base64 21

## 正反向shell:
# 正向Shell，服务器上使用ncat监听
ncat -l [port] -e /bin/bash
# 开发机上连接
ncat [remote_ip] [port]

#反向Shell，开发机上使用ncat监听
ncat -l [port]
# 服务器上连接
ncat [local_ip] [port] -e /bin/bash

## send file with GUI
yum install lrzsz
# send file from linux to remote machine
sz file.txt
# receive file from remote machine
rz

####################### Trifle #################
# ES query
curl -H "Content-Type: application/json" -u elastic:gRYppQtdTpP3nFzH -XPOST 'http://172.16.5.146:9200/imissu/_doc/_search?pretty' -d '{
  "query": {
            "bool": {
              "should": [
                { "term": { "user_id": "10" } },
                { "term": { "stage_name": "10" } }
              ]
            }
  },
  "sort": {
    "_script": {
      "type": "number",
      "script": {
        "lang": "painless",
        "source": " if (doc.uk.value.contains(params.q)) { if (doc.sk.value.contains(params.q)) {return Math.min(doc.sk.value.length()*10+1,
doc.uk.value.length()*10);} else {return doc.uk.value.length()*10;}} else {return doc.sk.value.length()*10+1;}",
        "params": {
          "q": "10"
        }
      },
      "order": "asc"
    }
  }
}'

# proto generation
protoc --micro_out=. --go_out=. ./customer.proto

####################### Mac OS #################
# Mac下使用了zsh会不执行/etc/profile文件，当然，如果用原始的是会执行。
# 转而执行的是这两个文件，每次登陆都会执行：~/.zshrc与/etc/zshenv、/etc/zshrc

# clear screen
cmd + k
# list processes
ps -lef
# show all files' size in current directory
du -sh *

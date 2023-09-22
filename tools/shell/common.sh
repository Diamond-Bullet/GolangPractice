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
# search by port
lsof -i:1234

# search by process id
lsof -p 1234
## awk
# reference https://www.runoob.com/
# BEGIN{execute at the beginning, for just 1 time};
# {process each line};
# END{execute in the end, just once}


# Syntax:
# `-F`: claim delimiter
#

# refer to https://juejin.cn/post/7055242139222949924
grep "parent" log.txt | awk -F "child\":\"" '{print $2}' | awk -F "\":\"" '{print $1" "$5}' | awk -F "\",\"name" '{print $1}' | sort -k 2 -r | head -n 25

awk '{for(i=1;i<=NF;i++){if(NR==1){res[i]=$i;}else{res[i]=res[i]" "$i;}}};
    END{for(i=1;i<=NF;i++){print res[i]}}' file.txt # transposing

awk '{if(NR==10){print $0}}' file.txt # output line 10

awk '{for(i=1;i<=NF;i++){res[$i]++}};
    END{for(w in res){print w, res[w]};}' file.txt | sort -nr -k 2 # count words frequency TODO

awk '$0 ~ /^([0-9]{3}\-){2}[0-9]{4}$|^\([0-9]{3}\)\s[0-9]{3}\-[0-9]{4}$/' file.txt # phone number match

awk '{if($0 ~ /github\.com/){print $1"@"$2}}' go.mod # download github dependencies in `go.mod`

ls -l | awk '$1 !~ /^d.*/  {print $9}' | xargs wc -l | sort -nr -k 1 # order non-directory files by lines amount of their content
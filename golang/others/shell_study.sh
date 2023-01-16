#!/bin/bash

# shell脚本执行的都是linux命令


# 变量命名
foo="123"
foo_name="ddd"
foo1="111"
echo -e "变量命名：${foo}\n"

# 引用传入的参数，必须是具体的数字
echo -e "执行文件名：${0}\n"
# $# 传入的参数个数
if [ $# -gt 0 ]; then
	echo -e "传入参数1: ${1}\n"
else
  echo -e "1 more args expected, get 1\n"
fi

# while 循环
index=0
while [ ${index} -le ${#} ]; do 
	echo -e "循环传入参数：${index}\n"
	# shellcheck disable=SC2219
	let "index++"
done	

# 数组, for 循环
# shellcheck disable=SC2206
foo_arr=(a b c d ${foo})
for item in ${foo_arr[*]}; do
	echo "${item}"
done

# 数组长度
echo -e "\n 数组foo_arr的长度是：${#foo_arr[*]}"

# case 语法
while true ; do    
    echo -n "输入1-5之间的数字：" # -n 不换行
    # shellcheck disable=SC2162
    read aNum # 将输入读取到 aNum
    case ${aNum} in
        1|2|3) echo "输入数字小于4: ${aNum}"
        ;;
        4|5) echo "输入数字大于等于4：${aNum}"
        ;;
        *) echo "无效数字，游戏结束"
        break # continue关键字，跳出当前循环
        ;;
    esac
done

# 函数
function testFunc() {
# 关系运算符：eq ne gt lt ge le
    if [ $# -eq 1 ]; then
        echo "参数数量为1个"
        echo "${0}" # 函数内部也是通过 $n 获取传递的参数
    fi

    if [ 1 -lt $# ]; then
        echo "参数数量大于1个，获得参数2：${1}"
    fi
}

testFunc 1 2

# awk
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


# .kinit_auto.sh脚本，开机自动kinit
fail=false
for((i=0;i<5;i++)); do
    kdestroy && kinit --password-file=/root/.kinit_password -l 86400 root@666.com
    if [ "$(klist | wc -l)" -ge 4 ]; then
        break
    fi
    fail=true
    echo "$(date "+%Y-%m-%d %H:%M:%S"), kinit failed in ${i}" >> /root/.kinit_log.txt
    sleep 1s
done
if [ $fail == true ]; then
    echo >> /root/.kinit_log.txt
fi

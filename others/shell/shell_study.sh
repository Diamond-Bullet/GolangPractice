#!/bin/bash
# Simply add `#!/bin/bash` on the top of file, and execute `chmod +x [file]`.
# In this way, you can avoid using `source [file]`.
# All commands shell script executes are linux commands.

# variable naming
foo="123"           # no space allowed.
foo1_name="${PATH}" # expressions work in double quotes.
foo2_name='$PATH'   # expressions don't work in single quotes. plain text.
foo1="$(which go)"  # store the output of command in variable

echo -e "variable naming: ${foo}, ${foo1_name}, ${foo1}\n" # reference variable in other place.

export foo # export variable as an environmental variable.
export foo2="golang"

unset foo # cancel exporting variable.

# 引用传入的参数，必须是具体的数字
echo -e "执行文件名：${0}\n"
# $# 传入的参数个数
if [ $# -gt 0 ]; then
  echo -e "传入参数1: ${1}\n"
else
  echo -e "1 more args expected, get 1\n"
fi

# while-loop
index=0
while [ ${index} -le ${#} ]; do
  echo -e "循环传入参数：${index}\n"
  ((index++))
done

# for-loop
foo_arr=(a b c d "${foo}")
for item in ${foo_arr[*]}; do
  echo "${item}"
done

# 数组长度
echo -e "\n 数组foo_arr的长度是：${#foo_arr[*]}"

# case 语法
while true; do
  echo -n "输入1-5之间的数字：" # -n 不换行
  read -r aNum          # 将输入读取到 aNum
  case ${aNum} in
  1 | 2 | 3)
    echo "输入数字小于4: ${aNum}"
    ;;
  4 | 5)
    echo "输入数字大于等于4：${aNum}"
    ;;
  *)
    echo "无效数字，游戏结束"
    break # continue关键字，跳出当前循环
    ;;
  esac
done

# function
function testFunc() {
  # 关系运算符：eq ne gt lt ge le
  if [ $# -eq 1 ]; then
    echo "参数数量为1个"
    echo "${1}" # 函数内部也是通过 $n 获取传递的参数
  fi

  if [ 1 -lt $# ]; then
    echo "参数数量大于1个，获得参数2：${2}"
  fi
}

testFunc 1 2

## environment variables. https://www.linuxprobe.com/environment-variable-configuration.html
#	.profile文件只在用户登录的时候读取一次，而.bashrc会在每次运行Shell脚本的时候读取一次。
#	可以自定义一个环境变量文件，比如在某个项目下定义uusama.profile，在这个文件中使用export定义一系列变量，
#	然后在~/.profile文件后面加上：source uusama.profile，这样你每次登陆都可以在Shell脚本中使用自己定义的一系列变量。
# 系统级
cat /etc/environment /etc/profile /etc/bash.bashrc /etc/profile.d/test.sh
# 用户级
cat ~/.profile ~/.bashrc

####################### Examples #################
# .kinit_auto.sh脚本，开机自动kinit
fail=false
for ((i = 0; i < 5; i++)); do
  kdestroy && kinit --password-file=/root/.kinit_password -l 86400 root@666.com
  if [ "$(klist | wc -l)" -ge 4 ]; then
    break
  fi
  fail=true
  echo "$(date "+%Y-%m-%d %H:%M:%S"), kinit failed in ${i}" >>/root/.kinit_log.txt
  sleep 1s
done
if [ $fail == true ]; then
  echo >>/root/.kinit_log.txt
fi

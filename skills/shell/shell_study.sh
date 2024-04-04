#!/bin/bash
# Simply add `#!/bin/bash` on the top of file, and execute `chmod +x [file]`.
# In this way, you can avoid using `source [file]`.
# All commands shell script executes are linux commands.

###################### variable naming ##############################
foo="123"           # no space allowed.
foo1_name="${PATH}" # expressions work in double quotes.
foo2_name='$PATH'   # expressions don't work in single quotes. plain text.
foo1="$(which go)"  # store the output of command in variable

# set variable by typing
read foo3

# give a var an alias
alias bar1=foo1
# unset the alias
unalias bar1

echo ${foo}
echo -e "variable naming: ${foo}, ${foo1_name}, ${foo1}\n" # reference variable in other place.

export foo # export variable as an environmental variable.
export foo2="golang"

unset foo # cancel exporting variable.

# print all env variables
env
# print all shell variables
set

## environment variables. https://www.linuxprobe.com/environment-variable-configuration.html
#	.profile is read only when user logs in. While .bashrc is read every time a shell script runs.

#	可以自定义一个环境变量文件，比如在某个项目下定义example.profile，在这个文件中使用export定义一系列变量，
#	然后在~/.profile文件后面加上：source example.profile，这样你每次登陆都可以在Shell脚本中使用自己定义的一系列变量。

# system-level
cat /etc/environment /etc/profile /etc/bash.bashrc /etc/profile.d/test.sh
# user-level
cat ~/.profile ~/.bashrc

###################### commandline parameters ##############################
# invoke arguments passed to, 必须是具体的数字
echo -e "executable name：${0}\n"
# `$#` arguments amount
if [ $# -gt 0 ]; then
  echo -e "argument 1: ${1}\n"
else
  echo -e "1 more args expected, get 1\n"
fi

###################### while-loop ##############################
index=0
while [ ${index} -le ${#} ]; do
  echo -e "argument $index：${index}\n"
  ((index++))
done

###################### for-loop ##############################
foo_arr=(a b c d "${foo}")
for item in ${foo_arr[*]}; do
  echo "${item}"
done

# array length
echo -e "length of array 'foo_arr' is: ${#foo_arr[*]}\n"

###################### switch-case ##############################
while true; do
  echo -n "input a number between 1 and 5: " # `-n` disable switching to a new line
  read -r aNum          # read input and store in aNum
  case ${aNum} in
  1 | 2 | 3)
    echo "input a number less than 4: ${aNum}"
    ;;
  4 | 5)
    echo "input a number greater than or equal to 4：${aNum}"
    ;;
  *)
    echo "invalid number. game over."
    break # `continue` jump to next execution of the loop.
    ;;
  esac
done

###################### function ##############################
function testFunc() {
  # 关系运算符：eq ne gt lt ge le
  if [ $# -eq 1 ]; then
    echo "arguments amount is 1"
    echo "${1}" # use `$num` to get arguments in the function as well.
  fi

  if [ 1 -lt $# ]; then
    echo "arguments amount greater than 1, get argument 2: ${2}"
  fi
}

testFunc 1 2

####################### Examples #################
# .kinit_auto.sh, kinit automatically when powering on.
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

# check whether file exists.
if [ -f "/data/filename" ]; then
  echo "file exists"
else
  echo "file doesn't exist"
fi
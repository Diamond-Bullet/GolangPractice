if [ "$SHELL" -eq "" ]; then
  echo "TODO"
fi

sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

function getOS() {
    if [ $# -eq 1 ]; then
        echo "参数数量为1个"
        echo "${0}" # 函数内部也是通过 $n 获取传递的参数
    fi
}
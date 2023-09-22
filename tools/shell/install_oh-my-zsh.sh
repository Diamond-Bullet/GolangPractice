function install_zsh() {
  if [[ "$(which zsh)" == *"not found" ]]; then
    apt update
    apt install zsh
  fi

  if [[ "$SHELL" != *"zsh" ]]; then
    chsh -s "$(which zsh)"
  fi
}

function get_os() {
  if [ $# -eq 1 ]; then
    echo "参数数量为1个"
    echo "${0}" # 函数内部也是通过 $n 获取传递的参数
  fi
}

sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" || exit 1

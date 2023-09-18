# 安装：pip install pyperclip pyautogui
import pyperclip
import pyautogui
import time

# 读取剪贴板文本内容
text = pyperclip.paste()

time.sleep(3)

# 模拟按键输入这段剪贴板文本内容
pyautogui.typewrite(text, interval=0.02)

# 使用方式：
# 先复制一段你需要的文本，然后运行该脚本，等待3秒，然后就会模拟键盘输入这段文本内容
# 注意：
# 等待3秒是为了确保你已经切换到了需要输入的地方，比如 notepad plus, terminal 等
# 如果文本中包含非 ASCII 字符，需要使用 base64 编码，然后复制，再运行该脚本


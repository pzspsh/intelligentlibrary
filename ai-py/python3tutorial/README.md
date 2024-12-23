# @python3 教程(python3 tutorial)

<details>
  <summary><h2>目录</h2></summary>

[介绍](#介绍)

- [Python 语言]()

[基础](#基础)

- [类型]()
- [变量]()
- [字符串与编码]()
- [条件判断]()
- [循环]()

[函数](#函数)

- [函数定义]()

[面向对象](#面向对象)

- [类与实例]()

[正则表达式](#正则表达式)

[进程和线程](#进程和线程)

- [进程和线程介绍]()
- [多进程]()
- [多线程]()

[网络编程](#网络编程)

- [TCP 编程]()
- [UDP 编程]()

</details>

```python
print("hello world")

```

python项目导包
```python
from rootdir.包文件夹名称...包文件夹名称 import 功能程序所在的.py文件名 as 重命名1
from rootdir.包文件夹名称...包文件夹名称.功能程序所在的.py文件名 import 类 as 重命名2, 功能函数名 as 重命名3


调用方法：
重命名1.功能函数名()
重命名1.类()

重命名2()
重命名3()

print(重命名2.字段)

例如：项目DemoProject
DemoProject:
  └── configs
  |  └── configs.py
  └── views
  |  └── utils.py
  |  └── views.py  
  └── main.py


config.py
class Config:
  fiel1 = "hello, world!"
  fiel2 = "你好, 世界!"

utils.py
def demo():
  print("hello demo...")

views.py
from view.utils import demo
from view import utils 
from view import utils as ut
from configs.configs import Config as cf

def ViewsRun():
  demo()
  utils.demo()
  print(ut.demo())
  print(cf.fiel1)
  print(cf.fiel2)

main.py
from views.views import ViewsRun


def main()
  ViewsRun()

if __name__ == "__main__":
  main()

```


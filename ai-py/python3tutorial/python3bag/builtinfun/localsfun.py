# -*- encoding: utf-8 -*-
"""
@File   : localsfun.py
@Time   : 2023-12-07 10:01:17
@Author : pan
"""


def pan(arg):
    p = 1
    print(locals())


""" 
locals() 函数会以字典类型返回当前位置的全部局部变量。

对于函数, 方法, lambda 函式, 类, 以及实现了 __call__ 方法的类实例, 它都返回 True。 
"""

if __name__ == "__main__":
    pan(4)

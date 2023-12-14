# -*- encoding: utf-8 -*-
"""
@File   : callablefun.py
@Time   : 2023-12-07 09:59:29
@Author : pan
"""
""" 
callable() 函数用于检查一个对象是否是可调用的。如果返回 True，object 仍然可能调用失败；但如果返回 False，调用对象object
绝对不会成功。

对于函数、方法、lambda 函式、 类以及实现了 __call__ 方法的类实例, 它都返回 True。
"""

print(callable(0))


def add(a, b):
    return a + b


print(callable(add))


class A:
    def method(self):
        return 0


print(callable(A))
a = A()
print(a)

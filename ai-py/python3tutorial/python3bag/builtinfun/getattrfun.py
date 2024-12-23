# -*- encoding: utf-8 -*-
"""
@File   : getattrfun.py
@Time   : 2023-12-07 10:00:26
@Author : pan
"""


class A(object):
    bar = 1


a = A()
res = getattr(a, "bar")  # 获取属性 bar 值
print(res)

res1 = getattr(a, "bar2", 3)  # 属性 bar2 不存在，但设置了默认值
print(res1)

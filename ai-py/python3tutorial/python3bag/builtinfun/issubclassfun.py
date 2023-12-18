# -*- encoding: utf-8 -*-
"""
@File   : issubclassfun.py
@Time   : 2023-12-07 10:01:04
@Author : pan
"""
""" 如果 class 是 classinfo 的子类返回 True，否则返回 False。 """


class A:
    pass


class B(A):
    pass


print(issubclass(B, A))  # 返回 True

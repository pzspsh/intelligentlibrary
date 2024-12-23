# -*- encoding: utf-8 -*-
"""
@File   : isinstancefun.py
@Time   : 2023-12-07 10:01:00
@Author : pan
"""


class A:
    pass


class B(A):
    pass


isinstance(A(), A)  # returns True
type(A()) == A  # returns True
isinstance(B(), A)  # returns True
type(B()) == A  # returns False

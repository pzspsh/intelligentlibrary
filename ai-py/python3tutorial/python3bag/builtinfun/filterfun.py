# -*- encoding: utf-8 -*-
"""
@File   : filterfun.py
@Time   : 2023-12-07 10:00:11
@Author : pan
"""


def is_odd(n):
    return n % 2 == 1


tmplist = filter(is_odd, [1, 2, 3, 4, 5, 6, 7, 8, 9, 10])
newlist = list(tmplist)
print(newlist)


import math


def is_sqr(x):
    return math.sqrt(x) % 1 == 0


tmplist = filter(is_sqr, range(1, 101))
newlist = list(tmplist)
print(newlist)

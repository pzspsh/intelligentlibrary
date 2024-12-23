# -*- encoding: utf-8 -*-
"""
@File   : zipfun.py
@Time   : 2023-12-07 10:03:01
@Author : pan
"""
a = [1, 2, 3]
b = [4, 5, 6]
c = [4, 5, 6, 7, 8]
zipped = zip(a, b)
print(zipped)

lis = list(zip(a, c))
print(lis)

a1, a2 = zip(*zip(a, b))
print(list(a1))
print(list(a2))

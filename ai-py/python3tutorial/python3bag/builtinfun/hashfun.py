# -*- encoding: utf-8 -*-
"""
@File   : hashfun.py
@Time   : 2023-12-07 10:00:36
@Author : pan
"""
res = hash("test")  # 字符串
print(res)

res = hash(str([1, 2, 3]))  # 集合
print(res)

res = hash(str(sorted({"1": 1})))  # 字典
print(res)

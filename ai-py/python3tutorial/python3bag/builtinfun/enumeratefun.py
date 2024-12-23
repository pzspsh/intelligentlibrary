# -*- encoding: utf-8 -*-
"""
@File   : enumeratefun.py
@Time   : 2023-12-07 10:00:01
@Author : pan
"""
seasons = ["Spring", "Summer", "Fall", "Winter"]
res = list(enumerate(seasons))
print(res)

res = list(enumerate(seasons, start=1))  # 下标从 1 开始
print(res)

i = 0
seq = ["one", "two", "three"]
for element in seq:
    print(i, seq[i])
    i += 1


seq = ["one", "two", "three"]
for i, element in enumerate(seq):
    print(i, element)

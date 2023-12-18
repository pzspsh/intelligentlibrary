# -*- encoding: utf-8 -*-
"""
@File   : reversedfun.py
@Time   : 2023-12-07 10:02:10
@Author : pan
"""
# 字符串
seqString = "panzhong"
print(list(reversed(seqString)))

# 元组
seqTuple = ("a", "b", "c", "d", "e", "f")
print(list(reversed(seqTuple)))

# range
seqRange = range(5, 9)
print(list(reversed(seqRange)))

# 列表
seqList = [1, 2, 4, 3, 5]
print(list(reversed(seqList)))

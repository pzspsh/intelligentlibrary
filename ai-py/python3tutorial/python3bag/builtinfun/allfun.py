# -*- encoding: utf-8 -*-
"""
@File   : allfun.py
@Time   : 2023-12-07 09:58:57
@Author : pan
"""
re = all(["a", "b", "c", "d"])  # 列表list，元素都不为空或0
print(re)

re = all(["a", "b", "", "d"])  # 列表list，存在一个为空的元素
print(re)

re = all(("a", "b", "c", "d"))  # 元组tuple，元素都不为空或0
print(re)


re = all(("a", "b", "", "d"))  # 元组tuple，存在一个为空的元素
print(re)


re = all((0, 1, 2, 3))  # 元组tuple，存在一个为0的元素
print(re)
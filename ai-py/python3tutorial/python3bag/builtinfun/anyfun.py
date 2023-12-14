# -*- encoding: utf-8 -*-
"""
@File   : anyfun.py
@Time   : 2023-12-07 09:58:51
@Author : pan
"""
res = any(["a", "b", "c", "d"])  # 列表list，元素都不为空或0
print(res)

res = any(["a", "b", "", "d"])  # 列表list，存在一个为空的元素
print(res)

res = any([0, "", False])  # 列表list,元素全为0,'',false
print(res)

res = any(("a", "b", "", "d"))  # 元组tuple，存在一个为空的元素
print(res)

res = any((0, "", False))  # 元组tuple，元素全为0,'',false
print(res)

res = any([])  # 空列表
print(res)

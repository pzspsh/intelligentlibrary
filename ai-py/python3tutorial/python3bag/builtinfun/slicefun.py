# -*- encoding: utf-8 -*-
"""
@File   : slicefun.py
@Time   : 2023-12-07 10:02:26
@Author : pan
"""
myslice = slice(5)  # 设置截取5个元素的切片
print(myslice)

arr = range(10)
print(arr)

arr = arr[myslice]
print(arr)

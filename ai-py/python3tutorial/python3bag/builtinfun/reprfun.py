# -*- encoding: utf-8 -*-
"""
@File   : reprfun.py
@Time   : 2023-12-07 10:02:07
@Author : pan
"""
""" 返回一个对象的 string 格式。 """
s = "物品\t单价\t数量\n包子\t1\t2"
print(s)
print(repr(s))

s = "panzhong"
print(repr(s))

dic = {"runoob": "runoob.com", "google": "google.com"}
print(repr(dic))

# -*- encoding: utf-8 -*-
"""
@File   : helpfun.py
@Time   : 2023-12-07 10:00:40
@Author : pan
"""
re = help("sys")  # 查看 sys 模块的帮助
print(re)

re = help("str")
print(re)

a = [1, 2, 3]
re = help(a)
print(re)

re = help(a.append)
print(re)

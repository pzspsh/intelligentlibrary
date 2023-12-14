# -*- encoding: utf-8 -*-
"""
@File   : compilefun.py
@Time   : 2023-12-07 09:59:40
@Author : pan
"""
str = "for i in range(0,10): print(i)"
c = compile(str, "", "exec")  # 编译为字节代码对象
print(c)

res = exec(c)
print(res)

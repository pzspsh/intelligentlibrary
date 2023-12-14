# -*- encoding: utf-8 -*-
"""
@File   : execfun.py
@Time   : 2023-12-07 10:00:08
@Author : pan
"""
res = exec('print("Hello World")')
print(res)


x = 10
expr = """
z = 30
sum = x + y + z
print(sum)
"""


def func():
    y = 20
    exec(expr)
    exec(expr, {"x": 1, "y": 2})
    exec(expr, {"x": 1, "y": 2}, {"y": 3, "z": 4})


func()

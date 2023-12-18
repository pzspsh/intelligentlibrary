# -*- encoding: utf-8 -*-
"""
@File   : hasattrfun.py
@Time   : 2023-12-07 10:00:33
@Author : pan
"""


class Coordinate:
    x = 10
    y = -5
    z = 0


point1 = Coordinate()
print(hasattr(point1, "x"))
print(hasattr(point1, "y"))
print(hasattr(point1, "z"))
print(hasattr(point1, "no"))  # 没有该属性

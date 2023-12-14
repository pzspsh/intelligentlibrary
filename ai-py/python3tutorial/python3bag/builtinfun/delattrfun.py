# -*- encoding: utf-8 -*-
"""
@File   : delattrfun.py
@Time   : 2023-12-07 09:59:47
@Author : pan
"""


class Coordinate:
    x = 10
    y = -5
    z = 0


point1 = Coordinate()

print("x = ", point1.x)
print("y = ", point1.y)
print("z = ", point1.z)

delattr(Coordinate, "z")

print("--删除 z 属性后--")
print("x = ", point1.x)
print("y = ", point1.y)

# 触发错误
print("z = ", point1.z)

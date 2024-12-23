# -*- encoding: utf-8 -*-
"""
@File   : roundfun.py
@Time   : 2023-12-07 10:02:14
@Author : pan
"""
""" 
round( x [, n]  )

x -- 数字表达式。
n -- 表示从小数点位数，其中 x 需要四舍五入，默认值为 0。
"""
print("round(70.23456) : ", round(70.23456))
print("round(56.659,1) : ", round(56.659, 1))
print("round(80.264, 2) : ", round(80.264, 2))
print("round(100.000056, 3) : ", round(100.000056, 3))
print("round(-100.000056, 3) : ", round(-100.000056, 3))

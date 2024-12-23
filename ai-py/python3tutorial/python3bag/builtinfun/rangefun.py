# -*- encoding: utf-8 -*-
"""
@File   : rangefun.py
@Time   : 2023-12-07 10:01:59
@Author : pan
"""
for number in range(1, 6):
    print(number)

for number in range(6):
    print(number)


for number in range(1, 6, 2):
    print(number)


for number in range(6, 1, -1):
    print(number)


numbers = list(range(1, 6))
print(numbers)


numbers = tuple(range(1, 6))
print(numbers)


a = range(5)  # 默认从 0 开始, 在 5 以内, 到不了 5
print(list(a))  # [0, 1, 2, 3, 4]

a = range(1, 5)  # 给出起点, 终点(到不了的)
print(list(a))  # [1, 2, 3, 4]

a = range(1, 11, 2)  #  给出起点, 终点(到不了的), 步长(正整数,负整数)
print(list(a))  # [1, 3, 5, 7, 9]

a = range(2, 2468, 2)  # 最大下标这样获得:len(a)-1
print(a[0], a[1], a[2], a[len(a) - 1])  #  2 4 6 2466

a = range(100, 80, -2)  # 前大后小的,步长应该是负数
print(list(a))  # [100, 98, 96, 94, 92, 90, 88, 86, 84, 82]

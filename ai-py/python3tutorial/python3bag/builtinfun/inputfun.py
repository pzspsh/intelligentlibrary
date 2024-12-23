# -*- encoding: utf-8 -*-
"""
@File   : inputfun.py
@Time   : 2023-12-07 10:00:54
@Author : pan
"""
a, b, c = input("请输入三角形三边的长：").split()
a = int(a)
b = int(b)
c = int(c)

# 计算三角形的半周长p
p = (a + b + c) / 2

# 计算三角形的面积s
s = (p * (p - a) * (p - b) * (p - c)) ** 0.5

# 输出三角形的面积
print("三角形面积为：", format(s, ".2f"))


str = input()
list = str.split(" ")
# len(list)取得list的长度
for i in range(0, len(list)):
    print("转换前，", list[i], "的类型是", type(list[i]))
    list[i] = int(list[i])
    print("转换后，", list[i], "的类型是", type(list[i]))

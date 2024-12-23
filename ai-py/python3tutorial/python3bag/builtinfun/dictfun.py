# -*- encoding: utf-8 -*-
"""
@File   : dictfun.py
@Time   : 2023-12-07 09:59:50
@Author : pan
"""
numbers = dict(x=5, y=0)
print("numbers =", numbers)
print(type(numbers))

empty = dict()
print("empty =", empty)
print(type(empty))


print(dict(zip(["one", "two", "three"], [1, 2, 3])))  # 映射函数方式来构造字典
print(dict([("one", 1), ("two", 2), ("three", 3)]))  # 可迭代对象方式来构造字典


# 没有设置关键字参数
numbers1 = dict([("x", 5), ("y", -5)])
print("numbers1 =", numbers1)

# 设置关键字参数
numbers2 = dict([("x", 5), ("y", -5)], z=8)
print("numbers2 =", numbers2)

# zip() 创建可迭代对象
numbers3 = dict(dict(zip(["x", "y", "z"], [1, 2, 3])))
print("numbers3 =", numbers3)


numbers1 = dict({"x": 4, "y": 5})
print("numbers1 =", numbers1)

# 以下代码不需要使用 dict()
numbers2 = {"x": 4, "y": 5}
print("numbers2 =", numbers2)

# 关键字参数会被传递
numbers3 = dict({"x": 4, "y": 5}, z=8)
print("numbers3 =", numbers3)

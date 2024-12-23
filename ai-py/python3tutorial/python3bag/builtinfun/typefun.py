# -*- encoding: utf-8 -*-
"""
@File   : typefun.py
@Time   : 2023-12-07 10:02:54
@Author : pan
"""
"""
type() 函数如果你只有第一个参数则返回对象的类型，三个参数返回新的类型对象。

isinstance() 与 type() 区别：

type() 不会认为子类是一种父类类型，不考虑继承关系。

isinstance() 会认为子类是一种父类类型，考虑继承关系。

如果要判断两个类型是否相同推荐使用 isinstance()。

语法
以下是 type() 方法的语法:

type(object)
type(name, bases, dict)
参数
name -- 类的名称。
bases -- 基类的元组。
dict -- 字典，类内定义的命名空间变量。
"""


class A:
    pass


class B(A):
    pass


isinstance(A(), A)  # returns True
type(A()) == A  # returns True
isinstance(B(), A)  # returns True
type(B()) == A  # returns False

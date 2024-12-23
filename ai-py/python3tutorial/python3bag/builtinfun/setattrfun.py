# -*- encoding: utf-8 -*-
"""
@File   : setattrfun.py
@Time   : 2023-12-07 10:02:18
@Author : pan
"""


class A(object):
    bar = 1


class B:
    name = "panzhong"


if __name__ == "__main__":
    """对已存在的属性进行赋值"""
    a = A()
    getattr(a, "bar")
    setattr(a, "bar", 5)
    print(a.bar)

    """ 如果属性不存在会创建一个新的对象属性，并对属性赋值 """
    b = B()
    setattr(b, "age", 38)
    print(b.age)

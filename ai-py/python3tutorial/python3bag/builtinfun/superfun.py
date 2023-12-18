# -*- encoding: utf-8 -*-
"""
@File   : superfun.py
@Time   : 2023-12-07 10:02:47
@Author : pan
"""


class A:
    def add(self, x):
        y = x + 1
        print(y)


class B(A):
    def add(self, x):
        super().add(x)


b = B()
b.add(2)  # 3


class FooParent(object):
    def __init__(self):
        self.parent = "I'm the parent."
        print("Parent")

    def bar(self, message):
        print("%s from Parent" % message)


class FooChild(FooParent):
    def __init__(self):
        # super(FooChild,self) 首先找到 FooChild 的父类（就是类 FooParent），然后把类 FooChild 的对象转换为类 FooParent 的对象
        super(FooChild, self).__init__()
        print("Child")

    def bar(self, message):
        super(FooChild, self).bar(message)
        print("Child bar fuction")
        print(self.parent)


if __name__ == "__main__":
    fooChild = FooChild()
    fooChild.bar("HelloWorld")

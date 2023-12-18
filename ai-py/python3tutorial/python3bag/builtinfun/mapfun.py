# -*- encoding: utf-8 -*-
"""
@File   : mapfun.py
@Time   : 2023-12-07 10:01:21
@Author : pan
"""


def square(x):
    return x**2


if __name__ == "__main__":
    map(square, [1, 2, 3, 4, 5])
    list(map(square, [1, 2, 3, 4, 5]))
    list(map(lambda x: x**2, [1, 2, 3, 4, 5]))

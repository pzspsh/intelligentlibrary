# -*- encoding: utf-8 -*-
"""
@File   : memoryviewfun.py
@Time   : 2023-12-07 10:01:28
@Author : pan
"""
v = memoryview("abcefg")
print(v[1])
print(v[-1])
print(v[1:4].tobytes())


v1 = memoryview(bytearray("abcefg", "utf-8"))
print(v1[1])
print(v1[-1])

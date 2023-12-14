# -*- encoding: utf-8 -*-
"""
@File   : main.py
@Time   : 2023-12-07 09:53:00
@Author : pan
"""
n = -37
print(bin(n))


def bit_length(self):
    s = bin(self)
    s = s.lstrip("-0b")
    print(len(s))


if __name__ == "__main__":
    bit_length(67)

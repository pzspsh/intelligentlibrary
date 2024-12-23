# -*- encoding: utf-8 -*-
'''
@File   : main.py
@Time   : 2023-12-07 09:53:00
@Author : pan
'''
import os

dir = "path1/path2/path3/file.txt"
filepath = os.path.dirname(dir)  # path1/path2/path3
print(filepath)

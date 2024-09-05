# -*- encoding: utf-8 -*-
"""
@File   : main.py
@Time   : 2024-09-05 15:35:16
@Author : pan
"""
import os


def GetAllDir(filepath):
    for root, dirs, files in os.walk(top=filepath, topdown=False):
        for name in files:
            print(os.path.join(root, name))
        for name in dirs:
            print(os.path.join(root, name))


if __name__ == "__main__":
    GetAllDir("filepath")

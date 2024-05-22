# -*- encoding: utf-8 -*-
"""
@File   : jsonconfig.py
@Time   : 2023-06-09 16:56:24
@Author : pan
"""
import json


def ReadJson(filename):
    with open(filename, "r", encoding="utf-8") as f:
        data = json.load(f)
    print(data)


if __name__ == "__main__":
    filename = "path/configs/config.json"
    ReadJson(filename)

# -*- encoding: utf-8 -*-
"""
@File   : example.py
@Time   : 2023-05-31 22:08:06
@Author : pan
"""
import re

target = "https://data.reinventory.co.uk"
mate_target = "((http|https)://|)(((\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}|[a-zA-Z0-9\u4e00-\u9fa5-@!*%?^+=~#.]+)(:(\d+)|)))"
result = re.match(mate_target, target, re.S)
print(result.groups())

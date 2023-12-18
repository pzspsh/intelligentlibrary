# -*- encoding: utf-8 -*-
"""
@File   : varsfun.py
@Time   : 2023-12-07 10:02:58
@Author : pan
"""
from datetime import timedelta
from enum import Enum


class Period(timedelta, Enum):
    "different lengths of time"
    _ignore_ = "Period i"
    Period = vars()
    for i in range(367):
        Period["day_%d" % i] = i


print(Period.day_5.value)
# 5 days, 0:00:00

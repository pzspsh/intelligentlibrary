# -*- encoding: utf-8 -*-
"""
@File   : main.py
@Time   : 2023-12-07 09:53:00
@Author : pan
"""
import sys

sys.path[0:0] = [
    "/XXX/path1",
    "/XXX/path2",
    "/XXX/path3",
]

import site

site.main()


# main main.py
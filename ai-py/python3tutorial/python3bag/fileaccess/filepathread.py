# -*- encoding: utf-8 -*-
"""
@File   : filepathread.py
@Time   : 2024-05-22 17:05:19
@Author : pan
"""
import os


def filespath(filePath) -> list:
    file_path_list = []
    for root, _, files in os.walk(filePath, topdown=False):
        for name in files:
            pocPath = os.path.join(root, name)
            if ".yaml" in pocPath:
                file_path_list.append(pocPath)
    return file_path_list

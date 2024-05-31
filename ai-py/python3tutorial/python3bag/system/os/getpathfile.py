# -*- encoding: utf-8 -*-
'''
@File   : getpathfile.py
@Time   : 2024-05-31 14:51:59
@Author : pan
'''
import os

# 文件路径
file_path = "/path/to/directory/filename.txt"
# 使用os.path.basename获取文件名
file_name = os.path.basename(file_path)
print(file_name)  # 输出: filename.txt

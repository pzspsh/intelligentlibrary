# -*- encoding: utf-8 -*-
"""
@File   : example.py
@Time   : 2023-12-14 11:14:38
@Author : pan
"""
import codecs
import io

buffer = io.StringIO()
stream = codecs.getwriter("rot_13")(buffer)

text = "abcdefghijklmnopqrstuvwxyz"

stream.write(text)
stream.flush()

print("Original:", text)
print("ROT-13 :", buffer.getvalue())

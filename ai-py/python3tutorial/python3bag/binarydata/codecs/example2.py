# -*- encoding: utf-8 -*-
"""
@File   : example2.py
@Time   : 2023-12-14 11:15:00
@Author : pan
"""
import codecs
import io

buffer = io.BytesIO()
stream = codecs.getwriter("zlib")(buffer)

text = b"abcdefghijklmnopqrstuvwxyz\n" * 50

stream.write(text)
stream.flush()

print("Original length :", len(text))
compressed_data = buffer.getvalue()
print("ZIP compressed :", len(compressed_data))

buffer = io.BytesIO(compressed_data)
stream = codecs.getreader("zlib")(buffer)

first_line = stream.readline()
print("Read first line :", repr(first_line))

uncompressed_data = first_line + stream.read()
print("Uncompressed :", len(uncompressed_data))
print("Same :", text == uncompressed_data)

# -*- encoding: utf-8 -*-
"""
@File   : example3.py
@Time   : 2023-12-14 11:16:10
@Author : pan
"""
import codecs
import base64

original_data = b"hello world!"
encoded_data = base64.b64encode(original_data)
decoded_data = base64.b64decode(encoded_data)
print("Original data: ", original_data)
print("Encoded data: ", encoded_data)
print("Decoded data: ", decoded_data)


utf8_str = "你好，世界！".encode("utf-8")
print("utf8_str: ", utf8_str)
utf16_str = codecs.encode("你好，世界！", "utf-16")
print("utf16_str: ", utf16_str)
gb2312_str = "你好，世界！".encode("gb2312")
print("gb2312_str: ", gb2312_str)

decoded_utf8_str = utf8_str.decode("utf-8")
print("decoded_utf8_str: ", decoded_utf8_str)

decoded_utf16_str = codecs.decode(utf16_str, "utf-16")
print("decoded_utf16_str: ", decoded_utf16_str)

decoded_gb2312_str = gb2312_str.decode("gb2312")
print("decoded_gb2312_str: ", decoded_gb2312_str)


import zlib

original_data = b"Python is a powerful programming language, widely used in data analysis, artificial intelligence, machine learning and other fields."
compressed_data = zlib.compress(original_data)
decompressed_data = zlib.decompress(compressed_data)
print("Original data: ", original_data)
print("Compressed data: ", compressed_data)
print("Decompressed data: ", decompressed_data)


unicode_str = "\u212b"
print("unicode_str: ", unicode_str)
print("Normalized unicode_str: ", codecs.normalize("NFKC", unicode_str))

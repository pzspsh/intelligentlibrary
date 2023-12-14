# -*- encoding: utf-8 -*-
"""
@File   : main.py
@Time   : 2023-12-07 09:53:00
@Author : pan
"""
from struct import *
from collections import namedtuple

pa = pack(">bhl", 1, 2, 3)
print(pa)  # b'\x01\x00\x02\x00\x00\x00\x03'

re = unpack(">bhl", b"\x01\x00\x02\x00\x00\x00\x03")
print(re)

# 尝试打包一个对于所定义字段来说过大的整数
# pa = pack(">h", 99999)  # 会报错
# print(pa)
""" 显示 's' and 'c' 格式字符之间的差异 """
pa = pack("@ccc", b"1", b"2", b"3")
print(pa)  # b'123'
pa = pack("@3s", b"123")
print(pa)  # b'123'

""" 解包的字段可通过将它们赋值给变量或将结果包装为一个具名元组来命名 """
record = b"raymond   \x32\x12\x08\x01\x08"
name, serialnum, school, gradelevel = unpack("<10sHHb", record)
print(name, serialnum, school, gradelevel)


Student = namedtuple("Student", "name serialnum school gradelevel")
res = Student._make(unpack("<10sHHb", record))
print(res)

""" 
格式字符的顺序可能会因为填充是隐式的而对在原生模式中的大小产生影响。 在标准模式下，用户要负责插入任何必要的填充。
请注意下面的第一个 pack 调用中在已打包的 '#' 之后添加了三个 NUL 字节以便在四字节边界上对齐到下面的整数。 在这个例子中，
输出是在一台小端序的机器上产生的 
"""
pa = pack("@ci", b"#", 0x12131415)
print(pa)
pa = pack("@ic", 0x12131415, b"#")
print(pa)

res = calcsize("@ci")
print(res)
res = calcsize("@ic")
print(res)
""" 以下格式 'llh0l' 将会在末尾添加两个填充字节，假定平台的 long 类型按 4 个字节的边界对齐的话 """
pa = pack("@llh0l", 1, 2, 3)
print(pa)

""" 请看这两个简单的示例（在 64 位的小端序机器上 """
res = calcsize("@lhl")
print(res)

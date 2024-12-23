# -*- encoding: utf-8 -*-
"""
@File   : main.py
@Time   : 2023-12-07 09:53:00
@Author : pan
"""
import string
from string import Template

s = Template("$who likes $what")
res = s.substitute(who="tim", what="kung pao")
print(res)


s = "abcde"
# 放入print中使用
print(s.upper())
print(s.lower())
print("abcdef".upper())
print("QWERT".lower())
# 这里注意是通过生成新的字符串而不是更改原来字符串
s.upper()
print(s)


s = "abcde,qweRTY"
t = "abcde qweRTY"
# 以，隔开的单词
print(s.title())
print(s.capitalize())
# 以空格隔开的单词
print(t.title())
print(t.capitalize())


# 1234 全是数字 为True
print("1234".isdecimal())
# asdf4 中4是数字不是字母 为False
print("asdf4".isdigit())
# qwe12@ 中@既不是数字 也不是字母为False
print("qwe12@".isalnum())
# asdf全是小写 为True
print("asdf".islower())
# ADS全是大写 为True
print("ADS".isupper())
# Wshd，qwe中 虽然W大写 但是第二个单词qwe中q小写 不符合title()所以为False
print("Wshd，qwe".istitle())
# n为换行 是空白符 为True
print("n".isspace())
# t为制表符 不可打印 为False
print("t".isprintable())
# qe123 符合命名规则 为True
print("qe125".isidentifier())
print("qwer".ljust(10, "+"))


# 不加"+""-"纯数字，用填充物"0"将字符串前填充满
print("12345".zfill(10))
# 加"-"纯数字，越过"-"用填充物"0"将字符串前填充满
print("-125".zfill(10))
# 加"+"数字字母组合，越过"+"用填充物"0"将字符串前填充满
print("+qwe125".zfill(10))
# 加其他符号，用填充物"0"将字符串前填充满
print("#qwe12".zfill(10))


# 全部字符串内 搜索qwe 出现的次数
print("qwertasdqwezxcqwe".count("qwe"))
# 由于字符串从0开始计数，1为字符串第二个，相当于从w开始
print("qwertasdqwezxcqwe".count("qwe", 1))
# 从字符串第 2个开始到第15个截止，共出现qwe的次数
print("qwertasdqwezxcqwe".count("qwe", 1, 14))


# 搜索开头位置为qwe 符合条件，为True
print("qwertasdqwezxcqwe".startswith("qwe"))
# 开头位置为字符串下标为1开始，也就是说开头为wer与qwe不同为False
print("qwertasdqwezxcqwe".startswith("qwe", 1))
# 结尾位置为qwe符合条件 为True
print("qwertasdqwezxcqwe".endswith("qwe", "asd"))


s = "qweraqwesfgzqweop"
print(s.find("qwe"))
print(s.rfind("qwe"))
print(s.index("qwe"))
print(s.rindex("qwe"))

s = "qweraqwesfgzqweop"
# 将字符串全部的qwe  换为asd
print(s.replace("qwe", "asd"))
# 将字符串前两个qwe  换为asd
print(s.replace("qwe", "asd", 2))
# 将字符串全部的qew  换为asd 没有则输出原字符串
print(s.replace("qew", "asd"))


t = "qwetqwertqasdsdftas"
print(t.expandtabs(4))


t = "qwertyuasdfghjkl"
print(t.partition("yua"))
print(t.partition("asqw"))
print(t.rpartition("asqw"))


t = input().split()
print(t)


# 字符串类型
a = "qwer"
print("_".join(a))

# 元组类型
b = ("a", "b", "c", "d")
print("=".join(b))

# 集合类型
c = {"qwe", "asd", "zxc"}
print(" ".join(c))


a = "    qweasdzxcrtqwe    "
print(a.strip())

b = "qweasdzxcrtqwe    "
print(b.lstrip("q"))

c = "   qweasdzxcrtqwe"
print(c.rstrip("qew"))

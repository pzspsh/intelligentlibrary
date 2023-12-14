# -*- encoding: utf-8 -*-
"""
@File   : main.py
@Time   : 2023-12-07 09:53:00
@Author : pan
"""
import os
import codecs
import configparser

# 获取当前文件目录的绝对路径
current_dir = os.path.abspath(os.path.dirname(__file__))
# 获取配置文件的绝对路径
config_path = os.path.join(os.path.dirname(current_dir), "config.ini")


class ReadBaseConfig:
    def __init__(self):
        fd = open(config_path, encoding="utf-8")
        data = fd.read()

        # 判断是否带有BOM文件，BOM文件通常是Excel格式，如果发现BOM_UTF8，则直接改写文件内容
        if data[:3] == codecs.BOM_UTF8:
            data = data[3:]
            file = codecs.open(config_path, "w")
            file.write(data)
            file.close()
        fd.close()

        # 实例化一个对象
        self.config = configparser.ConfigParser()
        # 读取文件
        self.config.read(config_path, encoding="utf-8")
        # 打印出配置文件的sections
        print(self.config.sections())


if __name__ == "__main__":
    pass

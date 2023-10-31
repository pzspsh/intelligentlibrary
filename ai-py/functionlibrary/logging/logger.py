# -*- encoding: utf-8 -*-
'''
@File   : logger.py
@Time   : 2023-06-01 17:09:46
@Author : pan
'''
import logging
import os
import sys
import threading
import time
import random
from colorama import Fore, Style, init
from logging.handlers import RotatingFileHandler

init(autoreset=True)

LogLevel = {
    "DEBUG": logging.DEBUG,
    "INFO": logging.INFO,
    "WARN": logging.WARN,
    "WARNING": logging.WARNING,
    "ERROR": logging.ERROR,
}

loggerLevel = LogLevel[str.upper("配置DEBUG")]


class LOGGER(object):

    def __init__(self, log_name, log_path=None):
        # 创建logger，如果参数为空则返回root logger
        self.path = log_path
        self.logger_ch = logging.getLogger(log_name)  # 控制台输出
        self.logger_ch.setLevel(loggerLevel)  # 设置logger日志等级
        if self.path:
            self.logger_fh = logging.getLogger(log_name + "_log")  # 日志记录
            self.logger_fh.setLevel(loggerLevel)  # 设置logger日志等级

        self.lock = threading.Lock()

        if not self.logger_ch.handlers:
            # 创建handler
            ch = logging.StreamHandler()

            # 设置输出日志格式
            formatter_ch = logging.Formatter(Fore.WHITE + "\r[%(asctime)s]  %(message)s", "%H:%M:%S")
            ch.setFormatter(formatter_ch)

            # 为logger添加的日志处理器
            self.logger_ch.addHandler(ch)

        if self.path and not self.logger_fh.handlers:
            # fh = logging.FileHandler(log_path, encoding="utf-8")
            handler = RotatingFileHandler(log_path, "a", 100 * 1024 * 1024, 5, "utf-8")

            time_f = "%Y-%m-%d %H:%M:%S" if log_name == "pan" else "%Y-%m-%d %H:%M:%S"
            formatter_fh = logging.Formatter("[%(asctime)s]  %(message)s", time_f)
            handler.setFormatter(formatter_fh)
            # 为handler指定输出格式
            self.logger_fh.addHandler(handler)

    def STDOUT(self, msg):
        m = ['-', '/', '\\', '|']
        pre = time.strftime(f'[%H:%M:%S] [{random.choice(m)}] -', time.localtime(time.time()))
        sys.stdout.write(pre + msg + "\r")
        sys.stdout.flush()

    def DEBUG(self, msg, write=True):
        self.lock.acquire()
        self.logger_ch.debug(Fore.WHITE + "[^] - " + str(msg) + Style.RESET_ALL)
        if self.path and write:
            self.logger_fh.debug("[^] - " + str(msg))
        self.lock.release()

    def INFO(self, msg, write=True):
        self.lock.acquire()
        self.logger_ch.info(Fore.CYAN + "[*] - " + str(msg) + Style.RESET_ALL)
        if self.path and write: self.logger_fh.info("[*] - " + str(msg))
        self.lock.release()

    def SUCCESS(self, msg):
        self.lock.acquire()
        self.logger_ch.warning(Fore.GREEN + "[+] - " + str(msg) + Style.RESET_ALL)
        if self.path: self.logger_fh.warning("[+] - " + str(msg))
        self.lock.release()

    def ERROR(self, msg):
        self.lock.acquire()
        self.logger_ch.error(Fore.RED + "[x] - " + str(msg) + Style.RESET_ALL)
        if self.path: self.logger_fh.error("[x] - " + str(msg))
        self.lock.release()

    def CRITICAL(self, msg):
        self.lock.acquire()
        self.logger_ch.critical(Fore.RED + "[!] - " + str(msg) + Style.RESET_ALL)
        if self.path: self.logger_fh.critical("[!] - " + str(msg))
        self.lock.release()


log_path = os.path.join("配置log文件路径", 'pan.log')
LOG = LOGGER('pan', log_path)
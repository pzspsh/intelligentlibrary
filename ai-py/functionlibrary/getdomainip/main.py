# -*- encoding: utf-8 -*-
"""
@File   : main.py
@Time   : 2025-03-10 16:45:07
@Author : pan
"""
import urllib.request

opener = urllib.request.build_opener()
opener.addheaders = [("User-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")]
try:
    response = opener.open("http://tpk.mwr.gov.cn", timeout=10)
    ip_port = response.fp.raw._sock.getpeername()
    print(ip_port)
except:
    print("解析不了")

# -*- encoding: utf-8 -*-
"""
@File   : verifyproxy.py
@Time   : 2024-11-15 14:19:29
@Author : pan
"""

import requests

proxystr = "36.151.192.236:41076"
proxys = f"http://{proxystr}"
url1 = "https://www.baidu.com"
url = "https://httpbin.org/get"
try:
    response = requests.get(url=url, proxies={"http": proxys, "https": proxys}, timeout=10)
    print(response.status_code)
    if response.status_code == 200:
        print("代理有效")
except Exception as err:
    print(err)

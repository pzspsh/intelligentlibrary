# -*- encoding: utf-8 -*-
"""
@File   : noproxy.py
@Time   : 2024-07-23 12:01:14
@Author : pan
"""
import socks
import socket
import requests
from requests.adapters import HTTPAdapter

requests.packages.urllib3.disable_warnings()
original_socket = socket.socket
# 保存原始的socket.socket，一定要保存原始的socket.socket再设置全局代理，然后在用原始socket恢复不调用代理


# 自定义一个不使用代理的HTTPAdapter
class UnproxiedAdapter(requests.adapters.HTTPAdapter):
    def init_poolmanager(self, *args, **kwargs):
        # 这里我们不需要特殊的代理设置，所以直接调用父类的init_poolmanager
        return super().init_poolmanager(*args, **kwargs)

    def proxy_headers(self, proxy):
        # 返回空的代理头，因为我们不使用代理
        return {}


def GlobalProxySet():
    try:
        socks.set_default_proxy(socks.SOCKS5, "223.113.54.165", 40092)  # 例如设置全局代理
        socket.socket = socks.socksocket
    except:
        pass


GlobalProxySet()

# 在需要不使用代理的请求之前，恢复原始的socket
socket.socket = original_socket  # 设置上面的socks5全局代理时，该设置有效

# 创建一个不使用代理的Session, 初始化Session，不恢复原始socket该请求还是会调用全局代理
s = requests.Session()
s.proxies = {}  # 设置上面的socks5全局代理时，该设置无效
s.trust_env = False
s.mount("http://", HTTPAdapter(max_retries=3))
s.mount("https://", HTTPAdapter(max_retries=3))
headers = {"Connection": "close"}
# 现在你可以使用这个session来发送不经过代理的请求
response_no_proxy = s.get("https://www.baidu.com", headers=headers, verify=False)
print(response_no_proxy.text)

# 在请求之后，如果你还想继续使用全局代理，可以重新设置socks.socksocket
socket.socket = socks.socksocket

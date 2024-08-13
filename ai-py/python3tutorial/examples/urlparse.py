# -*- encoding: utf-8 -*-
"""
@File   : urlparse.py
@Time   : 2024-08-12 15:38:52
@Author : pan
"""
from urllib.parse import urlparse, urlunparse, parse_qs

url = "http://www.example.com:8080/path/to/myfile.html?name=ferret&age=4"
parsed_url = urlparse(url)

print(parsed_url)
# 输出类似：ParseResult(scheme='http', netloc='www.example.com', path='/path/to/myfile.html', params='', query='name=ferret&age=4', fragment='')

# 访问各个部分
print(parsed_url.scheme)  # http
print(parsed_url.netloc)  # www.example.com
print(parsed_url.path)  # /path/to/myfile.html
print(parsed_url.query)  # name=ferret&age=4
print(parsed_url.port)

query_params = parse_qs(parsed_url.query)
print(query_params)
# 输出类似：{'name': ['ferret'], 'age': ['4']}

# 访问具体的参数
print(query_params["name"][0])  # ferret

# 构造新的URL
new_url_parts = ("https", "new.example.com", "/new/path", "", "key=value", "fragment")
new_url = urlunparse(new_url_parts)
print(new_url)  # https://new.example.com/new/path?key=value#fragment

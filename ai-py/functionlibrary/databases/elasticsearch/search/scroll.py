# -*- encoding: utf-8 -*-
"""
@File   : essearch.py
@Time   : 2024-06-14 14:20:00
@Author : pan
"""

# 数据量大于10000时，使用scroll
# https://www.elastic.co/guide/cn/elasticsearch/guide/current/_scroll_search.html
from elasticsearch import Elasticsearch


ES = Elasticsearch(["url"], http_auth=("user", "pass"))

index = "target-index"
scroll_duration = "1s"  # 滚动超时时间, 设置为1秒
size = 10000  # 每次滚动返回的文档数

# 初始搜索请求，并启用滚动
search_body = {
    "query": {"match_phrase": {"target-field": "search-taget"}},
    "size": size,
}

response = ES.search(index=index, body=search_body, scroll=scroll_duration)
scroll_id = response["_scroll_id"]
hits = response["hits"]["hits"]
count = 0
while len(hits):  # 滚动请求，直到没有文档返回
    for hit in hits:
        print(hit)
        count += 1
    response = ES.scroll(scroll_id=scroll_id, scroll=scroll_duration)
    hits = response["hits"]["hits"]
    # scroll_id = response["_scroll_id"]

print(count)
